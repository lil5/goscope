package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/averageflow/goscope/src/types"

	"github.com/averageflow/goscope/src/utils"
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

func GetAppName(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"applicationName": utils.Config.ApplicationName,
	})
}

// ShowSystemInfo is the controller to show system information of the current host in GoScope API.
func ShowSystemInfo(c *gin.Context) {
	cpuStatus, _ := cpu.Info()
	firstCPU := cpuStatus[0]
	memoryStatus, _ := mem.VirtualMemory()
	swapStatus, _ := mem.SwapMemory()
	hostStatus, _ := host.Info()
	diskStatus, _ := disk.Usage("/")

	environment := make(map[string]string)

	env := os.Environ()
	for i := range env {
		variable := strings.SplitN(env[i], "=", 2)
		environment[variable[0]] = variable[1]
	}

	responseBody := types.SystemInformationResponse{
		ApplicationName: utils.Config.ApplicationName,
		CPU: types.SystemInformationResponseCPU{
			CoreCount: fmt.Sprintf("%d Cores", firstCPU.Cores),
			ModelName: firstCPU.ModelName,
		},
		Memory: types.SystemInformationResponseMemory{
			Available: fmt.Sprintf("%.2f GB", float64(memoryStatus.Available)/BytesInOneGigabyte),
			Total:     fmt.Sprintf("%.2f GB", float64(memoryStatus.Total)/BytesInOneGigabyte),
			UsedSwap:  fmt.Sprintf("%.2f%%", swapStatus.UsedPercent),
		},
		Host: types.SystemInformationResponseHost{
			HostOS:        hostStatus.OS,
			HostPlatform:  hostStatus.Platform,
			Hostname:      hostStatus.Hostname,
			KernelArch:    hostStatus.KernelArch,
			KernelVersion: hostStatus.KernelVersion,
			Uptime:        fmt.Sprintf("%.2f hours", float64(hostStatus.Uptime)/SecondsInOneMinute/SecondsInOneMinute),
		},
		Disk: types.SystemInformationResponseDisk{
			FreeSpace:     fmt.Sprintf("%.2f GB", float64(diskStatus.Free)/BytesInOneGigabyte),
			MountPath:     diskStatus.Path,
			PartitionType: diskStatus.Fstype,
			TotalSpace:    fmt.Sprintf("%.2f GB", float64(diskStatus.Total)/BytesInOneGigabyte),
		},
		Environment: environment,
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, responseBody)
}
