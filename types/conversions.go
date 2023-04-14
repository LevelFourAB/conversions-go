package types

import (
	"reflect"

	"github.com/levelfourab/conversions-go"
)

type Option = conversions.Option[reflect.Type, any]

type Converter = conversions.Converter[any]

// Conversion is a type alias for a Conversions instance that is used for
// converting between Go types.
type Conversions = conversions.Conversions[reflect.Type, any]

// New creates a new Conversions instance configured to convert between Go
// types.
func New(opts ...Option) (*Conversions, error) {
	defaultOptions := []Option{
		conversions.WithTypeExtractor(TypeExtractor),
	}

	return conversions.New(append(defaultOptions, opts...)...)
}

// WithDefaultConversions makes default type conversions available, this will
// add converters between the built-in types, such as string to int. The
// following types are supported: string, int, int8, in16, int32, int64, uint,
// uint8, uint16, uint32, uint64, float32, float64, bool.
//
// These converters run in safe mode where they will error if they can not
// convert the value, such as turning a non-numeric string into an int. Or
// uint64 into an int32 if the value is too large.
func WithDefaultConversions() Option {
	return conversions.With(
		withStringConversions(),
		withIntConversions(),
		withInt8Conversions(),
		withInt16Conversions(),
		withInt32Conversions(),
		withInt64Conversions(),
		withUintConversions(),
		withUint8Conversions(),
		withUint16Conversions(),
		withUint32Conversions(),
		withUint64Conversions(),
		withFloat32Conversions(),
		withFloat64Conversions(),
	)
}
