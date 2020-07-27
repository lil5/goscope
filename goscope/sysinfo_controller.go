package goscope

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

const (
	BytesInOneGigabyte = 1073741824
	SecondsInOneMinute = 60
)

// Show system information of the current host.
func ShowSystemInfo(c *gin.Context) {
	cpuStatus, _ := cpu.Info()
	firstCPU := cpuStatus[0]
	memoryStatus, _ := mem.VirtualMemory()
	swapStatus, _ := mem.SwapMemory()
	hostStatus, _ := host.Info()
	diskStatus, _ := disk.Usage("/")
	variables := gin.H{
		"applicationName": os.Getenv("APPLICATION_NAME"),
		"cpu": gin.H{
			"coreCount": fmt.Sprintf("%d Cores", firstCPU.Cores),
			"modelName": firstCPU.ModelName,
		},
		"disk": gin.H{
			"freeSpace":     fmt.Sprintf("%.2f GB", float64(diskStatus.Free)/BytesInOneGigabyte),
			"partitionType": diskStatus.Fstype,
			"mountPath":     diskStatus.Path,
			"totalSpace":    fmt.Sprintf("%.2f GB", float64(diskStatus.Total)/BytesInOneGigabyte),
		},
		"host": gin.H{
			"kernelArch":    hostStatus.KernelArch,
			"kernelVersion": hostStatus.KernelVersion,
			"hostname":      hostStatus.Hostname,
			"hostOS":        hostStatus.OS,
			"hostPlatform":  hostStatus.Platform,
			"uptime":        fmt.Sprintf("%.2f hours", float64(hostStatus.Uptime)/SecondsInOneMinute/SecondsInOneMinute),
		},
		"memory": gin.H{
			"availableMemory": fmt.Sprintf("%.2f GB", float64(memoryStatus.Available)/BytesInOneGigabyte),
			"totalMemory":     fmt.Sprintf("%.2f GB", float64(memoryStatus.Total)/BytesInOneGigabyte),
			"usedSwap":        fmt.Sprintf("%.2f%%", swapStatus.UsedPercent),
		},
	}
	c.JSON(http.StatusOK, variables)
}
