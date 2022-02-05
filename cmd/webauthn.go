package cmd

import (
	"github.com/kkwon1/webauthn-playground-service/model"
	"github.com/kkwon1/webauthn-playground-service/protocol"
	"github.com/segmentio/ksuid"
)

func CreatePublicKeyCredentialCreationOption(id string, username string) model.PublicKeyCredentialCreationOptions {
	rp := model.RelyingPartyEntity{
		Name: "localhost",
		ID:   "localhost",
	}

	userEntity := model.UserEntity{
		Name:        "Kevin",
		DisplayName: "Kevin",
		ID:          ksuid.New().String(),
	}

	options := model.PublicKeyCredentialCreationOptions{
		Challenge:    protocol.CreateChallenge(),
		RelyingParty: rp,
		User:         userEntity,
		Parameters:   defaultRegistrationCredentialParameters(),
	}

	return options
}

func defaultRegistrationCredentialParameters() []protocol.CredentialParameter {
	return []protocol.CredentialParameter{
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgES256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgES384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgES512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgRS256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgRS384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgRS512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgPS256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgPS384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgPS512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: protocol.AlgEdDSA,
		},
	}
}