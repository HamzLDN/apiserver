module examples/prometheus-metrics

go 1.16

replace github.com/yubo/apiserver => ../..

require (
	github.com/prometheus/client_golang v1.12.1
	github.com/yubo/apiserver v0.0.0-00010101000000-000000000000
	github.com/yubo/golib v0.0.3-0.20220902030005-7f15ca001a44
)
