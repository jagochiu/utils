package webtools

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"golang.org/x/net/proxy"
)

type Protocol string

const (
	TCP Protocol = "tcp"
	UDP          = "udp"
)

type Socket struct {
	Used     bool
	Protocol Protocol
	Host     string
	Port     int
}

func Socks5Dialer(protocol Protocol, host string, port int) proxy.Dialer {
	dialer, err := proxy.SOCKS5(string(protocol), fmt.Sprintf(`%s:%d`, host, port), nil, proxy.Direct)
	if err != nil {
		return nil
	}
	return dialer
}
func Socks5Transport(protocol Protocol, host string, port int) *http.Transport {
	dialer := Socks5Dialer(protocol, host, port)
	if dialer == nil {
		return nil
	}
	return &http.Transport{
		DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.Dial(network, address)
		},
		DisableKeepAlives: true,
	}
}
func Socks5Client(protocol Protocol, host string, port int) *http.Client {
	return &http.Client{
		Transport: Socks5Transport(protocol, host, port),
	}
}
