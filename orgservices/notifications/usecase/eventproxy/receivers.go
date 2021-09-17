package eventproxy

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"github.com/timoth-y/chainmetric-network/orgservices/notifications/infrastructure/repository"
	"github.com/timoth-y/chainmetric-network/orgservices/notifications/infrastructure/services"
	"github.com/timoth-y/chainmetric-network/orgservices/notifications/model/intention"
	"github.com/timoth-y/chainmetric-network/orgservices/shared/core"
	"go.mongodb.org/mongo-driver/mongo"
)

type event struct {
	intention.EventConcern
	payload []byte
}

var eventsPipe chan event

func init() {
	eventsPipe = make(chan event, viper.GetInt("api.notifications.events_buffer_size"))
}

func spawnReceivers() {
	var receiverCount = viper.GetInt("api.notifications.event_receivers_count")

	for i := 0; i < receiverCount; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case e := <-eventsPipe:
					_ = e // TODO: redirect to subscribed users
				case <-ctx.Done():
					return
				}
			}
		}(ctx)
	}
}

func redirect(e event) error {
	var fcmService = services.NewNotificationsFirebase(core.Firebase)

	subs, err := repository.NewSubscriptionsMongo(core.MongoDB).GetByToken(e.SubscriptionToken())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Revoke(e.SubscriptionToken())
		}

		return fmt.Errorf("failed to retrive subscribtion tickets for '%s' concern: %w",
			e.SubscriptionToken(), err)
	}

	for i, sub := range subs {
		notification, err := e.NotificationFor(sub.UserID, e.payload)
		if err != nil {
			return fmt.Errorf("failed to form notification for '%s' concern with '%s' topic: %w",
				e.OfTopic(), e.SubscriptionToken(), err)
		}

		if err = fcmService.Push(notification); err != nil {
			core.Logrus.WithError(err).
				WithField("topic", e.OfTopic()).
				WithField("sub_id", e.SubscriptionToken()).
				WithField("user_id", sub.UserID).
				Errorf("failed to push notification via FCM")

			continue
		}

		subs[i].ReceivedTimes++
	}

	return nil
}