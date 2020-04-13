package common

import (
	"encoding/json"
	"strconv"
)

func ToInt(data interface{}) int {
	switch data.(type) {
	case string:
		switch data.(string) {
		case "true":
			fallthrough
		case "True":
			fallthrough
		case "TRUE":
			return 1
		case "false":
			fallthrough
		case "False":
			fallthrough
		case "FALSE":
			return 0
		case "":
			return 0
		default:
			val, err := strconv.Atoi(data.(string))
			if err != nil {
				return 0
			}
			return val
		}
	case int:
		return data.(int)
	case int64:
		return int(data.(int64))
	case int32:
		return int(data.(int32))
	case int16:
		return int(data.(int16))
	case int8:
		return int(data.(int8))
	case uint64:
		return int(data.(uint32))
	case uint16:
		return int(data.(uint16))
	case uint8:
		return int(data.(uint8))
	case float64:
		return int(data.(float64))
	case float32:
		return int(data.(float32))
	case bool:
		if data.(bool) {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func ToBool(data interface{}) bool {
	switch data.(type) {
	case bool:
		return data.(bool)
	case string:
		switch data.(string) {
		case "1":
			fallthrough
		case "true":
			fallthrough
		case "True":
			fallthrough
		case "TRUE":
			return true
		case "false":
			fallthrough
		case "False":
			fallthrough
		case "FALSE":
			return false
		default:
			return false
		}
	case int:
		return data.(int) == 1
	case int64:
		return data.(int64) == 1
	case int32:
		return data.(int32) == 1
	case int16:
		return data.(int16) == 1
	case int8:
		return data.(int8) == 1
	case uint64:
		return data.(uint32) == 1
	case uint16:
		return data.(uint16) == 1
	case uint8:
		return data.(uint8) == 1
	case float64:
		return int(data.(float64)) == 1
	case float32:
		return int(data.(float32)) == 1
	default:
		return false
	}
}

func ToFloat(data interface{}) float64 {
	switch data.(type) {
	case bool:
		if data.(bool) {
			return 1
		}
		return 0
	case string:
		val, err := strconv.ParseFloat(data.(string), 64)
		if err != nil {
			return 0
		}
		return val
	case int:
		return float64(data.(int))
	case int64:
		return float64(data.(int64))
	case int32:
		return float64(data.(int32))
	case int16:
		return float64(data.(int16))
	case int8:
		return float64(data.(int8))
	case uint64:
		return float64(data.(uint32))
	case uint16:
		return float64(data.(uint16))
	case uint8:
		return float64(data.(uint8))
	case float64:
		return data.(float64)
	case float32:
		return float64(data.(float32))
	default:
		return 0
	}
}

func ToString(data interface{}) string {
	switch data.(type) {
	case nil:
		return ""
	case string:
		return data.(string)
	case bool:
		if data.(bool) {
			return "1"
		}
		return "0"
	case int:
		return strconv.Itoa(data.(int))
	case int64:
		return strconv.Itoa(int(data.(int64)))
	case int32:
		return strconv.Itoa(int(data.(int32)))
	case int16:
		return strconv.Itoa(int(data.(int16)))
	case int8:
		return strconv.Itoa(int(data.(int8)))
	case uint64:
		return strconv.Itoa(int(data.(uint64)))
	case uint32:
		return strconv.Itoa(int(data.(uint32)))
	case uint16:
		return strconv.Itoa(int(data.(uint16)))
	case uint8:
		return strconv.Itoa(int(data.(uint8)))
	case float64:
		return strconv.FormatFloat(data.(float64), 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(data.(float32)), 'f', -1, 64)
	default:
		s, _ := json.Marshal(data)
		return string(s)
	}
}
