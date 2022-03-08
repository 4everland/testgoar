package testgoar

import (
	"github.com/everFinance/goar"
	"io/ioutil"
	"testing"
)

func TestInit(t *testing.T)  {
	arNode := "http://127.0.0.1:1984"
	testKeyData, _ := ioutil.ReadFile("./testKey.json")
	wallet, err := goar.NewWallet(testKeyData, arNode)
	if err != nil {
		t.Logf("%v\n", err)
	}
	_, err = Init(wallet, false)
	if err != nil {
		t.Logf("%v\n", err)
	}
}