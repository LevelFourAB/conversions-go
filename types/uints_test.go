package types_test

import (
	"context"
	"reflect"

	"github.com/levelfourab/conversions-go/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert from uints", func() {
	var c *types.Conversions
	BeforeEach(func() {
		var err error
		c, err = types.New(types.WithDefaultConversions())
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("from uint", func() {
		It("should be able to convert to string", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123"))
		})

		It("should be able to convert to int", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should fail to convert positive out of range int", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint(9223372036854775808), reflect.TypeOf(int(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int8", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert positive out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint(256), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should fail to convert positive out of range int16", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint(65536), reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should fail to convert positive out of range int32", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint(4294967296), reflect.TypeOf(int32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint8", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should fail to convert positive out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint(256), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint16", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should fail to convert positive out of range uint16", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint(65536), reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should fail to convert positive out of range uint32", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint(4294967296), reflect.TypeOf(uint32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint64(123)))
		})

		It("should be able to convert to float32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123)))
		})

		It("should be able to convert to float64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(123), reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float64(123)))
		})

		It("should be able to convert to true bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(1), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(true))
		})

		It("should be able to convert to false bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint(0), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(false))
		})

		It("should fail to convert to bool", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint(2), reflect.TypeOf(true))
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("from uint8", func() {
		It("should be able to convert to string", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123"))
		})

		It("should be able to convert to int", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should be able to convert to int8", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert positive out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint8(128), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should be able to convert to int32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should be able to convert to int64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should be able to convert to uint16", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should be able to convert to uint32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should be able to convert to uint64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint64(123)))
		})

		It("should be able to convert to float32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123)))
		})

		It("should be able to convert to float64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(123), reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float64(123)))
		})

		It("should be able to convert to true bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(1), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(true))
		})

		It("should be able to convert to false bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint8(0), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(false))
		})

		It("should fail to convert to bool", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint8(2), reflect.TypeOf(true))
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("from uint16", func() {
		It("should be able to convert to string", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123"))
		})

		It("should be able to convert to int", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should be able to convert to int8", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert positive out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint16(128), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should be able to convert to int32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should be able to convert to int64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should be able to convert to uint8", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should fail to convert positive out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint16(256), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should be able to convert to uint64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint64(123)))
		})

		It("should be able to convert to float32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123)))
		})

		It("should be able to convert to float64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(123), reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float64(123)))
		})

		It("should be able to convert to true bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(1), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(true))
		})

		It("should be able to convert to false bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint16(0), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(false))
		})

		It("should fail to convert to bool", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint16(2), reflect.TypeOf(true))
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("from uint32", func() {
		It("should be able to convert to string", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123"))
		})

		It("should be able to convert to int", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should be able to convert to int8", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert positive out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint32(128), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should fail to convert positive out of range int16", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint32(32768), reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should be able to convert to int64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should be able to convert to uint8", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should fail to convert positive out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint32(256), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint16", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should fail to convert positive out of range uint16", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint32(65536), reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint64(123)))
		})

		It("should be able to convert to float32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123)))
		})

		It("should be able to convert to float64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(123), reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float64(123)))
		})

		It("should be able to convert to true bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(1), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(true))
		})

		It("should be able to convert to false bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint32(0), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(false))
		})

		It("should fail to convert to bool", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint32(2), reflect.TypeOf(true))
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("from uint64", func() {
		It("should be able to convert to string", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123"))
		})

		It("should be able to convert to int", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should be able to convert to int8", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert positive out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint64(128), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should fail to convert positive out of range int16", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint64(32768), reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should fail to convert positive out of range int32", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint64(2147483648), reflect.TypeOf(int32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should be able to convert to uint8", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should fail to convert positive out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint64(256), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint16", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should fail to convert positive out of range uint16", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint64(65536), reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should fail to convert positive out of range uint32", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint64(4294967296), reflect.TypeOf(uint32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to float32", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123)))
		})

		It("should be able to convert to float64", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(123), reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float64(123)))
		})

		It("should be able to convert to true bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(1), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(true))
		})

		It("should be able to convert to false bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, uint64(0), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(false))
		})

		It("should fail to convert to bool", func(ctx context.Context) {
			_, err := c.Convert(ctx, uint64(2), reflect.TypeOf(true))
			Expect(err).To(HaveOccurred())
		})
	})
})
