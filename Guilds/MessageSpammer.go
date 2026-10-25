/*
really basic, will get ratelimited but works.
likely to not get flagged
*/

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func postMessage(wg *sync.WaitGroup, message string, token string, channelid string) {
	defer wg.Done()
	body := `{"content": "`+message+  `"}`
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	); defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://discord.com/api/v9/channels/" + channelid + "/messages",
		bytes.NewBuffer([]byte(body)),
	); if err != nil { log.Fatal(err) }
	req.Header.Set("Authorization", token)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.6 Safari/605.1.15")
	res, err := http.DefaultClient.Do(req); if err != nil { log.Fatal(err) }
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Printf("[+] Post Made # %d, %s", res.StatusCode, string(data))
}

func main() {
	var wg sync.WaitGroup
	var token string
	var message string
	var channelid string
	var amount int

	fmt.Println("[?] Token : ")
	fmt.Scanln(&token)
	fmt.Println("[?] Channel ID : ")
	fmt.Scanln(&channelid)
	fmt.Println("[?] Message : ")
	fmt.Scanln(&message)
	fmt.Println("[?] Amount : ")
	fmt.Scanln(&amount)

	for i := 0; i < amount; i+=1 {
		wg.Add(1)
		go postMessage(
			&wg,
			message,
			token,
			channelid,
		)
	}

	wg.Wait()

	return;
}
