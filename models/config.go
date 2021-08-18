package models

type Config struct {
	Settings Settings
	Remote   Remote
}

type Settings struct {
	Port uint
	Name string
}

type Remote struct {
	Host      string
	Port      uint
	SSL       bool
	AccessKey string
}
