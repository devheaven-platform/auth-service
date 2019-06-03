package transport

import (
	"github.com/devheaven-platform/auth-service/pkg/api/users"
	"github.com/devheaven-platform/auth-service/pkg/utils/transport"
)

// This object is used to group all the transport
// functions together.
type kafkaTransport struct {
	transport.BaseKafkaTransport
	service users.Service
}

// CreateKafkaTransport is used to intialize the transport
// layer. It takes an service as parameter
func CreateKafkaTransport(service users.Service) {
	transport := &kafkaTransport{
		service: service,
	}
}
