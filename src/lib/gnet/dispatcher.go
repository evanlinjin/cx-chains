package gnet

import (
	"errors"
	"fmt"
	"github.com/skycoin/encoder"
	"log"
	//"net"
	"reflect"
	//"time"
)

//per pool factory for dispatchers
type DispatcherManager struct {
	Dispatchers map[uint16]*Dispatcher
}

//optional callback for message handling
func (self *DispatcherManager) OnMessage(c *Connection, channel uint16,
	msg []byte) error {

	d, ok := self.Dispatchers[channel]

	if ok == false {
		log.Panic("channel doesnt exist") //channel doesnt exist
		return errors.New("channel does not have dispatcher")
	}

	message, err := d.convertToMessage(c, msg)
	if err != nil {
		log.Panic("message conversion failed") //cleanup
		return err
	}

	context := MessageContext{Conn: c}
	err = message.Handle(&context, nil)

	if err != nil {
		log.Panic() //why can handle function return error?
		return err
	}

	return nil
}

func NewDispatcherManager() *DispatcherManager {
	var dm DispatcherManager
	dm.Dispatchers = make(map[uint16]*Dispatcher)
	return &dm
}

//routes messages to services
type Dispatcher struct {
	Channel             uint16 //channel the dispatcher handles
	Pool                *ConnectionPool
	MessageIdMap        map[reflect.Type]MessagePrefix
	MessageIdReverseMap map[MessagePrefix]reflect.Type
}

//dispatchers have channels in and channels out
func (self *ConnectionPool) NewDispatcher(channel uint16) *Dispatcher {
	var d Dispatcher
	d.Pool = self

	d.MessageIdMap = make(map[reflect.Type]MessagePrefix)
	d.MessageIdReverseMap = make(map[MessagePrefix]reflect.Type)
	return &d
}

// Serializes a Message over a net.Conn

//func sendMessage(conn net.Conn, msg Message, timeout time.Duration) error {
//	m := encodeMessage(msg)
//	return sendByteMessage(conn, m, timeout)
//}

// Event handler that is called after a Connection sends a complete message
func (self *Dispatcher) convertToMessage(c *Connection, msg []byte) (Message, error) {
	msgId := [4]byte{}
	if len(msg) < len(msgId) {
		return nil, errors.New("Not enough data to read msg id")
	}
	copy(msgId[:], msg[:len(msgId)])
	msg = msg[len(msgId):]

	t, succ := self.MessageIdReverseMap[msgId]
	if !succ {
		logger.Debug("Connection %d sent unknown message id %s",
			c.Id, string(msgId[:]))
		return nil, fmt.Errorf("Unknown message %s received", string(msgId[:]))
	}
	logger.Debug("Message type %v is handling it", t)

	var m Message
	var v reflect.Value = reflect.New(t)
	logger.Debug("Giving %d bytes to the decoder", len(msg))
	used, err := deserializeMessage(msg, v)
	if err != nil {
		return nil, err
	}
	if used != len(msg) {
		return nil, errors.New("Data buffer was not completely decoded")
	}

	m, succ = (v.Interface()).(Message)
	if !succ {
		// This occurs only when the user registers an interface that does
		// match the Message interface.  They should have known about this
		// earlier via a call to VerifyMessages
		log.Panic("Message obtained from map does not match Message interface")
		return nil, errors.New("MessageIdMaps contain non-Message")
	}
	return m, nil
}

// Wraps encoder.DeserializeRawToValue and traps panics as an error
func deserializeMessage(msg []byte, v reflect.Value) (n int, e error) {
	//deserializer panic should not be allowed
	//deserializer panic means error in serializer
	/*
	   defer func() {
	       if r := recover(); r != nil {
	           logger.Debug("Recovering from deserializer panic: %v", r)
	           switch x := r.(type) {
	           case string:
	               e = errors.New(x)
	           case error:
	               e = x
	           default:
	               e = errors.New("Message deserialization failed")
	           }
	       }
	   }()
	*/
	n, e = encoder.DeserializeRawToValue(msg, v)
	return
}

// Packgs a Message into []byte containing length, id and data

/*
var encodeMessage = func(msg Message) []byte {
	t := reflect.ValueOf(msg).Elem().Type()
	msgId, succ := MessageIdMap[t]
	if !succ {
		txt := "Attempted to serialize message struct not in MessageIdMap: %v"
		log.Panicf(txt, msg)
	}
	bMsg := encoder.Serialize(msg)

	// message length
	bLen := encoder.SerializeAtomic(uint32(len(bMsg) + len(msgId)))
	m := make([]byte, 0)
	m = append(m, bLen...)     // length prefix
	m = append(m, msgId[:]...) // message id
	m = append(m, bMsg...)     // message bytes
	return m
}
*/
