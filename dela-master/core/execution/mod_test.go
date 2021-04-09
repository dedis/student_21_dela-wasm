package execution

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/dela/core/txn"
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
	step := Step{}
	args := map[string]interface{}{
		"counter":          10,
		"contractName":     "increaseCounter",
		"contractLanguage": "go",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
	}
	step.Current = fakeTx{json: marsh}

	srvc := WASMService{}

	var past = time.Now().Nanosecond()

	res, err := srvc.Execute(nil, step)
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, "11", res.Message)
	t.Log("Time in milliseconds :")
	t.Log(float32(time.Now().Nanosecond()-past) / float32(1000000))
}

func TestIncreaseCounterC(t *testing.T) {
	step := Step{}
	args := map[string]interface{}{
		"counter":          2,
		"contractName":     "increaseCounter",
		"contractLanguage": "c",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		t.Error(err)
	}
	step.Current = fakeTx{json: marsh}

	srvc := WASMService{}

	var past = time.Now().Nanosecond()

	res, err := srvc.Execute(nil, step)
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, "3", res.Message)
	t.Log("Time in milliseconds :")
	t.Log(float32(time.Now().Nanosecond()-past) / float32(1000000))
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
