package chaincode

import (
	"context"
	"fmt"

	"github.com/hyperledger/fabric/msp"

	"github.com/s7techlab/hlf-sdk-go/v2/api"
)

type Core struct {
	mspId         string
	name          string
	channelName   string
	endorsingMSPs []string
	endorsers []*api.HostEndpoint
	peerPool      api.PeerPool
	orderer       api.Orderer

	identity msp.SigningIdentity
}

func (c *Core) Endorsers() []*api.HostEndpoint {
	return c.endorsers
}

func (c *Core) Invoke(fn string) api.ChaincodeInvokeBuilder {
	return NewInvokeBuilder(c, fn)
}

func (c *Core) Query(fn string, args ...string) api.ChaincodeQueryBuilder {
	return NewQueryBuilder(c, c.identity, fn, args...)
}

func (c *Core) Install(version string) {
	panic("implement me")
}

func (c *Core) Subscribe(ctx context.Context) (api.EventCCSubscription, error) {
	peerDeliver, err := c.peerPool.DeliverClient(c.mspId, c.identity)
	if err != nil {
		return nil, fmt.Errorf(`initiate DeliverClient: %w`, err)
	}
	return peerDeliver.SubscribeCC(ctx, c.channelName, c.name)
}

func NewCore(
	mspId,
	ccName,
	channelName string,
	endorsingMSPs []string,
	endorsers []*api.HostEndpoint,
	peerPool api.PeerPool,
	orderer api.Orderer,
	identity msp.SigningIdentity,
) *Core {
	return &Core{
		mspId:         mspId,
		name:          ccName,
		channelName:   channelName,
		endorsingMSPs: endorsingMSPs,
		endorsers: endorsers,
		peerPool:      peerPool,
		orderer:       orderer,
		identity:      identity,
	}
}
