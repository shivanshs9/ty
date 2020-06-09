package ty

import (
	"reflect"
)

// TypeVariable is the underlying type of every type variable used in
// parametric types. It should not be used directly. Instead, use
//
//	type myOwnTypeVariable TypeVariable
//
// to create your own type variable. For your convenience, this package
// defines some type variables for you. (e.g., `A`, `B`, `C`, ...)
type TypeVariable struct {
	noImitation struct{}
}

// tyvarUnderlyingType is used to discover types that are type variables.
// Namely, any type variable must be convertible to `TypeVariable`.
var tyvarUnderlyingType = reflect.TypeOf(TypeVariable{})

// The A type is a dummy type. It is a stand-in for any Go type,
// but represents the same type for any given function invocation.
type A TypeVariable

// The B type is a dummy type. It is a stand-in for any Go type,
// but represents the same type for any given function invocation.
type B TypeVariable

// The C type is a dummy type. It is a stand-in for any Go type,
// but represents the same type for any given function invocation.
type C TypeVariable

// The D type is a dummy type. It is a stand-in for any Go type,
// but represents the same type for any given function invocation.
type D TypeVariable

// The E type is a dummy type. It is a stand-in for any Go type,
// but represents the same type for any given function invocation.
type E TypeVariable

// The F type is a dummy type. It is a stand-in for any Go type,
// but represents the same type for any given function invocation.
type F TypeVariable

// The G type is a dummy type. It is a stand-in for any Go type,
// but represents the same type for any given function invocation.
type G TypeVariable
