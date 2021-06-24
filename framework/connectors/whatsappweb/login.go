package whatsappconn

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	whatsapp "github.com/Rhymen/go-whatsapp"
)

// New whatsappweb connection
func New() *whatsapp.Conn {
	wac, err := whatsapp.NewConn(10 * time.Second)
	if err != nil {
		log.Fatalf("creating connection failed due %v\n", err)
	}
	return wac
}

// LoginOrRestore into whatsappweb web.
// Returns an error.
func LoginOrRestore(wac *whatsapp.Conn) {

	if err := login(wac); err != nil {
		log.Fatalf("logging failed due %v\n", err)
	}

	//verifies phone connectivity
	pong, err := wac.AdminTest()

	if !pong || err != nil {
		log.Fatalf("pinging failed due %v\n", err)
	}
}

//Disconect safe
func Disconect(wac *whatsapp.Conn) {
	session, err := wac.Disconnect()
	if err != nil {
		log.Fatalf("disconnecting failed due %v\n", err)
	}
	if err := writeSession(session); err != nil {
		log.Fatalf("saving session failed due %v", err)
	}
}

func login(wac *whatsapp.Conn) error {
	//load saved session
	session, err := readSession()
	if err == nil {
		//restore session
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return fmt.Errorf("restoring failed due %v", err)
		}
	} else {
		//no saved session -> regular login
		qr := make(chan string)
		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()
		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("fail during login due %v", err)
		}
	}

	//save session
	err = writeSession(session)
	if err != nil {
		return fmt.Errorf("saving session failed due %v", err)
	}
	return nil
}

func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

// WriteSession method save the Whatsapp session in /temp.
// Returns an error.
func writeSession(session whatsapp.Session) error {
	file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}
