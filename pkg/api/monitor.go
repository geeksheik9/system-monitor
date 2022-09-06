package api

import (
	"os/exec"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func MonitorLinux() {
	sleeper := 2
	for {
		time.Sleep(time.Duration(sleeper) * time.Second)
		cmd := exec.Command("./scripts/monitor.sh")
		out, err := cmd.Output()
		if err != nil {
			log.Errorf("Monitoring error: %s", err)
			sleeper *= 2
			continue
		}

		temp, err := strconv.Atoi(string(out[0:2]))
		if err != nil {
			log.Errorf("Conversion error: %s", err)
			sleeper *= 2
			continue
		}

		if temp > 40 {
			log.Errorf("Temperature is %v°C", temp)
			sleeper *= 2
			continue
		}
	}
}

// TODO: implement Monitor for windows
func MonitorWindows() {
	log.Info("Monitoring Windows Unimplmented")
}

// TODO: implement Monitor for darwin
func MonitorDarwin() {
	log.Info("Monitoring Darwin Unimplmented")
}

// TOOD: implement Monitor for arm
func MonitorArm() {
	log.Info("Monitoring ARM Unimplmented")
}

// TODO: implement alerting for texting or email
func Alert() {
	log.Info("Alert Unimplmented")
}
