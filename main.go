package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/manifoldco/promptui"
)

type Command struct {
	Cmd  string
	Args []string
}

var cmds = map[string]Command{
	"List simple":   {Cmd: "ls", Args: []string{}},
	"List advanced": {Cmd: "ls", Args: []string{"-lash"}},
	"System info":   {Cmd: "uname", Args: []string{"-a"}},
}

func main() {
	keys := make([]string, 0, len(cmds))
	for key := range cmds {
		keys = append(keys, key)
	}

	prompt := promptui.Select{
		Label: "Select Command",
		Items: keys,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	command, exists := cmds[result]

	if !exists {
		fmt.Printf("Command '%s' does not exist.\n", result)
		return
	}

	cmd := exec.Command(command.Cmd, command.Args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out.String())
}
