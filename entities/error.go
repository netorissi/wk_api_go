package entities

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type AppError struct {
	ID            string                 `json:"id,omitempty"`
	Message       string                 `json:"message,omitempty"`        // Message to be display to the end user without debugging information
	DetailedError string                 `json:"detailed_error,omitempty"` // Internal error string to help the developer
	RequestID     string                 `json:"request_id,omitempty"`     // The RequestId that's also set in the header
	StatusCode    int                    `json:"status_code,omitempty"`    // The http status code
	Where         string                 `json:"-"`                        // The function where it happened in the form of Struct.Func
	IsOAuth       bool                   `json:"is_oauth,omitempty"`       // Whether the error is OAuth specific
	Params        map[string]interface{} `json:"params"`
}

func (er *AppError) Error() string {
	return er.Where + ": " + er.Message + ", " + er.DetailedError
}

func (er *AppError) ToJson() string {
	b, err := json.Marshal(er)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func AppErrorFromJson(data io.Reader) *AppError {
	str := ""
	bytes, rerr := ioutil.ReadAll(data)
	if rerr != nil {
		str = rerr.Error()
	} else {
		str = string(bytes)
	}

	decoder := json.NewDecoder(strings.NewReader(str))
	var er AppError
	err := decoder.Decode(&er)
	if err == nil {
		return &er
	} else {
		return NewAppError("AppErrorFromJson", "entity.error.decode_json.app_error", nil, "body: "+str, http.StatusInternalServerError)
	}
}

func NewAppError(where string, ID string, params map[string]interface{}, details string, status int) *AppError {
	ap := &AppError{}
	ap.ID = ID
	ap.Params = params
	ap.Message = ID
	ap.Where = where
	ap.DetailedError = details
	ap.StatusCode = status
	ap.IsOAuth = false
	return ap
}
