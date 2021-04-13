package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"syscall/js"
)

var c chan bool

func init() {
	c = make(chan bool)
}

// inputs should only contain one element, which is a JSON in string format.
func increaseCounter(this js.Value, inputs []js.Value) interface{} {
	var args map[string]interface{}
	json.Unmarshal([]byte(inputs[0].String()), &args)
	counter, err := strconv.Atoi(fmt.Sprintf("%v", args["counter"]))
	if err != nil {
		return err
	}
	args["result"] = strconv.Itoa(counter + 1)
	return args
}

func main() {
	js.Global().Set("increaseCounter", js.FuncOf(increaseCounter))
	// Force the program to stay open by never sending to channel c
	<-c
}
