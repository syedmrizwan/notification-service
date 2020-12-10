package api

import (
	"notification_service_webapp/env"
	"notification_service_webapp/util"

	"github.com/nats-io/stan.go"
)

// PushMessageToNATS published message on the given channel name
func PushMessageToNATS(channelName string, message []byte) error {
	logger := util.GetLogger()
	logger.Info("Message received for PushMessageToNATS")
	conn, err := stan.Connect(
		env.Env.NatsCluster,
		env.Env.NatsClient,
		stan.NatsURL(env.Env.NatsAddress),
	)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer conn.Close()
	logger.Info("Pushing To Channel Name & Message")
	if err = conn.Publish(channelName, message); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
