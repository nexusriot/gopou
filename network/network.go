package network

import (
	"fmt"
	"net"
)

// Intersect is network1 intersects network2 takes net1, net2 *net.IPNet arguments
func Intersect(net1, net2 *net.IPNet) bool {
	return net2.Contains(net1.IP) || net1.Contains(net2.IP)
}

// IsPartOf calculates is net1 is a part of net2
func IsPartOf(net1, net2 *net.IPNet) bool {
	net1MaskSize, _ := net1.Mask.Size()
	net2MaskSize, _ := net2.Mask.Size()
	return net2.Contains(net1.IP) && net1MaskSize >= net2MaskSize
}

// NextIP gives the next Ip address of ip incremented by increment
func NextIP(ip net.IP, increment uint) net.IP {
	i4 := ip.To4()
	vb := uint(i4[0])<<24 + uint(i4[1])<<16 + uint(i4[2])<<8 + uint(i4[3]) + increment
	return net.IPv4(byte((vb>>24)&0xFF), byte((vb>>16)&0xFF), byte((vb>>8)&0xFF), byte(vb&0xFF))
}

func IsNetworkPrivate(networkAddr string) (bool, error) {
	var privateNetworks []*net.IPNet

	privateRanges := []string{
		"192.168.0.0/16",
		"172.16.0.0/12",
		"10.0.0.0/8",
	}
	for _, a := range privateRanges {
		_, addr, _ := net.ParseCIDR(a)
		privateNetworks = append(privateNetworks, addr)
	}
	_, addr, err := net.ParseCIDR(networkAddr)
	if err != nil {
		return false, fmt.Errorf("can not parse %s as network address", networkAddr)
	}
	for _, privateNet := range privateNetworks {
		if IsPartOf(addr, privateNet) {
			return true, nil
		}
	}
	return false, nil
}
