package rpc

import (
	"context"
	"net/http"

	"id-maker/internal/controller/rpc/proto"
	"id-maker/internal/entity"
	"id-maker/internal/usecase"
	"id-maker/pkg/grpcserver"
	"id-maker/pkg/logger"

	"github.com/golang/protobuf/ptypes/empty"
)

type SegmentRpc struct {
	t usecase.Segment
	l logger.Interface
}

func newSegmentRoutes(s usecase.Segment, l logger.Interface) {
	proto.RegisterGidServer(grpcserver.RpcServer, &SegmentRpc{s, l})
}

func (r *SegmentRpc) Ping(ctx context.Context, g *empty.Empty) (out *proto.PingReply, err error) {
	out = &proto.PingReply{
		Status: &proto.Status{
			Code: http.StatusOK,
		},
		Data: "pong",
	}
	return
}

func (r *SegmentRpc) GetId(ctx context.Context, in *proto.IdRequest) (out *proto.IdReply, err error) {
	var id int64

	out = &proto.IdReply{
		Status: &proto.Status{
			Code: http.StatusOK,
		},
	}

	tag := in.GetTag()
	if tag == "" {
		out.Status.Code = http.StatusInternalServerError
		out.Status.Msg = "param error"
		return
	}

	if id, err = r.t.GetId(tag); err != nil {
		out.Status.Code = http.StatusInternalServerError
		out.Status.Msg = err.Error()
		return
	}
	out.Id = id

	return
}

func (r *SegmentRpc) GetSnowId(ctx context.Context, g *empty.Empty) (out *proto.SnowIdReply, err error) {
	out = &proto.SnowIdReply{
		Status: &proto.Status{
			Code: http.StatusOK,
		},
		Id: r.t.SnowFlakeGetId(),
	}

	return
}

func (r *SegmentRpc) CreateTag(ctx context.Context, in *proto.CreateTagRequest) (out *proto.CreateTagReply, err error) {
	out = &proto.CreateTagReply{
		Status: &proto.Status{
			Code: http.StatusOK,
		},
	}
	if in.GetTag() == "" || in.GetStep() == 0 {
		out.Status.Code = http.StatusInternalServerError
		out.Status.Msg = "param error"
		return
	}

	if err = r.t.CreateTag(&entity.Segments{
		BizTag: in.GetTag(),
		MaxId:  in.GetMaxId(),
		Step:   in.GetStep(),
		Remark: in.GetRemark(),
	}); err != nil {
		out.Status.Code = http.StatusInternalServerError
		out.Status.Msg = err.Error()
		return
	}
	return
}
