/*
gets ratelimited unless proxies are involved
*/
package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
	"fmt"
)

func postMessage(message string, webhook string, wg *sync.WaitGroup) {
	defer wg.Done()
	body := `{
  "content": "`+message+`",
  "embeds": null,
  "username": "github.com/latepatches",
  "attachments": []
}`
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second); defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, webhook, bytes.NewBuffer([]byte(body))); if err != nil {log.Fatal(err)}
	req.Header.Set("content-type", "application/json")
	res, err := http.DefaultClient.Do(req); if err != nil {log.Fatal(err)}
	data, err := io.ReadAll(res.Body); if err != nil {log.Fatal(err)}
	if res.StatusCode < 300 {
		fmt.Println("[+] Success " + string(res.StatusCode))
	} else {
		fmt.Println("[-] " + string(data))
	}
}

func main() {
	var wg sync.WaitGroup
	var webhook string
	var message string
	var amt int

	fmt.Print("[.] Webhook : ")
	fmt.Scanln(&webhook)
	fmt.Print("[.] Message : ")
	fmt.Scanln(&message)
	fmt.Print("[.] Amount : ")
	fmt.Scanln(&amt)

	for i := 0; i < amt; i++ {
		wg.Add(1)
		go postMessage(
			message,
			webhook,
			&wg,
		)
	}

	wg.Wait()
}
