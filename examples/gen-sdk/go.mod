module examples/gen-sdk

go 1.16

replace github.com/yubo/apiserver => ../..

require (
	github.com/yubo/apiserver v0.0.0-00010101000000-000000000000
	github.com/yubo/golib v0.0.3-0.20220825061925-f4cd420e40b5
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c
)
