package main

import (
	"fmt"
)

type CommandConfig struct{
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	Database string
}

var Config CommandConfig

var Database map[string]string = map[string]string{
	"postgresql":fmt.Sprintf("pg_dump -U %s %s > backup.sql ", Config.DBUser, Config.DBName),
}