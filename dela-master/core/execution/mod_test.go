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

func TestExecute(t *testing.T) {
	step := Step{}
	args := map[string]interface{}{
		"counter": 0,
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
	require.Equal(t, "1", res.Message)
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
