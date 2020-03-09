package main

// SetDatabase sets the connection string of the database.
func (ws *Server) SetDatabase(host, port, name, user, pass, ssl string) {
	ws.dbhost = host
	ws.dbport = port
	ws.dbname = name
	ws.dbuser = user
	ws.dbpass = pass
	ws.ssl = ssl
}
