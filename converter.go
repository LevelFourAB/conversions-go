package conversions

import "context"

// Converter represents a function that can convert a value to another value.
type Converter[D any] func(ctx context.Context, value D) (D, error)

// TypedConversion represents a conversion between two types. It extends
// Converter with the types that the converter can convert between.
type TypedConversion[T comparable, D any] struct {
	// From is the type that the converter can convert from.
	From T
	// To is the type that the converter can convert to.
	To T
	// Converter is the converter function.
	Converter Converter[D]
}
