package models

type Config struct {
	AppPort  int16  `yaml:"app_port"`
	Username string `yaml:"db_username"`
	Password string `yaml:"db_password"`
	Name     string `yaml:"db_name"`
	Address  string `yaml:"db_address"`
	SSHMode  string `yaml:"db_ssh_mode"`
	Port     int16  `yaml:"db_port"`
}
