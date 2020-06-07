package work

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/phachon/mm-wiki/app/models"
	"github.com/phachon/mm-wiki/app/services"
	"github.com/phachon/mm-wiki/app/utils"
	"sync"
	"time"
)

var (
	DocSearchWorker = NewDocSearchWork()
)

const (
	// work 未启动或已停止
	RunStatusStop = 0
	// work 运行中
	RunStatusRunning = 1
)

type DocSearch struct {
	// 并发锁，理论上不存在并发的情况，为了安全
	lock sync.RWMutex
	// work 运行状态
	runStatus int
	// work 中是否有任务正在运行
	isTaskRunning bool
	// work 退出信号
	quit chan bool
}

func NewDocSearchWork() *DocSearch {
	return &DocSearch{
		runStatus:     RunStatusStop,
		isTaskRunning: false,
		quit:          make(chan bool, 1),
	}
}

// Start 开始 work
func (d *DocSearch) Start() {
	// 已经在运行
	if d.runStatus == RunStatusRunning {
		return
	}
	timer, ok := d.getFullTextSearchConf()
	if !ok {
		return
	}
	d.updateAllDocIndex()
	go func(d *DocSearch, t time.Duration) {
		defer func() {
			e := recover()
			if e != nil {
				logs.Info("[DocSearchWork] load all doc index panic: %v", e)
			}
			d.lock.Lock()
			d.runStatus = RunStatusStop
			d.isTaskRunning = false
			d.lock.Unlock()
		}()
		d.lock.Lock()
		d.runStatus = RunStatusRunning
		d.lock.Unlock()
		for {
			select {
			case <-time.Tick(t):
				if !d.isTaskRunning {
					d.updateAllDocIndex()
				}
			case <-d.quit:
				logs.Info("[DocSearchWork] stop doc index")
				return
			}
		}
	}(d, time.Duration(timer)*time.Second)
}

// Restart 重新启动 work
func (d *DocSearch) Restart() {
	d.Stop()
	time.Sleep(time.Millisecond)
	d.Start()
}

// Stop 停止 work
func (d *DocSearch) Stop() {
	d.quit <- true
}

// 查找是否开启全文索引并获取配置
func (d *DocSearch) getFullTextSearchConf() (timer int64, isOpen bool) {
	fulltextSearchOpen := models.ConfigModel.GetConfigValueByKey(models.ConfigKeyFulltextSearch, "0")
	docSearchTimer := models.ConfigModel.GetConfigValueByKey(models.ConfigKeyDocSearchTimer, "3600")
	timer = utils.Convert.StringToInt64(docSearchTimer)
	// 默认 3600 s
	if timer <= 0 {
		timer = int64(3600)
	}
	if fulltextSearchOpen == "1" {
		return timer, true
	}
	return timer, false
}

func (d *DocSearch) updateAllDocIndex() {

	logs.Info("[DocSearchWork] start load all doc index")

	d.lock.Lock()
	d.isTaskRunning = true
	d.lock.Unlock()

	// 分批次更新，每批次 100
	batchUpdateDocNum, _ := beego.AppConfig.Int("search::batch_update_doc_num")
	if batchUpdateDocNum <= 0 {
		batchUpdateDocNum = 100
	}
	services.DocIndexService.UpdateAllDocIndex(batchUpdateDocNum)
	services.DocIndexService.Flush()

	d.lock.Lock()
	d.isTaskRunning = false
	d.lock.Unlock()

	logs.Info("[DocSearchWork] finish all doc index flush")

}
