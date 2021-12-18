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
		{"11.0.0.0/12", "0.0.0.0/0", true},
		{"192.168.0.0/24", "192.168.0.0/16", true},
	}
	for _, testCase := range cases {
		_, net1, _ := net.ParseCIDR(testCase.network1)
		_, net2, _ := net.ParseCIDR(testCase.network2)
		result := IsPartOf(net1, net2)
		n.Require().Equal(testCase.expected, result)
	}
}

func TestNetworkTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkTestSuite))
}
