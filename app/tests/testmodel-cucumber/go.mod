module testmodel

go 1.16

require (
	github.com/cucumber/godog v0.11.0-rc1
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.6.1
	watsonconn v1.0.0
)

replace watsonconn => ./../../../framework/connectors/watson
