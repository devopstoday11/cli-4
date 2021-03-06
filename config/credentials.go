package config

import (
	"encoding/json"

	"github.com/99designs/keyring"
)

type CredentialType struct {
	Label string
	Key   string
}

var (
	BuildkiteGraphQLToken = CredentialType{"Buildkite GraphQL Token", "graphql-token"}
	GithubOAuthToken      = CredentialType{"Github OAuth Token", "rest-token"}
)

func ListCredentials() []CredentialType {
	return nil
}

func StoreCredential(kr keyring.Keyring, t CredentialType, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return kr.Set(keyring.Item{
		Key:   t.Key,
		Label: t.Label,
		Data:  data,
	})
}

func RetrieveCredential(kr keyring.Keyring, t CredentialType, into interface{}) error {
	item, err := kr.Get(t.Key)
	if err != nil {
		return err
	}
	return json.Unmarshal(item.Data, into)
}
