package accounts

import (
	"crypto/rand"
	"database/sql"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

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
	// Password is the result of bcrypt encoding.
	Password string
	// Fullname is optional.
	Fullname string
	// Admin is true if this user is a server administrator.
	Admin bool
	// MustChange password on next login.
	MustChange bool
}

var validchars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789#_-.,")

// randpass generates a fairly safe, simple password string of the desired length.
func randpass(size int) string {
	// Somewhat password-friendly.
	pw := make([]byte, size)
	for i := 0; i < size; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(validchars))))
		if err != nil {
			return ""
		}
		c := validchars[n.Int64()]
		pw[i] = c
	}
	return string(pw)
}

// CreateRootUser creates a system administrator and returns the generated password.
func (m *Manager) CreateRootUser(name, email string) (string, error) {
	pw := randpass(14)
	id, err := m.CreateUser(name, email, pw)
	if err != nil {
		return "", err
	}

	err = m.SetAdmin(id, true)
	return pw, err
}

// CreateUser creates a persistent user.
// The password will be truncated to 50 characters if longer due to bcrypt limitations.
func (m *Manager) CreateUser(name, email, password string) (int64, error) {
	pw := password
	if len(pw) > 50 {
		pw = pw[:50]
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	var id int64
	err = m.QueryRow("INSERT INTO public.users (name,email,password) VALUES ($1,$2,$3) RETURNING id;", name, email, string(hash)).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateUser sets the user details for the user ID in the specified struct.
func (m *Manager) UpdateUser(user *User) error {
	_, err := m.Exec("UPDATE TABLE public.users SET name=$1,email=$2,password=$3,fullname=$4,admin=$5,mustchange=$6;",
		user.Name, user.Email, user.Password, user.Fullname, user.Admin, user.MustChange,
	)
	return err
}

// GetUsers gets up to max users.
func (m *Manager) GetUsers(max int) ([]User, error) {
	rows, err := m.Query("SELECT id,name,email,password,fullname,admin,mustchange FROM public.users LIMIT $1;", max)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	list := []User{}
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Fullname, &u.Admin, &u.MustChange)
		if err != nil {
			return nil, err
		}

		list = append(list, u)
	}

	return list, nil
}

// GetUser account.
func (m *Manager) GetUser(name string) (*User, error) {
	u := &User{}
	err := m.QueryRow("SELECT id,name,email,password,fullname,admin,mustchange WHERE name=$1;", name).Scan(
		&u.ID, &u.Name, &u.Email, &u.Password, &u.Fullname, &u.Admin, &u.MustChange,
	)
	return u, err
}

// GetUserByID instead of name.
func (m *Manager) GetUserByID(id int64) (*User, error) {
	u := &User{}
	err := m.QueryRow("SELECT id,name,email,password,fullname,admin,mustchange WHERE id=$1;", id).Scan(
		&u.ID, &u.Name, &u.Email, &u.Password, &u.Fullname, &u.Admin, &u.MustChange,
	)
	return u, err
}

func (m *Manager) CreateProfile(p *Profile) error {
	return nil
}

// GetProfile for a user/site combination.
func (m *Manager) GetProfile(name, site string) (*Profile, error) {
	return nil, nil
}

// GetProfileByID instead of name.
func (m *Manager) GetProfileByID(id, site string) (*Profile, error) {
	return nil, nil
}

// SetAMustchange makes it so the user must change password on next login.
func (m *Manager) SetMustchange(userid int64, b bool) error {
	u, err := m.GetUserByID(userid)
	if err != nil {
		return err
	}

	u.MustChange = b
	return m.UpdateUser(u)
}

// SetAdmin sets superuser status for a user.
func (m *Manager) SetAdmin(userid int64, b bool) error {
	u, err := m.GetUserByID(userid)
	if err != nil {
		return err
	}

	u.Admin = b
	return m.UpdateUser(u)
}

// SetSiteAdmin sets site admin status for a user.
func (m *Manager) SetSiteAdmin(userid int64, site string, b bool) error {
	return nil
}

// IsAdmin returns whether a user is a superuser with full access to the server.
func (m *Manager) IsAdmin(userid int64, site string) bool {
	var b sql.NullBool
	err := m.QueryRow("SELECT admin FROM public.users WHERE id=$1;", userid).Scan(&b)
	if err != nil {
		return false
	}

	return b.Bool
}

// IsSiteAdmin returns whether a user is a superuser on the specified site.
func (m *Manager) IsSiteAdmin(userid int64, site string) bool {
	return false
}

// VerifyPassword returns true if the supplied passsord is a match for the user.
func (m *Manager) VerifyPassword(u *User, pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
	if err != nil {
		return false
	}

	return true
}
