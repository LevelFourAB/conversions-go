package types_test

import (
	"reflect"

	"github.com/levelfourab/conversions-go/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert from floats", func() {
	var c *types.Conversions
	BeforeEach(func() {
		var err error
		c, err = types.New(types.WithDefaultConversions())
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("from float32", func() {
		It("should be able to convert to string", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123.45"))
		})

		It("should be able to convert to int", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should be able to convert to int8", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert negative out of range int8", func() {
			_, err := c.Convert(float32(-129.45), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int8", func() {
			_, err := c.Convert(float32(128.45), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should fail to convert negative out of range int16", func() {
			_, err := c.Convert(float32(-32769.45), reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should fail to convert positive out of range int16", func() {
			_, err := c.Convert(float32(32768.45), reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int32", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should be able to convert to int64", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should be able to convert to uint8", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should fail to convert positive out of range uint8", func() {
			_, err := c.Convert(float32(256.45), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint16", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should fail to convert positive out of range uint16", func() {
			_, err := c.Convert(float32(65536.45), reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint32", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should be able to convert to uint64", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint64(123)))
		})

		It("should be able to convert to float64", func() {
			v, err := c.Convert(float32(123.45), reflect.TypeOf(float64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(BeNumerically("~", float64(123.45), 0.001))
		})
	})

	Describe("from float64", func() {
		It("should be able to convert to string", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal("123.45"))
		})

		It("should be able to convert to int", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(int(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int(123)))
		})

		It("should be able to convert to int8", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(int8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int8(123)))
		})

		It("should fail to convert positive out of range int8", func() {
			_, err := c.Convert(float64(128.45), reflect.TypeOf(int8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int16", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(int16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int16(123)))
		})

		It("should fail to convert positive out of range int16", func() {
			_, err := c.Convert(float64(32768.45), reflect.TypeOf(int16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int32", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(int32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int32(123)))
		})

		It("should fail to convert positive out of range int32", func() {
			_, err := c.Convert(float64(2147483648.45), reflect.TypeOf(int32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to int64", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(int64(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(int64(123)))
		})

		It("should be able to convert to uint", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(uint(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint(123)))
		})

		It("should be able to convert to uint8", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(uint8(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint8(123)))
		})

		It("should fail to convert positive out of range uint8", func() {
			_, err := c.Convert(float64(256.45), reflect.TypeOf(uint8(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint16", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(uint16(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint16(123)))
		})

		It("should fail to convert positive out of range uint16", func() {
			_, err := c.Convert(float64(65536.45), reflect.TypeOf(uint16(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to uint32", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(uint32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(uint32(123)))
		})

		It("should fail to convert positive out of range uint32", func() {
			_, err := c.Convert(float64(4294967296.45), reflect.TypeOf(uint32(0)))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to convert to float32", func() {
			v, err := c.Convert(float64(123.45), reflect.TypeOf(float32(0)))
			Expect(err).ToNot(HaveOccurred())
			Expect(v).To(Equal(float32(123.45)))
		})
	})
})
