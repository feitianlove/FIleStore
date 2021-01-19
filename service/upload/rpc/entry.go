package rpc

import (
	"context"
	"github.com/feitianlove/FIleStore/service/upload/config"
	"github.com/feitianlove/FIleStore/service/upload/proto"
)

type Upload struct {
}

func (upload *Upload) UploadEntry(ctx context.Context, req *proto.ReqEntry, res *proto.RespReqEntry) error {
	res.Entry = config.UploadEntry
	return nil
}
