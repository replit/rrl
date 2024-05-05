module github.com/replit/rrl

go 1.16

require (
	github.com/DataDog/datadog-go v4.4.0+incompatible // indirect
	github.com/coredns/caddy v1.1.1
	github.com/coredns/coredns v1.11.2
	github.com/miekg/dns v1.1.58
	github.com/prometheus/client_golang v1.17.0
	k8s.io/api v0.28.4
	k8s.io/apimachinery v0.28.4
	k8s.io/client-go v0.28.3
)

replace github.com/DataDog/dd-trace-go v0.6.1 => github.com/datadog/dd-trace-go v0.6.1
