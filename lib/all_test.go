package doci

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var value = "ping"
	var encoded = "pQFkZG9jaQIBEGRwaW5nGEBWERRXKYK7xPKe6Sri1lqe3CRT0skXDBg/AA=="

	ret, err := Encode(value, 0)
	if err != nil {
		t.Error(err)
		return
	}

	if ret != encoded {
		t.Errorf("Output of encoder does not match expected value: %s != %s [expected]", ret, encoded)
	}
}

func TestDecode(t *testing.T) {
	var value = "ping"
	var encoded = "pQFkZG9jaQIBEGRwaW5nGEBWERRXKYK7xPKe6Sri1lqe3CRT0skXDBg/AA=="

	ret, ts, err := Decode(encoded)
	if err != nil {
		t.Error(err)
		return
	}

	if ts != 0 {
		t.Errorf("Decoded Timestamp not as expected: %d != %d [expected]", ts, 0)
	}

	if ret != value {
		t.Errorf("Output of decoder does not match expected value: %s != %s [expected]", ret, encoded)
	}
}

func TestDNSDecode(t *testing.T) {
	var fqdn = "__elders__.net.taubyte.com"

	_, err := DecodeDnsTXT(fqdn, false)
	if err != nil {
		t.Error(err)
		return
	}

}
