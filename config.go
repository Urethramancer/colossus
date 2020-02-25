package main

// ColossusConfig points to file shares and pages to serve.
type ColossusConfig struct {
	// Static folder with HTML, CSS and Javascript. The structure is up to the user.
	Static string `json:"static"`
	// Shares are folders with or without password protection and expiry.
	// Each folder has a share.json file configuring its contents.
	Shares string `json:"shares"`
}

// Share folders share one or more files. Each has an optional subdirectory with access details.
type Share struct {
	// Folder set to true means share this folder.
	Folder bool `json:"folder"`
	// Path to the shared contents. This is a single file if Folder is false.
	// The folder can be anywhere in the filesystem.
	Path string `json:"path"`
	// Password for protecting this file or folder only with a password.
	// The user list won't be used if this is set. NOTE: This is a low-tech solution
	// with only a clear-text password.
	Password string `json:"password"`
	// Users will be "db" or "json". If configured to JSON a "users" folder will be used to hold accounts.
	Users string `json:"users"`
}
