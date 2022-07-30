package handler

import (
	"common"
	"game/server"
)

func HandleEcho(userid uint32, connid uint32, msgBody []byte) {
	server.SendResp(userid, MsgidEchoResp, common.ResultSuccess, msgBody)
}
