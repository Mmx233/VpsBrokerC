package util

type tool struct{}

var Tool tool

func (*tool) Try(e func() error, t uint) {
	for t > 0 {
		if e() != nil {
			t--
		} else {
			break
		}
	}
}
