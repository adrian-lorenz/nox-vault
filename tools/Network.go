package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"net/http"
)

func GetExtIP() (string, error) {
	url := "https://api.ipify.org"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil

}

func GetDnsIP(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return "", err
	}

	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String(), nil
		}
	}

	return "", fmt.Errorf("no IPv4 address found for %s", hostname)
}
func GetIP(c *gin.Context) string {
	i1 := c.Request.Header.Get("X-Forwarded-For")
	i2 := c.Request.RemoteAddr
	ip := i1
	if ip == "" {
		ip = i2
	}
	host, _, err := net.SplitHostPort(ip)
	if err != nil {
		return ip
	}
	return host
}
