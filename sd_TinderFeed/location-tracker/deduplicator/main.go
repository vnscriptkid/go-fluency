package deduplicator

type IDeduplicator interface {
	addToUserHistory(userID string, seenUserID string) error
	isUserInHistory(userID string, seenUserID string) (bool, error)
	getKey(userID string) string
}
