package testing

type KeyVal struct {
	Key   string
	Value interface{}
}

func KV(key string, val interface{}) KeyVal {
	return KeyVal{Key: key, Value: val}
}
