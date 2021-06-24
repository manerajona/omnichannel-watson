# Whatsapp Web Connector

## handlers.go
The code implements the Handler interface of the unofficial go-whatsapp library (https://github.com/Rhymen/go-whatsapp). The handlers manage the Whatsapp connection and the message routing.

## login.go
The code implements the Whatsapp Web session management, establishing a connection with our account by scanning a qr code. It also allows saving the session in the temporary files of the system, avoiding having to scan the qr code again.
