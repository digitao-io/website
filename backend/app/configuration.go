package app

type Configuration struct {
	Port       uint                      `json:"port"`
	Database   ConfigurationDatabase     `json:"database"`
	Objstorage ConfigurationObjstorage   `json:"objstorage"`
	Users      []ConfigurationCredential `json:"users"`
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

type ConfigurationObjstorage struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Bucket   string `json:"bucket"`
}

type ConfigurationCredential struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}
