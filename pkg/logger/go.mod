module github.com/antonioducs/wyd/logger

go 1.25.7

require (
	github.com/antonioducs/wyd/pkg/configs v0.0.0
	github.com/lmittmann/tint v1.0.0
	github.com/mattn/go-isatty v0.0.20
)

replace github.com/antonioducs/wyd/pkg/configs => ../configs
