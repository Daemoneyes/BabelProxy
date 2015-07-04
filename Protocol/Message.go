package Protocol

import "time"

var MsgTypeList = []string{"audio", "video", "text", "image","conf"}

type Message struct {
	msgBody       string
	sender        string
	receiver      string
	msgType       string
	msgCreateTime time.Time
}

func (msg *Message) GetMsgBody() string {
	return msg.msgBody
}

func (msg *Message) GetSender() string {
	return msg.sender
}

func (msg *Message) GetMsgType() string {
	return msg.msgType
}

func (msg *Message) GetReceiver() string {
	return msg.receiver
}

func (msg *Message) GetMsgCreateTime() time.Time {
	return msg.msgCreateTime
}

func CreateMsg(msgBody, sender, receiver, msgType string, msgCreateTime time.Time) *Message {
	return &Message{msgBody, sender, receiver, msgType, msgCreateTime}
}
