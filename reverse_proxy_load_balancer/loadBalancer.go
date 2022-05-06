package reverse_proxy_load_balancer

import "fmt"

type RemoteHost struct {
	Scheme string
	Host   string
	Port   int
}

type LoadBalancer struct {
	Hosts       []RemoteHost
	currentHost int
}

func (remoteHost RemoteHost) ToURI() string {
	return fmt.Sprintf("%s://%s:%d", remoteHost.Scheme, remoteHost.Host, remoteHost.Port)
}

func (loadBalancer *LoadBalancer) NextHost() RemoteHost {
	if len(loadBalancer.Hosts) == 0 {
		return RemoteHost{}
	}

	loadBalancer.currentHost = (loadBalancer.currentHost + 1) % len(loadBalancer.Hosts)
	return loadBalancer.Hosts[loadBalancer.currentHost]
}
