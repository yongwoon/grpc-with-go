package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hellopb "grpc-with-go/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"というレスポンスを返す
	message := fmt.Sprintf("Hello, %s!", req.GetName())

	return &hellopb.HelloResponse{
		Message: &message,
	}, nil
}

func main() {
	// 1. 8080番portのLisnterを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPC server を作成
	s := grpc.NewServer()

	// 3. grpc server に、hellopb.GreetingServiceServer を登録
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	// 4.server reflection (for grpcurl)　の設定
	reflection.Register(s)

	// 5. 作成した gRPC serve rを、8080 port で稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// 6. Ctrl+C が入力されたら Graceful shutdown されるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}