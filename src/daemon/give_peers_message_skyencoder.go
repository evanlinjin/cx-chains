// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package daemon

import (
	"errors"
	"math"

	"github.com/skycoin/skycoin/src/cipher/encoder"
)

// EncodeSizeGivePeersMessage computes the size of an encoded object of type GivePeersMessage
func EncodeSizeGivePeersMessage(obj *GivePeersMessage) int {
	i0 := 0

	// obj.Peers
	i0 += 4
	{
		i1 := 0

		// x.IP
		i1 += 4

		// x.Port
		i1 += 2

		i0 += len(obj.Peers) * i1
	}

	return i0
}

// EncodeGivePeersMessage encodes an object of type GivePeersMessage to the buffer in encoder.Encoder.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func EncodeGivePeersMessage(buf []byte, obj *GivePeersMessage) error {
	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Peers length check
	if len(obj.Peers) > math.MaxUint32 {
		return errors.New("obj.Peers length exceeds math.MaxUint32")
	}

	// obj.Peers length
	e.Uint32(uint32(len(obj.Peers)))

	// obj.Peers
	for _, x := range obj.Peers {

		// x.IP
		e.Uint32(x.IP)

		// x.Port
		e.Uint16(x.Port)

	}

	return nil
}

// DecodeGivePeersMessage decodes an object of type GivePeersMessage from the buffer in encoder.Decoder.
// Returns the number of bytes used from the buffer to decode the object.
func DecodeGivePeersMessage(buf []byte, obj *GivePeersMessage) (int, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Peers

		ul, err := d.Uint32()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}

		if length != 0 {
			obj.Peers = make([]IPAddr, length)

			for z1 := range obj.Peers {
				{
					// obj.Peers[z1].IP
					i, err := d.Uint32()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Peers[z1].IP = i
				}

				{
					// obj.Peers[z1].Port
					i, err := d.Uint16()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Peers[z1].Port = i
				}

			}
		}
	}

	return len(buf) - len(d.Buffer), nil
}
