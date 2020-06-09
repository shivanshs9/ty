package fun

import (
	"testing"
)

func TestKeys(t *testing.T) {
	m := map[string]int{
		"c": 0, "b": 0, "a": 0,
	}
	keys := Keys(m).([]string)

	scmp := func(a, b string) bool { return a < b }
	assertDeep(t, []string{"a", "b", "c"}, QuickSort(scmp, keys))
}

func TestValues(t *testing.T) {
	m := map[int]string{
		1: "c", 2: "b", 3: "a",
	}
	vals := Values(m).([]string)

	scmp := func(a, b string) bool { return a < b }
	assertDeep(t, []string{"a", "b", "c"}, QuickSort(scmp, vals))
}

func TestMergeVarargMaps(t *testing.T) {
	m1 := map[string]string{"x": "2"}
	m2 := map[string]string{"y": "whatever"}
	assertDeep(t, map[string]string{"x": "2", "y": "whatever"}, MergeMaps(m1, m2))
}

func TestMergeListMaps(t *testing.T) {
	m1 := map[string]string{"x": "2"}
	m2 := map[string]string{"y": "whatever"}
	assertDeep(t, map[string]string{"x": "2", "y": "whatever"}, MergeMaps([]map[string]string{m1, m2}))
}
