package apiCall

import (
  //f "fmt"
  "bytes"
  "io/ioutil"
  "net/http"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func Dial(url, method string, payload []byte) []byte {
  request, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
  check(err)
  response,err := http.DefaultClient.Do(request)
  check(err)
  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)
  check(err)
  return body
}
