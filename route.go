package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
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
	var q = r.URL.Query()
	var p = strings.TrimSpace(q.Get("p"))
	var m = strings.ToUpper(strings.TrimSpace(ps.ByName("mode")))

	//not-found
	v, ok := Formatters[m]
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("RAW-DATA: ", p)
	rescode, result := v.Format(v.Mode, p)
	s := rescode
	if s == "" {
		s = "Success"
	}
	jdata, _ := json.MarshalIndent(FormatterOutput{Result: result, Status: s}, "", "\t")
	fmt.Fprint(w, string(jdata))
}
