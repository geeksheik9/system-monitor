package api

import (
	"bytes"
	"os/exec"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func MonitorLinux() {
	sleeper := 2
	for {
		time.Sleep(time.Duration(sleeper) * time.Second)
		cmd := exec.Command(`cat sys/class/thermal/thermal_zone*/temp | column -s $'\t' -t | sed 's/\(.\)..$/.\1/'`)
		err := cmd.Run()
		if err != nil {
			log.Infof("Command issue: %v", cmd.Stdin)
			log.Infof("Error %v, ", cmd.Stdout)
			log.Errorf("Monitoring error: %s", err.Error())
			sleeper *= 2
			continue
		}

		var out bytes.Buffer
		cmd.Stdout = &out

		temp, err := strconv.Atoi(out.String())
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