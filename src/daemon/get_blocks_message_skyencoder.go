// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package daemon

import "github.com/skycoin/skycoin/src/cipher/encoder"

// EncodeSizeGetBlocksMessage computes the size of an encoded object of type GetBlocksMessage
func EncodeSizeGetBlocksMessage(obj *GetBlocksMessage) int {
	i0 := 0

	// obj.LastBlock
	i0 += 8

	// obj.RequestedBlocks
	i0 += 8

	return i0
}

// EncodeGetBlocksMessage encodes an object of type GetBlocksMessage to the buffer in encoder.Encoder.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func EncodeGetBlocksMessage(buf []byte, obj *GetBlocksMessage) error {
	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.LastBlock
	e.Uint64(obj.LastBlock)

	// obj.RequestedBlocks
	e.Uint64(obj.RequestedBlocks)

	return nil
}

// DecodeGetBlocksMessage decodes an object of type GetBlocksMessage from the buffer in encoder.Decoder.
// Returns the number of bytes used from the buffer to decode the object.
func DecodeGetBlocksMessage(buf []byte, obj *GetBlocksMessage) (int, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.LastBlock
		i, err := d.Uint64()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}
		obj.LastBlock = i
	}

	{
		// obj.RequestedBlocks
		i, err := d.Uint64()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}
		obj.RequestedBlocks = i
	}

	return len(buf) - len(d.Buffer), nil
}
