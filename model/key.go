package model

import (
	"encoding/json"
	"time"
)

type GithubPGPKeys []struct {
	ID           int         `json:"id"`
	PrimaryKeyID interface{} `json:"primary_key_id"`
	KeyID        string      `json:"key_id"`
	RawKey       string      `json:"raw_key"`
	PublicKey    string      `json:"public_key"`
	Emails       []struct {
		Email    string `json:"email"`
		Verified bool   `json:"verified"`
	} `json:"emails"`
	Subkeys []struct {
		ID                int           `json:"id"`
		PrimaryKeyID      int           `json:"primary_key_id"`
		KeyID             string        `json:"key_id"`
		RawKey            interface{}   `json:"raw_key"`
		PublicKey         string        `json:"public_key"`
		Emails            []interface{} `json:"emails"`
		Subkeys           []interface{} `json:"subkeys"`
		CanSign           bool          `json:"can_sign"`
		CanEncryptComms   bool          `json:"can_encrypt_comms"`
		CanEncryptStorage bool          `json:"can_encrypt_storage"`
		CanCertify        bool          `json:"can_certify"`
		CreatedAt         time.Time     `json:"created_at"`
		ExpiresAt         interface{}   `json:"expires_at"`
	} `json:"subkeys"`
	CanSign           bool      `json:"can_sign"`
	CanEncryptComms   bool      `json:"can_encrypt_comms"`
	CanEncryptStorage bool      `json:"can_encrypt_storage"`
	CanCertify        bool      `json:"can_certify"`
	CreatedAt         time.Time `json:"created_at"`
	ExpiresAt         time.Time `json:"expires_at"`
}

func (gpk GithubPGPKeys) JSON() string {
	gJSON, err := json.MarshalIndent(gpk, "", "  ")
	if err != nil {
		return ""
	}
	return string(gJSON)
}

func (gpk GithubPGPKeys) FILTER(email string) GithubPGPKeys {
	/*var githubResp model.GithubPGPKeys
	for _, key := range gpk {
		shouldInclude := false
		for _, emailItem := range key.Emails {
			if emailItem.Email == email {
				return key
			}
		}
	}*/
	return gpk
}
