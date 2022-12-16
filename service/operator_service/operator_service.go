package operator_service

import (
	"github.com/njtc406/server_engine/log"
	"github.com/njtc406/server_engine/node"
	"github.com/njtc406/server_engine/service"
)

func init() {
	//加载Daemon服务
	node.Setup(&OperatorService{})
}

type any interface{}

type OperatorService struct {
	service.Service
}

func (slf *OperatorService) OnInit() error {
	log.Release("start daemon node...")

	//获取配置文件

	return nil
}

func (slf *OperatorService) OnNodeConnected(nodeId int) {
	log.Release("regist daemon node success!")
	//TODO:节点连接事件
}

func (slf *OperatorService) OnNodeDisconnect(nodeId int) {
	log.Release("unRegist daemon node success!")
	//TODO:节点断开事件
}
