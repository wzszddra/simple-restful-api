package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shirou/gopsutil/v3/mem"

	"github.com/shirou/gopsutil/v3/disk"
)

type status struct {
	UsedMb      uint64  `json:"total"`
	TotalMb     uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func get_memory_status() status {
	v, _ := mem.VirtualMemory()

	var result = status{
		UsedMb:      v.Used,
		TotalMb:     v.Total,
		UsedPercent: v.UsedPercent,
	}

	return result
}

func get_disk_status() status {
	v, _ := disk.Usage("/")

	var result = status{
		UsedMb:      v.Used,
		TotalMb:     v.Total,
		UsedPercent: v.UsedPercent,
	}

	return result
}

func main() {
	router := gin.Default()
	router.GET("/status/mem", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, get_memory_status())
	})
	router.GET("/status/disk", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, get_disk_status())
	})
	router.GET("/status", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"mem":  get_memory_status(),
			"disk": get_disk_status(),
		})
	})

	router.Run("localhost:9090")
}

// 确实现学的go，可能有的写法不太聪明，比如查询/mem和/disk不知道能不能合成一个函数，不过go也没有多态？
// 代码也不是特别长，就不拆成好几个文件了qwq
// 基本功能是能实现的，尽量复现了截图的界面，两个小问题：
//// 1. 因为感觉传struct方便一点（如果传的map，利用率和其他两个类型不一样，还是说可以加个强转？），首字母就不能小写了
//// 2. 我是在Windows系统下VSCode自带的powershell里运行的，找不到JSON只输出Content的方法，不确定是不是环境问题（还没配其他系统的环境，并没有实验），不知道这样是不是满足需求了
