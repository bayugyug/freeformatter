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
	//"path"
	"strings"
	"time"
)

func b(s string) *bufio.Reader { return bufio.NewReader(strings.NewReader(s)) }

//handler entry of the app
func handler() {

	//get mime type
	if pMimeType != "" {
		v := Formatters["MIME-TYPE"]
		_, resmsg := v.Format(v.Mode, pMimeType)
		fmt.Println(resmsg)
		return
	}
	//get mime list
	if pMimeList {
		v := Formatters["MIME-LIST"]
		rescode, resmsg := v.Format(v.Mode, pMimeType)
		showStatus(rescode, resmsg)
		return
	}
	//encode
	if pEncData != "" {
		v := Formatters["ENC-DATA"]
		rescode, resmsg := v.Format(v.Mode, pEncData)
		showStatus(rescode, resmsg)
		return
	}
	//encode-url
	if pEncUrl != "" {
		v := Formatters["ENC-URL"]
		rescode, resmsg := v.Format(v.Mode, pEncUrl)
		showStatus(rescode, resmsg)
		return
	}
	//decode
	if pDecData != "" {
		v := Formatters["DEC-DATA"]
		rescode, resmsg := v.Format(v.Mode, pDecData)
		showStatus(rescode, resmsg)
		return
	}
	//decode-url
	if pDecUrl != "" {
		v := Formatters["DEC-URL"]
		rescode, resmsg := v.Format(v.Mode, pDecUrl)
		showStatus(rescode, resmsg)
		return
	}
	//encode:b64
	if pBase64Enc != "" {
		v := Formatters["B64-ENC-DATA"]
		rescode, resmsg := v.Format(v.Mode, pBase64Enc)
		showStatus(rescode, resmsg)
		return
	}
	//encode-url:b64
	if pBase64UrlEnc != "" {
		v := Formatters["B64-ENC-URL"]
		rescode, resmsg := v.Format(v.Mode, pBase64UrlEnc)
		showStatus(rescode, resmsg)
		return
	}
	//decode:b64
	if pBase64Dec != "" {
		v := Formatters["B64-DEC-DATA"]
		rescode, resmsg := v.Format(v.Mode, pBase64Dec)
		showStatus(rescode, resmsg)
		return
	}
	//decode-url:b64
	if pBase64UrlDec != "" {
		v := Formatters["B64-DEC-URL"]
		rescode, resmsg := v.Format(v.Mode, pBase64UrlDec)
		showStatus(rescode, resmsg)
		return
	}
	//html-esc
	if pHtmlEsc != "" {
		v := Formatters["HTML-ENC-DATA"]
		rescode, resmsg := v.Format(v.Mode, pHtmlEsc)
		showStatus(rescode, resmsg)
		return
	}
	//html-esc-url
	if pHtmlUrlEsc != "" {
		v := Formatters["HTML-ENC-URL"]
		rescode, resmsg := v.Format(v.Mode, pHtmlUrlEsc)
		if !pHttpServe {
			showStatus(rescode, resmsg)
			return
		}
		return
	}

	//qrcode
	if pQRCodeGen != "" {
		v := Formatters["QR-CODE-GEN"]
		rescode, resmsg := v.Format(v.Mode, pQRCodeGen)
		if !pHttpServe {
			showStatus(rescode, resmsg)
			return
		}
		return
	}
	//serve http
	if pHttpServe {
		initHttpRouters()
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

func fmtMime(mode, data string) (string, string) {
	mimelist := getMimes()
	var jdata []byte
	if strings.EqualFold(mode, "TYPE") {
		//type
		for _, v := range mimelist {
			if strings.EqualFold(v.Ext, data) {
				jdata, _ = json.MarshalIndent(v, "", "\t")
				break
			}
		}
	} else {
		//list
		jdata, _ = json.MarshalIndent(mimelist, "", "\t")
	}
	return "", string(jdata)
}

func fmtEnc(mode, data string) (string, string) {
	var s string
	if strings.EqualFold(mode, "URL") {
		if !strings.HasPrefix(data, "http") {
			return "Fail", "Invalid url parameter: " + data
		}
		_, s = getResult(data)
	} else {
		s = data
	}
	return "", url.QueryEscape(s)
}

func fmtDec(mode, data string) (string, string) {
	var s string
	if strings.EqualFold(mode, "URL") {
		if !strings.HasPrefix(data, "http") {
			return "Fail", "Invalid url parameter: " + data
		}

		_, s = getResult(data)
	} else {
		s = data
	}
	enc, _ := url.QueryUnescape(s)
	return "", enc
}

func fmtEncB64(mode, data string) (string, string) {
	var s string
	if strings.EqualFold(mode, "URL") {
		if !strings.HasPrefix(data, "http") {
			return "Fail", "Invalid url parameter: " + data
		}
		_, s = getResult(data)
	} else {
		s = data
	}
	enc := base64.StdEncoding.EncodeToString([]byte(s))
	return "", enc
}

func fmtDecB64(mode, data string) (string, string) {
	var s string

	if strings.EqualFold(mode, "URL") {
		if !strings.HasPrefix(data, "http") {
			return "Fail", "Invalid url parameter: " + data
		}
		_, s = getResult(data)
	} else {
		s = data
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "Fail", fmt.Sprintf("Decode error: %v", err)
	}
	return "", string(decoded)
}

func fmtHtmlEsc(mode, data string) (string, string) {
	var s string
	if strings.EqualFold(mode, "URL") {
		if !strings.HasPrefix(data, "http") {
			return "Fail", "Invalid url parameter: " + data
		}
		_, s = getResult(data)
	} else {
		s = data
	}
	return "", html.EscapeString(s)
}

func fmtHtmlUnEsc(mode, data string) (string, string) {
	var s string
	if strings.EqualFold(mode, "URL") {
		if !strings.HasPrefix(data, "http") {
			return "Fail", "Invalid url parameter: " + data
		}

		_, s = getResult(data)
	} else {
		s = data
	}
	return "", html.UnescapeString(s)
}

func fmtQRCode(mode, data string) (string, string) {

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
		return "Fail", fmt.Sprintf("QR error: Invalid parameters! %v", err)
	}
	//sanity
	if qrparams.Data == "" {
		return "Fail", "QR error: data is empty!"
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
	dst := f //path.Join(pQRCodeTmp, f)
	//gen
	err = qrcode.WriteFile(d, qrcode.Medium, s, dst)
	if err != nil {
		return "Fail", fmt.Sprintf("QR error: failed to generate code! %v", err)
	}
	return "", dst
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
