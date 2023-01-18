package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func curl(apiKey, refreshToken string) []byte {
	client := &http.Client{}
	jsonStr := fmt.Sprintf("{\"grant_type\":\"refresh_token\",\"refresh_token\":\"%s\"}", refreshToken)
	url := fmt.Sprintf("https://securetoken.googleapis.com/v1/token?key=%s", apiKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyText
}

func main() {
	apiKey := os.Args[1]
	refreshToken := os.Args[2]
	res := curl(apiKey, refreshToken)

	fmt.Println(string(res))

	/*
		var response Response

		json.Unmarshal(res, &response)

		fmt.Println(response.IdToken)
	*/
}

type Response struct {
	AccessToken  string `json:"access_token"`
	UserId       string `json:"user_id"`
	TokenType    string `json:"token_type"`
	ProjectId    string `json:"project_id"`
	IdToken      string `json:"id_token"`
	Registered   bool   `json:"registered"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
}
