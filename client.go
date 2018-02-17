package gosoauthstoragedynamodb

import "github.com/RangelReale/osin"

type Client struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	UserData     interface{}
}

func (c *Client) GetId() string {
	return c.ClientId
}

func (c *Client) GetSecret() string {
	return c.ClientSecret
}

func (c *Client) GetRedirectUri() string {
	return c.RedirectUri
}

func (c *Client) GetUserData() interface{} {
	return c.UserData
}

func (c *Client) CopyFrom(client osin.Client) {
	c.ClientId = client.GetId()
	c.ClientSecret = client.GetSecret()
	c.RedirectUri = client.GetRedirectUri()
	c.UserData = client.GetUserData()
}
