package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SimpleContract struct {
	contractapi.Contract
}

func (sc *SimpleContract) Create(ctx CustomTransactionContextInterface, key string, value string) error {
	// existing, err := ctx.GetStub().GetState(key)
	// if err != nil {
	// 	return errors.New("Unable to interact with World State")
	// }
	existing := ctx.GetData()

	if existing != nil {
		return fmt.Errorf("Cannot create world state pair with key %s, Already exists", key)
	}
	err := ctx.GetStub().PutState(key, []byte(value))
	if err != nil {
		return errors.New("Unable to interact with World State")
	}
	return nil
}

func (sc *SimpleContract) Update(ctx CustomTransactionContextInterface, key string, value string) error {
	// existing, err := ctx.GetStub().GetState(key)
	// if err != nil {
	// 	return errors.New("Unable to interact with World State")
	// }

	existing := ctx.GetData()

	if existing == nil {
		return fmt.Errorf("Cannot update world state pair with key %s. Does not exist", key)
	}
	err := ctx.GetStub().PutState(key, []byte(value))
	if err != nil {
		return errors.New("Unable to interact with World State")
	}
	return nil
}

func (sc *SimpleContract) Read(ctx CustomTransactionContextInterface, key string) (string, error) {
	existing, err := ctx.GetStub().GetState(key)
	if err != nil {
		return "", errors.New("Unable to interact with World State")
	}
	if existing == nil {
		return "", fmt.Errorf("Unable to read %s key. Does not exist", key)
	}
	return string(existing), nil
}
