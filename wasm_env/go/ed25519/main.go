package main

import (
	"encoding/base64"
	"encoding/json"
	"syscall/js"

	"go.dedis.ch/kyber/v4/suites"
)

// GOOS=js GOARCH=wasm go build -o main.wasm

// add : 46         both 10k ops
// mult : 6406

var c chan bool

func init() {
	c = make(chan bool)
}

// inputs should only contain one element, which is a JSON in string format.
func cryptoOp(this js.Value, inputs []js.Value) interface{} {
	var suite = suites.MustFind("Ed25519")
	var args map[string]interface{}
	json.Unmarshal([]byte(inputs[0].String()), &args)
	scalar := suite.Scalar()
	point1 := suite.Point()
	point2 := suite.Point()
	scalarB, _ := base64.StdEncoding.DecodeString(args["scalar"].(string))
	point1B, _ := base64.StdEncoding.DecodeString(args["point1"].(string))
	point2B, _ := base64.StdEncoding.DecodeString(args["point2"].(string))
	scalar.UnmarshalBinary(scalarB)
	point1.UnmarshalBinary(point1B)
	point2.UnmarshalBinary(point2B)
	point1.UnmarshalBinary([]byte(args["point1"].(string)))
	point2.UnmarshalBinary([]byte(args["point2"].(string)))
	var resultB []byte
	//var result kyber.Point
	for i := 0; i < 10000; i++ {
		//result = suite.Point().Mul(scalar, point1)
		suite.Point().Add(point1, point2)
	}
	resultB, _ = suite.Point().Mul(scalar, suite.Point().Add(point1, point2)).MarshalBinary()
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
