package accounts

// User accounts are the core of the login system.
// ID, Name, Email and Password are the minimum required fields.
// ID is filled in by the system on creation.
type User struct {
	// ID is unique across sites.
	ID int64
	// Name is the username used to log in.
	Name string
	// Emmail for password resets and system messages.
	Email string
	// First name is optional.
	First string
	// Last name is optional.
	Last string
	// Profiles for different sites.
	Profiles []Profile
	// Admin is true if this user is a server administrator.
	Admin bool
}

// CreateUser creates a persistent user.
func (m *Manager) CreateUser(name, email, password string) error {
	return nil
}

// GetUser account.
func (m *Manager) GetUser(name string) (*User, error) {
	return nil, nil
}

// GetUserByID instead of name.
func (m *Manager) GetUserByID(id string) (*User, error) {
	return nil, nil
}

// GetProfile for a user/site combination.
func (m *Manager) GetProfile(name, site string) (*Profile, error) {
	return nil, nil
}

// GetProfileByID instead of name.
func (m *Manager) GetProfileByID(name, site string) (*Profile, error) {
	return nil, nil
}

// SetAdmin sets superuser status for a user.
func (m *Manager) SetAdmin(userid int64, b bool) error {
	return nil
}

// SetSiteAdmin sets site admin status for a user.
func (m *Manager) SetSiteAdmin(userid int64, b bool) error {
	return nil
}

// IsAdmin returns whether a user is a superuser with full access to the server.
func (m *Manager) IsAdmin(userid int64, site string) bool {
	return false
}

// IsSiteAdmin returns whether a user is a superuser on the specified site.
func (m *Manager) IsSiteAdmin(userid int64, site string) bool {
	return false
}
