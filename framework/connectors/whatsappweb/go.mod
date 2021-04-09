module whatsappconn

go 1.16

require (
	github.com/Baozisoftware/qrcode-terminal-go v0.0.0-20170407111555-c0650d8dff0f
	github.com/Rhymen/go-whatsapp v0.1.1
	github.com/olivere/elastic/v7 v7.0.22
	monitor v1.0.0
	watsonconn v1.0.0
)

replace (
	watsonconn => ./../watson
	monitor => ./../../monitor
)
