package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	msg := `

		http://localhost:7777/mime-type?p=123445

		http://localhost:7777/mime-list

		http://localhost:7777/enc-data?p=123445

		http://localhost:7777/enc-url?p=htp://www.google.com.sg/

		http://localhost:7777/dec-data?p=123445

		http://localhost:7777/dec-url?p=htp://www.google.com.sg/

		http://localhost:7777/b64-enc-data?p=123445

		http://localhost:7777/b64-enc-url?p=htp://www.google.com.sg/

		http://localhost:7777/b64-dec-datap=123445

		http://localhost:7777/b64-dec-url?p=htp://www.google.com.sg/

		http://localhost:7777/html-enc-data?p=123445

		http://localhost:7777/html-enc-url?p=htp://www.google.com.sg/

`
	fmt.Fprint(w, "Welcome to Freeformatter!\n", fmt.Sprintf("Version: %s\n\n\n%s", pVersion, msg))
}

func formatHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//Content-Type: application/json
	var p = strings.TrimSpace(r.URL.Query().Get("p"))
	var m = strings.ToUpper(strings.TrimSpace(ps.ByName("mode")))

	//not-found
	v, ok := Formatters[m]
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	fmt.Println("RAW-DATA: ", p)
	rescode, result := v.Format(v.Mode, p)
	s := rescode
	if s == "" {
		s = "Success"
		//download-push-file
		if strings.EqualFold(m, "QR-CODE-GEN") {
			pushQrCode(w, r, pQRParams.Filename)
			return
		}
	}
	//good
	w.Header().Set("Content-Type", "application/json")
	jdata, _ := json.MarshalIndent(FormatterOutput{Result: result, Status: s}, "", "\t")
	fmt.Fprint(w, string(jdata))
}

func pushQrCode(w http.ResponseWriter, r *http.Request, filename string) {
	//Check if file exists and open
	fh, err := os.Open(filename)
	defer fh.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	//Get the Content-Type of the file
	fhdr := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	fh.Read(fhdr)
	//Get content type of file
	ftype := http.DetectContentType(fhdr)

	//Get the file size
	fstat, _ := fh.Stat()                        //Get info from file
	fsize := strconv.FormatInt(fstat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", ftype)
	w.Header().Set("Content-Length", fsize)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	fh.Seek(0, 0)
	io.Copy(w, fh) //'Copy' the file to the client
}
