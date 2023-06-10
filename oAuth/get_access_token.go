package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)
  
func main() {
	
	url := "https://mpc-provider.uk.auth0.com/oauth/token"
  
	payload := strings.NewReader("{\"client_id\":\"5sfLiQxSEWZh2twQDKFPS9JAQaxdnjnO\",\"client_secret\":\"A9MScXKcdvvOyqjLoa79sNRRdMCm76L3PCf5CfakuTWjwiNrhzLmwoar-P9gFE6X\",\"audience\":\"https://mpc-provider.uk.auth0.com/api/v2/\",\"grant_type\":\"client_credentials\"}")
  
	req, _ := http.NewRequest("POST", url, payload)
  
	req.Header.Add("content-type", "application/json")
  
	res, _ := http.DefaultClient.Do(req)
  
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
  
	fmt.Println(res)
	fmt.Println(string(body))
}