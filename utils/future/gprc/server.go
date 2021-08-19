package gprc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "unioj/utils/future/gprc/protobuf"
)

type server struct {
}

func (this *server) SayHello(ctx context.Context, in *pb.HelloRequest) (out *pb.HelloResponse, err error) {
	return &pb.HelloResponse{Msg: "hellooooooo"}, err
}

func (this *server) SayName(ctx context.Context, in *pb.NameRequest) (out *pb.NameResponse, err error) {
	return &pb.NameResponse{Msg: in.Name + "  good morning"}, err
}

func Server() {
	//grpc
	listen, err := net.Listen("tcp", ":8888")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("网络建立失败", err)
		}
	}()
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	pb.RegisterHelloServerServer(srv, &server{})
	err = srv.Serve(listen)
	if err != nil {
		panic(err)
		return
	}
}
