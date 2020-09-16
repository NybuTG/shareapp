package ft

import (
	"io"
	"net"
	"os"
	"share/util"
	"strconv"
	"strings"
)

const BUFFERSIZE = 1024

func Send(ipa string) {
	connection, err := net.Dial("tcp", ipa)
	util.Check(err)
	defer connection.Close()
	
	go sendFile(connection)
}

func sendFile(connection net.Conn) {
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

func Receive() {
	server, err := net.Listen("tcp", ":8829")
	util.Check(err)
	defer server.Close()
	for {
		connection, err := server.Accept()
		util.Check(err)
		go getFile(connection)
	}
}

func getFile(connection net.Conn) {
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	// Process packets into a file
	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")
	newFile, err := os.Create(fileName)
	util.Check(err)
	defer newFile.Close()

	var receivedBytes int64

	for {
		if (fileSize - receivedBytes) < BUFFERSIZE {
			io.CopyN(newFile, connection, BUFFERSIZE)
			receivedBytes += BUFFERSIZE
		}
	}
}
