package network

import (
	"net"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NetworkTestSuite struct {
	suite.Suite
}

func (n *NetworkTestSuite) TestIsPartOf() {

	cases := []struct {
		network1 string
		network2 string
		expected bool
	}{
		{"5.0.0.0/12", "0.0.0.0/0", true},
		{"192.168.0.0/24", "192.168.0.0/16", true},
		{"10.0.100.0/24", "10.0.0.0/8", true},
		{"10.0.10.0/24", "10.0.11.0/24", false},
		{"10.0.0.0/8", "10.1.0.0/12", false},
	}
	for _, testCase := range cases {
		_, net1, _ := net.ParseCIDR(testCase.network1)
		_, net2, _ := net.ParseCIDR(testCase.network2)
		result := IsPartOf(net1, net2)
		n.Require().Equal(testCase.expected, result)
	}
}

func (n *NetworkTestSuite) TestIntersect() {

	cases := []struct {
		network1 string
		network2 string
		expected bool
	}{
		{"192.168.0.0/24", "192.168.1.0/24", false},
		{"192.168.1.0/24", "192.168.1.12/32", true},
		{"192.168.0.0/16", "192.168.0.0/15", true},
		{"10.0.0.0/16", "10.0.0.1/32", true},
	}
	for _, testCase := range cases {
		_, net1, _ := net.ParseCIDR(testCase.network1)
		_, net2, _ := net.ParseCIDR(testCase.network2)
		result := Intersect(net1, net2)
		n.Require().Equal(testCase.expected, result)
	}
}

func (n *NetworkTestSuite) TestNextIP() {

	cases := []struct {
		ipAddress string
		inc       uint
		expected  net.IP
	}{
		{"192.168.0.1", 2, net.ParseIP("192.168.0.3")},
		{"10.0.0.10", 200, net.ParseIP("10.0.0.210")},
		{"10.0.0.47", 300, net.ParseIP("10.0.1.91")},
		{"192.168.0.1", 1024, net.ParseIP("192.168.4.1")},
		{"1.0.0.0", 16777216, net.ParseIP("2.0.0.0")},
	}
	for _, testCase := range cases {
		ip := net.ParseIP(testCase.ipAddress)
		result := NextIP(ip, testCase.inc)
		n.Require().Equal(testCase.expected, result)
	}
}

func (n *NetworkTestSuite) TestIsNetworkPrivate() {

	testCases := []struct {
		netAddr string
		result  bool
		err     bool
	}{
		{"127.0.0.0/16", true, false},
		{"169.254.1.0/24", true, false},
		{"10.0.0.1/32", true, false},
		{"10.0.0.0/7", false, false},
		{"192.168.20.0/24", true, false},
		{"8.132.1.0/24", false, false},
		{"192.168.0.0/14", false, false},
		{"4.1.22.0/35", false, true},
		{"256.1.2.3/32", false, true},
	}
	for _, t := range testCases {
		res, err := IsNetworkPrivate(t.netAddr)
		n.Require().Equal(t.result, res)
		if !t.err {
			n.Require().Nil(err)
		} else {
			n.Require().NotNil(err)
		}
	}
}

func TestNetworkTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkTestSuite))
}
