package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
 "encoding/json"
)

func main()  {
	result :=	health()
	fmt.Println(result)
  
}

func health() string {
	url := "http://10.16.17.43:9200/_cluster/health?pretty"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
  }
  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  result :=string(body)
  
  return result;
}