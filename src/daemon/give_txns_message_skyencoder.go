// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package daemon

import (
	"errors"
	"math"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"github.com/skycoin/skycoin/src/coin"
)

// EncodeSizeGiveTxnsMessage computes the size of an encoded object of type GiveTxnsMessage
func EncodeSizeGiveTxnsMessage(obj *GiveTxnsMessage) int {
	i0 := 0

	// obj.Transactions
	i0 += 4
	for _, x := range obj.Transactions {
		i1 := 0

		// x.Length
		i1 += 4

		// x.Type
		i1++

		// x.InnerHash
		i1 += 32

		// x.Sigs
		i1 += 4
		{
			i2 := 0

			// x
			i2 += 65

			i1 += len(x.Sigs) * i2
		}

		// x.In
		i1 += 4
		{
			i2 := 0

			// x
			i2 += 32

			i1 += len(x.In) * i2
		}

		// x.Out
		i1 += 4
		{
			i2 := 0

			// x.Address.Version
			i2++

			// x.Address.Key
			i2 += 20

			// x.Coins
			i2 += 8

			// x.Hours
			i2 += 8

			i1 += len(x.Out) * i2
		}

		i0 += i1
	}

	return i0
}

// EncodeGiveTxnsMessage encodes an object of type GiveTxnsMessage to the buffer in encoder.Encoder.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func EncodeGiveTxnsMessage(buf []byte, obj *GiveTxnsMessage) error {
	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Transactions maxlen check
	if len(obj.Transactions) > 256 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Transactions length check
	if len(obj.Transactions) > math.MaxUint32 {
		return errors.New("obj.Transactions length exceeds math.MaxUint32")
	}

	// obj.Transactions length
	e.Uint32(uint32(len(obj.Transactions)))

	// obj.Transactions
	for _, x := range obj.Transactions {

		// x.Length
		e.Uint32(x.Length)

		// x.Type
		e.Uint8(x.Type)

		// x.InnerHash
		e.CopyBytes(x.InnerHash[:])

		// x.Sigs maxlen check
		if len(x.Sigs) > 65535 {
			return encoder.ErrMaxLenExceeded
		}

		// x.Sigs length check
		if len(x.Sigs) > math.MaxUint32 {
			return errors.New("x.Sigs length exceeds math.MaxUint32")
		}

		// x.Sigs length
		e.Uint32(uint32(len(x.Sigs)))

		// x.Sigs
		for _, x := range x.Sigs {

			// x
			e.CopyBytes(x[:])

		}

		// x.In maxlen check
		if len(x.In) > 65535 {
			return encoder.ErrMaxLenExceeded
		}

		// x.In length check
		if len(x.In) > math.MaxUint32 {
			return errors.New("x.In length exceeds math.MaxUint32")
		}

		// x.In length
		e.Uint32(uint32(len(x.In)))

		// x.In
		for _, x := range x.In {

			// x
			e.CopyBytes(x[:])

		}

		// x.Out maxlen check
		if len(x.Out) > 65535 {
			return encoder.ErrMaxLenExceeded
		}

		// x.Out length check
		if len(x.Out) > math.MaxUint32 {
			return errors.New("x.Out length exceeds math.MaxUint32")
		}

		// x.Out length
		e.Uint32(uint32(len(x.Out)))

		// x.Out
		for _, x := range x.Out {

			// x.Address.Version
			e.Uint8(x.Address.Version)

			// x.Address.Key
			e.CopyBytes(x.Address.Key[:])

			// x.Coins
			e.Uint64(x.Coins)

			// x.Hours
			e.Uint64(x.Hours)

		}

	}

	return nil
}

