// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package daemon

import (
	"errors"
	"math"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"github.com/skycoin/skycoin/src/coin"
)

// EncodeSizeGiveBlocksMessage computes the size of an encoded object of type GiveBlocksMessage
func EncodeSizeGiveBlocksMessage(obj *GiveBlocksMessage) int {
	i0 := 0

	// obj.Blocks
	i0 += 4
	for _, x := range obj.Blocks {
		i1 := 0

		// x.Block.Head.Version
		i1 += 4

		// x.Block.Head.Time
		i1 += 8

		// x.Block.Head.BkSeq
		i1 += 8

		// x.Block.Head.Fee
		i1 += 8

		// x.Block.Head.PrevHash
		i1 += 32

		// x.Block.Head.BodyHash
		i1 += 32

		// x.Block.Head.UxHash
		i1 += 32

		// x.Block.Body.Transactions
		i1 += 4
		for _, x := range x.Block.Body.Transactions {
			i2 := 0

			// x.Length
			i2 += 4

			// x.Type
			i2++

			// x.InnerHash
			i2 += 32

			// x.Sigs
			i2 += 4
			{
				i3 := 0

				// x
				i3 += 65

				i2 += len(x.Sigs) * i3
			}

			// x.In
			i2 += 4
			{
				i3 := 0

				// x
				i3 += 32

				i2 += len(x.In) * i3
			}

			// x.Out
			i2 += 4
			{
				i3 := 0

				// x.Address.Version
				i3++

				// x.Address.Key
				i3 += 20

				// x.Coins
				i3 += 8

				// x.Hours
				i3 += 8

				i2 += len(x.Out) * i3
			}

			i1 += i2
		}

		// x.Sig
		i1 += 65

		i0 += i1
	}

	return i0
}

// EncodeGiveBlocksMessage encodes an object of type GiveBlocksMessage to the buffer in encoder.Encoder.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func EncodeGiveBlocksMessage(buf []byte, obj *GiveBlocksMessage) error {
	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Blocks maxlen check
	if len(obj.Blocks) > 128 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Blocks length check
	if len(obj.Blocks) > math.MaxUint32 {
		return errors.New("obj.Blocks length exceeds math.MaxUint32")
	}

	// obj.Blocks length
	e.Uint32(uint32(len(obj.Blocks)))

	// obj.Blocks
	for _, x := range obj.Blocks {

		// x.Block.Head.Version
		e.Uint32(x.Block.Head.Version)

		// x.Block.Head.Time
		e.Uint64(x.Block.Head.Time)

		// x.Block.Head.BkSeq
		e.Uint64(x.Block.Head.BkSeq)

		// x.Block.Head.Fee
		e.Uint64(x.Block.Head.Fee)

		// x.Block.Head.PrevHash
		e.CopyBytes(x.Block.Head.PrevHash[:])

		// x.Block.Head.BodyHash
		e.CopyBytes(x.Block.Head.BodyHash[:])

		// x.Block.Head.UxHash
		e.CopyBytes(x.Block.Head.UxHash[:])

		// x.Block.Body.Transactions maxlen check
		if len(x.Block.Body.Transactions) > 65535 {
			return encoder.ErrMaxLenExceeded
		}

		// x.Block.Body.Transactions length check
		if len(x.Block.Body.Transactions) > math.MaxUint32 {
			return errors.New("x.Block.Body.Transactions length exceeds math.MaxUint32")
		}

		// x.Block.Body.Transactions length
		e.Uint32(uint32(len(x.Block.Body.Transactions)))

		// x.Block.Body.Transactions
		for _, x := range x.Block.Body.Transactions {

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

		// x.Sig
		e.CopyBytes(x.Sig[:])

	}

	return nil
}

