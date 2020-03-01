package main

// ColossusConfig points to file shares and pages to serve.
type ColossusConfig struct {
	// Static folder with HTML, CSS and Javascript. The structure is up to the user.
	Static string `json:"static"`
	// Shares are folders with or without password protection and expiry.
	// Each folder has a share.json file configuring its contents.
	Shares string `json:"shares"`
}
