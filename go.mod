module github.com/replit/rrl

go 1.16

require (
	github.com/Azure/go-autorest/autorest/adal v0.9.22 // indirect
	github.com/coredns/caddy v1.1.1
	github.com/coredns/coredns v1.8.6
	github.com/miekg/dns v1.1.43
	github.com/prometheus/client_golang v1.11.0
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
)

replace github.com/DataDog/dd-trace-go v0.6.1 => github.com/datadog/dd-trace-go v0.6.1
