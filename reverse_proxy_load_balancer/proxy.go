package reverse_proxy_load_balancer

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

var loadBalancer = LoadBalancer{
	Hosts: []RemoteHost{
		{
			Scheme: "http",
			Host:   "localhost",
			Port:   8081,
		},
		{
			Scheme: "http",
			Host:   "localhost",
			Port:   8082,
		},
		{
			Scheme: "http",
			Host:   "localhost",
			Port:   8083,
		},
	},
}

func getProxyURL(c *gin.Context) (*url.URL, error) {
	remoteHost := loadBalancer.NextHost()
	proxy, err := url.Parse(remoteHost.ToURI())
	if err != nil {
		fmt.Printf("Error parsing reverse proxy address: %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Error parsing reverse proxy address",
			"error":   err.Error(),
		})
		return nil, err
	}
	return proxy, nil
}

func execProxyRequest(c *gin.Context, req *http.Request) (*http.Response, error) {
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Error making request",
			"error":   err.Error(),
		})
		return nil, err
	}
	return resp, nil
}

func returnReverseProxyResponse(c *gin.Context, resp *http.Response) {
	for headerKey, headerValues := range resp.Header {
		for _, headerValue := range headerValues {
			c.Header(headerKey, headerValue)
		}
	}
	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(c.Writer)
}

func Proxy(c *gin.Context) {
	req := c.Request
	proxy, err := getProxyURL(c)
	if err != nil {
		return
	}

	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host

	resp, err := execProxyRequest(c, req)
	if err != nil {
		return
	}

	returnReverseProxyResponse(c, resp)
	return
}
