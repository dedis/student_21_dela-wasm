// Package execution defines the service to execute a step in a validation
// batch.
//
// Documentation Last Review: 08.10.2020
//
package execution

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.dedis.ch/dela/core/store"

	"go.dedis.ch/dela/core/txn"
)

// Step is a context of execution. It allows for example a smart contract to
// execute a given transaction knowing what previous transactions have already
// been accepted and executed in a block.
type Step struct {
	Previous []txn.Transaction
	Current  txn.Transaction
}

// Result is the result of a transaction execution.
type Result struct {
	// Accepted is the success state of the transaction.
	Accepted bool

	// Message gives a change to the execution to explain why a transaction has
	// failed.
	Message string
}

// Service is the execution service that defines the primitives to execute a
// transaction.
type Service interface {
	// Execute must apply the transaction to the trie and return the result of
	// it.
	Execute(snap store.Snapshot, step Step) (Result, error)
}

func ExecuteBeta() (Result, error) {
	args := map[string]interface{}{
		"counter": 3,
		"xd":      "lol",
	}
	marsh, err := json.Marshal(args)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	responseBody := bytes.NewBuffer(marsh)
	resp, err := http.Post("http://127.0.0.1:3000/", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Default()
	json.Unmarshal(body, &args)
	return Result{true, fmt.Sprintf("%v", args["counter"])}, nil
}
