package smartstream

import (
	"log"
	"testing"

	"github.com/Gambitier/smartapigo/model"
)

var client *WebSocket

func TestSmartStream(t *testing.T) {
	client = New("A586457", "00998877")
	client.callbacks.onConnected = onConnected
	client.callbacks.onSnapquote = onSnapquote
	client.Connect()
}

func onConnected() {
	log.Printf("connected")
	err := client.Subscribe(model.SNAPQUOTE, []model.TokenInfo{{ExchangeType: model.NSECM, Token: "1594"}})
	if err != nil {
		log.Printf("error while subscribing")
	}
}

func onSnapquote(snapquote model.SnapQuote) {
	log.Printf("%d", snapquote.BestFiveSell[0])
}
