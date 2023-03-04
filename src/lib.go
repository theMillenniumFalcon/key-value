package main

import "strings"

// type of database
type Deleted int

var (
	NO   Deleted = 0
	SOFT Deleted = 1
	HARD Deleted = 2
)

type Record struct {
	rVolumes []string
	deleted  Deleted
	hash     string
}

func toRecord(data []byte) Record {
	var rec Record
	ss := string(data)
	rec.deleted = NO

	if strings.HasPrefix(ss, "DELETED") {
		rec.deleted = SOFT
		ss = ss[7:]
	}

	if strings.HasPrefix(ss, "HASH") {
		rec.hash = ss[4:36]
		ss = ss[36:]
	}

	rec.rVolumes = strings.Split(ss, ",")

	return rec
}

func fromRecord(rec Record) []byte {
	cc := ""

	if rec.deleted == HARD {
		panic("Can't put HARD delete in the database")
	}

	if rec.deleted == SOFT {
		cc = "DELETED"
	}

	if len(rec.hash) == 32 {
		cc += "HASH" + rec.hash
	}

	return []byte(cc + strings.Join(rec.rVolumes, ","))
}
