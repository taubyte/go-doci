package doci

import (
	"bytes"
	"errors"

	"github.com/multiformats/go-multihash"
)

func mhash_sha1(r string) []byte {
	mHashBuf, err := multihash.Sum([]byte(r), multihash.SHA1, -1)
	if err != nil {
		return nil
	}
	return mHashBuf
}

func hash(r string) ([]byte, error) {
	mHashBuf, err := multihash.Sum([]byte(r), multihash.SHA1, -1)
	if err != nil {
		return nil, err
	}
	return mHashBuf, nil
}

func check_hash(r string, h []byte) error {

	h0 := mhash_sha1(r)

	if bytes.Compare(h0, h) != 0 {
		return errors.New("Invalid signature")
	}

	return nil

}
