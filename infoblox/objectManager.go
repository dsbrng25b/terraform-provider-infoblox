package infoblox

import (
	"github.com/infobloxopen/infoblox-go-client"
)

var (
	cloudAttrs = false
	cmpType    = "terraform"
)

func NewObjectManager(connector ibclient.IBConnector, cmpType string, tenantID string) *ibclient.ObjectManager {
	if cloudAttrs {
		return ibclient.NewObjectManager(connector, cmpType, tenantID)
	} else {
		return ibclient.NewLocalObjectManager(connector)
	}
}
