package srv

const (
	// ENVHOST is the address to bind to.
	ENVHOST = "WEBHOST"
	// ENVPORT is the port to bind to.
	ENVPORT = "WEBPORT"
	// ENVDATA is where all configuration files and templates go, sorted per site.
	ENVDATA = "DATAPATH"
	// ENVSTATIC is the path to static files, which will be appended to each site's datapath.
	ENVSTATIC = "STATICPATH"
	// ENVSHARE is the location for shared files and folders.
	ENVSHARE = "SHAREPATH"
)
