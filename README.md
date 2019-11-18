# gobitkan
*基于Bitkan官方API文档的Golang实现*

## 官方API文档地址
> https://bitkan.pro/en/help/doc/open-api

## 安装
```shell script
go get -u github.com/api-sdk/gobitkan
```


## 函数列表

* 行情数据
    * PriceSymbols()
    * PriceKline(symbol string, kType string, params ...map[string]interface{})
    * PriceDepth(symbol string, params ...map[string]interface{})
    * PriceSourceDepth(symbol string)
    
* 账户相关
    * AccountBalance(wType string)
    * AccountCoinBalance(wType string, currency string, params ...map[string]interface{})
    * AccountTransfer(msgId string, transfer map[string]interface{})
    
* 币币交易
    * SpotAdd(symbol string, amount string, price string, orderType string, oType string, params ...map[string]interface{})
    * SpotCancel(id string)
    * SpotOpens(symbol string, params ...map[string]interface{})
    * SpotHistory(symbol string, params ...map[string]interface{})
    * SpotDetail(id string, params ...map[string]interface{})
    
## 函数返回
***所有函数返回类型为 ([]byte, error)***

## Demo
```golang

package main

import (
	"github.com/api-sdk/gobitkan"
	"fmt"
)

func main() {

	baseURL := ""
	accessKey := ""
	secretKey := ""
	privateKey := ``

	var bitkan = gobitkan.Bitkan{accessKey, secretKey, privateKey, baseURL} // baseURL可直接传空字符串，已内置接口地址。

	// 获取所有币种余额
	res, err := bitkan.AccountBalance("spot")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 资金划转
	data := map[string]interface{}{"currency": "USDT", "amount": "1", "chain": "omni"}
	res, err = bitkan.AccountTransfer("1006", data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 深度数据（可选参数示例）
	params := map[string]interface{}{}
	params["step"] = "2"
	res, err = bitkan.PriceDepth("BTC-USDT", params)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

}

```