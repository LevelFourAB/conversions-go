package conversions_test

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/levelfourab/conversions-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func testTypeExtractor(v any) (reflect.Type, error) {
	return reflect.TypeOf(v), nil
}

var strToInt conversions.Converter[any] = func(ctx context.Context, v any) (any, error) {
	i, err := strconv.ParseInt(v.(string), 0, 0)
	if err == nil {
		return int(i), nil
	}
	return 0, fmt.Errorf("unable to convert string to int")
}

var intToStr conversions.Converter[any] = func(ctx context.Context, v any) (any, error) {
	i := v.(int)
	return strconv.FormatInt(int64(i), 10), nil
}

var intToUint conversions.Converter[any] = func(ctx context.Context, v any) (any, error) {
	i := v.(int)
	if i >= 0 {
		return uint(i), nil
	}
	return 0, fmt.Errorf("unable to convert negative int to uint")
}

var uintToFloat32 conversions.Converter[any] = func(ctx context.Context, v any) (any, error) {
	i := v.(uint)
	return float32(i), nil
}

type MediaType string

type StringWithMediaType struct {
	MediaType MediaType
	Value     string
}

var _ = Describe("Conversions", func() {
	It("will fail if no type extractor is provided", func() {
		_, err := conversions.New[string, any]()
		Expect(err).To(HaveOccurred())
	})

	It("convert will fail if the type extractor returns an error", func(ctx context.Context) {
		conversions, err := conversions.New(
			conversions.WithTypeExtractor(func(v any) (reflect.Type, error) {
				return nil, fmt.Errorf("test error")
			}),
		)
		Expect(err).ToNot(HaveOccurred())

		conversions.AddConversion(reflect.TypeOf(""), reflect.TypeOf(int(0)), strToInt)

		_, err = conversions.Convert(ctx, "123", reflect.TypeOf(int(0)))
		Expect(err).To(HaveOccurred())
	})

	It("convert will fail if no conversion is found", func(ctx context.Context) {
		conversions, err := conversions.New(
			conversions.WithTypeExtractor(testTypeExtractor),
		)
		Expect(err).ToNot(HaveOccurred())

		_, err = conversions.Convert(ctx, "123", reflect.TypeOf(int(0)))
		Expect(err).To(HaveOccurred())
	})

	It("can convert a same type to same type", func(ctx context.Context) {
		conversions, err := conversions.New(
			conversions.WithTypeExtractor(testTypeExtractor),
		)
		Expect(err).ToNot(HaveOccurred())

		converted, err := conversions.Convert(ctx, "test", reflect.TypeOf(""))
		Expect(err).ToNot(HaveOccurred())
		Expect(converted).To(Equal("test"))
	})

	It("can convert directly", func(ctx context.Context) {
		conversions, err := conversions.New(
			conversions.WithTypeExtractor(testTypeExtractor),
		)
		Expect(err).ToNot(HaveOccurred())

		conversions.AddConversion(reflect.TypeOf(""), reflect.TypeOf(0), strToInt)

		converted, err := conversions.Convert(ctx, "123", reflect.TypeOf(0))
		Expect(err).ToNot(HaveOccurred())
		Expect(converted).To(Equal(123))
	})

	It("can convert indirectly", func(ctx context.Context) {
		conversions, err := conversions.New(
			conversions.WithTypeExtractor(testTypeExtractor),
		)
		Expect(err).ToNot(HaveOccurred())

		conversions.AddConversion(reflect.TypeOf(""), reflect.TypeOf(0), strToInt)
		conversions.AddConversion(reflect.TypeOf(0), reflect.TypeOf(uint(0)), intToUint)

		converted, err := conversions.Convert(ctx, "123", reflect.TypeOf(uint(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(converted).To(Equal(uint(123)))
	})

	It("can convert indirectly with 1 intermediate step", func(ctx context.Context) {
		conversions, err := conversions.New(
			conversions.WithTypeExtractor(testTypeExtractor),
		)
		Expect(err).ToNot(HaveOccurred())

		conversions.AddConversion(reflect.TypeOf(""), reflect.TypeOf(0), strToInt)
		conversions.AddConversion(reflect.TypeOf(0), reflect.TypeOf(uint(0)), intToUint)
		conversions.AddConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(float32(0)), uintToFloat32)

		converted, err := conversions.Convert(ctx, "123", reflect.TypeOf(float32(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(converted).To(Equal(float32(123)))
	})

	It("can convert indirectly with 2 intermediate steps", func(ctx context.Context) {
		conversions, err := conversions.New(
			conversions.WithTypeExtractor(testTypeExtractor),
		)
		Expect(err).ToNot(HaveOccurred())

		conversions.AddConversion(reflect.TypeOf(""), reflect.TypeOf(0), strToInt)
		conversions.AddConversion(reflect.TypeOf(0), reflect.TypeOf(uint(0)), intToUint)
		conversions.AddConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(float32(0)), uintToFloat32)
		conversions.AddConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(float64(0)), func(ctx context.Context, v any) (any, error) {
			return float64(v.(float32)), nil
		})

		converted, err := conversions.Convert(ctx, "123", reflect.TypeOf(float64(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(converted).To(Equal(float64(123)))
	})

	It("will handle loops in the conversion graph", func(ctx context.Context) {
		// This a special case where looking up a conversion requires traversing
		// a loop.
		//
		// A loop can be created by adding a conversion from A to B, and then
		// from B to A. In this test we force this by only adding a symmetrical
		// conversion.
		conversions, err := conversions.New(
			conversions.WithTypeExtractor(testTypeExtractor),
		)
		Expect(err).ToNot(HaveOccurred())

		conversions.AddConversion(reflect.TypeOf(""), reflect.TypeOf(0), strToInt)
		conversions.AddConversion(reflect.TypeOf(0), reflect.TypeOf(""), intToStr)

		// This should not get stuck in an infinite loop, but will error
		_, err = conversions.Convert(ctx, "123", reflect.TypeOf(uint(0)))
		Expect(err).To(HaveOccurred())
	})

	It("will handle custom type", func(ctx context.Context) {
		conversions, err := conversions.New(
			conversions.WithTypeExtractor(func(v StringWithMediaType) (MediaType, error) {
				return v.MediaType, nil
			}),
		)
		Expect(err).ToNot(HaveOccurred())

		conversions.AddConversion("text/html", "text/plain", func(ctx context.Context, v StringWithMediaType) (StringWithMediaType, error) {
			return StringWithMediaType{
				MediaType: "text/plain",
				Value:     v.Value,
			}, nil
		})

		converted, err := conversions.Convert(ctx, StringWithMediaType{
			MediaType: "text/html",
			Value:     "test",
		}, "text/plain")
		Expect(err).ToNot(HaveOccurred())
		Expect(converted).To(Equal(StringWithMediaType{
			MediaType: "text/plain",
			Value:     "test",
		}))
	})
})
