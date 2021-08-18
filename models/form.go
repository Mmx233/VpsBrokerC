package models

import (
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

type HeartBeat struct {
	Type     string
	TargetIp string
	Time     int64
}

type Stat struct {
	Cpu struct {
		Num  int     `json:"num"`
		Used float64 `json:"used"`
	} `json:"cpu"`
	Mem struct {
		Total uint64  `json:"total"`
		Used  float64 `json:"used"`
	} `json:"mem"`
	Swap struct {
		Total uint64  `json:"total"`
		Used  float64 `json:"used"`
	} `json:"swap"`
	Disk struct {
		Total uint64
		Used  uint64 `json:"used"`
	} `json:"disk"`
}

func (a Stat) Gen() *Stat {
	_ = util.Try(func() error {
		t0, e := cpu.Info()
		if e == nil {
			a.Cpu.Num = len(t0)
		}
		return e
	}, 3, func(e error) {})

	_ = util.Try(func() error {
		t1, e := cpu.Percent(time.Second, false)
		if e == nil {
			a.Cpu.Used = t1[0]
		}
		return e
	}, 3, func(e error) {})

	_ = util.Try(func() error {
		t2, e := mem.VirtualMemory()
		if e == nil {
			a.Mem.Used = t2.UsedPercent
			a.Mem.Total = t2.Total
		}
		return e
	}, 3, func(e error) {})

	_ = util.Try(func() error {
		t3, e := mem.SwapMemory()
		if e == nil {
			a.Swap.Used = t3.UsedPercent
			a.Swap.Total = t3.Total
		}
		return e
	}, 3, func(e error) {})

	_ = util.Try(func() error {
		t4, e := disk.Usage("/")
		if e == nil {
			a.Disk.Total = t4.Total
			a.Disk.Used = t4.Used
		}
		return e
	}, 3, func(e error) {})

	return &a
}
