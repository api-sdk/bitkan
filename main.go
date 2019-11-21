package gobitkan

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"strconv"
)

var baseUrl = "https://openapi.bitkan.pro"

type Bitkan struct {
	AccessKey, SecretKey, PrivateKey, BaseUrl string
}

func NewBitkan(accessKey, secretKey, privateKey, baseUrl string) *Bitkan {
	return &Bitkan{accessKey, secretKey, privateKey, baseUrl}
}

func (this *Bitkan) SetBaseUrl(url string) {
	baseUrl = url
	return
}

func (this *Bitkan) signData(input interface{}) ([]byte, error) {
	tonce := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	h := hmac.New(sha256.New, []byte(this.SecretKey))
	io.WriteString(h, tonce)
	skSign := fmt.Sprintf("%x", h.Sum(nil))

	//fmt.Println(skSign)
	eData, err := this.priEncrypt([]byte(skSign))
	if err != nil {
		return nil, err
	}
	eshsp := base64.StdEncoding.EncodeToString(eData)
	//fmt.Println(eshsp)
	var signData map[string]interface{}
	access := map[string]interface{}{"id": this.AccessKey, "type": 5}
	t, _ := strconv.ParseInt(tonce, 10, 64)
	auth := map[string]interface{}{"tonce": t, "ESHSP": eshsp, "access": access}
	signData = map[string]interface{}{"data": input, "auth": auth}
	sData, err := json.Marshal(signData)
	signedData := string(sData)
	//fmt.Println("signedData:", signedData)
	h1 := hmac.New(sha256.New, []byte(this.AccessKey))
	io.WriteString(h1, signedData)
	signature := fmt.Sprintf("%x", h1.Sum(nil))
	//fmt.Println("signature:", signature)
	eSign, err := this.priEncrypt([]byte(signature))
	if err != nil {
		return nil, err
	}
	esign := base64.StdEncoding.EncodeToString(eSign)
	a := map[string]interface{}{"signedData": signedData, "esign": esign}

	requestData, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return requestData, nil
}

//私钥签名
func (this *Bitkan) priEncrypt(data []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode([]byte(this.PrivateKey))
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	private, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(nil, private.(*rsa.PrivateKey), crypto.Hash(0), data)
}

func (this *Bitkan) httpRequest(method string, path string, data []byte) ([]byte, error) {
	method = strings.ToUpper(method)
	url := baseUrl + path
	if this.BaseUrl != "" {
		url = this.BaseUrl + path
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	if "GET" == method {
		q := req.URL.Query()
		m := make(map[string]interface{})
		err := json.Unmarshal(data, &m)
		if err != nil {
			return nil, err
		} else {
			for k, v := range m {
				q.Add(k, v.(string))
			}
		}
		req.URL.RawQuery = q.Encode()
		//fmt.Println("===========================")
		//fmt.Println(req.URL.String())
	}

	if "POST" == method {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//fmt.Println("response URL:", resp.Request.URL)
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	//return ioutil.ReadAll(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func (this *Bitkan) postRequest(path string, data map[string]interface{}) ([]byte, error) {
	requestData, _ := this.signData(data)
	body, err := this.httpRequest("post", path, requestData)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Bitkan) getRequest(path string, data map[string]interface{}) ([]byte, error) {
	requestData, _ := json.Marshal(data)
	body, err := this.httpRequest("get", path, requestData)
	if err != nil {
		return nil, err
	}
	return body, nil
}
