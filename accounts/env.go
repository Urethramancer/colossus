package accounts

const (
	// ENVDBHOST is the address of the database server.
	ENVDBHOST = "DBHOST"
	// ENVDBPORT is the port to connect to the DB.
	ENVDBPORT = "DBPORT"
	// ENVDBNAME is the name of the database schema to use.
	ENVDBNAME = "DBNAME"
	// ENVDBUSER is the user to connect.
	ENVDBUSER = "DBUSER"
	// ENVDBPASS is the password for the specified user.
	ENVDBPASS = "DBPASS"
	// ENVDBSSL is either enable or disable.
	ENVDBSSL = "DBSSL"
	// ENVCOST is the cost setting for bcrypt passwords. 12+ is recommended.
	ENVCOST = "BCRYPT_COST"
	// ENVEMAIL is the e-mail address of the administrator, which will be used
	// when creating the initial admin user.
	ENVEMAIL = "ADMIN_EMAIL"
)
