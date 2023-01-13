package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// メールアドレスとパスワードによるログイン
// https://cloud.google.com/identity-platform/docs/use-rest-api#section-sign-in-email-password

func curl(apiKey, email, pw string) []byte {
	client := &http.Client{}
	jsonStr := fmt.Sprintf("{\"email\":\"%s\",\"password\":\"%s\",\"returnSecureToken\":true}", email, pw)
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", apiKey)
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
	email := os.Args[2]
	pw := os.Args[3]
	res := curl(apiKey, email, pw)

	var response Response
	json.Unmarshal(res, &response)

	fmt.Println(response.IdToken)

}

type Response struct {
	Kind         string `json:"kind"`
	LocalId      string `json:"localId"`
	Email        string `json:"email"`
	DisplayName  string `json:"displayName"`
	IdToken      string `json:"idToken"`
	Registered   bool   `json:"registered"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}
