package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO : find out how to wrap it to the type interested in. Till then this is useless
// Creates a wrapper for all the Http parsing overheads
// reqObject : pass an instance of a struct to unmarshall the request to
// handler : func that will do further processing
func HttpWrapper(
	reqObject interface{},
	handler func(v interface{}) (interface{}, error),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := readReqBody(r, &reqObject)
		if err != nil {
			writeErrRes(w, err)
			return
		}

		res, err := handler(&reqObject) // TODO : what is the right way to invoke func argument

		if err != nil {
			writeErrRes(w, err)
			return
		}

		writeRes(w, res)
	}
}

type ErrRes struct {
	Error string `json:"error"`
}

func readReqBody(r *http.Request, reqBody interface{}) error {
	reqBodyJson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("unable to read request body. %s", err.Error())
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBodyJson, reqBody)
	if err != nil {
		return fmt.Errorf("unable to unmarshal request body. %s", err.Error())
	}

	return nil
}

func writeErrRes(w http.ResponseWriter, err error) {
	jsonErrRes, _ := json.Marshal(ErrRes{err.Error()})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonErrRes)
}

func writeRes(w http.ResponseWriter, content interface{}) {
	contentJson, err := json.Marshal(content)
	if err != nil {
		writeErrRes(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(contentJson)
}
