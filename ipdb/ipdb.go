package ipdb

import (
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

// parses the ~12MB dbdata.go (approx. 400 000 records)
func Init() {
	// This check is in case of multiple Init() calls in the same program
	if db != nil {
		return
	}

	lines := strings.Split(DBData, "\n")
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
