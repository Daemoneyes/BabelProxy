package BabelProxy

import (
	"fmt"
)

type WechatPlatformProvider struct {
	ip                   string
	port                 string
	platformProviderList []PlatformProvider
	meta                 map[string]string
	cCProviderInstance   CCProvider
	bot                  Bot
	contactRecordList    []Contact
}
