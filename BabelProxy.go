package BabelProxy

import (
	"BabelProxy/Bots"
	"BabelProxy/CCProvider"
	"BabelProxy/DataShare"
	"BabelProxy/PlatformProvider"
	"BabelProxy/Protocol"
	"BabelProxy/Utils"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	_ "log"
	"net/http"
	_ "os"
	_ "path/filepath"
)

type BabelProxy struct {
	ip, port           string
	cCProviderInstance CCProvider.CCProvider
	// Interface is a pointer , so dont need to use interface pointer to get struct pointer
	platformProviderInstanceList map[string]PlatformProvider.PlatformProvider
	meta                         map[string]string
	bot                          Bots.Bot
	contactRecordList            []Protocol.Contact
}

func (bp *BabelProxy) reloadCOnfiguration() {

}

func (bp *BabelProxy) createContact() {

}

//automachine to get msg from queue, pack it as Contact and send to different Provider
func (bp *BabelProxy) processMsg() {
	var msg *Protocol.Message
	for {
		msg = <-DataShare.MsgQ
		Utils.Logger.Println("Get Msg From Queue, type ", msg.GetMsgType())
		if msg.GetReceiver() == "wechat" {
			Utils.Logger.Println("Call Sent to Wechat Provider to Handle")
			bp.platformProviderInstanceList["wechat"].SendMsg(msg)

		}

	}
}

func (bp *BabelProxy) findContact(callId string) *Protocol.Contact {
	return nil
}

func (bp *BabelProxy) removeContact(callId string) bool {
	return true
}

func (bp *BabelProxy) GetIp() string {
	return bp.ip
}

func (bp *BabelProxy) Run() {
	fmt.Println("Server Start Running, Check Log to get More detail")
	Utils.Logger.Println("Start Listening at Port 10000")
	for _, pp := range bp.platformProviderInstanceList {
		url, _ := pp.GetMeta()["url"]
		http.Handle(url, pp)
	}
	err := http.ListenAndServe("0.0.0.0:10000", nil)
	Utils.Logger.Println(err)
}

func CreateProxy(f string) (*BabelProxy, error) {
	Utils.Logger.Println("Start Creating Proxy....")
	viper.SetConfigFile(f)
	err := viper.ReadInConfig()
	if err != nil {
		Utils.Logger.Println("Cannot Load Configure File at " + f)
		return nil, errors.New("Load Configuration Failed Error")
	}
	var bp = &BabelProxy{}
	bp.ip = viper.GetString("ip")
	bp.port = viper.GetString("port")
	bp.platformProviderInstanceList = make(map[string]PlatformProvider.PlatformProvider)
	wechatPlatformProvider, err := PlatformProvider.CreateWechatPlatformProvider("./wechat.json")
	bp.platformProviderInstanceList[wechatPlatformProvider.GetName()] = wechatPlatformProvider
	Utils.Logger.Println("Creating Proxy Finished")
	go bp.processMsg()
	Utils.Logger.Println("Processing Msg Start")
	return bp, nil
}
