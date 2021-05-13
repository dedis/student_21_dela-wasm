package execution

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/dela/core/txn"
	"go.dedis.ch/kyber/v3/suites"
)

/*func TestExecuteBeta(t *testing.T) {
	var past = time.Now().Nanosecond()
	res, err := ExecuteBeta()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, "1", res.Message)
	t.Log("Time in milliseconds :")
	t.Log(float32(time.Now().Nanosecond()-past) / float32(1000000))
}*/

func TestIncreaseCounterGo(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	step := Step{}
	var counter = rand.Intn(100)
	args := map[string]interface{}{
		"counter":          counter,
		"contractName":     "increaseCounter",
		"contractLanguage": "go",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
	}
	step.Current = fakeTx{json: marsh}

	srvc := WASMService{}

	var past = time.Now()

	res, err := srvc.Execute(nil, step)

	duration := time.Since(past).Nanoseconds()
	//t.Log(res)
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, strconv.Itoa(counter+1), res.Message)
	t.Log("Time in milliseconds :")
	t.Log(float32(duration) / float32(1000000))
}

func TestIncreaseCounterC(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	step := Step{}
	var counter = rand.Intn(100)
	args := map[string]interface{}{
		"counter":          counter,
		"contractName":     "increaseCounter",
		"contractLanguage": "c",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
	}
	step.Current = fakeTx{json: marsh}

	srvc := WASMService{}

	var past = time.Now()

	res, err := srvc.Execute(nil, step)
	duration := time.Since(past).Nanoseconds()
	t.Log(res)
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, strconv.Itoa(counter+1), res.Message)
	t.Log("Time in milliseconds :")
	t.Log(float32(duration) / float32(1000000))
}

func TestCryptoOp(t *testing.T) {
	var suite = suites.MustFind("Ed25519")
	step := Step{}
	scalar := suite.Scalar().Pick(suite.RandomStream())
	point1 := suite.Point().Pick(suite.RandomStream())
	point2 := suite.Point().Pick(suite.RandomStream())
	resultB, _ := suite.Point().Mul(scalar, suite.Point().Add(point1, point2)).MarshalBinary()
	scalarB, _ := scalar.MarshalBinary()
	point1B, _ := point1.MarshalBinary()
	point2B, _ := point2.MarshalBinary()
	// encoding to base64 because JSON does not support raw bytes
	args := map[string]interface{}{
		"scalar":           base64.StdEncoding.EncodeToString(scalarB),
		"point1":           base64.StdEncoding.EncodeToString(point1B),
		"point2":           base64.StdEncoding.EncodeToString(point2B),
		"contractName":     "ed25519",
		"contractLanguage": "go",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
	}
	step.Current = fakeTx{json: marsh}

	srvc := WASMService{}

	past := time.Now()

	res, err := srvc.Execute(nil, step)
	duration := time.Since(past).Nanoseconds()
	t.Log(res)
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, base64.StdEncoding.EncodeToString(resultB), res.Message)
	t.Log("Time in milliseconds :")
	t.Log(float32(duration) / float32(1000000))
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
