package svc

import (
	"strings"
	"time"
	"toakbut/pkg/bot/enum"
)

func IsInCommands(a string) bool {
	commands := []string{
		enum.BACK,
		enum.BREAK,
		enum.IN,
		enum.OUT,
		enum.HELP,
	}
	a = strings.ToUpper(a)
	itIs := false
	for _, cmd := range commands {
		if enum.PrefixCommand(cmd) == a {
			itIs = true
			break
		}
	}
	return itIs
}

func IsInWorkTypes(a string) bool {
	list := []string{
		enum.WFH,
		enum.WFO,
	}
	a = strings.ToUpper(a)
	itIs := false
	for _, t := range list {
		if (t) == a {
			itIs = true
			break
		}
	}
	return itIs
}

func parseTimeArgument(arg string) (time.Time, error) {
    if arg == "" {
        return time.Now(), nil
    }
    return time.Parse("15:04", arg)
}
