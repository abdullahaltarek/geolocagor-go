package main

import (
	"flag"
	"fmt"
	"net/http"
	"github.com/apex/log"
	"encoding/json"
	"io/ioutil"
	"time"
	"os"
	"github.com/alexsasharegan/dotenv"
)

type Info struct {
	Country string `json:"country"`
	IP      string `json:"query"`
	Status  string `json:"status"`
}

func main() {
	address := flag.String("a", "", "give the ip address to locate, putting me will locate you")
	flag.Parse()
	fmt.Println("Give me a second mercy...")
	time.Sleep(time.Second * 2)
	CheckConnection()
	fmt.Println("I'm trying my best to help you...")
	time.Sleep(time.Second * 2)
	if *address == "me" {
		info := Locate(*address)
		if info.Status == "fail" {
			fmt.Println("I believe you gave me a wrong IP :(")
			os.Exit(1)
		}
		fmt.Println("I got it :D You're trying to locate yourself, right? That's kewl :p")
		time.Sleep(time.Second * 2)
		fmt.Println("You're from " + info.Country + " and your IP is " + info.IP)
	} else {
		info := Locate(*address)
		if info.Status == "fail" {
			fmt.Println("I believe you gave a the wrong IP :(")
			os.Exit(1)
		}
		fmt.Println("Ok I got it :)")
		fmt.Println("The IP -> " + info.IP + " is from " + info.Country)
	}
}

func CheckConnection() {
	_, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("No internet buddy! :(")
		os.Exit(1)
	}
}

func Locate(address string) Info {
	dotenv.Load()
	if address == "me" {
		var info Info

		resp, err := http.Get(os.Getenv("LOCATION"))
		if err != nil {
			log.Fatal(err.Error())
		}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &info)
		return info
	} else {
		var info Info

		resp, err := http.Get(os.Getenv("LOCATION") + address)
		if err != nil {
			log.Fatal(err.Error())
		}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &info)
		return info
	}
}
