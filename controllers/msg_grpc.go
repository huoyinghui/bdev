package controllers

import (
	"bdev/logger"
	"bdev/models"
	pb "bdev/protos"
	context "golang.org/x/net/context"
)

type MsgSerController struct {
}

func (this *MsgSerController) MessageCreate(ctx context.Context, req *pb.Message) (*pb.MessageResponse, error) {
	mid, err := models.AddMessage(&models.Message{
		FromUserId: req.FromUserId,
		ToUserId:   req.ToUserId,
		Title:      req.Title,
		Message:    req.Message,
	})
	if err != nil {
		logger.Errorf("err:%v req:%v", err, req)
		return nil, err
	}
	return &pb.MessageResponse{
		Mid: mid,
	}, nil
}

func (this *MsgSerController) MessageDelete(ctx context.Context, req *pb.MessageDeleteRequest) (*pb.MessageResponse, error) {
	err := models.DeleteMessage(req.Id)
	if err != nil {
		logger.Errorf("err:%v req:%v", err, req)
		return nil, err
	}
	return &pb.MessageResponse{
		Mid: -1,
	}, nil
}

func (this *MsgSerController) MessageUpdate(ctx context.Context, req *pb.Message) (*pb.Message, error) {
	msg, err := models.UpdateMessage(req.Id, &models.Message{
		Id: req.Id,
		//只更新状态就行
		Status: req.Status,
	})
	if err != nil {
		logger.Errorf("err:%v req:%v", err, req)
		return nil, err
	}
	return &pb.Message{
		Id:         msg.Id,
		FromUserId: msg.FromUserId,
		ToUserId:   msg.ToUserId,
		CreatedAt:  msg.CreatedAt.Unix(),
		Title:      msg.Title,
		Message:    msg.Message,
		Status:     msg.Status,
		IsDelete:   msg.IsDelete,
	}, nil
}

func (this *MsgSerController) MessageRead(ctx context.Context, req *pb.MessageReadRequest) (resp *pb.MessageListResponse, err error) {
	nums := int64(0)
	msgList := []*models.Message{}
	// todo
	//models.ReadMessage()
	if req.FromUserId > 0 {
		nums, msgList, err = models.ReadMessageWithFromUserId(req.MsgType, req.Id, req.ToUserId, req.FromUserId, int(req.PageNumber), int(req.PageSize))
	} else {
		nums, msgList, err = models.ReadMessage(req.MsgType, req.Id, req.ToUserId, int(req.PageNumber), int(req.PageSize))
	}
	if err != nil {
		logger.Errorf("err:%v req:%v", err, req)
		return nil, err
	}
	pmsgList := []*pb.Message{}
	for _, msg := range msgList {
		logger.Debugf("msg:%v", msg)
		msg := pb.Message{
			Id:         msg.Id,
			FromUserId: msg.FromUserId,
			ToUserId:   msg.ToUserId,
			Title:      msg.Title,
			Message:    msg.Message,
			CreatedAt:  msg.CreatedAt.Unix(),
			Status:     msg.Status,
			IsDelete:   msg.IsDelete,
		}
		pmsgList = append(pmsgList, &msg)
	}

	return &pb.MessageListResponse{
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
		Nums:       nums,
		MsgList:    pmsgList,
	}, nil
}
