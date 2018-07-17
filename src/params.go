package alis

import (
	"strings"
	// "log"
)

type Params struct {
	Type string
	Args []string
}

func ParseParams(args []string) *Params {
	switch args[0] {
		case "map":
			return &Params{ Type: args[0], Args: args[1:] }
		case "search":
			return &Params{ Type: args[0], Args: args[1:] }
		default:
			return &Params{ Type: "exec", Args: []string{ args[0] } }
	}
}

func genArgs(str string) []string {
	strs := strings.Split(str, "=>")
	var result []string
	for _, str := range(strs) {
		result = append(result, strings.Trim(str, " "))
	}
	return result
}