package model

import (
	"github.com/kkwon1/webauthn-playground-service/protocol"
)

type PublicKeyCredentialCreationOptions struct {
	Challenge              protocol.Challenge       `json:"challenge"`
	RelyingParty           RelyingPartyEntity       `json:"rp"`
	User                   UserEntity               `json:"user"`
}
