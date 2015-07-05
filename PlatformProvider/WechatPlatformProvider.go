package PlatformProvider

import (
	"BabelProxy/DataShare"
	"BabelProxy/Protocol"
	"BabelProxy/Utils"
	"errors"
	_ "fmt"
	json "github.com/bitly/go-simplejson"
	_ "github.com/robfig/cron"
	"github.com/spf13/viper"
	_ "io/ioutil"
	"launchpad.net/xmlpath"
	"net/http"
	"strconv"
	"time"
)

type basicResponse{
	
}








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

func (wPP *WechatPlatformProvider) SendMsg(msg Protocol.Message) bool {
	switch msg.GetMsgType() {
	case "text":
		
	case "image":
	case "voice":
	case "video":
	case "news":
	case "resp":
	default:
		return true

	}
}

func (wPP *WechatPlatformProvider) UpdateToken() {
	for {
		Utils.Logger.Println("Update Access Token For WeChat")
		url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"
		updateUrl := url + "&appid=" + wPP.GetMeta()["appId"] + "&secret=" + wPP.GetMeta()["appsecret"]
		Utils.Logger.Println(updateUrl)
		resp, err := http.Get(updateUrl)
		if err == nil {
			js, _ := json.NewFromReader(resp.Body)
			value, flag := js.CheckGet("access_token")
			if flag {
				wPP.meta["access_token"] = value.MustString()
				Utils.Logger.Println("Update Access Token to ", wPP.meta["access_token"])
				time.Sleep(3600 * time.Second)
			}
		} else {
			Utils.Logger.Println("Get Token Failed", err)
			time.Sleep(3600 * time.Second)

		}
	}

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
		case "image":
			Utils.Logger.Println("Get a Image From ", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/PicUrl")
			msgBody, _ := msgBodyPath.String(content)
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "image", createTime)
			Utils.Logger.Println(newMsg.GetMsgBody())
			DataShare.MsgQ <- newMsg
		case "voice":
			Utils.Logger.Println("Get a Voice Msg From", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/MediaId")
			msgBody, _ := msgBodyPath.String(content)
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "voice", createTime)
			Utils.Logger.Println(newMsg.GetMsgBody())
			DataShare.MsgQ <- newMsg
		case "video":
			Utils.Logger.Println("Get a Video Msg From", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/MediaId")
			msgBody, _ := msgBodyPath.String(content)
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "voice", createTime)
			Utils.Logger.Println(newMsg.GetMsgBody())
			DataShare.MsgQ <- newMsg
		case "shortvideo":
			Utils.Logger.Println("Get a Short Video Msg From", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/MediaId")
			msgBody, _ := msgBodyPath.String(content)
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "voice", createTime)
			Utils.Logger.Println(newMsg.GetMsgBody())
			DataShare.MsgQ <- newMsg
		case "location":
			Utils.Logger.Println("Get a Location Msg From", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/Label")
			msgBody, _ := msgBodyPath.String(content)
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "Location", createTime)
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
	//	Utils.Logger.Println("Finish Creating WeChatPlatformProvider")
	//	c := cron.New()
	//	c.AddFunc("@every 1h30m", func() { wPP.UpdateToken() })
	//	c.Start()
	go wPP.UpdateToken()
	return wPP, nil
}
