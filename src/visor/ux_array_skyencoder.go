// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package visor

import (
	"errors"
	"math"

	"github.com/amherag/skycoin/src/cipher/encoder"
	"github.com/amherag/skycoin/src/coin"
)

// encodeSizeUxArray computes the size of an encoded object of type UxArray
func encodeSizeUxArray(obj *UxArray) uint64 {
	i0 := uint64(0)

	// obj.UxArray
	i0 += 4
	{
		i1 := uint64(0)

		// x.Head.Time
		i1 += 8

		// x.Head.BkSeq
		i1 += 8

		// x.Body.SrcTransaction
		i1 += 32

		// x.Body.Address.Version
		i1++

		// x.Body.Address.Key
		i1 += 20

		// x.Body.Coins
		i1 += 8

		// x.Body.Hours
		i1 += 8

		i0 += uint64(len(obj.UxArray)) * i1
	}

	return i0
}

// encodeUxArray encodes an object of type UxArray to a buffer allocated to the exact size
// required to encode the object.
func encodeUxArray(obj *UxArray) ([]byte, error) {
	n := encodeSizeUxArray(obj)
	buf := make([]byte, n)

	if err := encodeUxArrayToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeUxArrayToBuffer encodes an object of type UxArray to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeUxArrayToBuffer(buf []byte, obj *UxArray) error {
	if uint64(len(buf)) < encodeSizeUxArray(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.UxArray length check
	if uint64(len(obj.UxArray)) > math.MaxUint32 {
		return errors.New("obj.UxArray length exceeds math.MaxUint32")
	}

	// obj.UxArray length
	e.Uint32(uint32(len(obj.UxArray)))

	// obj.UxArray
	for _, x := range obj.UxArray {

		// x.Head.Time
		e.Uint64(x.Head.Time)

		// x.Head.BkSeq
		e.Uint64(x.Head.BkSeq)

		// x.Body.SrcTransaction
		e.CopyBytes(x.Body.SrcTransaction[:])

		// x.Body.Address.Version
		e.Uint8(x.Body.Address.Version)

		// x.Body.Address.Key
		e.CopyBytes(x.Body.Address.Key[:])

		// x.Body.Coins
		e.Uint64(x.Body.Coins)

		// x.Body.Hours
		e.Uint64(x.Body.Hours)

	}

	return nil
}

// decodeUxArray decodes an object of type UxArray from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeUxArray(buf []byte, obj *UxArray) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.UxArray

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length != 0 {
			obj.UxArray = make([]coin.UxOut, length)

			for z1 := range obj.UxArray {
				{
					// obj.UxArray[z1].Head.Time
					i, err := d.Uint64()
					if err != nil {
						return 0, err
					}
					obj.UxArray[z1].Head.Time = i
				}

				{
					// obj.UxArray[z1].Head.BkSeq
					i, err := d.Uint64()
					if err != nil {
						return 0, err
					}
					obj.UxArray[z1].Head.BkSeq = i
				}

				{
					// obj.UxArray[z1].Body.SrcTransaction
					if len(d.Buffer) < len(obj.UxArray[z1].Body.SrcTransaction) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.UxArray[z1].Body.SrcTransaction[:], d.Buffer[:len(obj.UxArray[z1].Body.SrcTransaction)])
					d.Buffer = d.Buffer[len(obj.UxArray[z1].Body.SrcTransaction):]
				}

				{
					// obj.UxArray[z1].Body.Address.Version
					i, err := d.Uint8()
					if err != nil {
						return 0, err
					}
					obj.UxArray[z1].Body.Address.Version = i
				}

				{
					// obj.UxArray[z1].Body.Address.Key
					if len(d.Buffer) < len(obj.UxArray[z1].Body.Address.Key) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.UxArray[z1].Body.Address.Key[:], d.Buffer[:len(obj.UxArray[z1].Body.Address.Key)])
					d.Buffer = d.Buffer[len(obj.UxArray[z1].Body.Address.Key):]
				}

				{
					// obj.UxArray[z1].Body.Coins
					i, err := d.Uint64()
					if err != nil {
						return 0, err
					}
					obj.UxArray[z1].Body.Coins = i
				}

				{
					// obj.UxArray[z1].Body.Hours
					i, err := d.Uint64()
					if err != nil {
						return 0, err
					}
					obj.UxArray[z1].Body.Hours = i
				}

			}
		}
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeUxArrayExact decodes an object of type UxArray from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeUxArrayExact(buf []byte, obj *UxArray) error {
	if n, err := decodeUxArray(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
