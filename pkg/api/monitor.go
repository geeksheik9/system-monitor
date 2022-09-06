package api

import (
	"os/exec"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func MonitorLinux() {
	sleeper := 30
	counter := 0
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

		log.Info(temp)

		if temp > 40 {
			alert("temp is very high")
			log.Errorf("Temperature is %v°C", temp)
			sleeper *= 2
			continue
		}

		if temp > 55 {
			alert("too hot")
			log.Errorf("Temperature is %v°C", temp)
			counter++
			sleeper *= 2
			if counter > 3 {
				exit()
			}
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
func alert(message string) {
	log.Info("Alert Unimplmented")
}

func exit() {
	log.Error("Counter has exceeded limit, shutting down system in 5 seconds")
	time.Sleep(5 * time.Second)
	log.Error("Shutting down system")
	exec.Command("shutdown", "-h", "now")
}
