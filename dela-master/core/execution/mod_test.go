package execution

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	var past = time.Now().Nanosecond()
	res, err := ExecuteBeta()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, "4", res.Message)
	t.Log("Time in milliseconds :")
	t.Log(float32(time.Now().Nanosecond()-past) / float32(1000000))
}
