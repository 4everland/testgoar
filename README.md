# TestWeave SDK

This is the SDK of the TestWeave. TestWeave is the testing environment of the Arweave.

The go implementation of [ArweaveTeam/testweave-sdk](https://github.com/ArweaveTeam/testweave-sdk).

### Install

```
go get github.com/4everland/testgoar
```

## MANDATORY PREREQUISITES

__To work with the TestWeave, you need to install a local testnet. To do so, follow the instructions here: [ArweaveTeam/testweave-docker](https://github.com/ArweaveTeam/testweave-docker) or [4everland/arweave-docker](https://github.com/4everland/arweave-docker)__

## Usage

use [github.com/everFinance/goar](https://github.com/everFinance/goar) deploying files on the Arweave

Firstly you need to create a TestWeave instance on the top on an Arweave node, as following:

```go
import (
    "github/4everland/testgoar/"
    "github.com/everFinance/goar"
)

// init arweave as usual
arNode := "http://127.0.0.1:1984"
testKeyData := []byte("test_wallet_key_content")
wallet, err := goar.NewWallet(testKeyData, arNode)

if err != nil {
    return err
}


// init TestWeave on the top of arweave
testGoAr, err := testgoar.Init(wallet, true)

if err != nil {
    return err
}

```

And here you go! Now you can use your arweave instance as usual, but every interaction will be performed on the test network!

For a fast bootstrap checkout the examples in the following sections.

To check all the useful helpers that the SDK supplies, checkout the XXX section.

## Example  - Submitting a data transaction

1. Initialize the arweave node and the TestWeave on it:

```go
import (
    "github/4everland/testgoar/"
    "github.com/everFinance/goar"
)

arNode := "http://127.0.0.1:1984"
testKeyData := []byte("test_key_content")
wallet, err := goar.NewWallet(testKeyData, arNode)

if err != nil {
return err
}

testGoAr, err := testgoar.Init(wallet, true)

if err != nil {
return err
}
```

2 Drop AR to wallet

```go
err = testGoAr.Drop(wallet.Address, utils.ARToWinston(big.NewFloat(10)))
//if automine is false
//_ = testGoAr.Mine()
```

3.Create a data transaction, sign and post it

```go

data := `
<html>
  <head>
    <meta charset="UTF-8">
    <title>Info about arweave</title>
  </head>
  <body>
    Arweave is the best web3-related thing out there!!!
  </body>
</html>`

tx, err := wallet.SendData([]byte(data),nil)
//if automine is false
//_ = testGoAr.Mine()
```

Thats it!

## SDK helpers

For easily test Arweave applications, the SDK supplies the helpers described in the following sections.
### drop(wallet, quantity)

Drops AR from the root wallet to another one. Use it as followings:

```go
err = testGoAr.Drop(wallet.Address, utils.ARToWinston(big.NewFloat(10)))
ar, err := wallet.Client.GetWalletBalance(wallet.Address)
```

### mine()

Mines the following block of the testnet and all the transactions contained in it.

```javascript
_ = testGoAr.Mine()
```

### getter rootJWK

Returns the root JWK, it has an initial balance of 10000000 and the address MlV6DeOtRmakDOf6vgOBlif795tcWimgyPsYYNQ8q1Y

```javascript
jkw := testGoAr.RootJWK
```
