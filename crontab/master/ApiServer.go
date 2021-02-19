package master

import (
	"crontab/common"
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"time"
)

// 任务的Http接口
type ApiServer struct {
	httpServer *http.Server
}

// 单例
var (
	// 单例对象
	G_apiServer *ApiServer
)

//保存任务接口
func handleJobSave(w http.ResponseWriter, r *http.Request) {
	// 保存到etcd中
	var (
		err     error
		postJob string
		job     common.Job
		oldJob  *common.Job
		bytes   []byte
	)
	// 解析POST表单
	if err = r.ParseForm(); err != nil {
		goto ERR
	}
	// 2.取表单的job字段
	postJob = r.PostForm.Get("job")
	// 反序列化job
	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}

	// 保存到etcd
	if oldJob, err = G_jobMgr.SaveJob(&job); err != nil {
		goto ERR
	}

	// 正确时应答
	if bytes, err = common.BuildResponse(0, "success", oldJob); err == nil {
		w.Write(bytes)
	}

	return
ERR:
	//返回异常应答
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

// 删除任务
func handleJobDelete(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		name   string
		oldJob *common.Job
		bytes  []byte
	)

	// POST：a=1&b=2
	if err = r.ParseForm(); err != nil {
		goto ERR
	}
	// 删除的任务名
	name = r.PostForm.Get("name")
	// 删除任务
	if oldJob, err = G_jobMgr.DeleteJob(name); err != nil {
		goto ERR
	}
	// 正常应答
	if bytes, err = common.BuildResponse(0, "success", oldJob); err == nil {
		w.Write(bytes)
	}
	return
ERR:
	// 异常应答
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		w.Write(bytes)
	}
}

// 列举所有任务
func handleJobList(resp http.ResponseWriter, req *http.Request) {
	var (
		jobList []*common.Job
		err     error
		bytes   []byte
	)
	// 获取任务列表
	if jobList, err = G_jobMgr.ListJobs(); err != nil {
		goto ERR
	}

	if bytes, err = common.BuildResponse(0, "success", jobList); err == nil {
		resp.Write(bytes)
	}
	return
ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}
}

// 杀死任务
func handJobKill(resp http.ResponseWriter, req *http.Request) {
	var (
		err   error
		name  string
		bytes []byte
	)
	// 解析Post表单
	if err = req.ParseForm(); err != nil {
		goto ERR
	}
	// 要杀死的任务
	name = req.PostForm.Get("name")
	// 杀死任务
	if err = G_jobMgr.KillJob(name); err != nil {
		goto ERR
	}
	// 正常应答
	if bytes, err = common.BuildResponse(0, "success", nil); err == nil {
		resp.Write(bytes)
	}
	return
ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}
}

func handJobLog(resp http.ResponseWriter, req *http.Request) {
	var (
		err error
		name string  // 任务名
		skipParam string // 从第几条开始
		limitParam string// 返回多少条
		skip int
		limit int
		logArr []*common.JobLog
		bytes []byte
	)
	// 解析Get参数
	if err = req.ParseForm();err!=nil {
		goto ERR
	}
	name = req.Form.Get("name")
	skipParam = req.Form.Get("skip")
	limitParam = req.Form.Get("limit")

	if skip,err = strconv.Atoi(skipParam);err != nil {
		skip = 0
	}

	if limit,err = strconv.Atoi(limitParam);err !=nil {
		limit = 20
	}

	if logArr, err = G_logMgr.ListLog(name,int64(skip),int64(limit)); err !=nil{
		goto ERR
	}
	// 正常应答
	if bytes, err = common.BuildResponse(0, "success", logArr); err == nil {
		resp.Write(bytes)
	}
	return
ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}
}

func handWorkList(resp http.ResponseWriter, req *http.Request) {
	var(
		workerArr []string
		err error
		bytes []byte
	)

	if workerArr, err = G_workerMgr.ListWorkers(); err != nil {
		goto ERR

	}

	// 正常应答
	if bytes, err = common.BuildResponse(0, "success", workerArr); err == nil {
		resp.Write(bytes)
	}
	return
ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}
}

// 初始化服务
func InitApiServer() (err error) {
	var (
		mux           *http.ServeMux
		listener      net.Listener
		httpServer    *http.Server
		staticDir     http.Dir     // 静态文件根目录
		staticHandler http.Handler // 静态文件的Http回调函数
	)

	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/jobs/save", handleJobSave)
	mux.HandleFunc("/jobs/delete", handleJobDelete)
	mux.HandleFunc("/jobs/list", handleJobList)
	mux.HandleFunc("/jobs/kill", handJobKill)
	mux.HandleFunc("/jobs/log", handJobLog)
	mux.HandleFunc("/worker/list", handWorkList)

	// 静态文件目录
	staticDir = http.Dir(G_config.Web)
	staticHandler = http.FileServer(staticDir)
	mux.Handle("/", http.StripPrefix("/",staticHandler))   // 静态文件注册路由,这里的StripPrefix，需要把/index.html去掉/后于根目录拼接

	// 启动Tcp监听
	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(G_config.ApiPort)); err != nil {
		return
	}

	// 创建一个Http服务
	httpServer = &http.Server{
		ReadTimeout:  time.Duration(G_config.ApiReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiWriteTimeout) * time.Millisecond,
		Handler:      mux, // 所谓的mux路由也是一个handler，mux相当于路由转发
	}

	// 赋值单例
	G_apiServer = &ApiServer{httpServer: httpServer}

	// 启动服务端
	go httpServer.Serve(listener)
	return
}
