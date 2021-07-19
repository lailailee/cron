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
