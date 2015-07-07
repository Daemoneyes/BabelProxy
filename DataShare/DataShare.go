package DataShare

import (
	"BabelProxy/Protocol"
)

var MsgQ = make(chan *Protocol.Message, 1000)
