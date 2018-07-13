package alis

import (
	"zl/alis/src/db"
	"zl/alis/src/utils"
	"log"
	"github.com/manifoldco/promptui"
	"fmt"
	"os/exec"
)

type Handler struct {
	Params
}

// Run ...
func (h *Handler) Run(p *Params) {
	switch p.Type {
		case "map":
			handleMap(p)
			break
		case "exec":
			handleExec(p)
			break
	}
}

func handleMap(p *Params) {
	preCmd, cmd := p.Args[0], p.Args[1]

	preCmdRows := db.Find(&db.Data{Cmd: preCmd})
	if len(preCmdRows) > 0 {
		panic(fmt.Sprintf("cmd: '%s' have exist!", preCmd))
	}

	data := &db.Data{ Cmd: cmd }
	rows := db.Find(data)
	length := len(rows)

	if length == 0 {
		result := db.Create(data)
		insertID, err := result.LastInsertId()
		utils.HandleErr(err)
		preData := &db.Data{ Cmd: preCmd, Next: insertID }
		db.Create(preData)
		return
	}
	
	index := getSelectedCmd(rows)
	preData := &db.Data{ Cmd: preCmd, Next: rows[index].ID }
	db.Create(preData)
	log.Println("map handle ok!")
}

func getSelectedCmd(rows []*db.Data) int {
	var cmdStrs []string
	cmdStrs = append(cmdStrs, "None")
	for index, row := range(rows) {
		cmdStrs = append(cmdStrs, fmt.Sprintf("%d: %s", index, row.Cmd))
	}
	prompt := promptui.Select{
		Label: "please select the index you want to map!",
		Items: cmdStrs,
	}

	i, _, err := prompt.Run()
	utils.HandleErr(err)
	return i - 1
}

func handleExec(p *Params) {
	data := &db.Data{ Cmd: p.Args[0] }
	rows := db.Find(data)
	length := len(rows)
	if length <= 0 {
		panic("no rows find!")
	}

	nextID := rows[0].Next
	nextRow := db.FindOriginOne(&db.Data{ ID: nextID })
	log.Printf("will to run: %s", nextRow.Cmd)
	// isSure := sure(nextRow.Cmd)
	// if !isSure {
	// 	return
	// }
	shell := exec.Command("/bin/bash", "-c", nextRow.Cmd)
	output, err := shell.Output()
	utils.HandleErr(err)
	log.Println(string(output))
}

func sure(cmd string) bool {
	prompt := promptui.Prompt{
		Label:    "sure to exec: " + cmd + "?(y/n)",
		Validate: func(input string) error { return nil },
	}

	result, _ := prompt.Run()
	if result == "y" || result == "Y" {
		return true
	}
	return false
}