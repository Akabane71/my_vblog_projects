package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

// 雪花算法库
// github.com/bwmarrin/snowflake
// github.com/sony/sonyflake   // 索尼公司提供的算法

var node *snowflake.Node

// Init 唯一标识 和 初始时间
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	// Go的时间表示
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}

// Demo 展示案例
func Demo() {
	if err := Init("2020-07-01", 1); err != nil {
		fmt.Println("init failed err", err)
	}

	id := GenID()
	fmt.Println(id) // 537620966713331712
}
