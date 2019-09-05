package config

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/gojekfarm/kat/testutil"
	"github.com/stretchr/testify/mock"
)

func TestAlter(t *testing.T) {
	admin := &testutil.MockClusterAdmin{}
	topics := []string{"topic1", "topic2"}
	config := "key1=val1"
	admin.On("AlterConfig", sarama.TopicResource, "topic1", mock.Anything, false).Return(nil).Times(1)
	admin.On("AlterConfig", sarama.TopicResource, "topic2", mock.Anything, false).Return(nil).Times(1)
	a := alter{admin: admin, topics: topics, config: config}
	a.alter()
	admin.AssertExpectations(t)
}
