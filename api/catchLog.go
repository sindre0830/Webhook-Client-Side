package api

import "time"

// inputLog stores the last catches.
var inputLog []Catch

// Catch keeps information about a catch.
//
// Functionality: add
type Catch struct {
	Time         string      `json:"time"`
	ErrorMessage error       `json:"error_message"`
	RawBody      interface{} `json:"raw_body"`
}

// add sets new content in structure and updates inputLog.
func (catch *Catch) add(errorMessage error, rawBody interface{}) {
	//set data in structure
	catch.Time = time.Now().Format("2006-01-02 15:04:05")
	catch.ErrorMessage = errorMessage
	catch.RawBody = rawBody
	//add catch to input log as the first element
	inputLog = append([]Catch{*catch}, inputLog...)
	//branch if there are more then 5 catches in log and discard oldest catch
	if len(inputLog) > 5 {
		inputLog = inputLog[:5]
	}
}
