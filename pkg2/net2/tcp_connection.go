package net2

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)

var (
	ConnectionClosedError  = errors.New("connection Closed")
	ConnectionBlockedError = errors.New("connection Blocked")

	globalConnectionId  uint64
	globalConnectionId2 uint64
)

type TcpConnection struct {
	name          string
	conn          net.Conn
	id            uint64
	codec         Codec
	sendChan      chan interface{}
	recvMutex     sync.Mutex
	sendMutex     sync.RWMutex
	closeFlag     int32
	closeChan     chan int
	closeMutex    sync.Mutex
	closeCallback closeCallback
	Context       interface{}
	debugString   string
	isWebsocket   bool
}

func NewTcpConnection(name string, conn net.Conn, sendChanSize int, codec Codec, cb closeCallback) *TcpConnection {
	// TODO(@benqi): globalConnectionId use
	if globalConnectionId >= (1 << 60) {
		atomic.StoreUint64(&globalConnectionId, 0)
	}

	conn2 := &TcpConnection{
		name:          name,
		conn:          conn,
		codec:         codec,
		closeChan:     make(chan int),
		id:            atomic.AddUint64(&globalConnectionId, 1),
		closeCallback: cb,
		isWebsocket:   false,
	}

	if sendChanSize > 0 {
		conn2.sendChan = make(chan interface{}, sendChanSize)
		go conn2.sendLoop()
	}
	return conn2
}

func NewTcpConnection2(name string, conn net.Conn, sendChanSize int, codec Codec, isWebsocket bool, cb closeCallback) *TcpConnection {
	conn2 := &TcpConnection{
		name:          name,
		conn:          conn,
		codec:         codec,
		closeChan:     make(chan int),
		id:            atomic.AddUint64(&globalConnectionId2, 1),
		closeCallback: cb,
		isWebsocket:   isWebsocket,
	}

	if sendChanSize > 0 {
		conn2.sendChan = make(chan interface{}, sendChanSize)
		go conn2.sendLoop()
	}
	return conn2
}

func (c *TcpConnection) String() string {
	if c.debugString == "" {
		c.debugString = fmt.Sprintf(`{"conn_id":"%d@(%s->%s)"}`, c.id, c.conn.RemoteAddr(), c.conn.LocalAddr())
	}
	return c.debugString
}

func (c *TcpConnection) LoadAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *TcpConnection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *TcpConnection) Name() string {
	return c.name
}

func (c *TcpConnection) GetConnID() uint64 {
	return c.id
}

func (c *TcpConnection) GetNetConn() net.Conn {
	return c.conn
}

func (c *TcpConnection) IsWebsocket() bool {
	return c.isWebsocket
}

func (c *TcpConnection) IsClosed() bool {
	return atomic.LoadInt32(&c.closeFlag) == 1
}

func (c *TcpConnection) Close() error {
	if atomic.CompareAndSwapInt32(&c.closeFlag, 0, 1) {
		if c.closeCallback != nil {
			c.closeCallback.OnConnectionClosed(c)
		}

		close(c.closeChan)

		if c.sendChan != nil {
			c.sendMutex.Lock()
			close(c.sendChan)
			if clear, ok := c.codec.(ClearSendChan); ok {
				clear.ClearSendChan(c.sendChan)
			}
			c.sendMutex.Unlock()
		}

		err := c.codec.Close()
		return err
	}
	return ConnectionClosedError
}

func (c *TcpConnection) Codec() Codec {
	return c.codec
}

func (c *TcpConnection) Receive() (interface{}, error) {
	c.recvMutex.Lock()
	defer c.recvMutex.Unlock()

	msg, err := c.codec.Receive()
	if err != nil {
		c.Close()
	}
	return msg, err
}

func (c *TcpConnection) sendLoop() {
	defer c.Close()
	for {
		select {
		case msg, ok := <-c.sendChan:
			if !ok || c.codec.Send(msg) != nil {
				return
			}
		case <-c.closeChan:
			return
		}
	}
}

func (c *TcpConnection) Send(msg interface{}) error {
	if c.sendChan == nil {
		if c.IsClosed() {
			return ConnectionClosedError
		}

		c.sendMutex.Lock()
		defer c.sendMutex.Unlock()

		err := c.codec.Send(msg)
		if err != nil {
			c.Close()
		}
		return err
	}

	c.sendMutex.RLock()
	if c.IsClosed() {
		c.sendMutex.RUnlock()
		return ConnectionClosedError
	}

	select {
	case c.sendChan <- msg:
		c.sendMutex.RUnlock()
		return nil
	default:
		c.sendMutex.RUnlock()
		c.Close()
		return ConnectionBlockedError
	}
}
