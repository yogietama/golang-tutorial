package main


import (
	"context"
	"log"
	"math/rand"
	"net"
	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang/org/grpc"
)


const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.Uni
}
