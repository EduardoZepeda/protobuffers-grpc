package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/EduardoZepeda/protobuffers-grpc/database"
	"github.com/EduardoZepeda/protobuffers-grpc/server"
	"github.com/EduardoZepeda/protobuffers-grpc/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	flagPort := flag.Int("port", 5060, "The port for the gRPC server")
	flag.Parse()
	port := fmt.Sprintf(":%d", *flagPort)
	list, err := net.Listen("tcp", port)
	log.Printf("Starting server on port %d \n", *flagPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to server")
	log.Println("Connecting to database on port 54321")
	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	server := server.NewStudentServer(repo)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)
	reflection.Register(s)
	err = s.Serve(list)
	log.Println("Starting gRPC server")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to gRPC server")
}
