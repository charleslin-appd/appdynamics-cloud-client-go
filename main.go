package main

import (
	"context"
	"fmt"
	"log"
	"os"

	appdcloudconnection_v1 "github.com/linchar/appd_cloud_api_client/connections_v1"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {

	// set up the struct according to clientcredentials package
	ccConfig := clientcredentials.Config{

		ClientID:     os.Getenv("APPD_CLOUD_CLIENT_ID"),
		ClientSecret: os.Getenv("APPD_CLOUD_CLIENT_SECRET"),
		TokenURL:     "https://lab1.observe.appdynamics.com/auth/dfdbdf71-7322-44d0-85fb-36c69a7c3789/default/oauth2/token",
	}

	// obtain the token
	token, err := ccConfig.Token(context.Background())

	if err != nil {
		log.Fatal(err.Error())
	}

	// Set up context with the token using API Context Key
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	auth := context.WithValue(context.Background(), appdcloudconnection_v1.ContextOAuth2, token)

	apiConfiguration := appdcloudconnection_v1.NewConfiguration()
	// Assign the Configures's HTTPClient to credential client, otherwise
	// it will default to the DefaultClient
	apiConfiguration.HTTPClient = ccConfig.Client(auth)
	// Create the API client with the configuration
	client := appdcloudconnection_v1.NewAPIClient(apiConfiguration)

	// make the call
	resp, httpRes, err := client.ConnectionsApi.GetConnections(auth, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if httpRes.StatusCode != 200 {
		fmt.Printf("Error: %v\n", httpRes)
	}

	fmt.Printf("%T: %d\n", resp, len(resp.Items))
}
