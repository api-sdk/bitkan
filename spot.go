package gobitkan

func (this *Bitkan) SpotAdd(symbol string, amount string, price string, orderType string, oType string, params ...map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{}
	data["symbol"] = symbol
	data["amount"] = amount
	data["price"] = price
	data["orderType"] = orderType
	data["type"] = oType
	if len(params) > 0 {
		for k, v := range params[0] {
			data[k] = v
		}
	}
	path := "/open_api/v1/spot/add"
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Bitkan) SpotCancel(id string) ([]byte, error) {
	data := map[string]interface{}{}
	data["id"] = id
	path := "/open_api/v1/spot/cancel"
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Bitkan) SpotOpens(symbol string, params ...map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{}
	data["symbol"] = symbol
	if len(params) > 0 {
		for k, v := range params[0] {
			data[k] = v
		}
	}
	path := "/open_api/v1/spot/opens"
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Bitkan) SpotHistory(symbol string, params ...map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{}
	data["symbol"] = symbol
	if len(params) > 0 {
		for k, v := range params[0] {
			data[k] = v
		}
	}
	path := "/open_api/v1/spot/history"
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Bitkan) SpotDetail(id string, params ...map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{}
	data["id"] = id
	if len(params) > 0 {
		for k, v := range params[0] {
			data[k] = v
		}
	}
	path := "/open_api/v1/spot/detail"
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
