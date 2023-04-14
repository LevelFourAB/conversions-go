package types

import "reflect"

// TypeString represents the string type.
var TypeString = reflect.TypeOf("")

// TypeInt represents the int type.
var TypeInt = reflect.TypeOf(0)

// TypeInt8 represents the int8 type.
var TypeInt8 = reflect.TypeOf(int8(0))

// TypeInt16 represents the int16 type.
var TypeInt16 = reflect.TypeOf(int16(0))

// TypeInt32 represents the int32 type.
var TypeInt32 = reflect.TypeOf(int32(0))

// TypeInt64 represents the int64 type.
var TypeInt64 = reflect.TypeOf(int64(0))

// TypeUint represents the uint type.
var TypeUint = reflect.TypeOf(uint(0))

// TypeUint8 represents the uint8 type.
var TypeUint8 = reflect.TypeOf(uint8(0))

// TypeUint16 represents the uint16 type.
var TypeUint16 = reflect.TypeOf(uint16(0))

// TypeUint32 represents the uint32 type.
var TypeUint32 = reflect.TypeOf(uint32(0))

// TypeUint64 represents the uint64 type.
var TypeUint64 = reflect.TypeOf(uint64(0))

// TypeFloat32 represents the float32 type.
var TypeFloat32 = reflect.TypeOf(float32(0))

// TypeFloat64 represents the float64 type.
var TypeFloat64 = reflect.TypeOf(float64(0))

// TypeBool represents the bool type.
var TypeBool = reflect.TypeOf(false)

// TypeExtractor is a function for use with conversions.WithTypeExtractor that
// returns the reflect.Type of the value passed in.
func TypeExtractor(v any) (reflect.Type, error) {
	return reflect.TypeOf(v), nil
}
