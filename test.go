package main

import(
	"fmt"
	"os"
	"github.com/bramvdbogaerde/go-scp/auth"
	"github.com/bramvdbogaerde/go-scp"
)

func main(){
	// Use SSH key authentication from the auth package
	clientConfig, _ := auth.PrivateKey("root", "pass.txt")

	// For other authentication methods see ssh.ClientConfig and ssh.AuthMethod
	// Create a new SCP client
	client := scp.NewClient("192.168.1.12:22", &clientConfig)

	// Connect to the remote server
	err := client.Connect()
	if err != nil{
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}

	// Open a file
	f, _ := os.Open("/Documents/readme.txt")

	// Close session after the file has been copied
	defer client.Session.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	client.CopyFile(f, "/home/root", "0655")
}
