package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

// 雪花算法
var node *snowflake.Node

// MakeSnow 可以根据需要转为 string类型返回 防止溢出
func MakeSnow(startTime string, machineId int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineId)
	return
}

func GenId() int64 {
	return node.Generate().Int64()
}