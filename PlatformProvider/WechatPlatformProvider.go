package PlatformProvider

import (
	"BabelProxy/DataShare"
	"BabelProxy/Protocol"
	"BabelProxy/Utils"
	ejson "encoding/json"
	"encoding/xml"
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

type TextResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
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

func (wPP *WechatPlatformProvider) SendMsg(msg *Protocol.Message) bool {
	switch msg.GetMsgType() {
	case "text":
		tR := &TextResponse{ToUserName: msg.GetSender(), FromUserName: wPP.GetMeta()["account"], CreateTime: time.Now().Unix(), MsgType: "text", Content: msg.GetMsgBody()}
		Utils.Logger.Println(tR)
		xmlstring, _ := xml.Marshal(tR)
		jsonString, _ := ejson.Marshal(tR)
		Utils.Logger.Println(string(xmlstring[:]))
		Utils.Logger.Println(string(jsonString[:]))
		return true
	case "image":
		tR := &TextResponse{ToUserName: msg.GetSender(), FromUserName: wPP.GetMeta()["account"], CreateTime: time.Now().Unix(), MsgType: "text", Content: msg.GetMsgBody()}
		xmlstring, _ := xml.Marshal(tR)
		Utils.Logger.Println(xmlstring)
		return true
	case "voice":
		tR := &TextResponse{ToUserName: msg.GetSender(), FromUserName: wPP.GetMeta()["account"], CreateTime: time.Now().Unix(), MsgType: "text", Content: msg.GetMsgBody()}
		xmlstring, _ := xml.Marshal(tR)
		Utils.Logger.Println(xmlstring)
		return true
	case "video":
		tR := &TextResponse{ToUserName: msg.GetSender(), FromUserName: wPP.GetMeta()["account"], CreateTime: time.Now().Unix(), MsgType: "text", Content: msg.GetMsgBody()}
		xmlstring, _ := xml.Marshal(tR)
		Utils.Logger.Println(xmlstring)
		return true
	case "news":
		tR := &TextResponse{ToUserName: msg.GetSender(), FromUserName: wPP.GetMeta()["account"], CreateTime: time.Now().Unix(), MsgType: "text", Content: msg.GetMsgBody()}
		xmlstring, _ := xml.Marshal(tR)
		Utils.Logger.Println(xmlstring)
		return true
	case "resp":
		tR := &TextResponse{ToUserName: msg.GetSender(), FromUserName: wPP.GetMeta()["account"], CreateTime: time.Now().Unix(), MsgType: "text", Content: msg.GetMsgBody()}
		xmlstring, _ := xml.Marshal(tR)
		Utils.Logger.Println(xmlstring)
		return true
	default:
		tR := &TextResponse{ToUserName: msg.GetSender(), FromUserName: wPP.GetMeta()["account"], CreateTime: time.Now().Unix(), MsgType: "text", Content: msg.GetMsgBody()}
		xmlstring, _ := xml.Marshal(tR)
		Utils.Logger.Println(xmlstring)
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
			msgMeta := make(map[string]string)
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "text", createTime, msgMeta)
			Utils.Logger.Println(newMsg.GetMsgBody())
			DataShare.MsgQ <- newMsg
			return
		case "image":
			Utils.Logger.Println("Get a Image From ", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/MediaId")
			msgBody, _ := msgBodyPath.String(content)
			msgMeta := make(map[string]string)
			picUrlPath := xmlpath.MustCompile("/xml/PicUrl")
			picUrl, _ := picUrlPath.String(content)
			msgMeta["PicUrl"] = picUrl
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "image", createTime, msgMeta)
			Utils.Logger.Println(newMsg.GetMsgBody())
			return
		case "voice":
			Utils.Logger.Println("Get a Voice Msg From", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/MediaId")
			msgBody, _ := msgBodyPath.String(content)
			msgMeta := make(map[string]string)
			formatPath := xmlpath.MustCompile("/xml/Format")
			format, _ := formatPath.String(content)
			msgMeta["Format"] = format
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "voice", createTime, msgMeta)
			Utils.Logger.Println(newMsg.GetMsgBody())
			return
		case "video":
			Utils.Logger.Println("Get a Video Msg From", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/MediaId")
			msgBody, _ := msgBodyPath.String(content)
			msgMeta := make(map[string]string)
			thumbMediaIdPath := xmlpath.MustCompile("/xml/ThumbMediaId")
			thumbMediaId, _ := thumbMediaIdPath.String(content)
			msgMeta["ThumbMediaId"] = thumbMediaId
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "voice", createTime, msgMeta)
			Utils.Logger.Println(newMsg.GetMsgBody())
			return
		case "shortvideo":
			Utils.Logger.Println("Get a Short Video Msg From", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/MediaId")
			msgBody, _ := msgBodyPath.String(content)
			msgMeta := make(map[string]string)
			thumbMediaIdPath := xmlpath.MustCompile("/xml/ThumbMediaId")
			thumbMediaId, _ := thumbMediaIdPath.String(content)
			msgMeta["ThumbMediaId"] = thumbMediaId
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "voice", createTime, msgMeta)
			Utils.Logger.Println(newMsg.GetMsgBody())
			return
		case "location":
			Utils.Logger.Println("Get a Location Msg From", r.URL)
			msgBodyPath := xmlpath.MustCompile("/xml/Label")
			msgBody, _ := msgBodyPath.String(content)
			msgMeta := make(map[string]string)
			Location_XPath := xmlpath.MustCompile("/xml/Location_X")
			Location_X, _ := Location_XPath.String(content)
			msgMeta["Location_X"] = Location_X
			Location_YPath := xmlpath.MustCompile("/xml/Location_Y")
			Location_Y, _ := Location_YPath.String(content)
			msgMeta["Location_Y"] = Location_Y
			newMsg := Protocol.CreateMsg(msgBody, sender, wPP.GetName(), "Location", createTime, msgMeta)
			Utils.Logger.Println(newMsg.GetMsgBody())
			return
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
