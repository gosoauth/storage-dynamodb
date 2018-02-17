package gosoauthstoragedynamodb

import (
	"time"

	"github.com/RangelReale/osin"
)

type AuthorizeData struct {
	Client              Client
	Code                string
	ExpiresIn           int32
	Scope               string
	RedirectUri         string
	State               string
	CreatedAt           time.Time
	UserData            interface{}
	CodeChallenge       string
	CodeChallengeMethod string
}

func (ad *AuthorizeData) CopyFrom(authorizeData *osin.AuthorizeData) {
	newClient := Client{}
	newClient.CopyFrom(authorizeData.Client)
	ad.Client = newClient

	ad.Code = authorizeData.Code
	ad.ExpiresIn = authorizeData.ExpiresIn
	ad.Scope = authorizeData.Scope
	ad.RedirectUri = authorizeData.RedirectUri
	ad.State = authorizeData.State
	ad.CreatedAt = authorizeData.CreatedAt
	ad.UserData = authorizeData.UserData
	ad.CodeChallenge = authorizeData.CodeChallenge
	ad.CodeChallengeMethod = authorizeData.CodeChallengeMethod
}

func (ad *AuthorizeData) Export() *osin.AuthorizeData {
	authorizeData := osin.AuthorizeData{}

	authorizeData.Client = &ad.Client
	authorizeData.Code = ad.Code
	authorizeData.ExpiresIn = ad.ExpiresIn
	authorizeData.Scope = ad.Scope
	authorizeData.RedirectUri = ad.RedirectUri
	authorizeData.State = ad.State
	authorizeData.CreatedAt = ad.CreatedAt
	authorizeData.UserData = ad.UserData
	authorizeData.CodeChallenge = ad.CodeChallenge
	authorizeData.CodeChallengeMethod = ad.CodeChallengeMethod

	return &authorizeData
}
