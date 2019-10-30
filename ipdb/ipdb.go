package ipdb

import (
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"strings"
)

var db []IPRecord

type IPRecord struct {
	FromIP  int64  `json:"From IP"`
	ToIP    int64  `json:"To IP"`
	Country string `json:"Country"`
}

// loads the 12MB file (approx. 400 000 records) into memory to serve as a database
func Init() {
    // In case of multiple Init() calls in the same program,
    // we should not load the 12MB file in memory too many times
	if db != nil {
		return
	}

	content, err := ioutil.ReadFile("ipdb/ip2country-v4.tsv")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		s := strings.Split(line, "\t")
		iprecord := IPRecord{
			FromIP:  ipToInt(s[0]),
			ToIP:    ipToInt(s[1]),
			Country: s[2],
		}
		db = append(db, iprecord)
	}
}

func GetCountry(ip string) string {
    if db == nil {
        log.Fatal("ERROR: no db found (ipdb.Init() not called?)")
    }
	ip_as_int := ipToInt(ip)
	for _, record := range db {
		if ip_as_int > record.FromIP && ip_as_int < record.ToIP {
			return record.Country
		}
	}
	return ""
}

func ipToInt(IPv4Address string) int64 {
    return big.NewInt(0).SetBytes(net.ParseIP(IPv4Address).To4()).Int64()
}

