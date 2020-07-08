package accounts

import "time"

func init() {
	authtokens = TokenMap{
		Tokens: make(map[string]*Token),
	}
}

var authtokens TokenMap

// TokenMap contains the tokens for logged in users, and can be persisted to disk.
type TokenMap struct {
	// Tokens currently available.
	Tokens map[string]*Token `json:"tokens"`
}

// Token for a logged in user.
type Token struct {
	// ID contains the token.
	UserID int64 `json:"userid"`
	// Expire timestamp.
	Expire time.Time `json:"expire"`
	// IP address the user last used. Re-login is required if connecting from a different one.
	IP string `json:"ip"`
}
