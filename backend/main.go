package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

type EthRPCRequest struct {
    JSONRPC string        `json:"jsonrpc"`
    Method  string        `json:"method"`
    Params  []interface{} `json:"params"`
    ID      int           `json:"id"`
}

type EthRPCResponse struct {
    Result string `json:"result"`
    ID     int    `json:"id"`
}



func loadEnvVar() string{
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Get the variable from .env file
    apiKey := os.Getenv("INFURA_API_KEY")
	return apiKey 
}




func queryGas(url string)*big.Float{
    request := EthRPCRequest{
        JSONRPC: "2.0",
        Method:  "eth_gasPrice",
        Params:  []interface{}{},
        ID:      1,
    }

    // Encode the request as JSON
    requestBody, err := json.Marshal(request)
    if err != nil {
        fmt.Println("Error encoding request:", err)
        return nil 
    }

    // Send the HTTP request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        fmt.Println("Error sending request:", err)
        return nil
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return nil
    }

    // Decode the JSON-RPC response
    var rpcResponse EthRPCResponse
    err = json.Unmarshal(body, &rpcResponse)
    if err != nil {
        fmt.Println("Error decoding response:", err)
        return nil
    }

	hex_res := strings.TrimPrefix(rpcResponse.Result, "0x")
	decimalValue := new(big.Int)
    decimalValue, success := decimalValue.SetString(hex_res, 16)
    if !success {
        log.Fatal("Invalid hexadecimal string")
    }

    weiFloat := new(big.Float).SetInt(decimalValue) 
    gweiFloat := new(big.Float).Quo(weiFloat, big.NewFloat(1e9)) 

    return gweiFloat 
}




func main() {
	
	API_KEY := loadEnvVar() 
    url := "https://mainnet.infura.io/v3/" + API_KEY 

    stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				gasPrice := queryGas(url)
				if gasPrice == nil {
					log.Printf("Error querying ETH node - Null price")
					continue
				}
				fmt.Printf("Gas Price: %f\n", gasPrice)
			}
		}
	}()

	<-stopChan // Wait for a signal to stop
	fmt.Println("Shutting down...")

    
}
