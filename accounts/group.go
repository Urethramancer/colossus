package accounts

// Group for a site.
type Group struct {
	// ID in the database.
	ID int64
}

// CreateGroup creates and persists a new group for a site.
func (s *Site) CreateGroup(name string) error {
	return nil
}
