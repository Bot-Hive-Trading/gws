package gws

// very simple handler builder to support event registration without creating separate struct for handling the events

type HandlerBuilder struct {
	onPongHandler    func(socket *Conn, message []byte)
	onPingHandler    func(socket *Conn, message []byte)
	onCloseHandler   func(socket *Conn, err error)
	onOpenHandler    func(socket *Conn)
	onMessageHandler func(socket *Conn, message *Message)
}

func NewHandler() *HandlerBuilder {
	return &HandlerBuilder{
		onPongHandler:    func(_ *Conn, _ []byte) {},
		onPingHandler:    func(_ *Conn, _ []byte) {},
		onCloseHandler:   func(_ *Conn, _ error) {},
		onOpenHandler:    func(_ *Conn) {},
		onMessageHandler: func(_ *Conn, _ *Message) {},
	}
}

// attach on pong handler
func (c *HandlerBuilder) OnPong(socket *Conn, message []byte) {
	c.onPongHandler(socket, message)
}
func (c *HandlerBuilder) SetOnPongHandler(h func(socket *Conn, message []byte)) {
	c.onPongHandler = h
}

// attach on ping handler
func (c *HandlerBuilder) OnPing(socket *Conn, message []byte) {
	c.onPingHandler(socket, message)
}
func (c *HandlerBuilder) SetOnPingHandler(h func(socket *Conn, message []byte)) {
	c.onPingHandler = h
}

// attach on close handler
func (c *HandlerBuilder) OnClose(socket *Conn, err error) {
	c.onCloseHandler(socket, err)
}
func (c *HandlerBuilder) SetOnCloseHandler(h func(socket *Conn, err error)) {
	c.onCloseHandler = h
}

// attach on open handler
func (c *HandlerBuilder) OnOpen(socket *Conn) {
	c.onOpenHandler(socket)
}
func (c *HandlerBuilder) SetOnOpenHandler(h func(socket *Conn)) {
	c.onOpenHandler = h
}

// attach on message handler
func (c *HandlerBuilder) OnMessage(socket *Conn, message *Message) {
	c.onMessageHandler(socket, message)
}
func (c *HandlerBuilder) SetOnMessageHandler(h func(socket *Conn, message *Message)) {
	c.onMessageHandler = h
}
