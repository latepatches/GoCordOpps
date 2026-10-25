package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func createRole(name string, token string, guildId string, wg *sync.WaitGroup) {
    defer wg.Done()
    body := `{"name":"` + name + `","color":0,"colors":{"primary_color":0,"secondary_color":null,"tertiary_color":null},"permissions":"0"}`
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    req, err := http.NewRequestWithContext(
        ctx,
        http.MethodPost,
        "https://discord.com/api/v9/guilds/"+guildId+"/roles",
        bytes.NewBuffer([]byte(body)),
    ); if err != nil { return }
    req.Header.Set("content-type", "application/json")
    req.Header.Set("authorization", token)
    req.Header.Set("user-agent", "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.6 Safari/605.1.15")
    res, err := http.DefaultClient.Do(req); if err != nil { return }
    defer res.Body.Close()
    data, err := io.ReadAll(res.Body)
    if res.StatusCode < 300 {
	fmt.Println("[+] Role Created Successfully ")
    } else {
	fmt.Println("[-] ", data)
    }
}

func main() {
	var wg sync.WaitGroup
	var token string
	var name string
	var guildId string
	var amt int
	fmt.Print("[?] Token : ")
	fmt.Scanln(&token)
	fmt.Print("[?] Name : ")
	fmt.Scanln(&name)
	fmt.Print("[?] Guild ID : ")
	fmt.Scanln(&guildId)
	fmt.Print("[?] Amount : ")
	fmt.Scanln(&amt)
	for i := 0; i < amt; i++ {
		wg.Add(1)
		go createRole(
			name,
			token,
			guildId,
			&wg,
		)
	}

	wg.Wait()
}
