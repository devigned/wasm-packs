package codec

import (
	"encoding"
	"strconv"
	"unsafe"
)

// DecodeNil calls DecodeNil on v if v implements NilDecoder.
func DecodeNil(v any) error {
	if v, ok := v.(NilDecoder); ok {
		return v.DecodeNil()
	}
	return nil
}

// DecodeBool decodes a boolean value into v.
// If *v is a pointer to a bool, then a bool will be allocated.
// If v implements BoolDecoder, then DecodeBool(b) is called.
func DecodeBool(v any, b bool) error {
	switch v := v.(type) {
	case *bool:
		*v = b
	case **bool:
		*Must(v) = b
	case BoolDecoder:
		return v.DecodeBool(b)
	}
	return nil
}

// DecodeNumber decodes a number encoded as a string into v.
// The following core types are supported:
// int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, and float64.
// Pointers to the above types are also supported, and will be allocated if necessary.
// The interface types IntDecoder, and FloatDecoder are also supported.
// If unable to decode into a numeric type, it will fall back to DecodeString.
func DecodeNumber(v any, n string) error {
	switch v := v.(type) {
	case *int:
		return decodeSignedValue(v, n)
	case **int:
		return decodeSignedValue(Must(v), n)
	case *int8:
		return decodeSignedValue(v, n)
	case **int8:
		return decodeSignedValue(Must(v), n)
	case *int16:
		return decodeSignedValue(v, n)
	case **int16:
		return decodeSignedValue(Must(v), n)
	case *int32:
		return decodeSignedValue(v, n)
	case **int32:
		return decodeSignedValue(Must(v), n)
	case *int64:
		return decodeSignedValue(v, n)
	case **int64:
		return decodeSignedValue(Must(v), n)

	case *uint:
		return decodeUnsignedValue(v, n)
	case **uint:
		return decodeUnsignedValue(Must(v), n)
	case *uint8:
		return decodeUnsignedValue(v, n)
	case **uint8:
		return decodeUnsignedValue(Must(v), n)
	case *uint16:
		return decodeUnsignedValue(v, n)
	case **uint16:
		return decodeUnsignedValue(Must(v), n)
	case *uint32:
		return decodeUnsignedValue(v, n)
	case **uint32:
		return decodeUnsignedValue(Must(v), n)
	case *uint64:
		return decodeUnsignedValue(v, n)
	case **uint64:
		return decodeUnsignedValue(Must(v), n)

	case *float32:
		return decodeFloatValue(v, n)
	case **float32:
		return decodeFloatValue(Must(v), n)
	case *float64:
		return decodeFloatValue(v, n)
	case **float64:
		return decodeFloatValue(Must(v), n)

	case IntDecoder[int]:
		return decodeSigned(v, n)
	case IntDecoder[int8]:
		return decodeSigned(v, n)
	case IntDecoder[int16]:
		return decodeSigned(v, n)
	case IntDecoder[int32]:
		return decodeSigned(v, n)
	case IntDecoder[int64]:
		return decodeSigned(v, n)

	case IntDecoder[uint]:
		return decodeUnsigned(v, n)
	case IntDecoder[uint8]:
		return decodeUnsigned(v, n)
	case IntDecoder[uint16]:
		return decodeUnsigned(v, n)
	case IntDecoder[uint32]:
		return decodeUnsigned(v, n)
	case IntDecoder[uint64]:
		return decodeUnsigned(v, n)

	case FloatDecoder[float32]:
		return decodeFloat(v, n)
	case FloatDecoder[float64]:
		return decodeFloat(v, n)
	}

	return DecodeString(v, n)
}

func decodeSignedValue[T Signed](v *T, n string) error {
	i, err := strconv.ParseInt(n, 10, int(unsafe.Sizeof(*v)))
	if err != nil {
		return err
	}
	*v = T(i)
	return nil
}

func decodeSigned[T Signed](v IntDecoder[T], n string) error {
	var x T
	i, err := strconv.ParseInt(n, 10, int(unsafe.Sizeof(x))*8)
	if err != nil {
		return err
	}
	return v.DecodeInt(T(i))
}

func decodeUnsignedValue[T Unsigned](v *T, n string) error {
	i, err := strconv.ParseUint(n, 10, int(unsafe.Sizeof(*v)))
	if err != nil {
		return err
	}
	*v = T(i)
	return nil
}

func decodeUnsigned[T Unsigned](v IntDecoder[T], n string) error {
	var x T
	i, err := strconv.ParseUint(n, 10, int(unsafe.Sizeof(x))*8)
	if err != nil {
		return err
	}
	return v.DecodeInt(T(i))
}

func decodeFloatValue[T Float](v *T, n string) error {
	f, err := strconv.ParseFloat(n, int(unsafe.Sizeof(*v)))
	if err != nil {
		return err
	}
	*v = T(f)
	return nil
}

func decodeFloat[T Float](v FloatDecoder[T], n string) error {
	var x T
	f, err := strconv.ParseFloat(n, int(unsafe.Sizeof(x))*8)
	if err != nil {
		return err
	}
	return v.DecodeFloat(T(f))
}

// DecodeString decodes s into v. The following types are supported:
// string, *string, and StringDecoder. It will fall back to DecodeBytes
// to attempt to decode into a byte slice or binary decoder.
func DecodeString(v any, s string) error {
	switch v := v.(type) {
	case *string:
		*v = s
		return nil
	case **string:
		*v = &s
		return nil
	case StringDecoder:
		return v.DecodeString(s)
	}
	return DecodeBytes(v, []byte(s))
}

// DecodeBytes decodes data into v. The following types are supported:
// []byte, BytesDecoder, encoding.BinaryUnmarshaler, and encoding.TextUnmarshaler.
func DecodeBytes(v any, data []byte) error {
	switch v := v.(type) {
	case *[]byte:
		Resize(v, len(data))
		copy(*v, data)
	case BytesDecoder:
		return v.DecodeBytes(data)
	case encoding.BinaryUnmarshaler:
		return v.UnmarshalBinary(data)
	case encoding.TextUnmarshaler:
		return v.UnmarshalText(data)
	}
	return nil
}

// DecodeSlice adapts slice s into an ElementDecoder and decodes it.
func DecodeSlice[T comparable](dec Decoder, s *[]T) error {
	return dec.Decode(Slice(s))
}

// DecodeMap adapts a string-keyed map m into a FieldDecoder and decodes it.
func DecodeMap[K ~string, V any](dec Decoder, m *map[K]V) error {
	return dec.Decode(Map(m))
}
