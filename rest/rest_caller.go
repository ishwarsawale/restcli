package rest

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)


func RestApiCaller(req *http.Request,caFile,certFile,keyFile string) ([]uint8,int){
	
  // Load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
  tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
    InsecureSkipVerify: true,
	}
  //add header
  req.Header.Add("content-type", "application/json")
  req.Header.Add("cache-control", "no-cache")
  
  //tls
  tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}
	resp, _ := client.Do(req)
  defer resp.Body.Close()
    StatusCode := resp.StatusCode
    body, _ := ioutil.ReadAll(resp.Body)
    return (body),StatusCode
}

