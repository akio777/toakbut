package enum

import "fmt"

const (
	WFH = "WFH"
	WFO = "WFO"

	IN  = "IN"
	OUT = "OUT"

	BREAK = "BREAK"
	BACK  = "BACK"

	HELP = "HELP"

	EMPTY_COMMANDS      = " type `!help` and enter for showing all commands"
	INVALID_COMMANDS    = " invalid command, please try again or check `!help`"
	INVALID_IN_COMMANDS = " invalid `!in` command, please try again or check `!help`"

	ERROR_CREATE_ATTENDACE = " something went wrong with attendance database, please contact IT support."
)

func PrefixCommand(command string) string {
	return fmt.Sprintf("!%s", command)
}
