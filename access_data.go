package gosoauthstoragedynamodb

import (
	"time"

	"github.com/RangelReale/osin"
)

type AccessData struct {
	Client        Client
	AuthorizeData *AuthorizeData
	AccessData    *AccessData
	AccessToken   string
	RefreshToken  string `globalIndex:"RefreshToken-index,hash"`
	ExpiresIn     int32
	Scope         string
	RedirectUri   string
	CreatedAt     time.Time
	UserData      interface{}
}

func (ad *AccessData) CopyFrom(accessData *osin.AccessData) {
	newClient := Client{}
	newClient.CopyFrom(accessData.Client)
	ad.Client = newClient

	newAuthorizeData := AuthorizeData{}
	newAuthorizeData.CopyFrom(accessData.AuthorizeData)
	ad.AuthorizeData = &newAuthorizeData

	if accessData.AccessData != nil {
		newAccessData := AccessData{}
		newAccessData.CopyFrom(accessData.AccessData)
		ad.AccessData = &newAccessData
	}

	ad.AccessToken = accessData.AccessToken
	ad.RefreshToken = accessData.RefreshToken
	ad.ExpiresIn = accessData.ExpiresIn
	ad.Scope = accessData.Scope
	ad.RedirectUri = accessData.RedirectUri
	ad.CreatedAt = accessData.CreatedAt
	ad.UserData = accessData.UserData
}

func (ad *AccessData) Export() *osin.AccessData {
	accessData := osin.AccessData{}

	accessData.Client = &ad.Client
	accessData.AuthorizeData = ad.AuthorizeData.Export()

	if ad.AccessData != nil {
		accessData.AccessData = ad.AccessData.Export()
	}

	accessData.AccessToken = ad.AccessToken
	accessData.RefreshToken = ad.RefreshToken
	accessData.ExpiresIn = ad.ExpiresIn
	accessData.Scope = ad.Scope
	accessData.RedirectUri = ad.RedirectUri
	accessData.CreatedAt = ad.CreatedAt
	accessData.UserData = ad.UserData

	return &accessData
}
