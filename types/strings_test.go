package types_test

import (
	"context"
	"reflect"

	"github.com/levelfourab/conversions-go/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert from string", func() {
	var c *types.Conversions
	BeforeEach(func() {
		var err error
		c, err = types.New(types.WithDefaultConversions())
		Expect(err).ToNot(HaveOccurred())
	})

	It("should be able to convert to int", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(0))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(123))
	})

	It("should be able to convert to int8", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(int8(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(int8(123)))
	})

	It("should fail to convert negative out of range int8", func(ctx context.Context) {
		_, err := c.Convert(ctx, "-129", reflect.TypeOf(int8(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range int8", func(ctx context.Context) {
		_, err := c.Convert(ctx, "128", reflect.TypeOf(int8(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to int16", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(int16(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(int16(123)))
	})

	It("should fail to convert negative out of range int16", func(ctx context.Context) {
		_, err := c.Convert(ctx, "-32769", reflect.TypeOf(int16(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range int16", func(ctx context.Context) {
		_, err := c.Convert(ctx, "32768", reflect.TypeOf(int16(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to int32", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(int32(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(int32(123)))
	})

	It("should fail to convert negative out of range int32", func(ctx context.Context) {
		_, err := c.Convert(ctx, "-2147483649", reflect.TypeOf(int32(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range int32", func(ctx context.Context) {
		_, err := c.Convert(ctx, "2147483648", reflect.TypeOf(int32(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to int64", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(int64(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(int64(123)))
	})

	It("should fail to convert negative out of range int64", func(ctx context.Context) {
		_, err := c.Convert(ctx, "-9223372036854775809", reflect.TypeOf(int64(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range int64", func(ctx context.Context) {
		_, err := c.Convert(ctx, "9223372036854775808", reflect.TypeOf(int64(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to uint", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(uint(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(uint(123)))
	})

	It("should be able to convert to uint8", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(uint8(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(uint8(123)))
	})

	It("should fail to convert negative out of range uint8", func(ctx context.Context) {
		_, err := c.Convert(ctx, "-1", reflect.TypeOf(uint8(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range uint8", func(ctx context.Context) {
		_, err := c.Convert(ctx, "256", reflect.TypeOf(uint8(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to uint16", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(uint16(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(uint16(123)))
	})

	It("should fail to convert negative out of range uint16", func(ctx context.Context) {
		_, err := c.Convert(ctx, "-1", reflect.TypeOf(uint16(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range uint16", func(ctx context.Context) {
		_, err := c.Convert(ctx, "65536", reflect.TypeOf(uint16(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to uint32", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(uint32(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(uint32(123)))
	})

	It("should fail to convert negative out of range uint32", func(ctx context.Context) {
		_, err := c.Convert(ctx, "-1", reflect.TypeOf(uint32(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range uint32", func(ctx context.Context) {
		_, err := c.Convert(ctx, "4294967296", reflect.TypeOf(uint32(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to uint64", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(uint64(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(uint64(123)))
	})

	It("should fail to convert negative out of range uint64", func(ctx context.Context) {
		_, err := c.Convert(ctx, "-1", reflect.TypeOf(uint64(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range uint64", func(ctx context.Context) {
		_, err := c.Convert(ctx, "18446744073709551616", reflect.TypeOf(uint64(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to float32", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(float32(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(float32(123)))
	})

	It("should be able to convert to float64", func(ctx context.Context) {
		v, err := c.Convert(ctx, "123", reflect.TypeOf(float64(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(float64(123)))
	})

	It("should be able to convert to true bool", func(ctx context.Context) {
		v, err := c.Convert(ctx, "true", reflect.TypeOf(false))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(true))
	})

	It("should be able to convert to false bool", func(ctx context.Context) {
		v, err := c.Convert(ctx, "true", reflect.TypeOf(false))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(true))
	})

	It("should fail to convert invalid bool", func(ctx context.Context) {
		_, err := c.Convert(ctx, "invalid", reflect.TypeOf(false))
		Expect(err).To(HaveOccurred())
	})
})
