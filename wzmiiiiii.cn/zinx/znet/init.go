package znet

import (
	"go.uber.org/zap"
	"wzmiiiiii.cn/zinx/zlog"
)

var sugaredLogger *zap.SugaredLogger

func init() {
	sugaredLogger = zlog.GetSugaredLogger()
}
