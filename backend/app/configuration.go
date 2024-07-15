package app

type Configuration struct {
	Port     uint                      `json:"port"`
	Database ConfigurationDatabase     `json:"database"`
	Users    []ConfigurationCredential `json:"users"`
}

type ConfigurationDatabase struct {
	Host               string `json:"host"`
	Port               int    `json:"port"`
	User               string `json:"user"`
	Password           string `json:"password"`
	Database           string `json:"database"`
	MaxIdleConnections int    `json:"maxIdleConnections"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
}

type ConfigurationCredential struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}
