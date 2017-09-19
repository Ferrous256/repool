package main

import(
	"fmt"
	"os"
	"golang.org/x/crypto/ssh"
	"github.com/bramvdbogaerde/go-scp"
	"net"
)

func main(){
	// Use SSH key authentication from the auth package
	clientConfig:= 	ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("admin"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

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
	f, _ := os.Open("poop.txt")

	fmt.Println("file opened")

	// Close session after the file has been copied
	defer client.Session.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	client.CopyFile(f, "poop.txt", "0655")
	fmt.Println("poop.txt copied")
}
