package main

import (
	"github.com/DarthPestilane/easytcp"
	"github.com/DarthPestilane/easytcp/examples/fixture"
	"github.com/DarthPestilane/easytcp/examples/tcp/proto_packet/common"
	"github.com/DarthPestilane/easytcp/message"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	srv := easytcp.NewServer(&easytcp.ServerOption{
		Packer: &common.CustomPacker{},
		Codec:  &easytcp.ProtobufCodec{},
	})

	srv.AddRoute(common.ID_FooReqID, handle, logTransmission(&common.FooReq{}, &common.FooResp{}))

	if err := srv.Serve(fixture.ServerAddr); err != nil {
		log.Errorf("serve err: %s", err)
	}
}

func handle(c *easytcp.Context) (*message.Entry, error) {
	var reqData common.FooReq
	if err := c.Bind(&reqData); err != nil {
		return nil, err
	}
	return c.Response(common.ID_FooRespID, &common.FooResp{
		Code:    2,
		Message: "success",
	})
}

func logTransmission(req, resp proto.Message) easytcp.MiddlewareFunc {
	return func(next easytcp.HandlerFunc) easytcp.HandlerFunc {
		return func(c *easytcp.Context) (*message.Entry, error) {
			if err := c.Bind(req); err == nil {
				log.Debugf("recv | id: %d; size: %d; data: %s", c.Message().ID, len(c.Message().Data), req)
			}

			respEntry, err := next(c)

			if err == nil && respEntry != nil {
				c.MustDecodeTo(respEntry.Data, resp)
				log.Infof("send | id: %d; size: %d; data: %s", respEntry.ID, len(respEntry.Data), resp)
			}
			return respEntry, err
		}
	}
}
