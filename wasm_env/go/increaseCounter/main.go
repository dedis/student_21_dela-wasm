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

// GOOS=js GOARCH=wasm go build -o main.wasm

// inputs should only contain one element, which is a JSON in string format.
func increaseCounter(this js.Value, inputs []js.Value) interface{} {
	var args map[string]interface{}
	json.Unmarshal([]byte(inputs[0].String()), &args)
	counter, err := strconv.Atoi(fmt.Sprintf("%v", args["counter"]))
	if err != nil {
		return err
	}
	/* for i := 1; i < 1000000; i++ {
		a := rand.Int()
		a *= rand.Int()
	} */
	args["Accepted"] = "true"
	args["result"] = strconv.Itoa(counter + 1)
	//print(args)
	return args
}

func main() {
	js.Global().Set("increaseCounter", js.FuncOf(increaseCounter))
	// Force the program to stay open by never sending to channel c
	<-c
}
