package transit_test

import (
	. "github.com/moleculer-go/goemitter"
	. "github.com/moleculer-go/moleculer/common"
	. "github.com/moleculer-go/moleculer/registry"
	log "github.com/sirupsen/logrus"
)

var logger = log.WithField("Unit Test", "Transit Pkg")

func CreateLogger(name string, value string) *log.Entry {
	return logger.WithField(name, value)
}

var localNode = CreateNode("unit-test-node")

func CreateBroker() *BrokerInfo {
	localBus := CreateEmitter()
	broker := &BrokerInfo{GetLocalNode: func() *Node { return &localNode }, GetLogger: CreateLogger, GetLocalBus: func() *Emitter { return localBus }}
	return broker
}
