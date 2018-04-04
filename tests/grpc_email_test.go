package test

import (
	pb "bdev/protos"
	_ "bdev/routers"
	"context"
	"testing"

	"google.golang.org/grpc"
)

var (
	recipient = "huoyinghui@apkpure.net"
	subject = "APKPure - Grpc"
	charSet = "UTF-8"
	htmlBody = `A password reset has been requested for the APKPure username: xxx
<br>
If you do actually want to reset your password, visit this link:
<br>
https://apkpure.com/account/resetpassword?k=xxx
<br>
Thank you for using the site!
https://apkpure.com`
	cc = []string{
		"285020853@qq.com", "hyhlinux@163.com",
	}
)

func init() {
	address = "192.168.0.96:5008"
	//address = "192.168.0.96:5009"
}

func TestGrpcEmail(t *testing.T) {
	req := pb.EmailRequest{
		//Sender: "",
		To: recipient,
		Cc: cc,
		Subject: subject,
		Charset: charSet,
		HtmlBody: htmlBody,
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//c := pb.NewTokenServiceClient(conn)
	c := pb.NewDeveloperServiceClient(conn)
	r, err := c.SendMail(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	t.Logf("resp:%v", r)
}
