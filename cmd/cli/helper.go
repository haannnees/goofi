package cli

import "strings"

func FlatArgs(args []string) string {
	length := len(args)
	sb := strings.Builder{}
	for i, arg := range args {
		sb.WriteString(arg)
		if i != (length - 1) {
			sb.WriteString(" ")
		}
	}
	return strings.ToLower(sb.String())
}

func FlatArgsSimple(args []string) string {
	sb := strings.Builder{}
	for _, arg := range args {
		sb.WriteString(arg)
	}
	return sb.String()
}