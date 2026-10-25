/*
Makes groups and leaves them instantly.
likely will not get flagged
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

var placeholderId = "1444366235222413386" /* required to make groups, put a random alt or persons id */

func groupSpam(userid string, token string, wg *sync.WaitGroup, workerId string) {
	defer wg.Done()

	type groupOut struct {
		ID   string `json:"id"`
	}

	var GroupOut groupOut
	body := `{"recipients":["` + userid + `","` + placeholderId + `"]}`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://discord.com/api/v9/users/@me/channels",
		bytes.NewBuffer([]byte(body)),
	); if err != nil { return }
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", token)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.6 Safari/605.1.15")
	res, err := http.DefaultClient.Do(req); if err != nil { return }
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		data, _ := io.ReadAll(res.Body)
		fmt.Printf("Channel creation failed: HTTP %d: %s\n",
			res.StatusCode, data)
		return
	}
	data, err := io.ReadAll(res.Body); if err != nil { return }
	json.Unmarshal([]byte(data), &GroupOut)
	groupId := GroupOut.ID

	if GroupOut.ID == "" {
		fmt.Println("No channel ID returned")
		return
	}

	req2, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		"https://discord.com/api/v10/channels/" + groupId,
		nil,
	); if err != nil { return }
	req2.Header.Set("content-type", "application/json")
	req2.Header.Set("authorization", token)
	req2.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.6 Safari/605.1.15")
	res2, err := http.DefaultClient.Do(req2); if err != nil { return }
	if res2.StatusCode < 300 {
		fmt.Printf("[+] WORKER ID: %s Group Created & Left Successfully.\n", workerId)
	} else {
		fmt.Printf("[-] WORKER ID: %s Failed\n", workerId)
	}
}

func main() {
	var token string
	var amount int
	var wg sync.WaitGroup
	var userid string

	fmt.Println("[?] Token : ")
	fmt.Scanln(&token)
	fmt.Println("[?] Victim ID : ")
	fmt.Scanln(&userid)
	fmt.Println("[?] Amount : ")
	fmt.Scanln(&amount)

	for i := 0; i < amount; i+=1 {
		wg.Add(1)
		fmt.Printf("\nWorker %d Has Started\n", i)
		go groupSpam(
			userid,
			token,
			&wg,
			string(i),
		)
	}

	wg.Wait()

	return;
}
