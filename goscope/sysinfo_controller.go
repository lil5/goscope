package goscope

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"os"
)

const (
	BytesInOneGigabyte = 1073741824
	SecondsInOneMinute = 60
)

func ShowSystemInfo(c *gin.Context) {
	// Markup
	sysinfoView, _ := Asset("../static/html/system_info.html")
	commonHeader, _ := Asset("../static/html/common_head.html")
	headerVariables := map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")}
	header := ReplaceVariablesInTemplate(string(commonHeader), headerVariables)
	// Styles
	highlightStyles, _ := Asset("../static/css/highlight.css")
	goscopeStyles, _ := Asset("../static/css/goscope.css")

	cpuStatus, _ := cpu.Info()
	firstCPU := cpuStatus[0]
	memoryStatus, _ := mem.VirtualMemory()
	swapStatus, _ := mem.SwapMemory()
	hostStatus, _ := host.Info()
	diskStatus, _ := disk.Usage("/")
	variables := map[string]string{
		"APPLICATION_NAME":    os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":       MinifyHTML(header),
		"CPU_CORE_COUNT":      fmt.Sprintf("%d Cores", firstCPU.Cores),
		"CPU_MODEL_NAME":      firstCPU.ModelName,
		"DISK_FREE":           fmt.Sprintf("%.2f GB", float64(diskStatus.Free)/BytesInOneGigabyte),
		"DISK_PARTITION_TYPE": diskStatus.Fstype,
		"DISK_PATH":           diskStatus.Path,
		"DISK_TOTAL":          fmt.Sprintf("%.2f GB", float64(diskStatus.Total)/BytesInOneGigabyte),
		"GOSCOPE_STYLES":      MinifyCSS(string(goscopeStyles)),
		"HIGHLIGHT_STYLES":    MinifyCSS(string(highlightStyles)),
		"HOST_KERNEL_ARCH":    hostStatus.KernelArch,
		"HOST_KERNEL_VERSION": hostStatus.KernelVersion,
		"HOST_NAME":           hostStatus.Hostname,
		"HOST_OS":             hostStatus.OS,
		"HOST_PLATFORM":       hostStatus.Platform,
		"HOST_UPTIME":         fmt.Sprintf("%.2f hours", float64(hostStatus.Uptime)/SecondsInOneMinute/SecondsInOneMinute),
		"MEMORY_AVAILABLE":    fmt.Sprintf("%.2f GB", float64(memoryStatus.Available)/BytesInOneGigabyte),
		"MEMORY_TOTAL":        fmt.Sprintf("%.2f GB", float64(memoryStatus.Total)/BytesInOneGigabyte),
		"SWAP_USED":           fmt.Sprintf("%.2f%%", swapStatus.UsedPercent),
	}
	ShowGoScopePage(c, MinifyHTML(string(sysinfoView)), variables)
}