// DecodeGiveTxnsMessage decodes an object of type GiveTxnsMessage from the buffer in encoder.Decoder.
// Returns the number of bytes used from the buffer to decode the object.
func DecodeGiveTxnsMessage(buf []byte, obj *GiveTxnsMessage) (int, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Transactions

		ul, err := d.Uint32()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}

		if length > 256 {
			return len(buf) - len(d.Buffer), encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Transactions = make([]coin.Transaction, length)

			for z1 := range obj.Transactions {
				{
					// obj.Transactions[z1].Length
					i, err := d.Uint32()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Transactions[z1].Length = i
				}

				{
					// obj.Transactions[z1].Type
					i, err := d.Uint8()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Transactions[z1].Type = i
				}

				{
					// obj.Transactions[z1].InnerHash
					if len(d.Buffer) < len(obj.Transactions[z1].InnerHash) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}
					copy(obj.Transactions[z1].InnerHash[:], d.Buffer[:len(obj.Transactions[z1].InnerHash)])
					d.Buffer = d.Buffer[len(obj.Transactions[z1].InnerHash):]
				}

				{
					// obj.Transactions[z1].Sigs

					ul, err := d.Uint32()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}

					length := int(ul)
					if length < 0 || length > len(d.Buffer) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}

					if length > 65535 {
						return len(buf) - len(d.Buffer), encoder.ErrMaxLenExceeded
					}

					if length != 0 {
						obj.Transactions[z1].Sigs = make([]cipher.Sig, length)

						for z3 := range obj.Transactions[z1].Sigs {
							{
								// obj.Transactions[z1].Sigs[z3]
								if len(d.Buffer) < len(obj.Transactions[z1].Sigs[z3]) {
									return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
								}
								copy(obj.Transactions[z1].Sigs[z3][:], d.Buffer[:len(obj.Transactions[z1].Sigs[z3])])
								d.Buffer = d.Buffer[len(obj.Transactions[z1].Sigs[z3]):]
							}

						}
					}
				}

				{
					// obj.Transactions[z1].In

					ul, err := d.Uint32()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}

					length := int(ul)
					if length < 0 || length > len(d.Buffer) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}

					if length > 65535 {
						return len(buf) - len(d.Buffer), encoder.ErrMaxLenExceeded
					}

					if length != 0 {
						obj.Transactions[z1].In = make([]cipher.SHA256, length)

						for z3 := range obj.Transactions[z1].In {
							{
								// obj.Transactions[z1].In[z3]
								if len(d.Buffer) < len(obj.Transactions[z1].In[z3]) {
									return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
								}
								copy(obj.Transactions[z1].In[z3][:], d.Buffer[:len(obj.Transactions[z1].In[z3])])
								d.Buffer = d.Buffer[len(obj.Transactions[z1].In[z3]):]
							}

						}
					}
				}

				{
					// obj.Transactions[z1].Out

					ul, err := d.Uint32()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}

					length := int(ul)
					if length < 0 || length > len(d.Buffer) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}

					if length > 65535 {
						return len(buf) - len(d.Buffer), encoder.ErrMaxLenExceeded
					}

					if length != 0 {
						obj.Transactions[z1].Out = make([]coin.TransactionOutput, length)

						for z3 := range obj.Transactions[z1].Out {
							{
								// obj.Transactions[z1].Out[z3].Address.Version
								i, err := d.Uint8()
								if err != nil {
									return len(buf) - len(d.Buffer), err
								}
								obj.Transactions[z1].Out[z3].Address.Version = i
							}

							{
								// obj.Transactions[z1].Out[z3].Address.Key
								if len(d.Buffer) < len(obj.Transactions[z1].Out[z3].Address.Key) {
									return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
								}
								copy(obj.Transactions[z1].Out[z3].Address.Key[:], d.Buffer[:len(obj.Transactions[z1].Out[z3].Address.Key)])
								d.Buffer = d.Buffer[len(obj.Transactions[z1].Out[z3].Address.Key):]
							}

							{
								// obj.Transactions[z1].Out[z3].Coins
								i, err := d.Uint64()
								if err != nil {
									return len(buf) - len(d.Buffer), err
								}
								obj.Transactions[z1].Out[z3].Coins = i
							}

							{
								// obj.Transactions[z1].Out[z3].Hours
								i, err := d.Uint64()
								if err != nil {
									return len(buf) - len(d.Buffer), err
								}
								obj.Transactions[z1].Out[z3].Hours = i
							}

						}
					}
				}
			}
		}
	}

	return len(buf) - len(d.Buffer), nil
}
