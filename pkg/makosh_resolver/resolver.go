package makosh_resolver

import (
	"encoding/json"
	"net/http"
	"sync"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/resolver"

	"github.com/godverv/makosh/pkg/makosh_be"
)

type UpdateAddresses func(addrs []string) error

type Resolver struct {
	getAddressesRequest *http.Request

	m sync.RWMutex

	doUpdate UpdateAddresses

	log logrus.StdLogger
}

func (r *Resolver) ResolveNow(_ resolver.ResolveNowOptions) {
	err := r.Resolve()
	if err != nil {
		r.log.Println("error resolving connections", err.Error())
	}
}

func (r *Resolver) Close() {}

func (r *Resolver) Resolve() error {
	httpResp, err := http.DefaultClient.Do(r.getAddressesRequest)
	if err != nil {
		return errors.Wrap(err, "error getting ")
	}

	defer httpResp.Body.Close()
	var endpointsResponse makosh_be.ListEndpoints_Response
	err = json.NewDecoder(httpResp.Body).Decode(&endpointsResponse)
	if err != nil {
		return errors.Wrap(err, "error parsing list endpoints response")
	}

	err = r.doUpdate(endpointsResponse.Urls)
	if err != nil {
		return errors.Wrap(err, "error updating connection state")
	}

	return nil
}