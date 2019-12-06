package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/laughtt/loginService/pkg/server"
	v1 "github.com/laughtt/loginService/pkg/service"
)

const (
	addDatabase = "INSERT INTO `members`(`username`, `password`) VALUES ('prueba2','mypasswor2d')"
)
const (
	database   = "service"
	name       = "sql"
	host       = "localhost"
	apiVersion = "v1"
	password   = "123"
	port       = "5505"
	username   = "root"
)

//Config Database connection
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string
	// DB Datastore parameters selsction
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

//AUTHservice Crea el servicio para la conexion de la base de datos

func RunServer() error {
	ctx := context.Background()
	cfg := Config{
		GRPCPort:            port,
		DatastoreDBHost:     host,
		DatastoreDBPassword: password,
		DatastoreDBSchema:   database,
		DatastoreDBUser:     username,
	}

	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	v1API := v1.NewAuthServiceServer(db)

	if err != nil {
		return fmt.Errorf("failed to coonect: %v", err)
	}

	defer db.Close()

	return grpc.RunServer(ctx, v1API, port)
}

func main() {
	fmt.Println("a")
	fmt.Println(RunServer())
	fmt.Println("a")
}
