package util

import "time"

func Try(e func() error, retry uint, retryMsg func(e error)) error {
	var num uint
	var err error
	for {
		num++
		err = e()
		if err != nil {
			if num < retry {
				retryMsg(err)
				time.Sleep(time.Second / 2)
				continue
			} else {
				return err
			}
		} else {
			break
		}
	}

	return nil
}
