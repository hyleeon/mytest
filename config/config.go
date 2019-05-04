package config

//Config is conf struct.
type Config struct {
}

//GetValue to get value.
func (config *Config) GetValue(key string) string {
	return get(key).(string)
}

func get(key string) interface{} {
	return key
}

//func set(key string, value interface{}) {
//
//}
