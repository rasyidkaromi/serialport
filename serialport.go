package serialport

import (
	"log"
)

type Connection struct{
	mode  *Mode
	port  Port
}

func NewConnection (portName string, mode *Mode) *Connection{
	connection := &Connection{}
	port, err := Open(portName, mode)

	if err == nil{
		connection.mode = mode
		connection.port = port

		connection.Write("Terhubung")
		log.Println ("Terhubung")
	}else{
		log.Fatal(err)
	}

	return connection
}

func (connection *Connection) Write (text string) bool{
	_, err := connection.port.Write([]byte(text))

	return err == nil
}

func (connection *Connection) Read (buffer *[]byte) (int, []byte, string, bool){
	n, err := connection.port.Read(*buffer)

	if err != nil{
		log.Fatal(err)
	}

	if n == 0{
		log.Println("\nEOF")
	}

	return n, *buffer, string(*buffer), err == nil
}

func (connection *Connection) Close (){
	connection.port.Close()
}
