// masterService
package master_service

// DisNodeConf 节点设置
type DisNodeConf struct {
	daemonId   int
	daemonName string
	nodeType   map[string]int //当前物理节点上允许master分配到该节点的node类型 比如：配置了 [gate, login],那么在master作节点分配时,就只会分配gate和login类型的节点在这个物理节点上
	nodeList   map[int]int    //在这个物理节点上启动的节点 map[nodeId]procId
}

// INodeDistributor 节点分配器
type INodeDistributor interface {
	//TODO:确定接口
	Init()
	Start()
	Stop()
}
