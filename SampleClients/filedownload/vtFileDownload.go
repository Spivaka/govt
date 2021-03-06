// vtFileDownload - fetches a sample from VirusTotal for the given resource. A resource can be MD5, SHA-1 or SHA-2 of a file.
//  vtFileDownload -rsrc=8ac31b7350a95b0b492434f9ae2f1cde
//
// This feature of the VirusTotal API is just available if you have a private API key.
// With a public API key you can not download samples.
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/slavikm/govt"
	"io/ioutil"
	"os"
)

var apikey string
var apiurl string
var rsrc string

// init - initializes flag variables.
func init() {
	flag.StringVar(&apikey, "apikey", os.Getenv("VT_API_KEY"), "Set environment variable VT_API_KEY to your VT API Key or specify on prompt")
	flag.StringVar(&apiurl, "apiurl", "https://www.virustotal.com/vtapi/v2/", "URL of the VirusTotal API to be used.")
	flag.StringVar(&rsrc, "rsrc", "8ac31b7350a95b0b492434f9ae2f1cde", "resource of file to retrieve report for. A resource can be md5, sha-1 or sha-2 sum of a file.")
}

// check - an error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()
	if rsrc == "" {
		fmt.Println("-rsrc=<md5|sha-1|sha-2> not given!")
		os.Exit(1)
	}
	c, err := govt.New(govt.SetApikey(apikey), govt.SetUrl(apiurl))
	check(err)

	// get a file report
	r, err := c.GetFile(rsrc)
	check(err)
	//fmt.Printf("r: %s\n", r)
	j, err := json.MarshalIndent(r, "", "    ")
	fmt.Printf("FileReport: ")
	os.Stdout.Write(j)
	//fmt.Printf("%d %s \t%s \t%s \t%d/%d\n", r.Status.ResponseCode, r.Status.VerboseMsg, r.Resource, r.ScanDate, r.Positives, r.Total)

	err = ioutil.WriteFile(rsrc, r.Content, 0600)
	fmt.Printf("file %s has been written.\n", rsrc)
	check(err)
}
