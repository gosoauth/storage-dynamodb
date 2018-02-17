package goauthstoragedynamodb

import (
	"fmt"

	"github.com/RangelReale/osin"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var Storage DynamoDBStorage

type DynamoDBStorage struct {
	db            *dynamo.DB
	clients       dynamo.Table
	authorizeData dynamo.Table
	accessData    dynamo.Table
}

func NewDynamoDBStorage() *DynamoDBStorage {
	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

	s := &DynamoDBStorage{
		db:            db,
		clients:       db.Table("OAuth-Clients"),
		authorizeData: db.Table("OAuth-AuthorizeData"),
		accessData:    db.Table("OAuth-AccessData"),
	}

	return s
}

func (s *DynamoDBStorage) Setup() {
	// TODO: Add Table Creation here
}

func (s *DynamoDBStorage) Clone() osin.Storage {
	return s
}

func (s *DynamoDBStorage) Close() {
}

func (s *DynamoDBStorage) GetClient(id string) (osin.Client, error) {
	fmt.Printf("GetClient: %s\n", id)

	var client Client

	err := s.clients.Get("ClientId", id).
		One(&client)

	if err != nil {
		fmt.Println(err.Error())
		return nil, osin.ErrNotFound
	}

	return &client, nil
}

func (s *DynamoDBStorage) SetClient(client osin.Client) error {
	fmt.Printf("SetClient: %s\n", client.GetId())

	newClient := Client{}
	newClient.CopyFrom(client)

	err := s.clients.Put(newClient).
		Run()

	if err != nil {
		fmt.Println(err.Error())
		return osin.ErrNotFound
	}

	return nil
}

func (s *DynamoDBStorage) SaveAuthorize(data *osin.AuthorizeData) error {
	fmt.Printf("SaveAuthorize: %s\n", data.Code)

	authorizeData := AuthorizeData{}
	authorizeData.CopyFrom(data)

	err := s.authorizeData.Put(authorizeData).
		Run()

	if err != nil {
		fmt.Println(err.Error())
		return osin.ErrNotFound
	}

	return nil
}

func (s *DynamoDBStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	fmt.Printf("LoadAuthorize: %s\n", code)

	var authorizeData AuthorizeData

	err := s.authorizeData.Get("Code", code).
		One(&authorizeData)

	if err != nil {
		fmt.Println(err.Error())
		return nil, osin.ErrNotFound
	}

	return authorizeData.Export(), nil
}

func (s *DynamoDBStorage) RemoveAuthorize(code string) error {
	fmt.Printf("RemoveAuthorize: %s\n", code)

	err := s.authorizeData.Delete("Code", code).
		Run()

	if err != nil {
		fmt.Println(err.Error())
		return osin.ErrNotFound
	}

	return nil
}

func (s *DynamoDBStorage) SaveAccess(data *osin.AccessData) error {
	fmt.Printf("SaveAccess: %s\n", data.AccessToken)

	accessData := AccessData{}
	accessData.CopyFrom(data)

	err := s.accessData.Put(accessData).
		Run()

	if err != nil {
		fmt.Println(err.Error())
		return osin.ErrNotFound
	}

	return nil
}

func (s *DynamoDBStorage) LoadAccess(code string) (*osin.AccessData, error) {
	fmt.Printf("LoadAccess: %s\n", code)

	var accessData AccessData

	err := s.accessData.Get("AccessToken", code).
		One(&accessData)

	if err != nil {
		fmt.Println(err.Error())
		return nil, osin.ErrNotFound
	}

	return accessData.Export(), nil
}

func (s *DynamoDBStorage) RemoveAccess(code string) error {
	fmt.Printf("RemoveAccess: %s\n", code)

	err := s.accessData.Delete("AccessToken", code).
		Run()

	if err != nil {
		fmt.Println(err.Error())
		return osin.ErrNotFound
	}

	return nil
}

func (s *DynamoDBStorage) LoadRefresh(code string) (*osin.AccessData, error) {
	fmt.Printf("LoadRefresh: %s\n", code)

	var accessData AccessData

	err := s.accessData.Get("RefreshToken", code).
		One(&accessData)

	if err != nil {
		fmt.Println(err.Error())
		return nil, osin.ErrNotFound
	}

	return s.LoadAccess(accessData.AccessToken)
}

func (s *DynamoDBStorage) RemoveRefresh(code string) error {
	fmt.Printf("RemoveRefresh: %s\n", code)

	accessData, err := s.LoadRefresh(code)

	if err != nil {
		return err
	}

	accessData.RefreshToken = ""

	err = s.SaveAccess(accessData)

	return err
}
