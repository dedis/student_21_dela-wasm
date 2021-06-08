package main

import (
	"testing"

	"encoding/base64"
	"encoding/json"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/dela/core/execution"
	"go.dedis.ch/dela/core/store"
	"go.dedis.ch/dela/core/txn"
	"go.dedis.ch/kyber/v3/suites"
)

// go test -bench=.

const iterations = 1

var suite = suites.MustFind("Ed25519")

// Increment benchmark

func BenchmarkNative_Increment(b *testing.B) {
	for i := 0; i < iterations; i++ {
		k := 0
		k++
	}
}

func BenchmarkWASM_Go_Increment(b *testing.B) {
	n := iterations
	var counter = 0
	args := map[string]interface{}{
		"counter":          counter,
		"contractName":     "increaseCounter",
		"contractLanguage": "go",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		b.Error(err)
	}
	step := execution.Step{}
	step.Current = fakeTx{json: marsh}

	srvc := execution.WASMService{}
	for i := 0; i < n; i++ {
		_, err = srvc.Execute(nil, step)
		if err != nil {
			b.Logf("failed to execute: %+v", err)
			b.FailNow()
		}
	}
}

func BenchmarkWASM_C_Increment(b *testing.B) {
	n := iterations

	var counter = 0
	args := map[string]interface{}{
		"counter":          counter,
		"contractName":     "increaseCounter",
		"contractLanguage": "c",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		b.Error(err)
	}
	step := execution.Step{}
	step.Current = fakeTx{json: marsh}
	srvc := execution.WASMService{}
	for i := 0; i < n; i++ {
		_, err = srvc.Execute(nil, step)

		if err != nil {
			b.Logf("failed to execute: %+v", err)
			b.FailNow()
		}
	}
}

// Simple crypto (Elliptic curve - EC) benchmarks

func BenchmarkNative_EC(b *testing.B) {
	scalar := suite.Scalar().Pick(suite.RandomStream())
	_, err := scalar.MarshalBinary()
	require.NoError(b, err)

	for i := 0; i < iterations; i++ {
		point := suite.Point().Mul(scalar, nil)
		_, err = point.MarshalBinary()
		require.NoError(b, err)
	}
}

func BenchmarkWASM_Go_EC(b *testing.B) {

	scalar := suite.Scalar().Pick(suite.RandomStream())
	scalarB, _ := scalar.MarshalBinary()
	args := map[string]interface{}{
		"scalar":           base64.StdEncoding.EncodeToString(scalarB),
		"contractName":     "simpleEC",
		"contractLanguage": "go",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		b.Error(err)
	}
	step := execution.Step{}
	step.Current = fakeTx{json: marsh}

	srvc := execution.WASMService{}
	for i := 0; i < iterations; i++ {
		_, err = srvc.Execute(nil, step)
		if err != nil {
			b.Logf("failed to execute: %+v", err)
			b.FailNow()
		}
	}
}

func BenchmarkWASM_C_EC(b *testing.B) {
	var suite = suites.MustFind("Ed25519")
	step := execution.Step{}
	point1 := suite.Point().Pick(suite.RandomStream())
	scalar := suite.Scalar().Pick(suite.RandomStream())
	point1B, _ := point1.MarshalBinary()
	scalarB, _ := scalar.MarshalBinary()
	// encoding to base64 because JSON does not support raw bytes
	args := map[string]interface{}{
		"point1":           base64.StdEncoding.EncodeToString(point1B),
		"scalar":           base64.StdEncoding.EncodeToString(scalarB),
		"contractName":     "ed25519_gen_mul",
		"contractLanguage": "c",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		b.Error(err)
	}
	step.Current = fakeTx{json: marsh}

	srvc := execution.WASMService{}
	for i := 0; i < iterations; i++ {
		_, err = srvc.Execute(nil, step)
	}
}

func BenchmarkNative_Ed25519_Add(b *testing.B) {
	var suite = suites.MustFind("Ed25519")
	point1 := suite.Point().Pick(suite.RandomStream())
	point2 := suite.Point().Pick(suite.RandomStream())
	for i := 0; i < iterations; i++ {
		suite.Point().Add(point1, point2)
	}

}

