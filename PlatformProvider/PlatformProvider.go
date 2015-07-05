package PlatformProvider

import (
	"BabelProxy/Protocol"
	"net/http"
)

type PlatformProvider interface {
	GetName() string
	GetMeta() map[string]string
	ReConfigure(f string) (bool, error)
	SendMsg(msg Protocol.Message) bool
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
