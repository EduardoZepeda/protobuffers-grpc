package server

import "github.com/EduardoZepeda/protobuffers-grpc/repository"

type Server struct {
	repo repository.Repository
}

func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}
