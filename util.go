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
