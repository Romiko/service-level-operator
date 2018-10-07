package prometheus

import (
	"sync"

	"github.com/prometheus/client_golang/api"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// ClientFactory knows how to get prometheus API clients.
type ClientFactory interface {
	// GetV1APIClient returns a new prometheus v1 API client.
	// address is the address of the prometheus.
	GetV1APIClient(address string) (promv1.API, error)
}

type factory struct {
	clis  map[string]api.Client
	climu sync.Mutex
}

// NewFactory returns a new client factory.
func NewFactory() ClientFactory {
	return &factory{
		clis: map[string]api.Client{},
	}
}

// GetV1APIClient satisfies ClientFactory interface.
func (f *factory) GetV1APIClient(address string) (promv1.API, error) {
	f.climu.Lock()
	f.climu.Unlock()

	cli, ok := f.clis[address]
	if !ok {
		cli, err := api.NewClient(api.Config{Address: address})
		if err != nil {
			return nil, err
		}
		f.clis[address] = cli
	}
	return promv1.NewAPI(cli), nil
}