package wsocket

import "time"

func ScheduleTaskStart() {

	go heartbeatOutTimeSchedule()

}

// 心跳检测异常断链的数据
func heartbeatOutTimeSchedule() {

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {

		select {
		case <-ticker.C:
			ClientManagerService.HeartbeatOutTimeClose()
		}

	}

}
