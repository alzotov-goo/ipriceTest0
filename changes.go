package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//Changes XML mapping
type Changes struct {
	//XMLName xml.Name `xml:"changes"`
	Changes []Change `xml:"change"`
}

//Change XML mapping
type Change struct {
	//XMLName  xml.Name `xml:"change"`
	ID       string `xml:"id,attr"`
	Version  string `xml:"version,attr"`
	Username string `xml:"username,attr"`
	Date     string `xml:"date,attr"`
	Href     string `xml:"href,attr"`
	WebURL   string `xml:"webUrl,attr"`
}

func main() {
	var url string
	var lgn string
	var bld string

	flag.StringVar(&url, "url", "", "TeamCity stem URL without 'https://'")
	flag.StringVar(&lgn, "login", "", "username:password")
	flag.StringVar(&bld, "buildId", "", "buildId")

	flag.Parse()

	fmt.Println("urlStem:", url)
	fmt.Println("login:", lgn)
	fmt.Println("buildId:", bld)

	if url == "" || lgn == "" || bld == "" {
		panic("all params are required")
	}

	var cred = strings.Split(lgn, ":")
	usr := cred[0]
	pwd := cred[1]
	fmt.Println("user:", usr)
	fmt.Println("password:", pwd)

	get := "https://" + url + "/httpAuth/app/rest/changes?locator=build%3A(id%3A" + bld + ")"
	fmt.Println("url:", get)

	req, _ := http.NewRequest("GET", get, nil)

	req.SetBasicAuth(usr, pwd)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println("___\n\n\n")

	var changes Changes

	xml.Unmarshal(body, &changes)

	p, _ := json.MarshalIndent(changes.Changes, "", "\t")

	fmt.Println(string(p))

}
