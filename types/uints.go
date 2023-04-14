package types

import (
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/levelfourab/conversions-go"
)

func withUintConversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(""), func(v any) (any, error) {
			return strconv.FormatUint(uint64(v.(uint)), 10), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(int(0)), func(v any) (any, error) {
			i := v.(uint)
			if i <= math.MaxInt64 {
				return int(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint to int")
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(int8(0)), func(v any) (any, error) {
			i := v.(uint)
			if i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(int16(0)), func(v any) (any, error) {
			i := v.(uint)
			if i <= math.MaxInt16 {
				return int16(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint to int16")
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(int32(0)), func(v any) (any, error) {
			i := v.(uint)
			if i <= math.MaxInt32 {
				return int32(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint to int32")
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(int64(0)), func(v any) (any, error) {
			i := v.(uint)
			if i <= math.MaxInt64 {
				return int64(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint to int64")
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(uint8(0)), func(v any) (any, error) {
			i := v.(uint)
			if i <= math.MaxUint8 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(uint16(0)), func(v any) (any, error) {
			i := v.(uint)
			if i <= math.MaxUint16 {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(uint32(0)), func(v any) (any, error) {
			i := v.(uint)
			if i <= math.MaxUint32 {
				return uint32(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint to uint32")
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(uint64(0)), func(v any) (any, error) {
			return uint64(v.(uint)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(float32(0)), func(v any) (any, error) {
			return float32(v.(uint)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(float64(0)), func(v any) (any, error) {
			return float64(v.(uint)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint(0)), reflect.TypeOf(true), func(v any) (any, error) {
			i := v.(uint)
			if i == 0 {
				return false, nil
			} else if i == 1 {
				return true, nil
			}
			return false, fmt.Errorf("unable to convert uint to bool")
		}),
	)
}

func withUint8Conversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(""), func(v any) (any, error) {
			return strconv.FormatUint(uint64(v.(uint8)), 10), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(int(0)), func(v any) (any, error) {
			return int(v.(uint8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(int8(0)), func(v any) (any, error) {
			i := v.(uint8)
			if i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint8 to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(int16(0)), func(v any) (any, error) {
			return int16(v.(uint8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(int32(0)), func(v any) (any, error) {
			return int32(v.(uint8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(int64(0)), func(v any) (any, error) {
			return int64(v.(uint8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(uint(0)), func(v any) (any, error) {
			return uint(v.(uint8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(uint16(0)), func(v any) (any, error) {
			return uint16(v.(uint8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(uint32(0)), func(v any) (any, error) {
			return uint32(v.(uint8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(uint64(0)), func(v any) (any, error) {
			return uint64(v.(uint8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(float32(0)), func(v any) (any, error) {
			return float32(v.(uint8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(float64(0)), func(v any) (any, error) {
			return float64(v.(uint8)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint8(0)), reflect.TypeOf(true), func(v any) (any, error) {
			i := v.(uint8)
			if i == 0 {
				return false, nil
			} else if i == 1 {
				return true, nil
			}
			return false, fmt.Errorf("unable to convert uint8 to bool")
		}),
	)
}

func withUint16Conversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(""), func(v any) (any, error) {
			return strconv.FormatUint(uint64(v.(uint16)), 10), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(int(0)), func(v any) (any, error) {
			return int(v.(uint16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(int8(0)), func(v any) (any, error) {
			i := v.(uint16)
			if i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint16 to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(int16(0)), func(v any) (any, error) {
			i := v.(uint16)
			if i <= math.MaxInt16 {
				return int16(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint16 to int16")
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(int32(0)), func(v any) (any, error) {
			return int32(v.(uint16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(int64(0)), func(v any) (any, error) {
			return int64(v.(uint16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(uint(0)), func(v any) (any, error) {
			return uint(v.(uint16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(uint8(0)), func(v any) (any, error) {
			i := v.(uint16)
			if i <= math.MaxUint8 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint16 to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(uint32(0)), func(v any) (any, error) {
			return uint32(v.(uint16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(uint64(0)), func(v any) (any, error) {
			return uint64(v.(uint16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(float32(0)), func(v any) (any, error) {
			return float32(v.(uint16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(float64(0)), func(v any) (any, error) {
			return float64(v.(uint16)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint16(0)), reflect.TypeOf(true), func(v any) (any, error) {
			i := v.(uint16)
			if i == 0 {
				return false, nil
			} else if i == 1 {
				return true, nil
			}
			return false, fmt.Errorf("unable to convert uint16 to bool")
		}),
	)
}

func withUint32Conversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(""), func(v any) (any, error) {
			return strconv.FormatUint(uint64(v.(uint32)), 10), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(int(0)), func(v any) (any, error) {
			return int(v.(uint32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(int8(0)), func(v any) (any, error) {
			i := v.(uint32)
			if i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint32 to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(int16(0)), func(v any) (any, error) {
			i := v.(uint32)
			if i <= math.MaxInt16 {
				return int16(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint32 to int16")
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(int32(0)), func(v any) (any, error) {
			i := v.(uint32)
			if i <= math.MaxInt32 {
				return int32(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint32 to int32")
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(int64(0)), func(v any) (any, error) {
			return int64(v.(uint32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(uint(0)), func(v any) (any, error) {
			return uint(v.(uint32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(uint8(0)), func(v any) (any, error) {
			i := v.(uint32)
			if i <= math.MaxUint8 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint32 to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(uint16(0)), func(v any) (any, error) {
			i := v.(uint32)
			if i <= math.MaxUint16 {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint32 to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(uint64(0)), func(v any) (any, error) {
			return uint64(v.(uint32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(float32(0)), func(v any) (any, error) {
			return float32(v.(uint32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(float64(0)), func(v any) (any, error) {
			return float64(v.(uint32)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint32(0)), reflect.TypeOf(true), func(v any) (any, error) {
			i := v.(uint32)
			if i == 0 {
				return false, nil
			} else if i == 1 {
				return true, nil
			}
			return false, fmt.Errorf("unable to convert uint32 to bool")
		}),
	)
}

func withUint64Conversions() Option {
	return conversions.With(
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(""), func(v any) (any, error) {
			return strconv.FormatUint(v.(uint64), 10), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(int(0)), func(v any) (any, error) {
			return int(v.(uint64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(int8(0)), func(v any) (any, error) {
			i := v.(uint64)
			if i <= math.MaxInt8 {
				return int8(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint64 to int8")
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(int16(0)), func(v any) (any, error) {
			i := v.(uint64)
			if i <= math.MaxInt16 {
				return int16(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint64 to int16")
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(int32(0)), func(v any) (any, error) {
			i := v.(uint64)
			if i <= math.MaxInt32 {
				return int32(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint64 to int32")
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(int64(0)), func(v any) (any, error) {
			return int64(v.(uint64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(uint(0)), func(v any) (any, error) {
			return uint(v.(uint64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(uint8(0)), func(v any) (any, error) {
			i := v.(uint64)
			if i <= math.MaxUint8 {
				return uint8(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint64 to uint8")
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(uint16(0)), func(v any) (any, error) {
			i := v.(uint64)
			if i <= math.MaxUint16 {
				return uint16(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint64 to uint16")
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(uint32(0)), func(v any) (any, error) {
			i := v.(uint64)
			if i <= math.MaxUint32 {
				return uint32(i), nil
			}
			return 0, fmt.Errorf("unable to convert uint64 to uint32")
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(float32(0)), func(v any) (any, error) {
			return float32(v.(uint64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(float64(0)), func(v any) (any, error) {
			return float64(v.(uint64)), nil
		}),
		conversions.WithConversion(reflect.TypeOf(uint64(0)), reflect.TypeOf(true), func(v any) (any, error) {
			i := v.(uint64)
			if i == 0 {
				return false, nil
			} else if i == 1 {
				return true, nil
			}
			return false, fmt.Errorf("unable to convert uint64 to bool")
		}),
	)
}
