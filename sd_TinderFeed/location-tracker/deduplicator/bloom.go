package deduplicator

import (
	redisbloom "github.com/RedisBloom/redisbloom-go"
)

type BloomDeduplicator struct {
	rdb *redisbloom.Client
}

func NewBloomDeduplicator() IDeduplicator {
	return &BloomDeduplicator{
		rdb: redisbloom.NewClient("localhost:6379", "", nil),
	}
}

func (b BloomDeduplicator) addToUserHistory(userID string, seenUserID string) error {
	// create a new Bloom filter with a capacity of 1000 and a false positive rate of 0.01 (1%)
	_, err := b.rdb.BfInsert(
		b.getKey(userID),
		10000, // desired capacity of the bloom filter, number of users that each user is likely to see (i.e., the number of users in their history)
		0.01,  // desired error ratio (false-positive rate) of the bloom filter
		0,     // expansion rate will be determined automatically based on the error ratio and capacity
		false, // create a new bloom filter if it doesn't exist
		false, // auto-scale bloom filter
		[]string{seenUserID})

	return err
}

func (b BloomDeduplicator) isUserInHistory(userID string, seenUserID string) (bool, error) {
	return b.rdb.Exists(b.getKey(userID), seenUserID)
}

func (b BloomDeduplicator) getKey(userID string) string {
	return "user_history_bloom:" + userID
}
