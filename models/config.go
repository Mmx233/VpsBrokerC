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
	Ip        string
	Port      uint
	SSL       bool
	AccessKey string
}
