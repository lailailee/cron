package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	msg := "5,10,11,15,20,30,45,55,58 * * 20 * 2022"
	c := New()
	c.AddFunc(msg, func() {
		fmt.Println(time.Now())
	})
	c.Start()
	fmt.Println(":LLLL")
}
func TestCron2(t *testing.T) {
	msg := "5,10,11,15,20,30,45,55,58"
	c := New()
	c.AddFunc(msg, func() {
		fmt.Println(time.Now())
	})
	c.Start()
}
func TestCron3(t *testing.T) {
	msg := "30 56"
	c := New()
	c.AddFunc(msg, func() {
		fmt.Println(time.Now())
	})
	c.Start()
	fmt.Println(":LLLL")
	select {}
}

// === RUN   TestCron
// 2021-07-20 00:31:15.772749 +0800 CST m=+21.002984191
// 2021-07-20 00:31:20.772793 +0800 CST m=+26.003178338
// 2021-07-20 00:31:30.771997 +0800 CST m=+36.002983774
// 2021-07-20 00:31:45.773642 +0800 CST m=+51.004888435
// 2021-07-20 00:32:15.77456 +0800 CST m=+81.005528902
// 2021-07-20 00:32:20.77205 +0800 CST m=+86.002938854
// 2021-07-20 00:32:30.772236 +0800 CST m=+96.002954091
// 2021-07-20 00:32:45.772588 +0800 CST m=+111.003037808
// 2021-07-20 00:33:15.773207 +0800 CST m=+141.003102382
// 2021-07-20 00:33:20.773244 +0800 CST m=+146.003045487
