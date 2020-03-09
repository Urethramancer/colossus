package main

// User config.
type User struct {
	// Name is the login name.
	Name string `json:"name"`
	// Password is a hash.
	Password string `json:"password"`
	// Admin is a list of shares where this user has upload privileges.
	Admin []string `json:"admin"`
}
