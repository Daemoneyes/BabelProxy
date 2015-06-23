package PlatformProvider

type PlatformProvider interface {
	init() *PlatformProvider
	getName() string
	getMeta() map[string]string
	reConfigure() bool
	sendMsg(msg string) bool
	getMsg() string
	getPlatform() *PlatformProvider
}
