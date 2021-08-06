package models

type VpsInit struct {
	Name string
	Port uint
}

type HeartBeat struct {
	Type     string
	TargetIp string
	Time     int64
}
