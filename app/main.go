package main

import (
  "os"
  "github.com/smarman85/catBurgler/app/pkg/userPass"
  "github.com/smarman85/catBurgler/app/pkg/secrets"
)

var BASE_URL string = "http://0.0.0.0:8200/"
var AUTH_URL string = "v1/auth/userpass/login/"
var USER string = os.Args[1]
var PASSWD string = os.Args[2]

func main() {
  clientToken := userPass.Auth(USER, PASSWD)
  secrets.GetSecrets(clientToken, "player_one")
}
