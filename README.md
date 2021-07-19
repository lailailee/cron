## Cron

```doc

type Plan struct {
	timeTick [][]int
	function func()
}

type Cron struct {
     planList []Plan
}

// New 新建Cron对象
func New() Cron

// AddFunc 添加监听事件
func (c *Cron) AddFunc(set string, function func())

// Start 开始监听
func (c *Cron) Start()

```
AddFunc set参数为cron表达式

(Cron表达式是一个字符串，字符串以5或6个空格隔开，分为6或7个域，每一个域代表一个含义)_

本Cron中的格式为

秒 分 时 日（每月的第几日） 月 年

每个域以逗号隔开代表响应含义的时间

*代表该域都匹配任意值（特殊符号暂只支持*）

详见 

https://www.cnblogs.com/junrong624/p/4239517.html

https://cron.qqe2.com/


