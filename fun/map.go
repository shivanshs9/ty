package fun

import (
	"reflect"

	"github.com/shivanshs9/ty"
)

// Keys has a parametric type:
//
//	func Keys(m map[A]B) []A
//
// Keys returns a list of the keys of `m` in an unspecified order.
func Keys(m interface{}) interface{} {
	chk := ty.Check(
		new(func(map[ty.A]ty.B) []ty.A),
		m)
	vm, tkeys := chk.Args[0], chk.Returns[0]

	vkeys := reflect.MakeSlice(tkeys, vm.Len(), vm.Len())
	for i, vkey := range vm.MapKeys() {
		vkeys.Index(i).Set(vkey)
	}
	return vkeys.Interface()
}

// Values has a parametric type:
//
//	func Values(m map[A]B) []B
//
// Values returns a list of the values of `m` in an unspecified order.
func Values(m interface{}) interface{} {
	chk := ty.Check(
		new(func(map[ty.A]ty.B) []ty.B),
		m)
	vm, tvals := chk.Args[0], chk.Returns[0]

	vvals := reflect.MakeSlice(tvals, vm.Len(), vm.Len())
	for i, vkey := range vm.MapKeys() {
		vvals.Index(i).Set(vm.MapIndex(vkey))
	}
	return vvals.Interface()
}

// MergeMaps either has a parametric type:
//
//	func MergeMaps(maps ...map[A]B) map[A]B
//	OR
//	func MergeMaps(maps []map[A]B) map[A]B
//
// MergeMaps returns the map obtained by merging all the maps
func MergeMaps(maps ...interface{}) interface{} {
	if len(maps) > 1 {
		// handle vararg case
		chk := ty.Check(new(func(...map[ty.A]ty.B) map[ty.A]ty.B), maps...)
		vxs, trs := chk.Args, chk.Returns[0]
		res := reflect.MakeMap(trs)
		xsLen := len(vxs)
		for i := 0; i < xsLen; i++ {
			iter := vxs[i].MapRange()
			for iter.Next() {
				res.SetMapIndex(iter.Key(), iter.Value())
			}
		}
		return res.Interface()
	} else if len(maps) == 1 {
		// handle slice of maps case
		chk := ty.Check(new(func([]map[ty.A]ty.B) map[ty.A]ty.B), maps[0])
		vxs, trs := chk.Args[0], chk.Returns[0]
		res := reflect.MakeMap(trs)
		xsLen := vxs.Len()
		for i := 0; i < xsLen; i++ {
			iter := vxs.Index(i).MapRange()
			for iter.Next() {
				res.SetMapIndex(iter.Key(), iter.Value())
			}
		}
		return res.Interface()
	} else {
		panic("Atleast provide one map to merge")
	}
}
