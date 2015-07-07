package Protocol

import "time"

var MsgTypeList = []string{"audio", "video", "text", "image","conf"}

type Message struct {
	msgBody       string
	sender        string
	receiver      string
	msgType       string
	msgCreateTime time.Time
	msgMeta       map[string]string
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


func (msg *Message) GetMsgMeta() map[string]string{
	return msg.msgMeta
}



func CreateMsg(msgBody, sender, receiver, msgType string, msgCreateTime time.Time,msgMeta map[string]string) *Message {
	return &Message{msgBody, sender, receiver, msgType, msgCreateTime,msgMeta}
}