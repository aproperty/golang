package rpcserver

import (
	"example/mailservice"
	"example/protodatasvr"

	"golang.org/x/net/context"
)

type SendEmailServer struct {
	protodatasvr.UnimplementedEmailServiceServer
}

// 发送邮件通知
func (t SendEmailServer) SendEmail(ctx context.Context, req *protodatasvr.SendEmailRequest) (*protodatasvr.SendEmailResponse, error) {
	sendEmail := &mailservice.SendEmail{}
	code, err := sendEmail.SendEmail(req.Address, req.Subject, req.Body)

	rsp := &protodatasvr.SendEmailResponse{}

	if code == 0 {
		rsp.Code = 0
		rsp.Msg = "success!"
		return rsp, nil
	}

	rsp.Code = 1
	rsp.Msg = "fail!"
	return rsp, err
}
