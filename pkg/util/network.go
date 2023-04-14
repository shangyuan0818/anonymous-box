package util

import (
	"errors"
	"net"
)

var (
	ErrNoLocalIPFound = errors.New("no local ip found")
)

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		switch address.(type) {
		case *net.IPNet:
			ip := address.(*net.IPNet).IP
			if ip.IsLoopback() {
				continue
			}

			if ip.To4() != nil {
				return ip.String(), nil
			}
		case *net.IPAddr:
			ip := address.(*net.IPAddr).IP
			if ip.IsLoopback() {
				continue
			}

			if ip.To4() != nil {
				return ip.String(), nil
			}
		default:
			continue
		}
	}

	return "", ErrNoLocalIPFound
}
