package timer2

import (
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestTimer(t *testing.T) {
	timer := NewTimer(100)
	tds := make([]*TimerData, 100)
	for i := 0; i < 100; i++ {
		tds[i] = timer.Add(time.Duration(i)*time.Second+5*time.Minute, nil)
	}
	printTimer(timer)
	for i := 0; i < 100; i++ {
		logx.Infof("td: %s, %s, %d", tds[i].Key, tds[i].ExpireString(), tds[i].index)
		timer.Del(tds[i])
	}
	printTimer(timer)
	for i := 0; i < 100; i++ {
		tds[i] = timer.Add(time.Duration(i)*time.Second+5*time.Minute, nil)
	}
	printTimer(timer)
	for i := 0; i < 100; i++ {
		timer.Del(tds[i])
	}
	printTimer(timer)
	timer.Add(time.Second, nil)
	time.Sleep(time.Second * 2)
	if len(timer.timers) != 0 {
		t.FailNow()
	}
}

func printTimer(timer *Timer) {
	logx.Infof("----------timers: %d ----------", len(timer.timers))
	for i := 0; i < len(timer.timers); i++ {
		logx.Infof("timer: %s, %s, index: %d", timer.timers[i].Key, timer.timers[i].ExpireString(), timer.timers[i].index)
	}
	logx.Infof("--------------------")
}
