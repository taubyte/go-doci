package doci

type dociEntry struct {
	Proto     string `cbor:"1,keyasint"`
	Version   int    `cbor:"2,keyasint"`
	Data      string `cbor:"16,keyasint"`
	Hash      []byte `cbor:"64,keyasint"`
	Timestamp int64  `cbor:"63,keyasint"`
}
