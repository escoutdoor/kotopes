package consumer

import (
	"github.com/IBM/sarama"

	"github.com/escoutdoor/kotopes/common/pkg/kafka"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
)

type GroupHandler struct {
	msgHandler kafka.Handler[sarama.ConsumerMessage]
}

// NewGroupHandler creates
func NewGroupHandler() *GroupHandler {
	return &GroupHandler{}
}

// Setup запускается в начале новой сессии до вызова ConsumeClaim
func (c *GroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup запускается в конце жизни сессии после того как все горутины ConsumeClaim завершаться
func (c *GroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim должен запустить потребительский цикл сообщений ConsumerGroupClaim().
// После закрытия канала Messages() обработчик должен завершить обработку
func (c *GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// Код ниже не стоит перемещать в горутину, так как ConsumeClaim
	// уже запускается в горутине, см.:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L869
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				logger.Logger().Info("message channel was closed")
				return nil
			}

			logger.Logger().Infof("message claimed: value = %s, timestamp = %v, topic = %s\n", string(message.Value), message.Timestamp, message.Topic)

			err := c.msgHandler(session.Context(), message)
			if err != nil {
				logger.Logger().Infof("error handling message: %v\n", err)
				continue
			}

			session.MarkMessage(message, "")

		// Должен вернуться, когда `session.Context()` завершен.
		// В противном случае возникнет `ErrRebalanceInProgress` или `read tcp <ip>:<port>: i/o timeout` при перебалансировке кафки. см.:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			logger.Logger().Info("session context done")
			return nil
		}
	}
}
