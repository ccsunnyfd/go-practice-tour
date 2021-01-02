package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net"

	pb "github.com/ccsunnyfd/practice/grpc-demo/proto"
	"google.golang.org/grpc"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

// GreeterServer is
type GreeterServer struct{}

// SayHello is
func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello.world"}, nil
}

// SayList is
func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.HelloReply{
			Message: "hello.list(" + r.GetName() + ")",
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// SayRecord is
func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: "hello.record"})
		}
		if err != nil {
			return err
		}
		log.Printf("hello.record(%v)", req)
	}
}

// SayRoute is
func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.HelloReply{
			Message: "hello.route",
		})
		if err != nil {
			return err
		}
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++
		log.Printf("receive request(%v)", req)
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
