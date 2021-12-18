package network

import (
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
