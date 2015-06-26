package PlatformProvider


import (
	_ "fmt"
	_ "net/http"
	_ "github.com/spf13/viper"
	"BabelProxy/Protocol"
)

type WechatPlatformProvider struct {
	name string
	meta map[string]string
	
}



func (wPP *WechatPlatformProvider)  GetName() string{
	return wPP.name
}

func (wPP *WechatPlatformProvider) GetMeta() map[string]string {
	return wPP.meta
}


func (wPP *WechatPlatformProvider) ReConfigure (f string) (bool,error){
	return true,nil
}


func (wPP *WechatPlatformProvider) SendMsg(msg Protocol.Message) (bool,error){
	return true,nil
}