/*
Package ipdb implements an IPv4 to Country lookup.
The database is embedded in the package.
*/
package ipdb

import (
	"math/big"
	"net"
	"strings"
)

var db []iPRecord

type iPRecord struct {
	FromIP  int64
	ToIP    int64
	Country string
}

// parses the ~12MB dbdata.go (approx. 400 000 records)
func loadDB() {
	// This check is in case of multiple loadDB() calls in the same program
	if db != nil {
		return
	}

	lines := strings.Split(DBData, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		s := strings.Split(line, "\t")
		iprecord := iPRecord{
			FromIP:  ipToInt(s[0]),
			ToIP:    ipToInt(s[1]),
			Country: s[2],
		}
		db = append(db, iprecord)
	}
}

func GetCountry(ip string) string {
	if db == nil {
		loadDB()
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
