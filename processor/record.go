package processor

type Record struct {
	key       string
	value     []byte
	timestamp int64
}

func (r *Record) Key() string {
	return r.key
}

func (r *Record) Value() []byte {
	return r.value
}

func (r *Record) Timestamp() int64 {
	return r.timestamp
}

func NewRecord(key string, value []byte, timestamp int64) *Record {
	return &Record{
		key:       key,
		value:     value,
		timestamp: timestamp,
	}
}
