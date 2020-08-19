package iprange

import "net"

func nextIP(ip net.IP, inc uint) net.IP {
	i := ip.To4()
	vprep0 := uint(i[0]) << 24
	vprep1 := uint(i[1]) << 16
	vprep2 := uint(i[2]) << 8
	vprep3 := uint(i[3])
	v := vprep0 + vprep1 + vprep2 + vprep3
	println("vprep0 ", vprep0, " vprep1 ", vprep1, " vprep2 ", vprep2, " vprep3 ", vprep3)
	println("v ", v)
	v += inc
	println(" inc v ", v)
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)

	return net.IPv4(v0, v1, v2, v3)
}
