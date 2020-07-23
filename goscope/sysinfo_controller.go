package goscope

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"os"
)

func ShowSystemInfo(c *gin.Context) {
	// Markup
	sysinfoView, _ := Asset("static/html/system_info.html")
	commonHeader, _ := Asset("static/html/common_head.html")
	header := ReplaceVariablesInTemplate(string(commonHeader), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	// Styles
	highlightStyles, _ := Asset("static/css/highlight.css")
	goscopeStyles, _ := Asset("static/css/goscope.css")

	cpuStatus, _ := cpu.Info()
	firstCpu := cpuStatus[0]

	memoryStatus, _ := mem.VirtualMemory()

	hostStatus, _ := host.Info()

	variables := map[string]string{
		"APPLICATION_NAME":    os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":       MinifyHtml(header),
		"HIGHLIGHT_STYLES":    MinifyCss(string(highlightStyles)),
		"GOSCOPE_STYLES":      MinifyCss(string(goscopeStyles)),
		"CPU_MODEL_NAME":      firstCpu.ModelName,
		"CPU_CORE_COUNT":      fmt.Sprintf("%d Cores", firstCpu.Cores),
		"MEMORY_TOTAL":        fmt.Sprintf("%.2f GB", float64(memoryStatus.Total)/1073741824),
		"MEMORY_AVAILABLE":    fmt.Sprintf("%.2f GB", float64(memoryStatus.Available)/1073741824),
		"HOST_NAME":           hostStatus.Hostname,
		"HOST_KERNEL_ARCH":    hostStatus.KernelArch,
		"HOST_KERNEL_VERSION": hostStatus.KernelVersion,
		"HOST_OS":             hostStatus.OS,
		"HOST_UPTIME":         fmt.Sprintf("%d", hostStatus.Uptime),
	}
	ShowGoScopePage(c, MinifyHtml(string(sysinfoView)), variables)
}
