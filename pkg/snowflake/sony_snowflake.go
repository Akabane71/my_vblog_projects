package snowflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func SonyInit(startTime string, machineId uint16) (err error) {
	sonyMachineID = machineId
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

func SonyGenID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}

func SonyDemo() {
	if err := SonyInit("2020-07-01", 1); err != nil {
		fmt.Println("Init  failed err:", err)
		return
	}
	id, _ := SonyGenID()
	fmt.Println("id:", id) // id: 215048360514551809
}
