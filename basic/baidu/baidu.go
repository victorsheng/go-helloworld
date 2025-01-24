package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
 "encoding/json"
//  "strings"
)

type Metric struct {
	Metric   string            `json:"metric"`    //sys.cpu.idle
	DataType string            `json:"data_type"` //COUNTER,GAUGE,DERIVE
	Value    interface{}       `json:"value"`     //99.00
	Tags     map[string]string `json:"tags"`      //{"product":"app01", "group":"dev02"}
}


const (
	COUNTER = "counter"
	GAUGE   = "gauge"
)

var infoKeys = map[string]string{
	"status":   GAUGE,
	"number_of_nodes" : GAUGE,
	"number_of_data_nodes" : GAUGE,
	"active_primary_shards" : GAUGE,
	"active_shards" : GAUGE,
	"relocating_shards" : GAUGE,
	"initializing_shards" : GAUGE,
	"unassigned_shards" : GAUGE,
	"delayed_unassigned_shards" : GAUGE,
	"number_of_pending_tasks" : GAUGE,
	"number_of_in_flight_fetch" : GAUGE,
	"task_max_waiting_in_queue_millis" : GAUGE,
  "active_shards_percent_as_number" : GAUGE}


func FormatBool(b bool) string {
    if b {
        return "true"
    }
    return "false"
}

func main()  {


  resultStr :=	health()
	
  var result map[string]interface{}
  json.Unmarshal([]byte(resultStr), &result)
  metrics := []Metric{}
  for key, value := range result {
  // Each value is an interface{} type, that is type asserted as a string
  
  if t, ok := infoKeys[key]; ok {
      
      metric := Metric{
				Metric:   fmt.Sprintf("es.%s", key),
				DataType:   t,
				Value:    value,
				Tags: map[string]string{
					"port": "1234",
				},
      }

      metrics = append(metrics, metric)
  }
fmt.Println(metrics)
}

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