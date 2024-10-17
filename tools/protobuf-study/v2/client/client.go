package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "main/v2/protoc"
	"os"
	"time"
)

const (
	address = "localhost:62019"
)

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 10秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	//req(c, ctx, name)
	for {
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("请输入: ")
		input.Scan()
		req(c, ctx, input.Text())
	}
}

func req(c pb.GreeterClient, ctx context.Context, name string) {
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
