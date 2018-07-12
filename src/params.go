package alis

import (
	"strings"
)

type Params struct {
	Type string
	Args []string
}

func ParseParams(str string) *Params {
	var p *Params
	if (strings.Contains(str, "=>")) {
		return &Params{ Type: "map", Args: genArgs(str) }
	}
	
	p = &Params{ Type: "exec", Args: []string{ str } }
	return p
}

func genArgs(str string) []string {
	strs := strings.Split(str, "=>")
	var result []string
	for _, str := range(strs) {
		result = append(result, strings.Trim(str, " "))
	}
	return result
}