package main

type canEqual interface {
	equals(j interface{}) bool
}

func arrayEquals[T canEqual](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !v.equals(b[i]) {
			return false
		}
	}
	return true
}

func takePtr[T any](v T) *T {
	return &v
}

type CompareResult int8

const (
	ERROR_CompareResult CompareResult = iota
	LESSER_THAN
	EQUAL_TO
	GREATER_THAN
)

func (compare_result CompareResult) invert() CompareResult {
	switch compare_result {
	case LESSER_THAN:
		return GREATER_THAN
	case GREATER_THAN:
		return LESSER_THAN
	default:
		return compare_result
	}
}
