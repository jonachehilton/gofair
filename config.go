package gofair

// Config holds login data
type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AppKey   string `json:"api_key"`
	CertFile string `json:"ssl_cert"`
	KeyFile  string `json:"ssl_key"`
	Locale   string
}
