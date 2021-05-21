package main

import (
	"encoding/base64"
	"encoding/json"
	"syscall/js"

	"go.dedis.ch/kyber/v4/suites"
)

// GOOS=js GOARCH=wasm go build -o ed25519_mul.wasm

var c chan bool

func init() {
	c = make(chan bool)
}

// inputs should only contain one element, which is a JSON in string format.
func ed25519_mul(this js.Value, inputs []js.Value) interface{} {
	var suite = suites.MustFind("Ed25519")
	var args map[string]interface{}
	json.Unmarshal([]byte(inputs[0].String()), &args)
	point1 := suite.Point()
	scalar := suite.Scalar()
	point1B, _ := base64.StdEncoding.DecodeString(args["point1"].(string))
	scalarB, _ := base64.StdEncoding.DecodeString(args["scalar"].(string))
	point1.UnmarshalBinary(point1B)
	scalar.UnmarshalBinary(scalarB)
	var resultB []byte
	//var result kyber.Point
	for i := 0; i < 1; i++ {
		resultB, _ = suite.Point().Mul(scalar, point1).MarshalBinary()
	}
	args["result"] = base64.StdEncoding.EncodeToString(resultB)
	//args["resultTest"] = result.String()
	args["Accepted"] = "true"
	return args
}

func main() {
	js.Global().Set("ed25519_mul", js.FuncOf(ed25519_mul))
	// Force the program to stay open by never sending to channel c
	<-c
}
