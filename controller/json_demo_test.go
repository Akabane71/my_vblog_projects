package controller

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"testing"
)

// 解决方法2 : `json:"string"` ; 默认Marshal成string
type MyDate struct {
	ID   int64  `json:"id,string"` // 当json序列化的时候将其转为string
	Name string `json:"name"`
}

// 解决方法1: 写个 序列化/反序列化  的方法

func (d *MyDate) Unmarshal() {

}

func (d *MyDate) Marshal() {

}

func TestJsonDemo(t *testing.T) {
	// 序列化:  后端的数据---> JSON格式的数据
	d1 := MyDate{
		ID:   math.MaxInt64,
		Name: "LiShun",
	}
	fmt.Println(d1.ID)
	b, err := json.Marshal(d1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// 由于前端可能接收不了int64这么大的数据;数据会出现失真
	id := string(d1.Name)

	id = strconv.FormatInt(d1.ID, 10)
	fmt.Println(id)

}
