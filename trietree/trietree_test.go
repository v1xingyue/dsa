package trietree

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestIpMatch(t *testing.T) {

	dbfile := "db.txt"
	ipTree := new(IpTree)
	c, err := ioutil.ReadFile(dbfile)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(c), "\n")
	log.Println("load ip list from ", dbfile, len(lines))
	for _, l := range lines {
		if l != "" {
			//log.Println("before add ip ", l)
			ipTree.addIp(l)
		}
	}
	for _, l := range lines {
		if l != "" {
			log.Println("find ip ", ipTree.findIp(l), l)
		}
	}

	log.Println("-")
}