func BenchmarkWASM_Go_Ed25519_Add(b *testing.B) {
	var suite = suites.MustFind("Ed25519")
	step := execution.Step{}
	point1 := suite.Point().Pick(suite.RandomStream())
	point2 := suite.Point().Pick(suite.RandomStream())
	point1B, _ := point1.MarshalBinary()
	point2B, _ := point2.MarshalBinary()
	// encoding to base64 because JSON does not support raw bytes
	args := map[string]interface{}{
		"point1":           base64.StdEncoding.EncodeToString(point1B),
		"point2":           base64.StdEncoding.EncodeToString(point2B),
		"contractName":     "ed25519",
		"contractLanguage": "go",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		b.Error(err)
	}
	step.Current = fakeTx{json: marsh}

	srvc := execution.WASMService{}
	for i := 0; i < iterations; i++ {
		_, err = srvc.Execute(nil, step)
	}
}

func BenchmarkWASM_C_Ed25519_Add(b *testing.B) {
	var suite = suites.MustFind("Ed25519")
	step := execution.Step{}
	point1 := suite.Point().Pick(suite.RandomStream())
	point2 := suite.Point().Pick(suite.RandomStream())
	point1B, _ := point1.MarshalBinary()
	point2B, _ := point2.MarshalBinary()
	// encoding to base64 because JSON does not support raw bytes
	args := map[string]interface{}{
		"point1":           base64.StdEncoding.EncodeToString(point1B),
		"point2":           base64.StdEncoding.EncodeToString(point2B),
		"contractName":     "ed25519",
		"contractLanguage": "c",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		b.Error(err)
	}
	step.Current = fakeTx{json: marsh}

	srvc := execution.WASMService{}
	for i := 0; i < iterations; i++ {
		_, err = srvc.Execute(nil, step)
	}
}

func BenchmarkNative_Ed25519_Mul(b *testing.B) {

	var suite = suites.MustFind("Ed25519")
	point1 := suite.Point().Pick(suite.RandomStream())
	scalar := suite.Scalar().Pick(suite.RandomStream())
	for i := 0; i < iterations; i++ {
		suite.Point().Mul(scalar, point1)
	}
}

func BenchmarkWASM_Go_Ed25519_Mul(b *testing.B) {

	var suite = suites.MustFind("Ed25519")
	step := execution.Step{}
	point1 := suite.Point().Pick(suite.RandomStream())
	scalar := suite.Scalar().Pick(suite.RandomStream())
	point1B, _ := point1.MarshalBinary()
	scalarB, _ := scalar.MarshalBinary()
	// encoding to base64 because JSON does not support raw bytes
	args := map[string]interface{}{
		"point1":           base64.StdEncoding.EncodeToString(point1B),
		"scalar":           base64.StdEncoding.EncodeToString(scalarB),
		"contractName":     "ed25519_mul",
		"contractLanguage": "go",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		b.Error(err)
	}
	step.Current = fakeTx{json: marsh}

	srvc := execution.WASMService{}
	for i := 0; i < iterations; i++ {
		_, err = srvc.Execute(nil, step)
	}
}

func BenchmarkWASM_C_Ed25519_Mul(b *testing.B) {

	var suite = suites.MustFind("Ed25519")
	step := execution.Step{}
	point1 := suite.Point().Pick(suite.RandomStream())
	scalar := suite.Scalar().Pick(suite.RandomStream())
	point1B, _ := point1.MarshalBinary()
	scalarB, _ := scalar.MarshalBinary()
	// encoding to base64 because JSON does not support raw bytes
	args := map[string]interface{}{
		"point1":           base64.StdEncoding.EncodeToString(point1B),
		"scalar":           base64.StdEncoding.EncodeToString(scalarB),
		"contractName":     "ed25519_mul",
		"contractLanguage": "c",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		b.Error(err)
	}
	step.Current = fakeTx{json: marsh}

	srvc := execution.WASMService{}
	for i := 0; i < iterations; i++ {
		_, err = srvc.Execute(nil, step)
	}
}

type inmemory struct {
	store.Readable
	store.Writable

	data map[string][]byte
}

func newInmemory() inmemory {
	return inmemory{
		data: make(map[string][]byte),
	}
}

func (i inmemory) Get(key []byte) ([]byte, error) {
	return i.data[string(key)], nil
}

func (i inmemory) Set(key []byte, value []byte) error {
	i.data[string(key)] = value
	return nil
}

func (i inmemory) Delete(key []byte) error {
	delete(i.data, string(key))
	return nil
}

type tx struct {
	txn.Transaction
	args map[string][]byte
}

func (t tx) GetArg(key string) []byte {
	return t.args[key]
}

// -----------------------------------------------------------------------------
// Utility functions

type fakeExec struct {
	err error
}

type fakeTx struct {
	txn.Transaction
	json []byte
}

func (tx fakeTx) GetArg(key string) []byte {
	return []byte(tx.json)
}
