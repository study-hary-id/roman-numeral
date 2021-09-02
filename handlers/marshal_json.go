package handlers

import (
	"encoding/json"
	"net/http"
)

// marshalJsonHandler returns the JSON encoding of v with []byte type,
// and handle if there is an error when performing json.Marshal.
func marshalJsonHandler(w http.ResponseWriter, obj interface{}) []byte {
	js, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return js
}