// DecodeGiveBlocksMessage decodes an object of type GiveBlocksMessage from the buffer in encoder.Decoder.
// Returns the number of bytes used from the buffer to decode the object.
func DecodeGiveBlocksMessage(buf []byte, obj *GiveBlocksMessage) (int, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Blocks

		ul, err := d.Uint32()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}

		if length > 128 {
			return len(buf) - len(d.Buffer), encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Blocks = make([]coin.SignedBlock, length)

			for z1 := range obj.Blocks {
				{
					// obj.Blocks[z1].Block.Head.Version
					i, err := d.Uint32()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Blocks[z1].Block.Head.Version = i
				}

				{
					// obj.Blocks[z1].Block.Head.Time
					i, err := d.Uint64()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Blocks[z1].Block.Head.Time = i
				}

				{
					// obj.Blocks[z1].Block.Head.BkSeq
					i, err := d.Uint64()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Blocks[z1].Block.Head.BkSeq = i
				}

				{
					// obj.Blocks[z1].Block.Head.Fee
					i, err := d.Uint64()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Blocks[z1].Block.Head.Fee = i
				}

				{
					// obj.Blocks[z1].Block.Head.PrevHash
					if len(d.Buffer) < len(obj.Blocks[z1].Block.Head.PrevHash) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}
					copy(obj.Blocks[z1].Block.Head.PrevHash[:], d.Buffer[:len(obj.Blocks[z1].Block.Head.PrevHash)])
					d.Buffer = d.Buffer[len(obj.Blocks[z1].Block.Head.PrevHash):]
				}

				{
					// obj.Blocks[z1].Block.Head.BodyHash
					if len(d.Buffer) < len(obj.Blocks[z1].Block.Head.BodyHash) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}
					copy(obj.Blocks[z1].Block.Head.BodyHash[:], d.Buffer[:len(obj.Blocks[z1].Block.Head.BodyHash)])
					d.Buffer = d.Buffer[len(obj.Blocks[z1].Block.Head.BodyHash):]
				}

				{
					// obj.Blocks[z1].Block.Head.UxHash
					if len(d.Buffer) < len(obj.Blocks[z1].Block.Head.UxHash) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}
					copy(obj.Blocks[z1].Block.Head.UxHash[:], d.Buffer[:len(obj.Blocks[z1].Block.Head.UxHash)])
					d.Buffer = d.Buffer[len(obj.Blocks[z1].Block.Head.UxHash):]
				}

				{
					// obj.Blocks[z1].Block.Body.Transactions

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
						obj.Blocks[z1].Block.Body.Transactions = make([]coin.Transaction, length)

						for z5 := range obj.Blocks[z1].Block.Body.Transactions {
							{
								// obj.Blocks[z1].Block.Body.Transactions[z5].Length
								i, err := d.Uint32()
								if err != nil {
									return len(buf) - len(d.Buffer), err
								}
								obj.Blocks[z1].Block.Body.Transactions[z5].Length = i
							}

							{
								// obj.Blocks[z1].Block.Body.Transactions[z5].Type
								i, err := d.Uint8()
								if err != nil {
									return len(buf) - len(d.Buffer), err
								}
								obj.Blocks[z1].Block.Body.Transactions[z5].Type = i
							}

							{
								// obj.Blocks[z1].Block.Body.Transactions[z5].InnerHash
								if len(d.Buffer) < len(obj.Blocks[z1].Block.Body.Transactions[z5].InnerHash) {
									return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
								}
								copy(obj.Blocks[z1].Block.Body.Transactions[z5].InnerHash[:], d.Buffer[:len(obj.Blocks[z1].Block.Body.Transactions[z5].InnerHash)])
								d.Buffer = d.Buffer[len(obj.Blocks[z1].Block.Body.Transactions[z5].InnerHash):]
							}

							{
								// obj.Blocks[z1].Block.Body.Transactions[z5].Sigs

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
									obj.Blocks[z1].Block.Body.Transactions[z5].Sigs = make([]cipher.Sig, length)

									for z7 := range obj.Blocks[z1].Block.Body.Transactions[z5].Sigs {
										{
											// obj.Blocks[z1].Block.Body.Transactions[z5].Sigs[z7]
											if len(d.Buffer) < len(obj.Blocks[z1].Block.Body.Transactions[z5].Sigs[z7]) {
												return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
											}
											copy(obj.Blocks[z1].Block.Body.Transactions[z5].Sigs[z7][:], d.Buffer[:len(obj.Blocks[z1].Block.Body.Transactions[z5].Sigs[z7])])
											d.Buffer = d.Buffer[len(obj.Blocks[z1].Block.Body.Transactions[z5].Sigs[z7]):]
										}

									}
								}
							}

							{
								// obj.Blocks[z1].Block.Body.Transactions[z5].In

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
									obj.Blocks[z1].Block.Body.Transactions[z5].In = make([]cipher.SHA256, length)

									for z7 := range obj.Blocks[z1].Block.Body.Transactions[z5].In {
										{
											// obj.Blocks[z1].Block.Body.Transactions[z5].In[z7]
											if len(d.Buffer) < len(obj.Blocks[z1].Block.Body.Transactions[z5].In[z7]) {
												return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
											}
											copy(obj.Blocks[z1].Block.Body.Transactions[z5].In[z7][:], d.Buffer[:len(obj.Blocks[z1].Block.Body.Transactions[z5].In[z7])])
											d.Buffer = d.Buffer[len(obj.Blocks[z1].Block.Body.Transactions[z5].In[z7]):]
										}

									}
								}
							}

							{
								// obj.Blocks[z1].Block.Body.Transactions[z5].Out

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
									obj.Blocks[z1].Block.Body.Transactions[z5].Out = make([]coin.TransactionOutput, length)

									for z7 := range obj.Blocks[z1].Block.Body.Transactions[z5].Out {
										{
											// obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Address.Version
											i, err := d.Uint8()
											if err != nil {
												return len(buf) - len(d.Buffer), err
											}
											obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Address.Version = i
										}

										{
											// obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Address.Key
											if len(d.Buffer) < len(obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Address.Key) {
												return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
											}
											copy(obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Address.Key[:], d.Buffer[:len(obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Address.Key)])
											d.Buffer = d.Buffer[len(obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Address.Key):]
										}

										{
											// obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Coins
											i, err := d.Uint64()
											if err != nil {
												return len(buf) - len(d.Buffer), err
											}
											obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Coins = i
										}

										{
											// obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Hours
											i, err := d.Uint64()
											if err != nil {
												return len(buf) - len(d.Buffer), err
											}
											obj.Blocks[z1].Block.Body.Transactions[z5].Out[z7].Hours = i
										}

									}
								}
							}
						}
					}
				}

				{
					// obj.Blocks[z1].Sig
					if len(d.Buffer) < len(obj.Blocks[z1].Sig) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}
					copy(obj.Blocks[z1].Sig[:], d.Buffer[:len(obj.Blocks[z1].Sig)])
					d.Buffer = d.Buffer[len(obj.Blocks[z1].Sig):]
				}

			}
		}
	}

	return len(buf) - len(d.Buffer), nil
}
