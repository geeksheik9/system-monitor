package main

import (
	"runtime"

	"github.com/geeksheik9/system-monitor/pkg/api"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Running on %s", runtime.GOOS)
	switch runtime.GOOS {
	case "linux":
		api.MonitorLinux()
	case "windows":
		api.MonitorWindows()
	case "darwin":
		api.MonitorDarwin()
	case "arm":
		api.MonitorArm()
	default:
		log.Fatal("Unsupported OS")
	}

}
