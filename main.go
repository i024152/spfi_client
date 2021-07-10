package main

import (
	"net/http"
	"time"
	"fmt"
	spfi_api "github.wdf.sap.corp/i024152/spfi_client/api"
)

func main() {

	client := &spfi_api.RestClient{
		Endpoint: "http://www.google.com",
		HttpClient: &http.Client{
			Timeout: time.Minute,
		},
	} 

	msg, err := client.Ping()

    if err != nil {
        fmt.Println(err)
		return
    }
	fmt.Println(msg)

}
