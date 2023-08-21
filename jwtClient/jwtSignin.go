// https://stackoverflow.com/questions/51452148/how-can-i-make-a-request-with-a-bearer-token-in-go
// jwtSignin
// program that signs in a user
// author prr, azul software
// date: 21 August 2023
// copyright 2023 prr, azul software
//

package main

import (
	"os"
	"fmt"
    "io/ioutil"
    "log"
	"bytes"
    "net/http"

	"github.com/goccy/go-json"
 	util "github.com/prr123/utility/utilLib"

)

type userAuth struct {
	User string `json:"username"`
	Pwd string `json:"password"`
}

func main() {

    numarg := len(os.Args)
    dbg := false
    flags:=[]string{"dbg","user", "pass"}

    // default file
    usrStr := ""
	pwdStr := ""

    useStr := "./jwtSignin /user=usr /pass=passwd [/dbg]"
    helpStr := "Client to test the signin process\n"

    if numarg > 4 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s", useStr)
        os.Exit(-1)
    }

    if numarg > 1 && os.Args[1] == "help" {
        fmt.Printf("help: %s\n", helpStr)
        fmt.Printf("usage is: %s\n", useStr)
        os.Exit(1)
    }

    flagMap, err := util.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    _, ok := flagMap["dbg"]
    if ok {dbg = true}
    if dbg {
        fmt.Printf("dbg -- flag list:\n")
        for k, v :=range flagMap {
            fmt.Printf("  flag: /%s value: %s\n", k, v)
        }
    }

    val, ok := flagMap["user"]
    if !ok {
        log.Fatalf(" error no user provided!\n")
    } else {
        if val.(string) == "none" {log.Fatalf("error: no user name provided!\n")}
        usrStr = val.(string)
    }

    pwdval, ok := flagMap["pass"]
    if !ok {
        log.Fatalf(" error no pass flag provided!\n")
    } else {
        if pwdval.(string) == "none" {log.Fatalf("error: no password provided!\n")}
        pwdStr = pwdval.(string)
    }

    log.Printf("debug: %t\n", dbg)
    log.Printf("usr: %s password: %s\n", usrStr, pwdStr)

    url := "http://89.116.30.49:12001/signin"
	log.Printf("Destination: %s\n", url)

	usrData := userAuth {
		User: usrStr,
		Pwd: pwdStr,
	}

	if dbg {fmt.Printf("userData %v\n", usrData)}

    // Create a Bearer string by appending string access token
//	tokenStr := "abcdefghijklmnop"
//    bearer := "Bearer " + tokenStr

	jsonBody, err := json.Marshal(usrData)
	if err != nil {
		log.Fatalf("error marshal jsn: %v\n", err)
	}
 	bodyReader := bytes.NewReader(jsonBody)

    // Create a new request using http
    req, err := http.NewRequest("GET", url, bodyReader)

    // add authorization header to the req
//    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}

	// synchronous client is blocked until response comes
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Error on response: %v\n", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes: %v", err)
    }
    log.Printf("resp body: %s\n", string([]byte(body)))
}
