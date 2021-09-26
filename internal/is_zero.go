package internal

func IsZero(v interface{}) bool {
	switch v := v.(type) {
	case error:
		return v == nil
	case int:
		return v == 0
	case int8:
		return v == 0
	case int16:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0
	case uint:
		return v == 0
	case uint8:
		return v == 0
	case uint16:
		return v == 0
	case uint32:
		return v == 0
	case uint64:
		return v == 0
	case uintptr:
		return v == 0
	case float32:
		return v == 0
	case float64:
		return v == 0
	case string:
		return v == ""
	case bool:
		return !v
	case interface{}:
		return v == nil
	}
	return v == nil
}
