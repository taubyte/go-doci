package doci

import (
	"encoding/base64"
	"fmt"

	"github.com/fxamacker/cbor/v2"
)

func Encode(data string, timestamp int64) (entry string, err error) {

	h, err := hash(data)

	if err != nil {
		return "", err
	}

	e := dociEntry{
		Proto:     protocolName,
		Version:   protocolVersion,
		Data:      data,
		Hash:      h,
		Timestamp: timestamp,
	}

	fmt.Println(e)

	b, err := cbor.Marshal(e)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func Decode(raw string) (data string, timestamp int64, err error) {
	b, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return "", -1, err
	}

	var e dociEntry
	err = cbor.Unmarshal(b, &e)
	if err != nil {
		return "", -1, err
	}

	fmt.Println(e)

	err = check_hash(e.Data, e.Hash)
	if err != nil {
		return "", -1, err
	}

	return e.Data, e.Timestamp, nil

}
