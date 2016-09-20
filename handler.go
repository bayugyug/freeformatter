package main

import (
	"bufio"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"html"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func b(s string) *bufio.Reader { return bufio.NewReader(strings.NewReader(s)) }

//handler entry of the app
func handler() {

	//get mime type
	if pMimeType != "" {
		v := FormatterList["MIME-TYPE"]
		v.Format(v.Mode, pMimeType)
		return
	}
	//get mime list
	if pMimeList {
		v := FormatterList["MIME-LIST"]
		v.Format(v.Mode, pMimeType)
		return
	}
	//encode
	if pEncData != "" {
		v := FormatterList["ENC-DATA"]
		v.Format(v.Mode, pEncData)
		return
	}
	//encode-url
	if pEncUrl != "" {
		v := FormatterList["ENC-URL"]
		v.Format(v.Mode, pEncUrl)
		return
	}
	//decode
	if pDecData != "" {
		v := FormatterList["DEC-DATA"]
		v.Format(v.Mode, pDecData)
		return
	}
	//decode-url
	if pDecUrl != "" {
		v := FormatterList["DEC-URL"]
		v.Format(v.Mode, pDecUrl)
		return
	}
	//encode:b64
	if pBase64Enc != "" {
		v := FormatterList["B64-ENC-DATA"]
		v.Format(v.Mode, pBase64Enc)
		return
	}
	//encode-url:b64
	if pBase64UrlEnc != "" {
		v := FormatterList["B64-ENC-URL"]
		v.Format(v.Mode, pBase64UrlEnc)
		return
	}
	//decode:b64
	if pBase64Dec != "" {
		v := FormatterList["B64-DEC-DATA"]
		v.Format(v.Mode, pBase64Dec)
		return
	}
	//decode-url:b64
	if pBase64UrlDec != "" {
		v := FormatterList["B64-DEC-URL"]
		v.Format(v.Mode, pBase64UrlDec)
		return
	}
	//html-esc
	if pHtmlEsc != "" {
		v := FormatterList["HTML-ENC-DATA"]
		v.Format(v.Mode, pHtmlEsc)
		return
	}
	//html-esc-url
	if pHtmlUrlEsc != "" {
		v := FormatterList["HTML-ENC-URL"]
		v.Format(v.Mode, pHtmlUrlEsc)
		return
	}

	//qrcode
	if pQRCodeGen != "" {
		v := FormatterList["QR-CODE-GEN"]
		v.Format(v.Mode, pQRCodeGen)
		return
	}
}

func getMimes() []MIMETypes {
	var mimes []MIMETypes
	err := json.Unmarshal(jsonStr, &mimes)
	if err != nil {
		fmt.Println("error:", err)
	}
	return mimes
}

func fmtMime(mode, data string) {
	mimelist := getMimes()
	if strings.EqualFold(mode, "TYPE") {
		//type
		for _, v := range mimelist {
			if strings.EqualFold(v.Ext, data) {
				jdata, _ := json.MarshalIndent(v, "", "\t")
				fmt.Println(string(jdata))
				break
			}
		}
	} else {
		//list
		jdata, _ := json.MarshalIndent(mimelist, "", "\t")
		fmt.Println(string(jdata))
	}
}

func fmtEnc(mode, data string) {
	var s string
	if strings.EqualFold(mode, "URL") {
		_, s = getResult(data)
	} else {
		s = data
	}
	showStatus("", url.QueryEscape(s))
}

func fmtDec(mode, data string) {
	var s string
	if strings.EqualFold(mode, "URL") {
		_, s = getResult(data)
	} else {
		s = data
	}
	enc, _ := url.QueryUnescape(s)
	showStatus("", enc)
}

func fmtEncB64(mode, data string) {
	var s string
	if strings.EqualFold(mode, "URL") {
		_, s = getResult(data)
	} else {
		s = data
	}
	enc := base64.StdEncoding.EncodeToString([]byte(s))
	showStatus("", enc)
}

func fmtDecB64(mode, data string) {
	var s string

	if strings.EqualFold(mode, "URL") {
		_, s = getResult(data)
	} else {
		s = data
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		showStatus("Fail", fmt.Sprintf("Decode error: %v", err))
		return
	}
	showStatus("", string(decoded))
}

func fmtHtmlEsc(mode, data string) {
	var s string
	if strings.EqualFold(mode, "URL") {
		_, s = getResult(data)
	} else {
		s = data
	}
	showStatus("", html.EscapeString(s))
}

func fmtHtmlUnEsc(mode, data string) {
	var s string
	if strings.EqualFold(mode, "URL") {
		_, s = getResult(data)
	} else {
		s = data
	}
	showStatus("", html.UnescapeString(s))
}

func fmtQRCode(mode, data string) {

	//-qr-code-gen='{"data": "https://www.google.com.sg/","filename":"qrcode.png","size":256}'
	type QRCodeData struct {
		Data     string `json:"data,ommitempty`
		Filename string `json:"filename,ommitempty`
		Size     int    `json:"size,ommitempty`
	}
	var qrparams QRCodeData
	err := json.Unmarshal([]byte(data), &qrparams)

	//sanity
	if err != nil {
		showStatus("Fail", fmt.Sprintf("QR error: Invalid parameters! %v", err))
		return
	}
	//sanity
	if qrparams.Data == "" {
		showStatus("Fail", "QR error: data is empty!")
		return
	}
	var d, f string
	var s int

	//default
	s = qrparams.Size
	if qrparams.Size < 50 || qrparams.Size > 512 {
		s = 128
	}
	//default
	d = strings.TrimSpace(qrparams.Data)
	if len(d) > 4196 {
		d = d[:4196]
	}
	//default
	f = strings.TrimSpace(qrparams.Filename)
	if f == "" {
		f = "qrcode.png"
	}
	//gen
	err = qrcode.WriteFile(d, qrcode.Medium, s, f)
	if err != nil {
		showStatus("Fail", fmt.Sprintf("QR error: failed to generate code! %v", err))
		return
	}
	showStatus("", f)
}

func showStatus(status, result string) {
	s := status
	if s == "" {
		s = "Success"
	}
	jdata, _ := json.MarshalIndent(FormatterOutput{Result: result, Status: s}, "", "\t")
	fmt.Println(string(jdata))
}

//getResult http req a url
func getResult(url string) (int, string) {
	//client
	c := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, RootCAs: pool},
		Dial: (&net.Dialer{
			Timeout: 300 * time.Second,
		}).Dial,
	},
	}
	res, err := c.Get(url)
	if err != nil {
		log.Println("ERROR: getResult:", err)
		return 0, ""
	}
	//get response
	robots, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Println("ERROR: getResult:", err)
		return 0, ""
	}
	//give
	return res.StatusCode, string(robots)
}
