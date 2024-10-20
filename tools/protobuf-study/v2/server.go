package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	pb "main/v2/protoc"
	"net"
)

const (
	port = ":62019"
)

type Server struct {
	pb.UnimplementedGreeterServer
} //服务对象

// SayHello 实现服务的接口 在proto中定义的所有服务都是接口
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("reply with Hello " + in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //起一个服务
	pb.RegisterGreeterServer(s, &Server{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("listen and server with: ", port)
}
