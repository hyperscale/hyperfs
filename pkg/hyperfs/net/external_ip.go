// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package net

import (
	"errors"
	stdnet "net"

	"github.com/rs/zerolog/log"
)

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() (stdnet.IP, error) {
	conn, err := stdnet.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return []byte{}, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*stdnet.UDPAddr)

	return localAddr.IP, nil
}

// LocalIP returns external ip
func LocalIP() (stdnet.IP, error) {
	addrs, err := stdnet.InterfaceAddrs()
	if err != nil {
		return []byte{}, err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*stdnet.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP, nil
			}
		}
	}

	return []byte{}, errors.New("are you connected to the network?")
}

// ExternalIP returns external ip
func ExternalIP() (stdnet.IP, error) {
	ifaces, err := stdnet.Interfaces()
	if err != nil {
		return []byte{}, err
	}

	for _, iface := range ifaces {
		if iface.Flags&stdnet.FlagUp == 0 {
			continue // interface down
		}

		if iface.Flags&stdnet.FlagLoopback != 0 {
			continue // loopback interface
		}

		log.Info().Msgf("Interface Name: %s", iface.Name)

		addrs, err := iface.Addrs()
		if err != nil {
			return []byte{}, err
		}

		for _, addr := range addrs {
			var ip stdnet.IP
			switch v := addr.(type) {
			case *stdnet.IPNet:
				ip = v.IP
				log.Info().Msgf("IPNet: %+v", v)
			case *stdnet.IPAddr:
				ip = v.IP
				log.Info().Msgf("IPAddr: %+v", v)
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}

			return ip, nil
		}
	}

	return []byte{}, errors.New("are you connected to the network?")
}
