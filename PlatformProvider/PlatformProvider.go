package PlatformProvider

import "BabelProxy/Protocol"



type PlatformProvider interface {
	GetName() string
	GetMeta() map[string]string
	ReConfigure(f string) (bool,error)
	SendMsg(msg Protocol.Message) (bool,error)
	GetMsg() string
	GetPlatform() *PlatformProvider
}




