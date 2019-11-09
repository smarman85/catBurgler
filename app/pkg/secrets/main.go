package secrets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//var baseEndpoint string = "http://0.0.0.0:8200/v1/secret/data/users/"
var baseEndpoint string = "https://0.0.0.0:33284/v1/secret/data/users/"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func apiCall(url, token string) []byte {
	//fmt.Println(url)
	//fmt.Println(token)
	request, err := http.NewRequest("GET", url, nil)
	check(err)
	request.Header.Set("X-Vault-Token", token)

	resp, err := http.DefaultClient.Do(request)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	return body
}

type Payload struct {
	Data Info `json:"data"`
}

type Info struct {
	Envelope map[string]string `json:"data"`
}

func decodeData(payload []byte) map[string]string {
	var data Payload
	err := json.Unmarshal(payload, &data)
	check(err)
	//fmt.Println(data.Data.Envelope)
	return data.Data.Envelope
}

func wikiLeak(payload map[string]string) {
	for k, v := range payload {
		fmt.Printf("export %s=\"%s\"\n", k, v)
	}
}

func GetSecrets(clientToken, secret string) {
	//fmt.Println(clientToken)
	rawData := apiCall(baseEndpoint+secret, clientToken)
	//fmt.Println(string(rawData))
	secrets := decodeData(rawData)
	wikiLeak(secrets)
}
