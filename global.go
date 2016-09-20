package main

import (
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

const (
	usageShowConsole  = "use to enable the output in console"
	usageMimeList     = "use to show the list of MIME types"
	usageDataFormat   = "use to encode/decode a data"
	usageBase64Format = "use to base64 encode/decode a data"
	usageHtmlFormat   = "use to escape/unescape an HTML"
	usageQRCode       = "use to generate a PNG QR Code"
)

type Formatter struct {
	Mode   string
	Format func(mode string, data string)
}

type FormatterOutput struct {
	Result string `json:"result,ommitempty`
	Status string `json:"status,ommitempty`
}

var (
	pLogDir = "."
	//loggers
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger

	//signal flag
	pStillRunning = true

	pBuildTime = "0"
	pVersion   = "0.1.0" + "-" + pBuildTime
	//console
	pShowConsole = true
	//envt
	pEnvVars = map[string]*string{
		"GMONGERS_LDIR": &pLogDir,
	}

	//print_format
	pPrintFormat  = "json"
	pHelp         = false
	pMimeList     = false
	pMimeType     = ""
	pEncData      = ""
	pDecData      = ""
	pEncUrl       = ""
	pDecUrl       = ""
	pBase64Enc    = ""
	pBase64Dec    = ""
	pBase64UrlEnc = ""
	pBase64UrlDec = ""
	pHtmlEsc      = ""
	pHtmlUrlEsc   = ""
	pQRCodeGen    = ""
	//ssl certs
	pool *x509.CertPool

	FormatterList map[string]Formatter
)

func init() {
	//uniqueness
	rand.Seed(time.Now().UnixNano())
	//recovery
	initRecov()
	//evt
	initEnvParams()
	//loggers
	initLogger(os.Stdout, os.Stdout, os.Stderr)

	//init certs
	pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(pemCerts)

	//in
	FormatterList = map[string]Formatter{
		"MIME-TYPE":     {"TYPE", fmtMime},
		"MIME-LIST":     {"LIST", fmtMime},
		"ENC-DATA":      {"DATA", fmtEnc},
		"ENC-URL":       {"URL", fmtEnc},
		"DEC-DATA":      {"DATA", fmtDec},
		"DEC-URL":       {"URL", fmtDec},
		"B64-ENC-DATA":  {"DATA", fmtEncB64},
		"B64-ENC-URL":   {"URL", fmtEncB64},
		"B64-DEC-DATA":  {"DATA", fmtDecB64},
		"B64-DEC-URL":   {"URL", fmtDecB64},
		"HTML-ENC-DATA": {"DATA", fmtHtmlEsc},
		"HTML-ENC-URL":  {"URL", fmtHtmlEsc},
		"QR-CODE-GEN":   {"DATA", fmtQRCode},
	}

}

//initRecov is for dumpIng segv in
func initRecov() {
	//might help u
	defer func() {
		recvr := recover()
		if recvr != nil {
			fmt.Println("MAIN-RECOV-INIT: ", recvr)
		}
	}()
}

//os.Stdout, os.Stdout, os.Stderr
func initLogger(i, w, e io.Writer) {
	//just in case
	if !pShowConsole {
		infoLog = makeLogger(i, pLogDir, "freeformatter", "INFO: ")
		warnLog = makeLogger(w, pLogDir, "freeformatter", "WARN: ")
		errorLog = makeLogger(e, pLogDir, "freeformatter", "ERROR: ")
	} else {
		infoLog = log.New(i,
			"INFO: ",
			log.Ldate|log.Ltime|log.Lmicroseconds)
		warnLog = log.New(w,
			"WARN: ",
			log.Ldate|log.Ltime|log.Lshortfile)
		errorLog = log.New(e,
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}
}

//initEnvParams enable all OS envt vars to reload internally
func initEnvParams() {
	//just in-case, over-write from ENV
	for k, v := range pEnvVars {
		if os.Getenv(k) != "" {
			*v = os.Getenv(k)
		}
	}
	//get options
	flag.BoolVar(&pMimeList, "mime-list", pMimeList, usageMimeList)
	flag.StringVar(&pMimeType, "mime-type", pMimeType, usageMimeList)

	flag.StringVar(&pEncData, "encode", pEncData, usageDataFormat)
	flag.StringVar(&pDecData, "decode", pDecData, usageDataFormat)

	flag.StringVar(&pEncUrl, "encode-url", pEncUrl, usageDataFormat)
	flag.StringVar(&pDecUrl, "decode-url", pDecUrl, usageDataFormat)

	flag.StringVar(&pBase64Enc, "b64encode", pBase64Enc, usageBase64Format)
	flag.StringVar(&pBase64Dec, "b64decode", pBase64Dec, usageBase64Format)

	flag.StringVar(&pBase64UrlEnc, "b64encode-url", pBase64UrlEnc, usageBase64Format)
	flag.StringVar(&pBase64UrlDec, "b64decode-url", pBase64UrlDec, usageBase64Format)

	flag.StringVar(&pHtmlEsc, "html-esc", pHtmlEsc, usageHtmlFormat)
	flag.StringVar(&pHtmlUrlEsc, "html-esc-url", pHtmlUrlEsc, usageHtmlFormat)

	flag.StringVar(&pQRCodeGen, "qr-code-gen", pQRCodeGen, usageQRCode)
	flag.Parse()

	//either 1 should be present
	if pHelp {
		showUsage()
	}

	//either 1 should be present
	if !pMimeList &&
		pMimeType == "" &&
		pEncData == "" &&
		pDecData == "" &&
		pEncUrl == "" &&
		pDecUrl == "" &&
		pBase64Enc == "" &&
		pBase64Dec == "" &&
		pBase64UrlEnc == "" &&
		pBase64UrlDec == "" &&
		pHtmlEsc == "" &&
		pHtmlUrlEsc == "" &&
		pQRCodeGen == "" {
		showUsage()
	}

}

//formatLogger try to init all filehandles for logs
func formatLogger(fdir, fname, pfx string) string {
	t := time.Now()
	r := regexp.MustCompile("[^a-zA-Z0-9]")
	p := t.Format("2006-01-02") + "-" + r.ReplaceAllString(strings.ToLower(pfx), "")
	s := path.Join(pLogDir, fdir)
	if _, err := os.Stat(s); os.IsNotExist(err) {
		//mkdir -p
		os.MkdirAll(s, os.ModePerm)
	}
	return path.Join(s, p+"-"+fname+".log")
}

//makeLogger initialize the logger either via file or console
func makeLogger(w io.Writer, ldir, fname, pfx string) *log.Logger {
	logFile := w
	if !pShowConsole {
		var err error
		logFile, err = os.OpenFile(formatLogger(ldir, fname, pfx), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			log.Println(err)
		}
	}
	//give it
	return log.New(logFile,
		pfx,
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

}

//dumpW log into warning
func dumpW(s ...interface{}) {
	warnLog.Println(s...)
}

//dumpWF log into warning w/ fmt
func dumpWF(format string, s ...interface{}) {
	warnLog.Println(fmt.Sprintf(format, s...))
}

//dumpE log into error
func dumpE(s ...interface{}) {
	errorLog.Println(s...)
}

//dumpE log into error w/ fmt
func dumpEF(format string, s ...interface{}) {
	errorLog.Println(fmt.Sprintf(format, s...))
}

//dumpI log into info
func dumpI(s ...interface{}) {
	infoLog.Println(s...)
}

//dumpIF log into info
func dumpIF(format string, s ...interface{}) {
	infoLog.Println(fmt.Sprintf(format, s...))
}

//showUsage
func showUsage() {

	msg := `
-------------------------------------------------------


	./freeformatter  --encode "test data here"

	./freeformatter  --decode "test+data+here"

	./freeformatter  --mime-type '.avi'

	./freeformatter  --mime-list

	./freeformatter  --qr-code-gen='{"data": "https://www.google.com.sg/","filename":"qrcode.png","size":256}'
-------------------------------------------------------
`
	log.Println("Version: ", pVersion, "\n")
	flag.PrintDefaults()
	fmt.Println(msg)
	os.Exit(0)
}
