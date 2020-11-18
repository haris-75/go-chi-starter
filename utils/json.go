package utils

import (
	"datumbrain/my-project/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ParseJson gets json for request and fills the target model
func ParseJson(r *http.Request, target interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

// RespondJson makes the response with payload as json format
func RespondJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Error.Println(err)
		RespondError(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(response)
}

// RespondCustomError makes the error response with given message in json format
func RespondCustomError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf(`{"code":%v,"message":"%v"}`, code, message)))
}

// RespondCustomError makes the error response with default message in json format
func RespondError(w http.ResponseWriter, code int) {
	var message string
	switch code {
	case http.StatusBadRequest:
		message = "The request had invalid inputs or otherwise cannot be served."
	case http.StatusUnauthorized:
		message = "Authorization information is missing or invalid."
	case http.StatusNotFound:
		message = "Unable to find requested record."
	case http.StatusRequestTimeout:
		message = "Request took too long to process."
	case http.StatusRequestedRangeNotSatisfiable:
		message = "No resource available, unable to fulfill the request."
	case http.StatusTooManyRequests:
		message = "Request rate too high, requests from this this user are throttled."
	case http.StatusInternalServerError:
		message = "An error was encountered."
	case http.StatusServiceUnavailable:
		message = "The service is unavailable, please try again later."
	case http.StatusGatewayTimeout:
		message = "The service timed out waiting for an upstream response. Try again later."
	}

	RespondCustomError(w, code, message)
}

// FillFields marshal `src` to Json and then unmarshal into target.
func FillFields(target interface{}, src interface{}) error {
	jsonString, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonString, target)
	return err
}
