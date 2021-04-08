package poker

import (
	"fmt"
	"os"
	"time"
)

// BlindAlerter schedules alerts for blind amounts.
// 定时弹出盲注的数量
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// BlindAlerterFunc allows you to implement BlindAlerter with a function.
// 通过函数的方式实现接口
type BlindAlerterFunc func(duration time.Duration, amount int)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter.
// BlindAlerterFunc 对于 BlindAlerter 的具体实现
func (a BlindAlerterFunc)ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

// StdOutAlerter will schedule alerts and print them to os.Stdout.
// 会将打印输出在 os.Stdout
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		_, _ = fmt.Fprintf(os.Stdout, "Blind is not %d\n", amount)
	})
}

/*NOTE
blind 在德州扑克中指的是盲注的意思
SB(small blind) = 小盲注
BB(Big blind) = 大盲注
*/