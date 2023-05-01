package deduplicator

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type SetDeduplicator struct {
	rdb *redis.Client
}

func NewSetDeduplicator(client *redis.Client) IDeduplicator {
	return &SetDeduplicator{
		rdb: client,
	}
}

func (d SetDeduplicator) addToUserHistory(userID string, seenUserID string) error {
	_, err := d.rdb.SAdd(context.Background(), d.getKey(userID), seenUserID).Result()

	return err
}

func (d SetDeduplicator) isUserInHistory(userID string, seenUserID string) (bool, error) {
	return d.rdb.SIsMember(context.Background(), d.getKey(userID), seenUserID).Result()
}

func (d SetDeduplicator) getKey(userID string) string {
	return "user_history:" + userID
}
