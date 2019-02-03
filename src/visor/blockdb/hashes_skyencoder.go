// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package blockdb

import (
	"errors"
	"math"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
)

// EncodeSizeHashes computes the size of an encoded object of type Hashes
func EncodeSizeHashes(obj *Hashes) int {
	i0 := 0

	// obj.Hashes
	i0 += 4
	{
		i1 := 0

		// x
		i1 += 32

		i0 += len(obj.Hashes) * i1
	}

	return i0
}

// EncodeHashes encodes an object of type Hashes to the buffer in encoder.Encoder.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func EncodeHashes(buf []byte, obj *Hashes) error {
	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Hashes length check
	if len(obj.Hashes) > math.MaxUint32 {
		return errors.New("obj.Hashes length exceeds math.MaxUint32")
	}

	// obj.Hashes length
	e.Uint32(uint32(len(obj.Hashes)))

	// obj.Hashes
	for _, x := range obj.Hashes {

		// x
		e.CopyBytes(x[:])

	}

	return nil
}

// DecodeHashes decodes an object of type Hashes from the buffer in encoder.Decoder.
// Returns the number of bytes used from the buffer to decode the object.
func DecodeHashes(buf []byte, obj *Hashes) (int, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Hashes

		ul, err := d.Uint32()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}

		if length != 0 {
			obj.Hashes = make([]cipher.SHA256, length)

			for z1 := range obj.Hashes {
				{
					// obj.Hashes[z1]
					if len(d.Buffer) < len(obj.Hashes[z1]) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}
					copy(obj.Hashes[z1][:], d.Buffer[:len(obj.Hashes[z1])])
					d.Buffer = d.Buffer[len(obj.Hashes[z1]):]
				}

			}
		}
	}

	return len(buf) - len(d.Buffer), nil
}
