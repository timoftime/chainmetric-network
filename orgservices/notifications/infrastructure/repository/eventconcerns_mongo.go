package repository

import (
	"context"

	"github.com/spf13/viper"
	"github.com/timoth-y/chainmetric-network/orgservices/notifications/model/events"
	"github.com/timoth-y/chainmetric-network/orgservices/notifications/model/intention"
	"github.com/timoth-y/chainmetric-network/smartcontracts/shared/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// EventConcernsMongo defines subscriptions repository for MongoDB database.
type EventConcernsMongo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewEventConcernsMongo constructs new EventConcernsMongo repository instance.
func NewEventConcernsMongo(client *mongo.Client) *EventConcernsMongo {
	return &EventConcernsMongo{
		client:     client,
		collection: client.Database(viper.GetString("mongo_database")).Collection("event_concerns"),
	}
}

// GetAll retrieves all intention.EventConcern from the collection.
func (r *EventConcernsMongo) GetAll() ([]intention.EventConcern, error) {
	var (
		results     []intention.EventConcern
		ctx, cancel = context.WithTimeout(context.Background(), viper.GetDuration("mongo_query_timeout"))
	)


	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var (
			record intention.EventConcern
			topic  = intention.EventTopic(cursor.Current.Lookup("event_topic").String())
		)

		switch topic {
		case events.RequirementsViolationTopic:
			record = new(events.RequirementsViolationEvent)
		default:
			core.Logrus.WithField("topic", topic).
				Warn("unknown topic: document cannot be decoded")
			break
		}

		if err = cursor.Decode(record); err != nil {
			results = append(results, record)
		}
	}

	return results, err
}

// Insert stores intention.EventConcern in the database.
func (r *EventConcernsMongo) Insert(concerns ...intention.EventConcern) error {
	var (
		docs []interface{}
		ctx, cancel = context.WithTimeout(context.Background(), viper.GetDuration("mongo_query_timeout"))
	)

	defer cancel()

	for i := range concerns {
		docs = append(docs, concerns[i])
	}

	_, err := r.collection.InsertMany(ctx, docs)

	return err
}

// DeleteByTokens removes intention.EventConcern from the database by given `ids`.
func (r *EventConcernsMongo) DeleteByTokens(tokens ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("mongo_query_timeout"))

	defer cancel()

	_, err := r.collection.DeleteMany(ctx, bson.M{
		"subscription_token": bson.M{
			"$in": tokens,
		},
	})

	return err
}