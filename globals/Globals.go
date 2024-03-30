package globals

var (
	Mode               = ""
	MasterKey          = ""
	JWTKey             = ""
	SystemWhitelist    = []string{"127.0.0.1"}
	SystemWhitelistDNS []string
	Whitelist          []string
	WhitelistDNS       []string
	Read               []string
	Write              []string
	Internal                = []string{"admin"}
	Look               bool = true
)
