package main

import (
	"fmt"
	"github.com/chilts/sid"
	"github.com/kjk/betterguid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	//"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
	"log"
	"math/rand"
	"time"
)

func genXid() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:           %s\n", id.String())
}

func genKsuid() {
	id := ksuid.New()
	fmt.Printf("github.com/segmentio/ksuid:  %s\n", id.String())
}

func genBetterGUID() {
	id := betterguid.New()
	fmt.Printf("github.com/kjk/betterguid:   %s\n", id)
}

func genUlid() {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("github.com/oklog/ulid:       %s\n", id.String())
}

func getMachineID() (uint16, error) {
    var machineID uint16 = 6
    return machineID, nil
}

func checkMachineID(machineID uint16) bool {
    existsMachines := []uint16{1, 2, 3, 4, 5}
    for _, v := range existsMachines {
        if v == machineID {
            return false
        }
    }
    return true
}

func genSonyflake() {
        //st := time.Unix(0,0)
         var st time.Time
    st, err := time.Parse("2006-01-02", "2014-01-01")
    if err != nil {
        fmt.Printf("flake.init() failed with %s\n", err.Error())
        panic(err)
    }

	flake := sonyflake.NewSonyflake(sonyflake.Settings{
	     StartTime: st,
	      MachineID: getMachineID,
        CheckMachineID: checkMachineID,
	
	
})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err.Error())
	}
	fmt.Printf("github.com/sony/sonyflake:   %x\n", id)
	fmt.Printf("github.com/sony/sonyflake num:   %d\n", int64(id))
}

func genSid() {
	id := sid.Id()
	fmt.Printf("github.com/chilts/sid:       %s\n", id)
}

/*func genUUIDv4() {
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("get uuid error [%s]",err)
	}
	fmt.Printf("github.com/satori/go.uuid:   %s\n", id)
}*/

func main() {
	genXid()
	genKsuid()
	genBetterGUID()
	genUlid()
	genSonyflake()
	genSid()
//	genUUIDv4()
}
