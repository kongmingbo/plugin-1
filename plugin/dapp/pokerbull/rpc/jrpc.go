// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"context"
	"encoding/hex"

	"github.com/33cn/chain33/types"
	pb "github.com/33cn/plugin/plugin/dapp/pokerbull/types"
)

func (c *Jrpc) PokerBullStartTx(parm *pb.PBStartTxReq, result *interface{}) error {
	if parm == nil {
		return types.ErrInvalidParam
	}
	head := &pb.PBGameStart{
		Value:     parm.Value,
		PlayerNum: parm.PlayerNum,
	}
	reply, err := c.cli.Start(context.Background(), head)
	if err != nil {
		return err
	}
	*result = hex.EncodeToString(reply.Data)
	return nil
}

func (c *Jrpc) PokerBullContinueTx(parm *pb.PBContinueTxReq, result *interface{}) error {
	if parm == nil {
		return types.ErrInvalidParam
	}

	head := &pb.PBGameContinue{
		GameId: parm.GameId,
	}

	reply, err := c.cli.Continue(context.Background(), head)
	if err != nil {
		return err
	}

	*result = hex.EncodeToString(reply.Data)
	return nil
}

func (c *Jrpc) PokerBullQuitTx(parm *pb.PBQuitTxReq, result *interface{}) error {
	if parm == nil {
		return types.ErrInvalidParam
	}

	head := &pb.PBGameQuit{
		GameId: parm.GameId,
	}
	reply, err := c.cli.Quit(context.Background(), head)
	if err != nil {
		return err
	}

	*result = hex.EncodeToString(reply.Data)
	return nil
}

func (c *Jrpc) PokerBullQueryTx(parm *pb.PBQueryReq, result *interface{}) error {
	if parm == nil {
		return types.ErrInvalidParam
	}
	head := &pb.PBGameQuery{
		GameId: parm.GameId,
	}
	reply, err := c.cli.Show(context.Background(), head)
	if err != nil {
		return err
	}

	*result = hex.EncodeToString(reply.Data)
	return nil
}
