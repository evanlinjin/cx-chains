// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package historydb

import (
	"errors"
	"math"

	"github.com/amherag/skycoin/src/cipher"
	"github.com/amherag/skycoin/src/cipher/encoder"
	"github.com/amherag/skycoin/src/coin"
)

// encodeSizeTransaction computes the size of an encoded object of type Transaction
func encodeSizeTransaction(obj *Transaction) uint64 {
	i0 := uint64(0)

	// obj.Txn.Length
	i0 += 4

	// obj.Txn.Type
	i0++

	// obj.Txn.InnerHash
	i0 += 32

	// obj.Txn.Sigs
	i0 += 4
	{
		i1 := uint64(0)

		// x
		i1 += 65

		i0 += uint64(len(obj.Txn.Sigs)) * i1
	}

	// obj.Txn.In
	i0 += 4
	{
		i1 := uint64(0)

		// x
		i1 += 32

		i0 += uint64(len(obj.Txn.In)) * i1
	}

	// obj.Txn.Out
	i0 += 4
	{
		i1 := uint64(0)

		// x.Address.Version
		i1++

		// x.Address.Key
		i1 += 20

		// x.Coins
		i1 += 8

		// x.Hours
		i1 += 8

		i0 += uint64(len(obj.Txn.Out)) * i1
	}

	// obj.BlockSeq
	i0 += 8

	return i0
}

// encodeTransaction encodes an object of type Transaction to a buffer allocated to the exact size
// required to encode the object.
func encodeTransaction(obj *Transaction) ([]byte, error) {
	n := encodeSizeTransaction(obj)
	buf := make([]byte, n)

	if err := encodeTransactionToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeTransactionToBuffer encodes an object of type Transaction to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeTransactionToBuffer(buf []byte, obj *Transaction) error {
	if uint64(len(buf)) < encodeSizeTransaction(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Txn.Length
	e.Uint32(obj.Txn.Length)

	// obj.Txn.Type
	e.Uint8(obj.Txn.Type)

	// obj.Txn.InnerHash
	e.CopyBytes(obj.Txn.InnerHash[:])

	// obj.Txn.Sigs maxlen check
	if len(obj.Txn.Sigs) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Txn.Sigs length check
	if uint64(len(obj.Txn.Sigs)) > math.MaxUint32 {
		return errors.New("obj.Txn.Sigs length exceeds math.MaxUint32")
	}

	// obj.Txn.Sigs length
	e.Uint32(uint32(len(obj.Txn.Sigs)))

	// obj.Txn.Sigs
	for _, x := range obj.Txn.Sigs {

		// x
		e.CopyBytes(x[:])

	}

	// obj.Txn.In maxlen check
	if len(obj.Txn.In) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Txn.In length check
	if uint64(len(obj.Txn.In)) > math.MaxUint32 {
		return errors.New("obj.Txn.In length exceeds math.MaxUint32")
	}

	// obj.Txn.In length
	e.Uint32(uint32(len(obj.Txn.In)))

	// obj.Txn.In
	for _, x := range obj.Txn.In {

		// x
		e.CopyBytes(x[:])

	}

	// obj.Txn.Out maxlen check
	if len(obj.Txn.Out) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Txn.Out length check
	if uint64(len(obj.Txn.Out)) > math.MaxUint32 {
		return errors.New("obj.Txn.Out length exceeds math.MaxUint32")
	}

	// obj.Txn.Out length
	e.Uint32(uint32(len(obj.Txn.Out)))

	// obj.Txn.Out
	for _, x := range obj.Txn.Out {

		// x.Address.Version
		e.Uint8(x.Address.Version)

		// x.Address.Key
		e.CopyBytes(x.Address.Key[:])

		// x.Coins
		e.Uint64(x.Coins)

		// x.Hours
		e.Uint64(x.Hours)

	}

	// obj.BlockSeq
	e.Uint64(obj.BlockSeq)

	return nil
}

// decodeTransaction decodes an object of type Transaction from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeTransaction(buf []byte, obj *Transaction) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Txn.Length
		i, err := d.Uint32()
		if err != nil {
			return 0, err
		}
		obj.Txn.Length = i
	}

	{
		// obj.Txn.Type
		i, err := d.Uint8()
		if err != nil {
			return 0, err
		}
		obj.Txn.Type = i
	}

	{
		// obj.Txn.InnerHash
		if len(d.Buffer) < len(obj.Txn.InnerHash) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.Txn.InnerHash[:], d.Buffer[:len(obj.Txn.InnerHash)])
		d.Buffer = d.Buffer[len(obj.Txn.InnerHash):]
	}

	{
		// obj.Txn.Sigs

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Txn.Sigs = make([]cipher.Sig, length)

			for z2 := range obj.Txn.Sigs {
				{
					// obj.Txn.Sigs[z2]
					if len(d.Buffer) < len(obj.Txn.Sigs[z2]) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Txn.Sigs[z2][:], d.Buffer[:len(obj.Txn.Sigs[z2])])
					d.Buffer = d.Buffer[len(obj.Txn.Sigs[z2]):]
				}

			}
		}
	}

	{
		// obj.Txn.In

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Txn.In = make([]cipher.SHA256, length)

			for z2 := range obj.Txn.In {
				{
					// obj.Txn.In[z2]
					if len(d.Buffer) < len(obj.Txn.In[z2]) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Txn.In[z2][:], d.Buffer[:len(obj.Txn.In[z2])])
					d.Buffer = d.Buffer[len(obj.Txn.In[z2]):]
				}

			}
		}
	}

	{
		// obj.Txn.Out

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Txn.Out = make([]coin.TransactionOutput, length)

			for z2 := range obj.Txn.Out {
				{
					// obj.Txn.Out[z2].Address.Version
					i, err := d.Uint8()
					if err != nil {
						return 0, err
					}
					obj.Txn.Out[z2].Address.Version = i
				}

				{
					// obj.Txn.Out[z2].Address.Key
					if len(d.Buffer) < len(obj.Txn.Out[z2].Address.Key) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Txn.Out[z2].Address.Key[:], d.Buffer[:len(obj.Txn.Out[z2].Address.Key)])
					d.Buffer = d.Buffer[len(obj.Txn.Out[z2].Address.Key):]
				}

				{
					// obj.Txn.Out[z2].Coins
					i, err := d.Uint64()
					if err != nil {
						return 0, err
					}
					obj.Txn.Out[z2].Coins = i
				}

				{
					// obj.Txn.Out[z2].Hours
					i, err := d.Uint64()
					if err != nil {
						return 0, err
					}
					obj.Txn.Out[z2].Hours = i
				}

			}
		}
	}

	{
		// obj.BlockSeq
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.BlockSeq = i
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeTransactionExact decodes an object of type Transaction from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeTransactionExact(buf []byte, obj *Transaction) error {
	if n, err := decodeTransaction(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
