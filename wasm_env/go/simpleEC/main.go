package main

import (
	"encoding/base64"
	"encoding/json"
	"syscall/js"

	"go.dedis.ch/kyber/v4/suites"
)

// GOOS=js GOARCH=wasm go build -o simpleEC.wasm

var c chan bool

func init() {
	c = make(chan bool)
}

// inputs should only contain one element, which is a JSON in string format.
func simpleEC(this js.Value, inputs []js.Value) interface{} {
	var suite = suites.MustFind("Ed25519")
	var args map[string]interface{}
	json.Unmarshal([]byte(inputs[0].String()), &args)
	scalar := suite.Scalar()
	scalarB, _ := base64.StdEncoding.DecodeString(args["scalar"].(string))
	scalar.UnmarshalBinary(scalarB)
	var resultB []byte
	for i := 0; i < 1; i++ {
		resultB, _ = suite.Point().Mul(scalar, nil).MarshalBinary()
	}
	args["result"] = base64.StdEncoding.EncodeToString(resultB)
	//args["resultTest"] = result.String()
	args["Accepted"] = "true"
	return args
}

func main() {
	js.Global().Set("simpleEC", js.FuncOf(simpleEC))
	// Force the program to stay open by never sending to channel c
	<-c
}
