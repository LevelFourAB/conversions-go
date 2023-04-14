package types

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/levelfourab/conversions-go"
)

func withStringConversions() conversions.Option[reflect.Type, any] {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(int(0)), func(v any) (any, error) {
			i, err := strconv.ParseInt(v.(string), 0, 0)
			if err == nil {
				return int(i), nil
			}
			return 0, fmt.Errorf("unable to convert string to int")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(int8(0)), func(v any) (any, error) {
			i, err := strconv.ParseInt(v.(string), 0, 8)
			if err == nil {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert string to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(int16(0)), func(v any) (any, error) {
			i, err := strconv.ParseInt(v.(string), 0, 16)
			if err == nil {
				return int16(i), nil
			}
			return 0, fmt.Errorf("unable to convert string to int16")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(int32(0)), func(v any) (any, error) {
			i, err := strconv.ParseInt(v.(string), 0, 32)
			if err == nil {
				return int32(i), nil
			}
			return 0, fmt.Errorf("unable to convert string to int32")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(int64(0)), func(v any) (any, error) {
			i, err := strconv.ParseInt(v.(string), 0, 64)
			if err == nil {
				return i, nil
			}
			return 0, fmt.Errorf("unable to convert string to int64")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(uint(0)), func(v any) (any, error) {
			i, err := strconv.ParseUint(v.(string), 0, 0)
			if err == nil {
				return uint(i), nil
			}
			return 0, fmt.Errorf("unable to convert string to uint")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(uint8(0)), func(v any) (any, error) {
			i, err := strconv.ParseUint(v.(string), 0, 8)
			if err == nil {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert string to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(uint16(0)), func(v any) (any, error) {
			i, err := strconv.ParseUint(v.(string), 0, 16)
			if err == nil {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert string to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(uint32(0)), func(v any) (any, error) {
			i, err := strconv.ParseUint(v.(string), 0, 32)
			if err == nil {
				return uint32(i), nil
			}
			return 0, fmt.Errorf("unable to convert string to uint32")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(uint64(0)), func(v any) (any, error) {
			i, err := strconv.ParseUint(v.(string), 0, 64)
			if err == nil {
				return i, nil
			}
			return 0, fmt.Errorf("unable to convert string to uint64")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(float32(0)), func(v any) (any, error) {
			i, err := strconv.ParseFloat(v.(string), 32)
			if err == nil {
				return float32(i), nil
			}
			return 0, fmt.Errorf("unable to convert string to float32")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(float64(0)), func(v any) (any, error) {
			i, err := strconv.ParseFloat(v.(string), 64)
			if err == nil {
				return i, nil
			}
			return 0, fmt.Errorf("unable to convert string to float64")
		}),
		conversions.WithConversion(reflect.TypeOf(""), reflect.TypeOf(false), func(v any) (any, error) {
			i, err := strconv.ParseBool(v.(string))
			if err == nil {
				return i, nil
			}
			return false, fmt.Errorf("unable to convert string to bool")
		}),
	)
}
