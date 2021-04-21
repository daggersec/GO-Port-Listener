// Will Listen on a spcified range of ports with HTTP service
// on connect will echo time, server ip:port, client ip:port
// to compile into windows or linux binary:
// 	env GOOS=target-OS GOARCH=target-architecture go build package-import-path
// 	GOOS = linux, windows, android, darwin
// 	GOARCH = 386, amd64 arm, arm64

package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func createServer(name string, port int) *http.Server {

	xvar := http.NewServeMux()
	xvar.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t := time.Now()
		stamp := t.Format("2006-01-02 15:04:05")
		fmt.Fprint(res, stamp+" - server: "+req.Host+" - client: "+req.RemoteAddr)
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: xvar,
	}

	return &server
}

func main() {
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}

	//startport:=8000
	//endport:=8030
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// Intro

	fmt.Print(" _____ _____    _____         _      __    _     _                   \n")
	fmt.Print("|   __|     |  |  _  |___ ___| |_   |  |  |_|___| |_ ___ ___ ___ ___ \n")
	fmt.Print("|  |  |  |  |  |   __| . |  _|  _|  |  |__| |_ -|  _| -_|   | -_|  _|\n")
	fmt.Print("|_____|_____|  |__|  |___|_| |_|    |_____|_|___|_| |___|_|_|___|_|  \n\n")
	fmt.Print("GO Port Listener\n")
	fmt.Print("Ports below 1024 will require admin/sudo priveleges\n")

	// get port numbers
	fmt.Print("Enter First port: ")
	var startport int
	fmt.Scan(&startport)
	fmt.Print("Enter Last port: ")
	var endport int
	fmt.Scan(&endport)
	var portnum int
	portnum = endport - startport

	// error checking
	if startport > endport {
		fmt.Print("end port must be greater than start port\n")
		os.Exit(3)
	} else if endport > 65535 {
		fmt.Print("end port cannot be greater than 65535\n")
		os.Exit(3)
	} else if portnum > 9867 {
		fmt.Print("cannot open more than 9868 ports")
		os.Exit(3)
	} else if portnum > 1000 {
		fmt.Print("Opening more than a 1000 ports\n")
		fmt.Print("Are you sure ? (y/n)")
		var portnumok string
		fmt.Scan(&portnumok)
		switch portnumok {
		case "n":
			os.Exit(3)
		case "y":
			fmt.Print("Opening more than 1000 ports")
		}
	}

	fmt.Print("opening TCP ports " + strconv.Itoa(startport) + " through " + strconv.Itoa(endport) + "\n")
	fmt.Print("On ip address: " + ip + "\n")
	fmt.Print("Ctrl+C to exit\n")

	for i := startport; i < endport+1; i++ {
		httport := i
		go func() {
			server := createServer(string(httport), httport)
			fmt.Println(server.ListenAndServe())
			wg.Done()
		}()
	}
	wg.Wait()

}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
