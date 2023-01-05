package main

import (
	"context"
	"fmt"
	"log"
	"os"

	appdcloudconnection_v1 "github.com/linchar/appd_cloud_api_client/connections_v1"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	ctx := context.Background()
	conf := clientcredentials.Config{
		ClientID:     os.Getenv("APPD_CLOUD_CLIENT_ID"),
		ClientSecret: os.Getenv("APPD_CLOUD_CLIENT_SECRET"),
		TokenURL:     "https://lab1.observe.appdynamics.com/auth/dfdbdf71-7322-44d0-85fb-36c69a7c3789/default/oauth2/token",
		AuthStyle:    oauth2.AuthStyleAutoDetect,
	}

	token, err := conf.Token(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	// fmt.Println("oauth successful here is the access token ", token.AccessToken)

	config := appdcloudconnection_v1.NewConfiguration()
	fmt.Printf("%T\n", config)
	config.AddDefaultHeader("Authorization", "Bearer "+token.AccessToken)

	client := appdcloudconnection_v1.NewAPIClient(config)
	fmt.Printf("%T\n", client)

	// resp, httpRes, err := client.ConnectionsApi.GetConnections(ctx).Execute()
	resp, httpRes, err := client.ConnectionsApi.GetConnections(ctx, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if httpRes.StatusCode != 200 {
		fmt.Printf("Error: %v\n", httpRes)
	}

	fmt.Printf("%T: %d\n", resp, len(resp.Items))
}
