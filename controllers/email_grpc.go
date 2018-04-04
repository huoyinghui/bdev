package controllers

import (
	"bdev/utils"
	pb "bdev/protos"
	"fmt"
	context "golang.org/x/net/context"
)

type EmailController struct {
}

func (email *EmailController) SendMail(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {
	ret, err := utils.SendMail(req.Sender, req.To, req.Subject, req.Charset, req.HtmlBody, req.TextBody, req.Cc)
	if err != nil {
		fmt.Printf("\nSendMail err:%v\nreq:%v\n", err, req)
		return nil, err
	}

	return &pb.EmailResponse{
		MessageId: *ret.MessageId,
	}, nil
}
