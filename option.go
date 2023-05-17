package db

type Options struct {
	Driver     string `default:"mysql"`
	DataSource string
	DBName     string `json:"dbname"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       int    `json:"port"`

	MaxIdleConns    int `json:"maxidleconns"`
	MaxOpenConns    int `json:"maxopenconns"`
	ConnMaxLifetime int `json:"connMaxLifetime"`

	ReadTimeout  int `json:"readtimeout"`
	WriteTimeout int `json:"writetimeout"`
}
