package types_test

import (
	"context"
	"reflect"

	"github.com/levelfourab/conversions-go/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert from ints", func() {
	var c *types.Conversions
	BeforeEach(func() {
		var err error
		c, err = types.New(types.WithDefaultConversions())
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("from int", func() {
		It("should be able to convert to string", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123"))
		})

		It("should be able to convert to int", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(0))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(123))
		})

		It("should be able to convert to int8", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert negative out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, -129, reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, 128, reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should fail to convert negative out of range int16", func(ctx context.Context) {
			_, err := c.Convert(ctx, -32769, reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int16", func(ctx context.Context) {
			_, err := c.Convert(ctx, 32768, reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int32", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should fail to convert negative out of range int32", func(ctx context.Context) {
			_, err := c.Convert(ctx, -2147483649, reflect.TypeOf(int32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int32", func(ctx context.Context) {
			_, err := c.Convert(ctx, 2147483648, reflect.TypeOf(int32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int64", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should fail to convert negative out of range uint", func(ctx context.Context) {
			_, err := c.Convert(ctx, -1, reflect.TypeOf(uint(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint8", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should fail to convert negative out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, -1, reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, 256, reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint16", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should fail to convert negative out of range uint16", func(ctx context.Context) {
			_, err := c.Convert(ctx, -1, reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range uint16", func(ctx context.Context) {
			_, err := c.Convert(ctx, 65536, reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint32", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should fail to convert negative out of range uint32", func(ctx context.Context) {
			_, err := c.Convert(ctx, -1, reflect.TypeOf(uint32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range uint32", func(ctx context.Context) {
			_, err := c.Convert(ctx, 4294967296, reflect.TypeOf(uint32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint64", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint64(123)))
		})

		It("should fail to convert negative out of range uint64", func(ctx context.Context) {
			_, err := c.Convert(ctx, -1, reflect.TypeOf(uint64(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to float32", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123)))
		})

		It("should be able to convert to float64", func(ctx context.Context) {
			v, err := c.Convert(ctx, 123, reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float64(123)))
		})

		It("should be able to convert to true bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, 1, reflect.TypeOf(true))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(true))
		})

		It("should be able to convert to false bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, 0, reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(false))
		})

		It("should fail to convert to bool", func(ctx context.Context) {
			_, err := c.Convert(ctx, 2, reflect.TypeOf(false))
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("from int8", func() {
		It("should be able to convert to string", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123"))
		})

		It("should be able to convert to int", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should be able to convert to int8", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should be able to convert to int16", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should be able to convert to int32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should be able to convert to int64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should be able to convert to uint8", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should be able to convert to uint16", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should be able to convert to uint32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should be able to convert to uint64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint64(123)))
		})

		It("should be able to convert to float32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123)))
		})

		It("should be able to convert to float64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(123), reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float64(123)))
		})

		It("should be able to convert to true bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(1), reflect.TypeOf(true))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(true))
		})

		It("should be able to convert to false bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, int8(0), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(false))
		})

		It("should fail to convert to bool", func(ctx context.Context) {
			_, err := c.Convert(ctx, int8(2), reflect.TypeOf(false))
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("from int16", func() {
		It("should be able to convert to string", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123"))
		})

		It("should be able to convert to int", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should be able to convert to int8", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert negative out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int16(-1234), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int16(1234), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should be able to convert to int32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should be able to convert to int64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should be able to convert to uint8", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should fail to convert negative out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int16(-123), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int16(1234), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint16", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should be able to convert to uint32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should be able to convert to uint64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint64(123)))
		})

		It("should be able to convert to float32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123)))
		})

		It("should be able to convert to float64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(123), reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float64(123)))
		})

		It("should be able to convert to true bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(1), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(true))
		})

		It("should be able to convert to false bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, int16(0), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(false))
		})

		It("should fail to convert to bool", func(ctx context.Context) {
			_, err := c.Convert(ctx, int16(2), reflect.TypeOf(true))
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("from int32", func() {
		It("should be able to convert to string", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123"))
		})

		It("should be able to convert to int", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should be able to convert to int8", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert negative out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int32(-129), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int32(128), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should fail to convert negative out of range int16", func(ctx context.Context) {
			_, err := c.Convert(ctx, int32(-32769), reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int16", func(ctx context.Context) {
			_, err := c.Convert(ctx, int32(32768), reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should be able to convert to int64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should be able to convert to uint8", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should fail to convert negative out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int32(-1), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int32(256), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint16", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should fail to convert negative out of range uint16", func(ctx context.Context) {
			_, err := c.Convert(ctx, int32(-1), reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range uint16", func(ctx context.Context) {
			_, err := c.Convert(ctx, int32(65536), reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should fail negative out of range uint32", func(ctx context.Context) {
			_, err := c.Convert(ctx, int32(-1), reflect.TypeOf(uint32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint64(123)))
		})

		It("should be able to convert to float32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123)))
		})

		It("should be able to convert to float64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(123), reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float64(123)))
		})

		It("should be able to convert to true bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(1), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(true))
		})

		It("should be able to convert to false bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, int32(0), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(false))
		})

		It("should fail to convert to bool", func(ctx context.Context) {
			_, err := c.Convert(ctx, int32(2), reflect.TypeOf(true))
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("from int64", func() {
		It("should be able to convert to string", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123"))
		})

		It("should be able to convert to int", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should be able to convert to int8", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert negative out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(-129), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(128), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should fail to convert negative out of range int16", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(-32769), reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int16", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(32768), reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should fail to convert negative out of range int32", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(-2147483649), reflect.TypeOf(int32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int32", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(2147483648), reflect.TypeOf(int32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should fail to convert negative out of range uint", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(-1), reflect.TypeOf(uint(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint8", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should fail to convert negative out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(-1), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range uint8", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(256), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint16", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should fail to convert negative out of range uint16", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(-1), reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range uint16", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(65536), reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should fail to convert negative out of range uint32", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(-1), reflect.TypeOf(uint32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range uint32", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(4294967296), reflect.TypeOf(uint32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint64(123)))
		})

		It("should fail to convert negative out of range uint64", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(-1), reflect.TypeOf(uint64(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to float32", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123)))
		})

		It("should be able to convert to float64", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(123), reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float64(123)))
		})

		It("should be able to convert to true bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(1), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(true))
		})

		It("should be able to convert to false bool", func(ctx context.Context) {
			v, err := c.Convert(ctx, int64(0), reflect.TypeOf(false))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(false))
		})

		It("should fail to convert to bool", func(ctx context.Context) {
			_, err := c.Convert(ctx, int64(2), reflect.TypeOf(true))
			Expect(err).To(HaveOccurred())
		})
	})
})
