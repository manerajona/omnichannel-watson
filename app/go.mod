module github.com/manerajona/omnichannel-watson

go 1.16

require (
	watsonconn v1.0.0
	whatsappconn v1.0.0
)

replace (
	whatsappconn => ./../framework/connectors/whatsappweb
	watsonconn => ./../framework/connectors/watson
	monitor => ./../framework/monitor
)
