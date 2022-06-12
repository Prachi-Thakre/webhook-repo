package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/webhook-repo/utilities"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		utilities.URLReturnResponseJson(w, getData(w, r))
	} else if r.Method == "POST" {
		utilities.URLReturnResponseJson(w, addData(w, r))
	} else {
		utilities.URLReturnResponseJson(w, "error")
	}

}

func getData(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	// no request validation right now

	return
}

func addData(w http.ResponseWriter, r *http.Request) (returnData utilities.ResponseJSON) {
	// no request validation right now

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ErrorResponse(&returnData, "Failure: unable to read")
		return
	}

	fmt.Println(string(body))
	return

}
