package lib

import (
	"strconv"
	"strings"
)

func Stringify(v interface{}) string {
	switch x := v.(type) {
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case int64:
		return strconv.FormatInt(x, 10)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(x)
	case []byte:
		return string(x)
	case []string:
		return strings.Join(x, "")
	default:
		// Fallback for unsupported types - could extend as needed
		return ""
	}
}
