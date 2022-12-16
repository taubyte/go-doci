package doci

import (
	"net"
)

func DecodeDnsTXT(fqdn string, safe bool) ([]string, error) {
	txtrecords, err := net.LookupTXT(fqdn)

	if err != nil {
		return nil, err
	}

	var data []string

	for _, txt := range txtrecords {
		d, _, err := Decode(txt)
		if safe == false && err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	return data, nil
}
