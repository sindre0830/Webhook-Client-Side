package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
)

// secret for decryption
var secret = []byte{43, 123, 65, 232, 4, 42, 35, 234, 21, 122, 214}

// InputHandler handles webhook executions and checks if it is encrypted correctly.
func InputHandler(w http.ResponseWriter, r *http.Request) {
	var catch Catch
	//get content and branch if an error occured
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		catch.add(err, nil)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//extract signature from header based on known key
	signature := r.Header.Get("Signature")
	//decode signature
	signatureByte, err := hex.DecodeString(signature)
	if err != nil {
		catch.add(err, string(content))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//hash content of body and try to decode message
	mac := hmac.New(sha256.New, secret)
	_, err = mac.Write(content)
	if err != nil {
		catch.add(err, string(content))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//compare HMAC with received request
	if hmac.Equal(signatureByte, mac.Sum(nil)) {
		catch.add(nil, string(content))
		http.Error(w, "Successfully invoked dummy web service.", http.StatusOK)
	} else {
		err := errors.New("invalid invocation: hmac is not equal to request")
		catch.add(err, string(content))
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
