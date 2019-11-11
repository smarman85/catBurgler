package userPass

import (
	//f "fmt"
	"encoding/json"
	"github.com/smarman85/catBurgler/app/pkg/apiCall"
)

var URL string = "v1/auth/userpass/login/"

type Password struct {
	Pass string `json:"password"`
}

type TokenData struct {
	Data AuthData `json:"auth"`
}

type AuthData struct {
	Token string `json:"client_token"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func encodePayload(userPassword string) []byte {
	password := Password{Pass: userPassword}
	data, err := json.Marshal(password)
	check(err)
	return data
}

func decodedToken(payload []byte) string {
	var data TokenData
	err := json.Unmarshal(payload, &data)
	check(err)
	return data.Data.Token
}

func Auth(baseUrl, username, password string) string {
  //f.Println(baseUrl+URL+username)
	dataPayload := encodePayload(password)
	rawData := apiCall.Dial(baseUrl+URL+username, "POST", dataPayload)
	return decodedToken(rawData)
}
