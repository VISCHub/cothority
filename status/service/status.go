package status

import (
	"gopkg.in/dedis/onet.v1"
	"gopkg.in/dedis/onet.v1/log"
	"gopkg.in/dedis/onet.v1/network"
)

// This file contains all the code to run a Stat service. The Stat receives takes a
// request for the Status reports of the server, and sends back the status reports for each service
// in the server.

// ServiceName is the name to refer to the Status service.
const ServiceName = "Status"

func init() {
	onet.RegisterNewService(ServiceName, newStatService)
	network.RegisterMessage(&Request{})
	network.RegisterMessage(&Response{})

}

// Stat is the service that returns the status reports of all services running on a server.
type Stat struct {
	*onet.ServiceProcessor
}

// Request is what the Status service is expected to receive from clients.
type Request struct{}

// Response is what the Status service will reply to clients.
type Response struct {
	Msg            map[string]*onet.Status
	ServerIdentity *network.ServerIdentity
}

// Request treats external request to this service.
func (st *Stat) Request(req *Request) (network.Message, onet.ClientError) {
	log.Lvl3("Returning", st.Context.ReportStatus())
	return &Response{
		Msg:            st.Context.ReportStatus(),
		ServerIdentity: st.ServerIdentity(),
	}, nil
}

// newStatService creates a new service that is built for Status
func newStatService(c *onet.Context) onet.Service {
	s := &Stat{
		ServiceProcessor: onet.NewServiceProcessor(c),
	}
	err := s.RegisterHandler(s.Request)
	if err != nil {
		log.ErrFatal(err, "Couldn't register message:")
	}

	return s
}
