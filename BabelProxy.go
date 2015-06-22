package BabelProxy

type BabelProxy struct {
	ip, port                     string
	cCProviderInstance           *CCProvider
	platformProviderInstanceList []*PlatformProvider
	meta                         map[string]string
	bot                          *Bot
	contactRecordList            []*Contact
}

func (bp *BabelProxy) reloadCOnfiguration() {

}

func (bp *BabelProxy) createContact() {

}

func (bp *BabelProxy) processMsg() {

}

func (bp *BabelProxy) findContact(callId string) *Contact {

}

func (bp *BabelProxy) removeContact(callId string) bool {

}
