package testgoar

import (
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/utils"
	"io/ioutil"
	"math/big"
	"testing"
	"time"
)

const arNode = "http://127.0.0.1:1984"

func TestInit(t *testing.T)  {
	testKeyData, _ := ioutil.ReadFile("./testKey.json")
	wallet, err := goar.NewWallet(testKeyData, arNode)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	_, err = Init(wallet, false)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	_, err = wallet.Client.GetInfo()
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestMine(t *testing.T)  {
	testSDK, _, err := newTestSDk()

	if err != nil {
		t.Errorf("%v\n", err)
	}

	err = testSDK.Mine()
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestReadyForMining(t *testing.T)  {
	testSDK, _, err := newTestSDk()

	if err != nil {
		t.Errorf("%v\n", err)
	}

	txs, err := testSDK.ReadyForMining()
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Logf("txs: %v\n", txs)
}

func TestDrop(t *testing.T)  {

	testSDK, wallet, err := newTestSDk()
	if err != nil {
		t.Errorf("%v\n", err)
	}

	arBefore, err := wallet.Client.GetWalletBalance(wallet.Address)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	t.Logf("before drop: %f\n", arBefore)

	err = testSDK.Drop(wallet.Address, utils.ARToWinston(big.NewFloat(10)))

	if err != nil {
		t.Errorf("%v\n", err)
	}

	time.Sleep(time.Second * 10)
	arAfter, err := wallet.Client.GetWalletBalance(wallet.Address)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Logf("after drop: %f\n", arAfter)

	if big.NewFloat(0).Add(big.NewFloat(10), arBefore).Cmp(arAfter) != 0 {
		t.Errorf("Drop error, before: %f, after: %f\n", arBefore, arAfter)
		t.Errorf("Drop error, before: %f, after: %f\n", arBefore, arAfter)
	}
}

func newTestSDk() (*TestGoAr, *goar.Wallet ,error) {
	testKeyData, _ := ioutil.ReadFile("./testKey.json")
	wallet, err := goar.NewWallet(testKeyData, arNode)
	if err != nil {
		return nil, nil, err
	}
	testSDK, err := Init(wallet, false)
	if err != nil {
		return nil, nil, err
	}

	return testSDK, wallet, nil
}