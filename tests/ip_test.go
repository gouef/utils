package tests

import (
	"fmt"
	"github.com/gouef/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIp2long(t *testing.T) {
	tests := []struct {
		ip       string
		expected uint32
		err      error
	}{
		{"192.168.1.1", 3232235777, nil},                                          // Standardní IPv4 adresa
		{"255.255.255.255", 4294967295, nil},                                      // Maximalní IPv4 adresa
		{"0.0.0.0", 0, nil},                                                       // Minimální IPv4 adresa
		{"invalid_ip", 0, fmt.Errorf("invalid IP address: invalid_ip")},           // Neplatná adresa
		{"256.256.256.256", 0, fmt.Errorf("invalid IP address: 256.256.256.256")}, // Neplatná adresa
		{"fe80::1", 0, fmt.Errorf("not an IPv4 address: fe80::1")},                // IPv6 adresa
	}

	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			result, err := utils.Ip2long(tt.ip)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestLong2ip(t *testing.T) {
	tests := []struct {
		ipLong   uint32
		expected string
	}{
		{3232235777, "192.168.1.1"},
		{4294967295, "255.255.255.255"},
		{0, "0.0.0.0"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.ipLong), func(t *testing.T) {
			result := utils.Long2ip(tt.ipLong)
			assert.Equal(t, tt.expected, result)
		})
	}
}
