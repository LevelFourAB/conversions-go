package conversions

type options[T comparable, D any] struct {
	typeExtractor func(D) (T, error)

	converters []TypedConversion[T, D]
}

// Option represents an option for Conversions.
type Option[T comparable, D any] func(*options[T, D]) error

// WithTypeExtractor sets the type extractor to use when converting values.
// The type extractor is asked to extract the type of a value, the type is
// then used to find suitable conversions.
func WithTypeExtractor[T comparable, D any](extractor func(D) (T, error)) Option[T, D] {
	return func(o *options[T, D]) error {
		o.typeExtractor = extractor
		return nil
	}
}

// WithConversion makes a conversion available for use.
func WithConversion[T comparable, D any](from T, to T, conv Converter[D]) Option[T, D] {
	return WithTypedConversion(TypedConversion[T, D]{from, to, conv})
}

// WithTypedConversion makes a typed conversion available for use.
func WithTypedConversion[T comparable, D any](conv TypedConversion[T, D]) Option[T, D] {
	return func(o *options[T, D]) error {
		o.converters = append(o.converters, conv)
		return nil
	}
}

// With combines multiple options into a single option.
func With[T comparable, D any](opts ...Option[T, D]) Option[T, D] {
	return func(o *options[T, D]) error {
		for _, opt := range opts {
			if err := opt(o); err != nil {
				return err
			}
		}
		return nil
	}
}
