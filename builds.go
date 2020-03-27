package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strings"
)

//Builds XML mapping
type Builds struct {
	//XMLName xml.Name `xml:"builds"`
	Builds []Build `xml:"build"`
}

//Build XML mapping
type Build struct {
	//XMLName  xml.Name `xml:"build"`
	ID          string `xml:"id,attr"`
	BuildTypeID string `xml:"buildTypeId,attr"`
	Number      string `xml:"number,attr"`
	Status      string `xml:"status,attr"`
	State       string `xml:"state,attr"`
	BranchName  string `xml:"branchName,attr"`
	Href        string `xml:"href,attr"`
	WebURL      string `xml:"webUrl,attr"`
}

func main() {
	var url string
	var lgn string
	var since string
	var status string

	flag.StringVar(&url, "url", "", "TeamCity stem URL without 'https://'")
	flag.StringVar(&lgn, "login", "", "username:password")
	flag.StringVar(&since, "since", "202001", "ex 20200100 for 2020/01 00:00")
	flag.StringVar(&status, "status", "SUCCESS", "build complete STATUS")

	flag.Parse()

	fmt.Println("urlStem:", url)
	fmt.Println("login:", lgn)
	fmt.Println("since Date:", since)
	fmt.Println("status:", status)

	if url == "" || lgn == "" || since == "" || status == "" {
		panic("url and login are required")
	}

	var cred = strings.Split(lgn, ":")
	usr := cred[0]
	pwd := cred[1]
	fmt.Println("user:", usr)
	fmt.Println("password:", pwd)

	get := "https://" + url + "/httpAuth/app/rest/builds?locator=" + neturl.QueryEscape("branch:default:any,sinceDate:"+since+"00T000000+0000,status:"+status)

	fmt.Println("url:", get)

	req, _ := http.NewRequest("GET", get, nil)

	req.SetBasicAuth(usr, pwd)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println("___\n\n\n")

	var builds Builds

	xml.Unmarshal(body, &builds)

	out, _ := json.MarshalIndent(builds.Builds, "", "\t")

	fmt.Println(string(out))
}
