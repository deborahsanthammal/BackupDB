package main

import (
	"os"
	"os/exec"
	"fmt"
)



func Backup(){
	command := exec.Command(Database[Config.Database])
	err := command.Run() 
	if err != nil {
		fmt.Printf("Backup failed with error: %s \n", err.Error())
		os.Exit(1)
	}
}