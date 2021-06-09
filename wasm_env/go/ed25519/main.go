package main

import (
	"encoding/base64"
	"encoding/json"
	"syscall/js"

	"go.dedis.ch/kyber/v4/suites"
)

// GOOS=js GOARCH=wasm go build -o main.wasm

// WASM
// add : 30         both 10k ops
// mult : 6406

// NATIVE
// add : 6
// mult : 1500

var c chan bool

func init() {
	c = make(chan bool)
}

// inputs should only contain one element, which is a JSON in string format.
func cryptoOp(this js.Value, inputs []js.Value) interface{} {
	suite := suites.MustFind("Ed25519")
	var args map[string]interface{}
	json.Unmarshal([]byte(inputs[0].String()), &args)
	stream := suite.RandomStream()
	point1 := suite.Point().Pick(stream)
	point2 := suite.Point().Pick(stream)
	//point1B, _ := base64.StdEncoding.DecodeString(args["point1"].(string))
	//point1.UnmarshalBinary(point1B)
	var resultB []byte
	//var result kyber.Point
	for i := 0; i < 500; i++ {
		point1 = suite.Point().Add(point1, point2)
	}
	args["result"] = base64.StdEncoding.EncodeToString(resultB)
	//args["resultTest"] = result.String()
	args["Accepted"] = "true"
	return args
}

func main() {
	js.Global().Set("cryptoOp", js.FuncOf(cryptoOp))
	// Force the program to stay open by never sending to channel c
	<-c
}
