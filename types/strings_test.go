package types_test

import (
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

	It("should be able to convert to int", func() {
		v, err := c.Convert("123", reflect.TypeOf(0))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(123))
	})

	It("should be able to convert to int8", func() {
		v, err := c.Convert("123", reflect.TypeOf(int8(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(int8(123)))
	})

	It("should fail to convert negative out of range int8", func() {
		_, err := c.Convert("-129", reflect.TypeOf(int8(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range int8", func() {
		_, err := c.Convert("128", reflect.TypeOf(int8(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to int16", func() {
		v, err := c.Convert("123", reflect.TypeOf(int16(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(int16(123)))
	})

	It("should fail to convert negative out of range int16", func() {
		_, err := c.Convert("-32769", reflect.TypeOf(int16(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range int16", func() {
		_, err := c.Convert("32768", reflect.TypeOf(int16(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to int32", func() {
		v, err := c.Convert("123", reflect.TypeOf(int32(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(int32(123)))
	})

	It("should fail to convert negative out of range int32", func() {
		_, err := c.Convert("-2147483649", reflect.TypeOf(int32(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range int32", func() {
		_, err := c.Convert("2147483648", reflect.TypeOf(int32(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to int64", func() {
		v, err := c.Convert("123", reflect.TypeOf(int64(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(int64(123)))
	})

	It("should fail to convert negative out of range int64", func() {
		_, err := c.Convert("-9223372036854775809", reflect.TypeOf(int64(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range int64", func() {
		_, err := c.Convert("9223372036854775808", reflect.TypeOf(int64(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to uint", func() {
		v, err := c.Convert("123", reflect.TypeOf(uint(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(uint(123)))
	})

	It("should be able to convert to uint8", func() {
		v, err := c.Convert("123", reflect.TypeOf(uint8(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(uint8(123)))
	})

	It("should fail to convert negative out of range uint8", func() {
		_, err := c.Convert("-1", reflect.TypeOf(uint8(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range uint8", func() {
		_, err := c.Convert("256", reflect.TypeOf(uint8(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to uint16", func() {
		v, err := c.Convert("123", reflect.TypeOf(uint16(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(uint16(123)))
	})

	It("should fail to convert negative out of range uint16", func() {
		_, err := c.Convert("-1", reflect.TypeOf(uint16(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range uint16", func() {
		_, err := c.Convert("65536", reflect.TypeOf(uint16(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to uint32", func() {
		v, err := c.Convert("123", reflect.TypeOf(uint32(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(uint32(123)))
	})

	It("should fail to convert negative out of range uint32", func() {
		_, err := c.Convert("-1", reflect.TypeOf(uint32(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range uint32", func() {
		_, err := c.Convert("4294967296", reflect.TypeOf(uint32(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to uint64", func() {
		v, err := c.Convert("123", reflect.TypeOf(uint64(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(uint64(123)))
	})

	It("should fail to convert negative out of range uint64", func() {
		_, err := c.Convert("-1", reflect.TypeOf(uint64(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should fail to convert positive out of range uint64", func() {
		_, err := c.Convert("18446744073709551616", reflect.TypeOf(uint64(0)))
		Expect(err).To(HaveOccurred())
	})

	It("should be able to convert to float32", func() {
		v, err := c.Convert("123", reflect.TypeOf(float32(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(float32(123)))
	})

	It("should be able to convert to float64", func() {
		v, err := c.Convert("123", reflect.TypeOf(float64(0)))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(float64(123)))
	})

	It("should be able to convert to true bool", func() {
		v, err := c.Convert("true", reflect.TypeOf(false))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(true))
	})

	It("should be able to convert to false bool", func() {
		v, err := c.Convert("true", reflect.TypeOf(false))
		Expect(err).ToNot(HaveOccurred())
		Expect(v).To(Equal(true))
	})

	It("should fail to convert invalid bool", func() {
		_, err := c.Convert("invalid", reflect.TypeOf(false))
		Expect(err).To(HaveOccurred())
	})
})
