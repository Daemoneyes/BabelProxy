package BabelProxy

import (
	"BabelProxy/Bots"
	"BabelProxy/CCProvider"
	"BabelProxy/PlatformProvider"
	"BabelProxy/Protocol"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	_ "log"
	"net/http"
	_ "os"
	_ "path/filepath"
)

type BabelProxy struct {
	ip, port                     string
	cCProviderInstance           CCProvider.CCProvider
	platformProviderInstanceList []PlatformProvider.PlatformProvider
	meta                         map[string]string
	bot                          Bots.Bot
	contactRecordList            []*Protocol.Contact
}

func (bp *BabelProxy) reloadCOnfiguration() {

}

func (bp *BabelProxy) createContact() {

}

func (bp *BabelProxy) processMsg() {

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

func createProxy(f string) (*BabelProxy, error) {
	viper.SetConfigFile(f)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Cannot Load Configure File at " + f)
		return nil, errors.New("Load Configuration Failed Error")
	}
	var bp = &BabelProxy{}
	bp.ip = viper.GetString("ip")
	bp.port = viper.GetString("port")
	bp.platformProviderInstanceList = make([]PlatformProvider.PlatformProvider, 10)
	bp.platformProviderInstanceList = append(bp.platformProviderInstanceList, PlatformProvider.WechatPlatformProviderInstance)
	return bp, nil
}

var Bp, err = createProxy("./proxy.json")

func (bp *BabelProxy) Run() {
	for _, pp := range bp.platformProviderInstanceList {
		url, _ := pp.GetMeta()["url"]
		http.HandleFunc(url, pp.GetMsg)
	}
	http.ListenAndServe(":10000", nil)
}
