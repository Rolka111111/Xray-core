package tcp

import (
	"github.com/luckyluke-a/xray-core/common"
	"github.com/luckyluke-a/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
