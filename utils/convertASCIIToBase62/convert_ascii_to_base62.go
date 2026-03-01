package convertasciitobase62

import "fmt"

func Base62CharToValue(c byte) (int, error) {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0'), nil
	case c >= 'A' && c <= 'Z':
		return int(c-'A') + 10, nil
	case c >= 'a' && c <= 'z':
		return int(c-'a') + 36, nil
	default:
		return 0, fmt.Errorf("invalid base62 character: %c", c)
	}
}

func Base62ValueToChar(c byte) (byte, error) {
	switch {
	case c >= 0 && c <= 9:
		return byte(c + '0'), nil
	case c >= 10 && c <= 36:
		return byte(c+'A') - 10, nil
	case c >= 37 && c <= 61:
		return byte(c+'a') - 36, nil
	default:
		return 0, fmt.Errorf("invalid base62 value: %c", c)
	}
}
