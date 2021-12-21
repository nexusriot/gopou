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

// getPrivateRanges gets slice of private ranges in the string format
func getPrivateRanges() []string {
	return []string{
		"127.0.0.0/8",
		"192.168.0.0/16",
		"172.16.0.0/12",
		"10.0.0.0/8",
		"169.254.0.0/16", // RFC3927
		"::1/128",        // IPv6 lo
		"fe80::/10",      // IPv6 link local
		"fc00::/7",
	}
}

func getPrivateNetworks() []*net.IPNet {
	var privateNetworks []*net.IPNet

	for _, a := range getPrivateRanges() {
		_, addr, _ := net.ParseCIDR(a)
		privateNetworks = append(privateNetworks, addr)
	}
	return privateNetworks
}

// IsPrivateNetwork gets is networkAddr in private address range
func IsPrivateNetwork(networkAddr string) (bool, error) {

	_, addr, err := net.ParseCIDR(networkAddr)
	if err != nil {
		return false, fmt.Errorf("can not parse %s as network address", networkAddr)
	}
	for _, privateNet := range getPrivateNetworks() {
		if IsPartOf(addr, privateNet) {
			return true, nil
		}
	}
	return false, nil
}

// IsPrivateIP gets is given net.IP private
func IsPrivateIP(ip net.IP) bool {
	if ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsLoopback() {
		return true
	}
	for _, block := range getPrivateNetworks() {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

// Hosts gets slice of hosts with str format by cidr in string format
func Hosts(cidr string) ([]string, error) {
	var ips []string
	ip, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	increment := func(ip net.IP) {
		for j := len(ip) - 1; j >= 0; j-- {
			ip[j]++
			if ip[j] > 0 {
				break
			}
		}
	}
	for ip := ip.Mask(network.Mask); network.Contains(ip); increment(ip) {
		ips = append(ips, ip.String())
	}
	return ips, nil
}
