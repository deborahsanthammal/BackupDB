package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"github.com/jackc/pgx/v5"
)

var db_port uint64
var db string

var Config pgx.ConnConfig


func init() {

	var required_arguments []string = []string{
		"--name",
		"--user",
		"--pass",
		"--host",
		"--db",
	} 

	flag.StringVar(&Config.Database, "name", "", "Database name")
	flag.StringVar(&Config.User, "user", "root", "Database user")
	flag.StringVar(&Config.Password, "password", "root", "Database password")
	flag.StringVar(&Config.Host, "host", "localhost", "Database host")
	flag.Uint64Var(&db_port, "port", 5432, "Database port")
	flag.StringVar(&db, "db", "", "Database")

	for _, element := range required_arguments {
		if !slices.Contains(os.Args, element) {
			fmt.Println("Error: Not enough arguments")
			fmt.Printf("Required Arguments: %v \n", required_arguments)
			os.Exit(1)
		}
	}

	

	Config.Port = uint16(db_port)

	flag.Parse()
}



