package gobitkan

func (this *Bitkan) AccountBalance(wType string) ([]byte, error) {
	data := map[string]interface{}{}
	data["type"] = wType
	path := "/open_api/v1/account/balance"
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Bitkan) AccountCoinBalance(wType string, currency string, params ...map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{}
	data["type"] = wType
	data["currency"] = currency
	if len(params) > 0 {
		for k, v := range params[0] {
			data[k] = v
		}
	}
	path := "/open_api/v1/account/coin_balance"
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Bitkan) AccountTransfer(msgId string, transfer map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{}
	data["msgId"] = msgId
	data["transfer"] = transfer
	path := "/open_api/v1/account/transfer"
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
