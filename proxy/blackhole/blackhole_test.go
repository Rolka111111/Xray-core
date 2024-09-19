package blackhole_test

import (
	"context"
	"testing"

	"github.com/luckyluke-a/xray-core/common"
	"github.com/luckyluke-a/xray-core/common/buf"
	"github.com/luckyluke-a/xray-core/common/serial"
	"github.com/luckyluke-a/xray-core/common/session"
	"github.com/luckyluke-a/xray-core/proxy/blackhole"
	"github.com/luckyluke-a/xray-core/transport"
	"github.com/luckyluke-a/xray-core/transport/pipe"
)

func TestBlackholeHTTPResponse(t *testing.T) {
	ctx := session.ContextWithOutbounds(context.Background(), []*session.Outbound{{}})
	handler, err := blackhole.New(ctx, &blackhole.Config{
		Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
	})
	common.Must(err)

	reader, writer := pipe.New(pipe.WithoutSizeLimit())

	var mb buf.MultiBuffer
	var rerr error
	go func() {
		b, e := reader.ReadMultiBuffer()
		mb = b
		rerr = e
	}()

	link := transport.Link{
		Reader: reader,
		Writer: writer,
	}
	common.Must(handler.Process(ctx, &link, nil))
	common.Must(rerr)
	if mb.IsEmpty() {
		t.Error("expect http response, but nothing")
	}
}
