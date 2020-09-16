package filetransfer

import (
	"io"
	"net"
	"os"
	"share/util"
	"strconv"
)

const BUFFERSIZE = 1024

func Run(ipa string) {
	connection, err := net.Dial("tcp", ipa)
	util.Check(err)
	defer connection.Close()
	
	go sendFileToClient(connection)
}

func sendFileToClient(connection net.Conn) {
	defer connection.Close()

	file, err := os.Open("dummy.dat")
	util.Check(err)

	fileInfo, err := file.Stat()
	util.Check(err)
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)

	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
	// TODO: connection.Write([]byte(fileHash))
	sendBuffer := make([]byte, BUFFERSIZE)

	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}
	return
}

func fillString(returnString string, toLength int) (string) {
	for {
		lengthString := len(returnString)
		if lengthString < toLength {
			returnString = returnString + ":"
			continue
		}
		break
	}
	return returnString
}

func getFile(ipa string) {
	server, err := net.Listen("tcp", ipa)
	util.Check(err)
}
