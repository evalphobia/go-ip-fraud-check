package ipfraudcheck

import (
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	defaultWhoisServer  = "whois.radb.net"
	defaultWhoisPort    = 43
	defaultWhoisTimeout = 30 * time.Second
)

type WhoisClient struct {
	Server  string
	Port    int
	Timeout time.Duration
}

func NewWhoisClient() WhoisClient {
	return WhoisClient{
		Server:  defaultWhoisServer,
		Port:    defaultWhoisPort,
		Timeout: defaultWhoisTimeout,
	}
}

func (c WhoisClient) GetRoutes(asn int64) ([]string, error) {
	byt, err := whois(c.Server, c.Port, c.Timeout, asn)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(byt), "\n")
	routes := make([]string, 0, len(lines)/8)
	for _, line := range lines {
		if !strings.HasPrefix(line, "route") {
			continue
		}
		route := strings.TrimSpace(
			strings.TrimPrefix(
				strings.TrimPrefix(line, "route6:"),
				"route:"))
		routes = append(routes, route)
	}
	return routes, nil
}

// codes are base on: https://github.com/likexian/whois
func whois(server string, port int, timeout time.Duration, asn int64) ([]byte, error) {
	dialer := &net.Dialer{
		Timeout: timeout,
	}

	now := time.Now()
	conn, err := dialer.Dial("tcp", net.JoinHostPort(server, strconv.Itoa(port)))
	if err != nil {
		return nil, fmt.Errorf("whois: connect to whois server failed: %w", err)
	}
	defer conn.Close()

	// send query
	_ = conn.SetWriteDeadline(time.Now().Add(dialer.Timeout - time.Since(now)))
	query := getWhoisQuery(server, asn)
	_, err = conn.Write([]byte(query + "\r\n"))
	if err != nil {
		return nil, fmt.Errorf("whois: send to whois server failed: %w", err)
	}

	// get response
	_ = conn.SetReadDeadline(time.Now().Add(dialer.Timeout - time.Since(now)))
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, fmt.Errorf("whois: read from whois server failed: %w", err)
	}

	return buf, nil
}

func getWhoisQuery(server string, asn int64) string {
	switch server {
	case "whois.radb.net":
		return fmt.Sprintf("-i origin AS%d", asn)
	}
	return fmt.Sprintf("AS%d", asn)
}
