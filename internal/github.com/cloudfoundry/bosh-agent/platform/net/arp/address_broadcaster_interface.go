package arp

import (
	boship "github.com/cloudfoundry/bosh-init/internal/github.com/cloudfoundry/bosh-agent/platform/net/ip"
)

type AddressBroadcaster interface {
	BroadcastMACAddresses([]boship.InterfaceAddress)
}
