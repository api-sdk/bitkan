package gobitkan

func (this *Bitkan) PriceSymbols() ([]byte, error) {
	data := map[string]interface{}{}
	path := "/open_api/v1/price/symbols"
	body, err := this.getRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Bitkan) PriceKline(symbol string, kType string, params ...map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{}
	data["symbol"] = symbol
	data["type"] = kType
	if len(params) > 0 {
		for k, v := range params[0] {
			data[k] = v
		}
	}
	path := "/open_api/v1/price/k_line"
	body, err := this.getRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Bitkan) PriceDepth(symbol string, params ...map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{}
	data["symbol"] = symbol
	if len(params) > 0 {
		for k, v := range params[0] {
			data[k] = v
		}
	}
	path := "/open_api/v1/price/depth"
	body, err := this.getRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Bitkan) PriceSourceDepth(symbol string) ([]byte, error) {
	data := map[string]interface{}{}
	data["symbol"] = symbol
	path := "/open_api/v1/price/source_depth"
	body, err := this.getRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
