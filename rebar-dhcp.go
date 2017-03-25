package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/gcfg.v1"
)

type Config struct {
	Network struct {
		Port     int
		Username string
		Password string
	}
}

var ignore_anonymus bool
var config_path, key_pem, cert_pem, data_dir string
var server_ip string
var net_name string

func init() {
	// path := os.Executable()
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	// fmt.Println(dir)
	flag.StringVar(&config_path, "config_path", dir+"/config/config.cfg", "Path to config file")
	flag.StringVar(&key_pem, "key_pem", "", "Path to key file")
	flag.StringVar(&cert_pem, "cert_pem", "", "Path to cert file")
	flag.StringVar(&data_dir, "data_dir", dir+"/data", "Path to store data")
	// flag.StringVar(&server_ip, "server_ip", "", "Server IP to return in packets (e.g. 10.10.10.1/24)")
	// flag.StringVar(&server_ip, "s", "", "Server IP to return in packets (e.g. 10.10.10.1/24)")
	flag.StringVar(&net_name, "n", "eth0", "the network interface name default is 'eth0'")
	flag.StringVar(&net_name, "net", "eth0", "the network interface name default is 'eth0'")
	flag.BoolVar(&ignore_anonymus, "ignore_anonymus", false, "Ignore unknown MAC addresses")
}

func main() {
	flag.Parse()

	var cfg Config
	cerr := gcfg.ReadFileInto(&cfg, config_path)
	if cerr != nil {
		log.Fatal(cerr)
	}
	fs, err := NewFileStore(data_dir + "/database.json")
	if err != nil {
		log.Fatal(err)
	}

	fe := NewFrontend(cert_pem, key_pem, cfg, fs)

	if err := StartDhcpHandlers(fe.DhcpInfo, net_name); err != nil {
		log.Fatal(err)
	}
	fe.RunServer(true)
}
