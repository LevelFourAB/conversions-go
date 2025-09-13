package types

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/levelfourab/conversions-go"
)

func withFloat32Conversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(""), func(_ context.Context, v any) (any, error) {
			return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32), nil
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(int(0)), func(_ context.Context, v any) (any, error) {
			return int(v.(float32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(int8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float32)
			if i >= math.MinInt8 && i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert float32 to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(int16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float32)
			if i >= math.MinInt16 && i <= math.MaxInt16 {
				return int16(i), nil
			}
			return 0, fmt.Errorf("unable to convert float32 to int16")
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(int32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float32)
			if i >= math.MinInt32 && i <= math.MaxInt32 {
				return int32(i), nil
			}
			return 0, fmt.Errorf("unable to convert float32 to int32")
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(int64(0)), func(_ context.Context, v any) (any, error) {
			return int64(v.(float32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(uint(0)), func(_ context.Context, v any) (any, error) {
			return uint(v.(float32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(uint8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float32)
			if i >= 0 && i <= math.MaxUint8 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert float32 to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(uint16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float32)
			if i >= 0 && i <= math.MaxUint16 {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert float32 to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(uint32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float32)
			if i >= 0 {
				return uint32(i), nil
			}
			return 0, fmt.Errorf("unable to convert float32 to uint32")
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(uint64(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float32)
			if i >= 0 {
				return uint64(i), nil
			}
			return 0, fmt.Errorf("unable to convert float32 to uint64")
		}),
		conversions.WithConversion(reflect.TypeOf(float32(0)), reflect.TypeOf(float64(0)), func(_ context.Context, v any) (any, error) {
			return float64(v.(float32)), nil
		}),
	)
}

func withFloat64Conversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(""), func(_ context.Context, v any) (any, error) {
			return strconv.FormatFloat(v.(float64), 'f', -1, 64), nil
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(int(0)), func(_ context.Context, v any) (any, error) {
			return int(v.(float64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(int8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float64)
			if i >= math.MinInt8 && i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert float64 to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(int16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float64)
			if i >= math.MinInt16 && i <= math.MaxInt16 {
				return int16(i), nil
			}
			return 0, fmt.Errorf("unable to convert float64 to int16")
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(int32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float64)
			if i >= math.MinInt32 && i <= math.MaxInt32 {
				return int32(i), nil
			}
			return 0, fmt.Errorf("unable to convert float64 to int32")
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(int64(0)), func(_ context.Context, v any) (any, error) {
			return int64(v.(float64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(uint(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float64)
			if i >= 0 {
				return uint(i), nil
			}
			return 0, fmt.Errorf("unable to convert float64 to uint")
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(uint8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float64)
			if i >= 0 && i <= math.MaxUint8 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert float64 to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(uint16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float64)
			if i >= 0 && i <= math.MaxUint16 {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert float64 to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(uint32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float64)
			if i >= 0 && i <= math.MaxUint32 {
				return uint32(i), nil
			}
			return 0, fmt.Errorf("unable to convert float64 to uint32")
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(uint64(0)), func(_ context.Context, v any) (any, error) {
			return uint64(v.(float64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(float64(0)), reflect.TypeOf(float32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(float64)
			if i >= math.SmallestNonzeroFloat32 && i <= math.MaxFloat32 {
				return float32(i), nil
			}
			return 0, fmt.Errorf("unable to convert float64 to float32")
		}),
	)
}
