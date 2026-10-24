/*
Creates Webhooks
Spams Each One
simple as that
*/
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var resp struct {
	URL string `json:"url"`
}

func extractWebhooks(token string, array []string, wg *sync.WaitGroup, cid string) []string {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := `{"name": "eHook"}`
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"https://discord.com/api/v9/channels/" + cid + "/webhooks",
		bytes.NewBuffer([]byte(body)),
	)
	if err != nil {
		return array
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", token)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.6 Safari/605.1.15")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return array
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return array
	}
	json.Unmarshal(data, &resp)
	array = append(array, resp.URL)

	return array
}

func postMessage(webhook string, wg *sync.WaitGroup, body string) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		webhook,
		bytes.NewBuffer([]byte(body)),
	)
	if err != nil {
		fmt.Println("Failed")
	}
	req.Header.Set("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Failed at res")
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed at readall")
	}
	fmt.Println("[+] Post Was Made " + string(data))
}

func main() {
	var wg sync.WaitGroup
	var array []string
	var opt string
	var body string
	var cid string
	var token string

	fmt.Println("[?] Message : ")
	fmt.Scanln(&opt)
	fmt.Println("[?] Channel ID : ")
	fmt.Scanln(&cid)
	fmt.Println("[?] Token : ")
	fmt.Scanln(&token)

	body = `{"content": "` + opt + `"}`

	for i := 0; i < 15; i += 1 {
		wg.Add(1)
		array = extractWebhooks(
			token,
			array,
			&wg,
			cid,
		)
	}
	wg.Wait()

	for _, line := range array {
		fmt.Println("[+] " + line)
	}

	for i := 0; i < 15; i+=1 {
		for _, line := range array {
			wg.Add(1)
			go postMessage(line, &wg, body)
		}

		wg.Wait()
	}

	return;
}
