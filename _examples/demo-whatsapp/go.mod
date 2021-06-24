module whatsapp

go 1.16

require (
	github.com/Rhymen/go-whatsapp v0.1.1
	whatsappconn v1.0.0
)

replace (
	watsonconn => ./../../framework/connectors/watson
	monitor => ./../../framework/monitor
	whatsappconn => ./../../framework/connectors/whatsappweb
)
