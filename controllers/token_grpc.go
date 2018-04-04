package controllers

import (
	"bdev/config"
	"bdev/logger"
	pb "bdev/protos"
	"bdev/utils"
	context "golang.org/x/net/context"
)

type JWTSerController struct {
}

func (jwts *JWTSerController) GenAccessToken(ctx context.Context, req *pb.GenToekenRequest) (*pb.GenTokenResponse, error) {
	secret := utils.SecSecret(req.Uid, config.AppConf.JwtSalt)
	logger.Debugf("sec:%v", secret)
	token, err := utils.CreateToken(req.Uid, secret, req.Exp)
	if err != nil {
		logger.Errorf("req:%v err:%v token:%v", req, err, token)
		//return nil, grpc.Errorf()
		return nil, err
	}
	return &pb.GenTokenResponse{
		Token: token,
	}, nil
}

func (jwts *JWTSerController) CheckToken(ctx context.Context, req *pb.CheckToekenRequest) (*pb.CheckTokenResponse, error) {
	uid, err := utils.GetUid(req.Token)
	if err != nil {
		logger.Errorf("req:%v err:%v uid:%v", req, err, uid)
		return nil, err
	}
	secret := utils.SecSecret(uid, config.AppConf.JwtSalt)
	// 覆盖uid
	logger.Debugf("sec:%v", secret)
	uid, err = utils.AuthToken(req.Token, secret)
	if err != nil {
		logger.Errorf("req:%v err:%v uid:%v", req, err, uid)
		return nil, err
	}

	return &pb.CheckTokenResponse{
		Uid: uid,
	}, nil
}
