package gprc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "unioj/utils/future/gprc/protobuf"
)

func Client() {
	//	连接服务器
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	cli := pb.NewHelloServerClient(conn)
	//	通过句柄调用函数
	re, err := cli.SayHello(context.Background(), &pb.HelloRequest{Name: "ethan"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("调用结果：", re.GetMsg())
	res, err := cli.SayName(context.Background(), &pb.NameRequest{
		Name: "alban",
	})
	fmt.Println("第二次调用结果:", res.GetMsg())
}
