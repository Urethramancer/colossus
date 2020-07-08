package accounts

// Profile contains the information to associate a User with a site in the database.
type Profile struct {
	// User this profile belongs to.
	User int64
	// Site this profile applies to.
	Site string
	// Groups this profile is a member of.
	Groups []string
	// Admin is true if this profile is a site administrator.
	Admin bool
}

// CreateProfile creates, persists and returns a new profile for a user, or an error.
func (u *User) CreateProfile(site string) error {
	return nil
}
