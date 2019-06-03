package transport

import (
	"github.com/devheaven-platform/auth-service/pkg/api/users"
	"github.com/devheaven-platform/auth-service/pkg/utils/transport"
	log "github.com/sirupsen/logrus"
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

	transport.Listen("db.personnel.create-employee", transport.createUser)
	transport.Listen("db.personnel.update-employee", transport.updateUser)
	transport.Listen("db.personnel.delete-employee", transport.deleteUser)
}

func (t *kafkaTransport) createUser(message interface{}) {
	log.Info("Create User")
}

func (t *kafkaTransport) updateUser(message interface{}) {
	log.Info("Update User")
}

func (t *kafkaTransport) deleteUser(message interface{}) {
	log.Info("Delete User")
}
