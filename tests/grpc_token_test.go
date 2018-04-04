package test

import (
	pb "bdev/protos"
	_ "bdev/routers"
	"context"
	"testing"

	"google.golang.org/grpc"
)

func init() {
	address = "192.168.0.96:5008"
	//address = "192.168.0.96:5009"
}

func TestGrpcToken(t *testing.T) {
	req := pb.GenToekenRequest{
			//Uid: "5ab4bef1c4cd748f32c6dff3",
		Uid: "1234567",
			//Exp: int64(1532541706),
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//c := pb.NewTokenServiceClient(conn)
	c := pb.NewDeveloperServiceClient(conn)
	r, err := c.GenAccessToken(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	t.Logf("resp:%v", r)
	req2 := pb.CheckToekenRequest{
		Token: r.Token,
	}
	rr, err := c.CheckToken(context.Background(), &req2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("check: %v", rr)
}

func TestGrpcTokenCheck(t *testing.T) {
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
