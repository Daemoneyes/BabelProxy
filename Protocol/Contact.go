package Protocol

import (
	"errors"
	"time"
)

type Contact struct {
	customerId   string
	customerName string
	agentId      string
	agentName    string
	msgList      []Message
	startTime    time.Time
	callId       string
}

func (contact *Contact) getMsg(index int) (msg *Message, err error) {
	if index < len(contact.msgList) {
		return &contact.msgList[index], nil
	} else {
		return nil, errors.New("Index Out of Range")
	}
}

func (contact *Contact) getMsgList() []Message {
	return contact.msgList
}

func (contact *Contact) getcustomerId() string {
	return contact.customerId
}

func (contact *Contact) getagentId() string {
	return contact.agentId
}

func (contact *Contact) getcustomerName() string {
	return contact.customerName
}

func (contact *Contact) getagentName() string {
	return contact.agentName
}

func (contact *Contact) getcallId() string {
	return contact.callId
}

func (contact *Contact) getstartTime() time.Time {
	return contact.startTime
}

func createContact(customerId, customerName, agentId, agentName, callId string) *Contact {
	startTime := time.Now()
	msgList := make([]Message, 10)
	contact := Contact{customerId, customerName, agentId, agentName, msgList, startTime, callId}
	return &contact
}
