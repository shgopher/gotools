package id

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	VERSION_1 = 1
	VERSION_2 = 2
	VERSION_3 = 3
	VERSION_4 = 4
	VERSION_5 = 5
)
// only version 3 5 need data
// version1 can be used in Distributed scenario
// version2 less to use
// version3 less to use based in MD5
// version4 based in random
// vesion 5 base in sha1 safety but slow.
func NewUUID(version int,data []byte)(uuid.UUID,error){
	switch version {
	case VERSION_1:
		return uuid.NewUUID()
	case VERSION_2:
		return uuid.NewDCEGroup()
	case VERSION_3:
		return uuid.NewMD5(uuid.NameSpaceDNS,data),nil
	case VERSION_4:
		return uuid.NewRandom()
	case VERSION_5:
		return uuid.NewSHA1(uuid.NameSpaceOID,data),nil
	default:
		return uuid.UUID{},fmt.Errorf("version number is wrong ,you just can use [1,5]")
	}
}

