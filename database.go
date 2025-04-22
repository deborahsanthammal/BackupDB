package main

import(
	"context"
	"github.com/jackc/pgx/v5"
)

func Initialize(config pgx.ConnConfig) *pgx.Conn {
	// var conn pgx.Conn 
	// var err error
	conn, err := pgx.ConnectConfig(context.Background(), &config)
	if err != nil {
		return nil
	}
	return conn
}

// func dump(conn *pgx.Conn){
// 	conn.Exec()
// }