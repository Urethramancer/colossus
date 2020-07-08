package accounts

import (
	"database/sql"
	"fmt"

	"github.com/Urethramancer/colossus/internal/settings"
	_ "github.com/lib/pq"
)

// Manager handles the user database and related groups/profiles.
type Manager struct {
	settings.Settings
	*sql.DB
	// address
	DBHost string
	// port number
	DBPort string
	// name to connect to
	DBName string
	// user to connect as
	DBUser string
	// password to authenticate with
	DBPass string
	// SSL enabled or disabled
	SSL string
}

// NewMnager creates an account manager using the supplied database.
func NewManager(options ...func(*Manager)) (*Manager, error) {
	m := &Manager{}
	m.InitVars(map[string]string{
		ENVDBHOST: "localhost",
		ENVDBPORT: "5432",
		ENVDBNAME: "colossus",
		ENVDBUSER: "colossus",
		ENVDBPASS: "",
		ENVDBSSL:  "disable",
	})

	for _, o := range options {
		o(m)
	}

	conn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		m.Get(ENVDBHOST), m.Get(ENVDBPORT), m.Get(ENVDBNAME), m.Get(ENVDBUSER), m.Get(ENVDBPASS), m.Get(ENVDBSSL),
	)
	println(conn)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	m.DB = db
	err = m.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return m, nil
}
