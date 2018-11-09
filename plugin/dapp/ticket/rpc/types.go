package rpc

import (
	ty "gitlab.33.cn/chain33/plugin/plugin/dapp/ticket/types"
	"gitlab.33.cn/chain33/chain33/rpc/types"
)

type Jrpc struct {
	cli *channelClient
}

type Grpc struct {
	*channelClient
}

type channelClient struct {
	types.ChannelClient
}

func Init(name string, s types.RPCServer) {
	cli := &channelClient{}
	grpc := &Grpc{channelClient: cli}
	cli.Init(name, s, &Jrpc{cli: cli}, grpc)
	ty.RegisterTicketServer(s.GRPC(), grpc)
}