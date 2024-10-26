package main

import (
	"fmt"
	"flag"
)

var db_name string
var db_host string
var db_port string
var db_user string
var db_password string

func init() {
	flag.StringVar(&db_name, "name", "", "Database name")
	flag.StringVar(&db_user, "user", "root", "Database user")
	flag.StringVar(&db_password, "password", "root", "Database password")
	flag.StringVar(&db_host, "host", "localhost", "Database host")
	flag.StringVar(&db_port, "port", "", "Database port")
	flag.Parse()
}

func main() {
	fmt.Println(`
			The values are as follows:
			Database name: ` + db_name + `
			Database user: ` + db_user + `
			Database password: ` + db_password + `
			Database host: ` + db_host + `
			Database port: ` + db_port + `
	`)
}


