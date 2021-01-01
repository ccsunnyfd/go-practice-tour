package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/ccsunnyfd/practice/grpc-demo/proto"
	"google.golang.org/grpc"
)

var port string

// init
func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	err := SayHello(client)
	if err != nil {
		log.Fatalf("SayHello err: %v", err)
	}
}

// SayHello is
func SayHello(client pb.GreeterClient) error {
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "eddycjy",
	})
	if err != nil {
		return err
	}

	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}
