package PlatformProvider

import (
	"BabelProxy/DataShare"
	"BabelProxy/Protocol"
	"BabelProxy/Utils"
	"errors"
	_ "fmt"
	"github.com/spf13/viper"
	_ "io/ioutil"
	"launchpad.net/xmlpath"
	"net/http"
	"strconv"
	"time"
)

var SupportMsgType = []string{"text", "image", "voice", "video", "shortvideo", "location"}

type WechatPlatformProvider struct {
	name string
	meta map[string]string
}

func (wPP *WechatPlatformProvider) GetName() string {
	return wPP.name
}

func (wPP *WechatPlatformProvider) GetMeta() map[string]string {
	return wPP.meta
}

func (wPP *WechatPlatformProvider) ReConfigure(f string) (bool, error) {
	return true, nil
}

func (wPP *WechatPlatformProvider) SendMsg(msg Protocol.Message) (bool, error) {
	return true, nil
}

func (wPP *WechatPlatformProvider) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content, err := xmlpath.Parse(r.Body)
	if err != nil {
		Utils.Logger.Println("Cann't Parse Message from ", r.URL)
	} else {
		senderPath := xmlpath.MustCompile("/xml/FromUserName")
		createTimePath := xmlpath.MustCompile("/xml/CreateTime")
		sender, _ := senderPath.String(content)
		createTimeStampStr, _ := createTimePath.String(content)
		createTimeStamp, _ := strconv.Atoi(createTimeStampStr)
		createTime := time.Unix(int64(createTimeStamp), 0)
		msgTypePath := xmlpath.MustCompile("/xml/MsgType")
		switch msgType, _ := msgTypePath.String(content); msgType {
		case "text":
			Utils.Logger.Println("Get a Text From ", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/Content")
			msgBody, _ := msgBodyPath.String(content)
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "text", createTime)
			Utils.Logger.Println(newMsg.GetMsgBody())
			DataShare.MsgQ <- newMsg
		default:
			Utils.Logger.Println("default choice")
		}
	}
}

func CreateWechatPlatformProvider(f string) (*WechatPlatformProvider, error) {
	Utils.Logger.Println("Start Creating WeChatPlatformProvider")
	viper.SetConfigFile(f)
	err := viper.ReadInConfig()
	if err != nil {
		Utils.Logger.Println("Cannot Load Wechat Configure File at " + f)
		return &WechatPlatformProvider{}, errors.New("Load Configuration Failed Error")
	}
	var wPP = &WechatPlatformProvider{}
	wPP.name = viper.GetString("name")
	wPP.meta = make(map[string]string)
	wPP.meta["account"] = viper.GetString("account")
	wPP.meta["appId"] = viper.GetString("appId")
	wPP.meta["appsecret"] = viper.GetString("appsecret")
	wPP.meta["url"] = viper.GetString("url")
	Utils.Logger.Println("Finish Creating WeChatPlatformProvider")

	return wPP, nil
}
