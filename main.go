package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	// "github.com/jackc/pgx/v5"
)



func init() {

	var required_arguments []string = []string{
		"--name",
		"--user",
		"--pass",
		"--host",
		"--db",
	}

	flag.StringVar(&Config.DBName, "name", "", "Database name")
	flag.StringVar(&Config.DBUser, "user", "root", "Database user")
	flag.StringVar(&Config.DBPass, "pass", "root", "Database password")
	flag.StringVar(&Config.DBHost, "host", "localhost", "Database host")
	flag.StringVar(&Config.DBPort, "port", "5432", "Database port")
	flag.StringVar(&Config.Database, "db", "postgres", "Database")

	for _, element := range required_arguments {
		if !slices.Contains(os.Args, element) {
			fmt.Println("Error: Not enough arguments")
			fmt.Printf("Required Arguments: %v \n", required_arguments)
			os.Exit(1)
		}
	}
	flag.Parse()
}

func main(){
	Backup()
}





