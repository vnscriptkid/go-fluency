package HashMap

type RequestLogger struct {
	// Write your code here
	limit     int
	keyToTime map[string]int
}

// requestLoggerInit initializes RequestLogger object
func (l *RequestLogger) requestLoggerInit(timeLimit int) {
	l.keyToTime = map[string]int{}

	l.limit = timeLimit
}

// messageRequestDicision accepts and deny message requests
func (l *RequestLogger) messageRequestDecision(timestamp int, req string) bool {
	if lastTime, ok := l.keyToTime[req]; !ok {
		l.keyToTime[req] = timestamp
		return true
	} else {
		accept := timestamp-lastTime >= l.limit

		// update lastTime
		l.keyToTime[req] = timestamp

		return accept
	}
}
