package goauthstoragedynamodb

import (
	"fmt"
	"os"

	"github.com/guregu/dynamo"
)

func setupOAuthClientTable(client *dynamo.DB) {
	err := client.CreateTable("OAuth-Clients", Client{}).
		Provision(1, 1).
		Run()

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created the table OAuth-Clients")
}

func setupOAuthAuthorizeDataTable(client *dynamo.DB) {
	err := client.CreateTable("OAuth-AuthorizeData", AuthorizeData{}).
		Provision(1, 1).
		Run()

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created the table OAuth-AuthorizeData")
}

func setupOAuthAccessDataTable(client *dynamo.DB) {
	err := client.CreateTable("OAuth-AccessData", AccessData{}).
		Provision(1, 1).
		Run()

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created the table OAuth-AccessData")
}
