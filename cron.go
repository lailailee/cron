package main

import (
	"strconv"
	"strings"
	"time"
)

type Cron struct {
	planList []Plan
}
type Plan struct {
	timeTick [][]int
	function func()
}

func New() Cron {
	var cron Cron
	return cron
}

func (c *Cron) AddFunc(set string, function func()) {
	timeTick := parseParameter(set)
	plan := Plan{
		timeTick: timeTick,
		function: function,
	}
	c.planList = append(c.planList, plan)
}

func (c *Cron) Start() {
	c.listenTime()
}

func (c *Cron) listenTime() {
	tc := time.NewTicker(time.Duration(1000) * time.Millisecond)
	for {
		select {
		case <-tc.C:
			{
				for _, plan := range c.planList {
					b := checkTime(plan.timeTick)
					// fmt.Print("checkTime:")
					// fmt.Println(b)
					if b {
						plan.function()
					}
				}
			}
		}
	}
}

func getCurrentTime() (current []int) {
	t := time.Now()
	current = append(current, t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year())
	return
}

func checkTime(timeTick [][]int) (b bool) {
	current := getCurrentTime()
	// fmt.Println(current)
	b = true
	for i, v := range current {
		if len(timeTick[i]) != 0 && !contain(timeTick[i], v) {
			b = false
			break
		}
	}
	return b
}

func contain(array interface{}, item interface{}) bool {
	switch value := array.(type) {
	case []string:
		for _, eachItem := range value {
			if eachItem == item {
				return true
			}
		}
	case []int:
		for _, eachItem := range value {
			if eachItem == item {
				return true
			}
		}
	}
	return false
}

// 	// msg := "30,45 0,10,20,30,40,50 * * * *"
func parseParameter(parameter string) [][]int {
	var timeTick [][]int
	timeList := strings.Split(parameter, " ")
	for _, v := range timeList {
		ir := []int{}
		if !strings.Contains(v, "*") {
			ir = reInt(strings.Split(v, ","))
		}
		timeTick = append(timeTick, ir)
	}
	return timeTick
}
func reInt(m []string) (ir []int) {
	ir = []int{}
	if len(m) == 0 {
		return
	}
	for _, v := range m {
		var err error
		s, err := strconv.Atoi(v)
		ir = append(ir, s)
		if err != nil {
			return
		}
	}
	return
}
