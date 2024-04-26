package gws

// very simple handler builder to support event registration without creating separate struct for handling the events

type Handler struct {
	onPongHandler    func(socket *Conn, message []byte)
	onPingHandler    func(socket *Conn, message []byte)
	onCloseHandler   func(socket *Conn, err error)
	onOpenHandler    func(socket *Conn)
	onMessageHandler func(socket *Conn, message *Message)
}

func NewHandler() *Handler {
	return &Handler{
		onPongHandler:    func(_ *Conn, _ []byte) {},
		onPingHandler:    func(_ *Conn, _ []byte) {},
		onCloseHandler:   func(_ *Conn, _ error) {},
		onOpenHandler:    func(_ *Conn) {},
		onMessageHandler: func(_ *Conn, _ *Message) {},
	}
}

// attach on pong handler
func (c *Handler) OnPong(socket *Conn, message []byte) {
	c.onPongHandler(socket, message)
}
func (c *Handler) SetOnPongHandler(h func(socket *Conn, message []byte)) {
	c.onPongHandler = h
}

// attach on ping handler
func (c *Handler) OnPing(socket *Conn, message []byte) {
	c.onPingHandler(socket, message)
}
func (c *Handler) SetOnPingHandler(h func(socket *Conn, message []byte)) {
	c.onPingHandler = h
}

// attach on close handler
func (c *Handler) OnClose(socket *Conn, err error) {
	c.onCloseHandler(socket, err)
}
func (c *Handler) SetOnCloseHandler(h func(socket *Conn, err error)) {
	c.onCloseHandler = h
}

// attach on open handler
func (c *Handler) OnOpen(socket *Conn) {
	c.onOpenHandler(socket)
}
func (c *Handler) SetOnOpenHandler(h func(socket *Conn)) {
	c.onOpenHandler = h
}

// attach on message handler
func (c *Handler) OnMessage(socket *Conn, message *Message) {
	c.onMessageHandler(socket, message)
}
func (c *Handler) SetOnMessageHandler(h func(socket *Conn, message *Message)) {
	c.onMessageHandler = h
}
