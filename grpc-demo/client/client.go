package main

import (
	"context"
	"flag"
	"io"
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

	// unary RPC
	err := SayHello(client, &pb.HelloRequest{
		Name: "AndySayHello",
	})
	if err != nil {
		log.Fatalf("SayHello err: %v", err)
	}

	// server-side streaming RPC
	err = SayList(client, &pb.HelloRequest{
		Name: "AndySayList",
	})
	if err != nil {
		log.Fatalf("SayList err: %v", err)
	}

	// client-side streaming RPC
	err = SayRecord(client, &pb.HelloRequest{
		Name: "AndySayRecord",
	})
	if err != nil {
		log.Fatalf("SayRecord err: %v", err)
	}
}

// SayHello is
func SayHello(client pb.GreeterClient, r *pb.HelloRequest) error {
	resp, err := client.SayHello(context.Background(), r)
	if err != nil {
		return err
	}

	log.Printf("client.SayHello resp: %s", resp.GetMessage())
	return nil
}

// SayList is
func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, err := client.SayList(context.Background(), r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("client.SayList resp: %v", resp.GetMessage())
	}

	return nil
}

// SayRecord is
func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, err := client.SayRecord(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("client.SayRecord resp: %v", resp.GetMessage())

	return nil
}
