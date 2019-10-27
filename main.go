package main

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "os"
  "net/http"
  "github.com/smarman85/catBurgler/pkg/secrets"
)

var BASE_URL string = "http://0.0.0.0:8200/"
var AUTH_URL string = "v1/auth/userpass/login/"
var USER string = os.Args[1]
var PASSWD string = os.Args[2]

func check(e error) {
  if e != nil {
    panic(e)
  }
}

// structs:
type Password struct {
  Pass string `json:"password"`
}

type AuthType struct {
  Data Authentication `json:"auth"`
}

type Authentication struct {
  Token string `json:"client_token"`
}

func apiCall(endpoint, user string, password []byte) []byte {
  req, err := http.NewRequest("POST", BASE_URL+endpoint+user, bytes.NewBuffer(password))
  check(err)
  resp, err := http.DefaultClient.Do(req)
  check(err)
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  check(err)
  return body
}

func dataPayload(pass string) []byte {
  password := Password{Pass:pass}
  data, err := json.Marshal(password)
  check(err)
  return data
}

func decodeByte(payload []byte) string {
  var data AuthType
  err := json.Unmarshal(payload, &data)
  check(err)
  //fmt.Println("Client Token:\t" + data.Data.Token)
  return data.Data.Token
}


func main() {
  dataPass := dataPayload(PASSWD)
  authResponse := apiCall(AUTH_URL, USER, dataPass)
  clientToken := decodeByte(authResponse)
  //fmt.Println("Client Token:\t", clientToken)
  secrets.GetSecrets(clientToken, "player_one")
}
