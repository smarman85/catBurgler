package main

import (
	f "fmt"
	"github.com/smarman85/catBurgler/app/pkg/userPass"
	"os"
	// "github.com/smarman85/catBurgler/app/pkg/secrets"
)

var BASE_URL string = "https://0.0.0.0:33301/"
//var BASE_URL string = "http://0.0.0.0:33297/"
var AUTH_URL string = "v1/auth/userpass/login/"
var USER string = os.Args[1]
var PASSWD string = os.Args[2]

func main() {
	clientToken := userPass.Auth(BASE_URL, USER, PASSWD)
	f.Println(clientToken)
	//secrets.GetSecrets(clientToken, USER)
}
