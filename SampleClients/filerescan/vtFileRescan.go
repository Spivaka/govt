// vtFileRescan - asks VirusTotal to rescan a given resource.
// Resource can be a MD5, SHA-1 or SHA-2of a file.
//  vtFileRescan -rsrc=8ac31b7350a95b0b492434f9ae2f1cde
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/slavikm/govt"
	"os"
)

var apikey string
var apiurl string
var rsrc string

// init - initializes flag variables.
func init() {
	flag.StringVar(&apikey, "apikey", os.Getenv("VT_API_KEY"), "Set environment variable VT_API_KEY to your VT API Key or specify on prompt")
	flag.StringVar(&apiurl, "apiurl", "https://www.virustotal.com/vtapi/v2/", "URL of the VirusTotal API to be used.")
	flag.StringVar(&rsrc, "rsrc", "8ac31b7350a95b0b492434f9ae2f1cde", "md5 sum of a file to as VT about.")
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
		fmt.Println("-rsrc=<md5|sha1|sha2> fehlt!")
		os.Exit(1)
	}
	c, err := govt.New(govt.SetApikey(apikey), govt.SetUrl(apiurl))
	check(err)

	r, err := c.RescanFile(rsrc)
	check(err)
	j, err := json.MarshalIndent(r, "", "    ")
	fmt.Printf("FileReport: ")
	os.Stdout.Write(j)
}
