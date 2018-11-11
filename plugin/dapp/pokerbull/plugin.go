package pokerbull

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/33cn/plugin/plugin/dapp/pokerbull/commands"
	"github.com/33cn/plugin/plugin/dapp/pokerbull/executor"
	"github.com/33cn/plugin/plugin/dapp/pokerbull/rpc"
	"github.com/33cn/plugin/plugin/dapp/pokerbull/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.PokerBullX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.PokerBullCmd,
		RPC:      rpc.Init,
	})
}
