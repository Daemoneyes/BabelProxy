package BabelProxy

import (
	"BabelProxy/Bots"
	"BabelProxy/CCProvider"
	"BabelProxy/PlatformProvider"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	_ "os"
)

type BabelProxy struct {
	ip, port                     string
	cCProviderInstance           CCProvider.CCProvider
	platformProviderInstanceList []PlatformProvider.PlatformProvider
	meta                         map[string]string
	bot                          Bots.Bot
	contactRecordList            []*Contact
}

func (bp *BabelProxy) reloadCOnfiguration() {

}

func (bp *BabelProxy) createContact() {

}

func (bp *BabelProxy) processMsg() {

}

func (bp *BabelProxy) findContact(callId string) *Contact {
	return nil
}

func (bp *BabelProxy) removeContact(callId string) bool {
	return true
}

func CreateProxy(f string) (*BabelProxy, error) {
	viper.SetConfigFile(f)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Cannot Load Configure File at " + f)
		return nil, errors.New("Load Configuration Failed Error")

	}
	fmt.Println(viper.GetString("test"))
	return nil, nil
}
