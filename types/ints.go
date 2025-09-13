package types

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/levelfourab/conversions-go"
)

func withIntConversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(""), func(_ context.Context, v any) (any, error) {
			return strconv.FormatInt(int64(v.(int)), 10), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(int8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int)
			if i >= math.MinInt8 && i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert int to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(int16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int)
			if i >= math.MinInt16 && i <= math.MaxInt16 {
				return int16(i), nil
			}
			return 0, fmt.Errorf("unable to convert int to int16")
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(int32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int)
			if i >= math.MinInt32 && i <= math.MaxInt32 {
				return int32(i), nil
			}
			return 0, fmt.Errorf("unable to convert int to int32")
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(int64(0)), func(_ context.Context, v any) (any, error) {
			return int64(v.(int)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(uint(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int)
			if i >= 0 {
				return uint(i), nil
			}
			return 0, fmt.Errorf("unable to convert int to uint")
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(uint8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int)
			if i >= 0 && i <= math.MaxUint8 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert int to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(uint16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int)
			if i >= 0 && i <= math.MaxUint16 {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert int to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(uint32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int)
			if i >= 0 && i <= math.MaxUint32 {
				return uint32(i), nil
			}
			return 0, fmt.Errorf("unable to convert int to uint32")
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(uint64(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int)
			if i >= 0 {
				return uint64(i), nil
			}
			return 0, fmt.Errorf("unable to convert int to uint64")
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(float32(0)), func(_ context.Context, v any) (any, error) {
			return float32(v.(int)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(float64(0)), func(_ context.Context, v any) (any, error) {
			return float64(v.(int)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int(0)), reflect.TypeOf(false), func(_ context.Context, v any) (any, error) {
			i := v.(int)
			switch i {
			case 0:
				return false, nil
			case 1:
				return true, nil
			default:
				return false, fmt.Errorf("unable to convert int to bool")
			}
		}),
	)
}

func withInt8Conversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(""), func(_ context.Context, v any) (any, error) {
			return strconv.FormatInt(int64(v.(int8)), 10), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(int(0)), func(_ context.Context, v any) (any, error) {
			return int(v.(int8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(int16(0)), func(_ context.Context, v any) (any, error) {
			return int16(v.(int8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(int32(0)), func(_ context.Context, v any) (any, error) {
			return int32(v.(int8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(int64(0)), func(_ context.Context, v any) (any, error) {
			return int64(v.(int8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(uint(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int8)
			if i >= 0 {
				return uint(i), nil
			}
			return 0, fmt.Errorf("unable to convert int8 to uint")
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(uint8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int8)
			if i >= 0 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert int8 to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(uint16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int8)
			if i >= 0 {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert int8 to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(uint32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int8)
			if i >= 0 {
				return uint32(i), nil
			}
			return 0, fmt.Errorf("unable to convert int8 to uint32")
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(uint64(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int8)
			if i >= 0 {
				return uint64(i), nil
			}
			return 0, fmt.Errorf("unable to convert int8 to uint64")
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(float32(0)), func(_ context.Context, v any) (any, error) {
			return float32(v.(int8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(float64(0)), func(_ context.Context, v any) (any, error) {
			return float64(v.(int8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int8(0)), reflect.TypeOf(false), func(_ context.Context, v any) (any, error) {
			i := v.(int8)
			switch i {
			case 0:
				return false, nil
			case 1:
				return true, nil
			default:
				return false, fmt.Errorf("unable to convert int8 to bool")
			}
		}),
	)
}

func withInt16Conversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(""), func(_ context.Context, v any) (any, error) {
			return strconv.FormatInt(int64(v.(int16)), 10), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(int(0)), func(_ context.Context, v any) (any, error) {
			return int(v.(int16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(int8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int16)
			if i >= math.MinInt8 && i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert int16 to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(int32(0)), func(_ context.Context, v any) (any, error) {
			return int32(v.(int16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(int64(0)), func(_ context.Context, v any) (any, error) {
			return int64(v.(int16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(uint(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int16)
			if i >= 0 {
				return uint(i), nil
			}
			return 0, fmt.Errorf("unable to convert int16 to uint")
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(uint8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int16)
			if i >= 0 && i <= math.MaxUint8 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert int16 to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(uint16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int16)
			if i >= 0 {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert int16 to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(uint32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int16)
			if i >= 0 {
				return uint32(i), nil
			}
			return 0, fmt.Errorf("unable to convert int16 to uint32")
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(uint64(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int16)
			if i >= 0 {
				return uint64(i), nil
			}
			return 0, fmt.Errorf("unable to convert int16 to uint64")
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(float32(0)), func(_ context.Context, v any) (any, error) {
			return float32(v.(int16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(float64(0)), func(_ context.Context, v any) (any, error) {
			return float64(v.(int16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int16(0)), reflect.TypeOf(false), func(_ context.Context, v any) (any, error) {
			i := v.(int16)
			switch i {
			case 0:
				return false, nil
			case 1:
				return true, nil
			default:
				return false, fmt.Errorf("unable to convert int16 to bool")
			}
		}),
	)
}

func withInt32Conversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(""), func(_ context.Context, v any) (any, error) {
			return strconv.FormatInt(int64(v.(int32)), 10), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(int(0)), func(_ context.Context, v any) (any, error) {
			return int(v.(int32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(int8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int32)
			if i >= math.MinInt8 && i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert int32 to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(int16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int32)
			if i >= math.MinInt16 && i <= math.MaxInt16 {
				return int16(i), nil
			}
			return 0, fmt.Errorf("unable to convert int32 to int16")
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(int64(0)), func(_ context.Context, v any) (any, error) {
			return int64(v.(int32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(uint(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int32)
			if i >= 0 {
				return uint(i), nil
			}
			return 0, fmt.Errorf("unable to convert int32 to uint")
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(uint8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int32)
			if i >= 0 && i <= math.MaxUint8 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert int32 to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(uint16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int32)
			if i >= 0 && i <= math.MaxUint16 {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert int32 to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(uint32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int32)
			if i >= 0 {
				return uint32(i), nil
			}
			return 0, fmt.Errorf("unable to convert int32 to uint32")
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(uint64(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int32)
			if i >= 0 {
				return uint64(i), nil
			}
			return 0, fmt.Errorf("unable to convert int32 to uint64")
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(float32(0)), func(_ context.Context, v any) (any, error) {
			return float32(v.(int32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(float64(0)), func(_ context.Context, v any) (any, error) {
			return float64(v.(int32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int32(0)), reflect.TypeOf(false), func(_ context.Context, v any) (any, error) {
			i := v.(int32)
			switch i {
			case 0:
				return false, nil
			case 1:
				return true, nil
			default:
				return false, fmt.Errorf("unable to convert int32 to bool")
			}
		}),
	)
}

func withInt64Conversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(""), func(_ context.Context, v any) (any, error) {
			return strconv.FormatInt(v.(int64), 10), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(int(0)), func(_ context.Context, v any) (any, error) {
			return int(v.(int64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(int8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int64)
			if i >= math.MinInt8 && i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert int64 to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(int16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int64)
			if i >= math.MinInt16 && i <= math.MaxInt16 {
				return int16(i), nil
			}
			return 0, fmt.Errorf("unable to convert int64 to int16")
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(int32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int64)
			if i >= math.MinInt32 && i <= math.MaxInt32 {
				return int32(i), nil
			}
			return 0, fmt.Errorf("unable to convert int64 to int32")
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(uint(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int64)
			if i >= 0 {
				return uint(i), nil
			}
			return 0, fmt.Errorf("unable to convert int64 to uint")
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(uint8(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int64)
			if i >= 0 && i <= math.MaxUint8 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert int64 to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(uint16(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int64)
			if i >= 0 && i <= math.MaxUint16 {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert int64 to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(uint32(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int64)
			if i >= 0 && i <= math.MaxUint32 {
				return uint32(i), nil
			}
			return 0, fmt.Errorf("unable to convert int64 to uint32")
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(uint64(0)), func(_ context.Context, v any) (any, error) {
			i := v.(int64)
			if i >= 0 {
				return uint64(i), nil
			}
			return 0, fmt.Errorf("unable to convert int64 to uint64")
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(float32(0)), func(_ context.Context, v any) (any, error) {
			return float32(v.(int64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(float64(0)), func(_ context.Context, v any) (any, error) {
			return float64(v.(int64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(int64(0)), reflect.TypeOf(false), func(_ context.Context, v any) (any, error) {
			i := v.(int64)
			switch i {
			case 0:
				return false, nil
			case 1:
				return true, nil
			default:
				return false, fmt.Errorf("unable to convert int64 to bool")
			}
		}),
	)
}
