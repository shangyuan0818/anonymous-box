package util

import "github.com/samber/lo"

func TypedMapToInterfaceMap[T any](m map[string]T) map[string]interface{} {
	return lo.MapEntries(m, func(k string, v T) (string, interface{}) {
		return k, v
	})
}

func InterfaceMapToTypedMap[T any](m map[string]interface{}) map[string]T {
	return lo.MapEntries(m, func(k string, v interface{}) (string, T) {
		if v, ok := v.(T); ok {
			return k, v
		} else {
			return k, lo.Empty[T]()
		}
	})
}
