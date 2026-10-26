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

func createChannel(token string, name string, guildId string, wg *sync.WaitGroup) {
	defer wg.Done()
	body := `{"type":0,"name":"`+name+`"}`
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second); defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost ,"https://discord.com/api/v9/guilds/"+guildId+"/channels", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", token)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.6 Safari/605.1.15")
	if err != nil { return }
	res, err := http.DefaultClient.Do(req); if err != nil { return }; defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	if res.StatusCode < 300 {
		fmt.Println("[+] Channel Created " + string(data))
	} else {
		fmt.Println("[-] " + string(data))
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
	fmt.Print("[?] Channel Name : ")
	fmt.Scanln(&name)
	fmt.Print("[?] Guild ID : ")
	fmt.Scanln(&guildId)
	fmt.Print("[?] Amount : ")
	fmt.Scanln(&amt)

	for i := 0; i < amt; i++ {
		wg.Add(1)
		go createChannel(
			token,
			name,
			guildId,
			&wg,
		)
	}

	wg.Wait()

}
