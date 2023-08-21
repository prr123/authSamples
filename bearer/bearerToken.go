// https://stackoverflow.com/questions/51452148/how-can-i-make-a-request-with-a-bearer-token-in-go

package main

import (
	"os"
	"fmt"
    "io/ioutil"
    "log"
    "net/http"

 	util "github.com/prr123/utility/utilLib"
)

func main() {

    numarg := len(os.Args)
    dbg := false
    flags:=[]string{"dbg","token"}

    // default file
    tokStr := ""

    useStr := "./bearerToken /token=tokStr [/dbg]"
    helpStr := "Client to test jwt tokens\n"

    if numarg > 3 {
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

    val, ok := flagMap["token"]
    if !ok {
        log.Fatalf(" error no token provided!\n")
    } else {
        if val.(string) == "none" {log.Fatalf("error: no token string provided!\n")}
        tokStr = val.(string)
    }

    log.Printf("token: %s debug: %t\n", tokStr, dbg)

    url := "http://89.116.30.49:12001"
	log.Printf("Destination: %s\n", url)

    // Create a Bearer string by appending string access token
    bearer := "Bearer " + tokStr

    // Create a new request using http
    req, err := http.NewRequest("GET", url, nil)

    // add authorization header to the req
    req.Header.Add("Authorization", bearer)

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
