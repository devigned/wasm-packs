package codec

// Codec is any type that can encode or decode itself or an associated type.
type Codec any

// Resolver is the interface implemented by types that return a codec for the value at v.
// Values returned by Resolver should implement one or more encode or decode methods.
type Resolver interface {
	ResolveCodec(v any) Codec
}

// Resolvers is a slice of [Resolver] values. It also implements the [Resolver] interface.
type Resolvers []Resolver

// ResolveCodec walks the slice of Resolvers, returning the first non-nil value or an error.
func (rs Resolvers) ResolveCodec(v any) Codec {
	for _, r := range rs {
		c := r.ResolveCodec(v)
		if c != nil {
			return c
		}
	}
	return nil
}

// Decoder is the interface implemented by types that can decode data into Go type(s).
type Decoder interface {
	Decode(v any) error
}

// NilDecoder is the interface implemented by types that can decode from nil.
type NilDecoder interface {
	DecodeNil() error
}

// BoolDecoder is the interface implemented by types that can decode from a bool.
type BoolDecoder interface {
	DecodeBool(bool) error
}

// BytesDecoder is the interface implemented by types that can decode from a byte slice.
// It is similar to [encoding.BinaryUnmarshaler] and [encoding.TextUnmarshaler].
type BytesDecoder interface {
	DecodeBytes([]byte) error
}

// StringDecoder is the interface implemented by types that can decode from a string.
type StringDecoder interface {
	DecodeString(string) error
}

// IntDecoder is the interface implemented by types that can decode
// from an integer value. See [Integer] for the list of supported types.
type IntDecoder[T Integer] interface {
	DecodeInt(T) error
}

// FloatDecoder is the interface implemented by types that can decode
// from a floating-point value. See [Float] for the list of supported types.
type FloatDecoder[T Float] interface {
	DecodeFloat(T) error
}

// FieldDecoder is the interface implemented by types that can decode
// fields, such as structs or maps.
type FieldDecoder interface {
	DecodeField(dec Decoder, name string) error
}

// ElementDecoder is the interface implemented by types that can decode
// 0-indexed elements, such as a slice or an array.
type ElementDecoder interface {
	DecodeElement(dec Decoder, i int) error
}
