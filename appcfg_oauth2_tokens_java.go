package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type User struct {
	AccessToken          string `json:"access_token"`
	ExpirationTimeMillis int    `json:"expiration_time_millis"`
	RefreshToken         string `json:"refresh_token"`
}

type Credential struct {
	UserName User `json:"ubuntu"`
}

type OAuth2Token struct {
	Credentials Credential `json:"credentials"`
}

func main() {
	json, err := CreateTokenJSONText()
	if err != nil {
		fmt.Println(err)
		return
	}

	CreateTokenFile(json)
}

func CreateTokenJSONText() (jsonText string, err error) {
	expirationTimeMillis, err := strconv.Atoi(os.Getenv("APP_CFG_EXPIRATION_TIME_MILLIS"))
	if err != nil {
		fmt.Println(err)
		return
	}

	user := User{
		os.Getenv("APP_CFG_ACCESS_TOKEN"),
		expirationTimeMillis,
		os.Getenv("APP_CFG_REFRESH_TOKEN"),
	}

	credential := Credential{user}

	value := OAuth2Token{credential}
	text, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
		return
	}

	return string(text), nil
}

func CreateTokenFile(jsonText string) {
	tokenFilePath := os.Getenv("HOME") + "/.appcfg_oauth2_tokens_java"
	text := []byte(jsonText + "\n")
	err := ioutil.WriteFile(tokenFilePath, text, 0660)
	if err != nil {
		fmt.Println(tokenFilePath, err)
		return
	}
}
