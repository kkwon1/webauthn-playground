package model

type RelyingPartyEntity struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type UserEntity struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName,omitempty"`
	ID          string `json:"id"`
}
