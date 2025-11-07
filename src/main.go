package main

import (
	"fmt"
	"os"
	"wool/tools"
)

func main() {
	var args = os.Args

	if len(args) < 3 {
		fmt.Println("wool <command> <parameter>")
	}

	var cmd string = args[1]
	var param string = args[2]

	switch cmd {
	//CREATE NEW PROJECT
	case "new":
		_, err := tools.CrateProject(param)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	//BUILD PROJECT
	case "build":
		_, err := tools.Build(param)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

	//DISPLAY HELP
	default:
		fmt.Println("Unknown command: ", cmd)
	}
}
