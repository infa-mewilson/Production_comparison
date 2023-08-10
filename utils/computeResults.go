package utils

import (
	"NEWGOLANG/config"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

func SendESRequest(apis string, responsedatamap map[string]config.Value, query string) {

	// Load the client certificate and key
	cert, err := tls.LoadX509KeyPair("my_clientqa.pem", "my_clientqa.key")
	if err != nil {
		log.Fatal(err)
	}

	// Create a TLS configuration with only the client certificate and key
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true, // Skip certificate verification if CA certificate is not provided
	}

	// Create a HTTP client with the custom TLS configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	// Build the Elasticsearch URL with the API Gateway endpoint
	apiGatewayEndpoint := "https://iics-intcloud-prod-apigw-es.ext.prod.elk.cloudtrust.rocks"
	elasticSearchURL := apiGatewayEndpoint + "/filebeat-*-intcloud-*/_search"
	requestBody := bytes.NewBuffer([]byte(query))
	// Send a POST request with the query to Elasticsearch
	request, err := http.NewRequest("POST", elasticSearchURL, requestBody)
	if err != nil {
		log.Fatal(err)
	}
	apiKey := "soF1_e8kTtOlk90rzi0Qpw"
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-api-key", apiKey)
	//log.Println(request)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read and process the Elasticsearch response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Elasticsearch Response:")
	//fmt.Println(string(responseBody))

	var result config.ESResponse
	if err := json.Unmarshal(responseBody, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON", err)
	}
	var responsetimerounded95 float64
	var avgresponsetimerounded float64
	log.Println(result)
	responsetime95th := result.Aggregations.Num0.Values.Nine51
	//responsetime90th := result.Aggregations.Num0.Values.Nine00
	totalhits := result.Hits.Total.Value
	avgResponsetime := result.Aggregations.Num1.Value
	log.Println(totalhits, avgResponsetime)
	log.Println(responsetime95th)
	if responsetime95th == 0 {
		responsetimerounded95 = 0.0
	}
	if avgResponsetime == 0 {
		avgresponsetimerounded = 0.0
	}
	responsetimerounded95 = math.Round(responsetime95th*100) / 100
	avgresponsetimerounded = math.Round(avgResponsetime*100) / 100
	log.Println(responsetimerounded95, avgresponsetimerounded)
	responsedatamap[apis] = config.Value{AverageResponseTime: avgresponsetimerounded, Responsetime95: responsetimerounded95, TotalHits: totalhits}
}
