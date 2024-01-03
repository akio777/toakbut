package utils

import "time"

func TimeFormat(_time time.Time) string {
	return _time.Format("02/01/2006 15:04")
}
