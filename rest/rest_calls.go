package rest

import (
    "fmt"
    "net/http"
    "encoding/json"
    "bytes"
    "strings"
    //"restcli/parser"
)


func Get_Rest(ip,port,ca,ce,key,endpoint string) {
    correctedEndpoint := correctEndpoint(endpoint) 
    url := "https://"+ip+":"+port+correctedEndpoint
    req, _ := http.NewRequest("GET", url, nil)
    value,_ := RestApiCaller(req,ca,ce,key)
    fmt.Println(string(value) )
}

func Post_Rest(ip,port,ca,ce,key,endpoint,data_p string) {
    correctedEndpoint := correctEndpoint(endpoint) 
    url := "https://"+ip+":"+port+correctedEndpoint
    data := json.RawMessage(data_p) 
    payloadBytes,_ := json.Marshal(data)
    body := bytes.NewReader(payloadBytes)
    req, _ := http.NewRequest("POST", url, body)
    value,_ := RestApiCaller(req,ca,ce,key)
    fmt.Println(string(value) )
    //parser.Parse_json(value)
}

type Payload struct {
	Query []interface{} `json:"query"`
}

func correctEndpoint(url string) (string){
	if !strings.HasPrefix(url,"/"){
       url = "/"+url
	}

	if  strings.HasSuffix(url,"/") {
	      url = url[:len(url)-len("/")]
	}
    return url
}
