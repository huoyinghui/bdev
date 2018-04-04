package test

import (
	pb "bdev/protos"
	_ "bdev/routers"
	"context"
	"testing"

	"google.golang.org/grpc"
)

var (
	address = "192.168.0.96:5008"
)

func TestGrpcCreate(t *testing.T) {
	req := pb.Message{
		FromUserId: int64(0),
		ToUserId:   int64(1),
		Title:      "APKPURE DEV",
		Message:    "Welcome to join us",
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//c := pb.NewMessageServiceClient(conn)
	c := pb.NewDeveloperServiceClient(conn)
	r, err := c.MessageCreate(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	t.Logf("resp:%v", r)
}

func TestGrpcRead(t *testing.T) {
	req := pb.MessageReadRequest{
		PageSize:   5,
		PageNumber: 1,
		//MsgType: "ALL",
		MsgType: "SEEN",
		//FromUserId: int64(0),
		ToUserId: int64(0),
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//c := pb.NewMessageServiceClient(conn)
	c := pb.NewDeveloperServiceClient(conn)
	r, err := c.MessageRead(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	t.Logf("resp:%v", r)
	for i, msg := range r.MsgList {
		t.Logf("i:%v msg:%v", i, msg)
	}
}

func TestGrpcUpdate(t *testing.T) {
	req := pb.Message{
		Id:     int64(24),
		Status: "SEEN",
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//c := pb.NewMessageServiceClient(conn)
	c := pb.NewDeveloperServiceClient(conn)
	r, err := c.MessageUpdate(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	t.Logf("resp:%v", r)
}

func TestGrpcDelete(t *testing.T) {
	req := pb.MessageDeleteRequest{
		Id: int64(24),
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//c := pb.NewMessageServiceClient(conn)
	c := pb.NewDeveloperServiceClient(conn)
	r, err := c.MessageDelete(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	t.Logf("resp:%v", r)
}
