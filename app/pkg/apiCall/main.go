package apiCall

import (
	//f "fmt"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	localCertFile = "/path/to/cert"
)

func Dial(url, method string, payload []byte) []byte {
	// cert info
	insecure := flag.Bool("insecure-ssl", false, "Accept/Ignore all server SSL certificates")
	flag.Parse()

	rootCAs, err := x509.SystemCertPool()
	check(err)
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	certs, err := ioutil.ReadFile(localCertFile)
	if err != nil {
		log.Fatalf("Failed to append %q to RootCAs: %v", localCertFile, err)
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}

	config := &tls.Config{
		InsecureSkipVerify: *insecure,
		RootCAs:            rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
  check(err)
	resp, err := client.Do(req)
  check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	return body

	//request, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	//check(err)
	//response,err := http.DefaultClient.Do(request)
	//check(err)
	//defer response.Body.Close()

	//body, err := ioutil.ReadAll(response.Body)
	//check(err)
	//return body
}
