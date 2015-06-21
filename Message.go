package BabelProxy

type Message struct {
	msgBody string
	sender  string
}

func (msg *Message) getmsgBody() string {
	return msg.msgBody
}

func (msg *Message) getsender() string {
	return msg.sender
}

func (msg *Message) setmsgBody(msbBody string) {
	msg.msgBody = msbBody
}

func (msg *Message) setsender(sender string) {
	msg.sender = sender
}
