package api

import (
	"net/http"
	"fmt"
	"io"
)

type RestClient struct {
	Endpoint    string
	HttpClient *http.Client
}

func (client *RestClient ) Ping() (string, error) {
	fmt.Printf("Getting " + client.Endpoint)
	resp, err := client.HttpClient.Get(client.Endpoint)
    if err != nil {
        fmt.Println(err)
		return "", err
    }

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
		return "", err
    }
	return string(body), nil

}
