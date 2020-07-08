package accounts

// Site is a domain with its own set of user profiles and web pages.
type Site struct {
	// ID is generated at creation.
	ID int64
	// Owner is the user ID of the main contact for this site.
	Owner int64
	// Domain is the FQDN of the site.
	Domain string
}
