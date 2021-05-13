package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"syscall/js"
)

var c chan bool

func init() {
	c = make(chan bool)
}

// GOOS=js GOARCH=wasm go build -o randFloatMul.wasm

// inputs should only contain one element, which is a JSON in string format.
func randFloatMul(this js.Value, inputs []js.Value) interface{} {
	var args map[string]interface{}
	json.Unmarshal([]byte(inputs[0].String()), &args)
	rand1, err := strconv.ParseFloat(fmt.Sprintf("%v", args["rand1"]), 64)
	rand2, err2 := strconv.ParseFloat(fmt.Sprintf("%v", args["rand2"]), 64)
	if err != nil {
		return err
	}
	if err2 != nil {
		return err
	}
	args["Accepted"] = "true"
	args["result"] = strconv.FormatFloat(rand1*rand2, 'E', -1, 64)
	//rand.Seed(time.Now().UTC().UnixNano())
	args["random"] = strconv.Itoa(rand.Intn(100))
	return args
}

func main() {
	js.Global().Set("randFloatMul", js.FuncOf(randFloatMul))
	// Force the program to stay open by never sending to channel c
	<-c
}
