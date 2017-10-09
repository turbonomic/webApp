package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	cpuPath          = "/cpuwork.php"
	memPath          = "/memwork.php"
	defaultThreadNum = 4
)

type requestConfig struct {
	host      string
	paths     []string
	kind      string
	threadNum int
	// the duration of this test, in seconds
	time int

	//how much cpu to be used, should be less than 1000
	cpu int

	//how much memory to be used, in MB
	memory int
	//how long to hold the memory, in milliseconds
	latency int
}

func do_send_request(url string, qdata url.Values) string {
	res, err := http.PostForm(url, qdata)
	if err != nil {
		glog.Warning(err)
		return "err"
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		glog.Warningf("bad http[%s] result: %+v", url, res)
		return "err2"
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		glog.Fatal(err)
		return "err3"
	}

	if len(content) < 0 {
		glog.Fatalf("content is empty:%s", content)
		return "err4"
	}

	return string(content[:])
}

func send_request(index int, req *requestConfig, stop chan struct{}) int {
	glog.Infof("worker[%d] begins", index)
	data := fmt.Sprintf("cpu=%d&memory=%d&value=%d", req.cpu, req.memory, req.latency)
	qdata, err := url.ParseQuery(data)
	if err != nil {
		glog.Fatal("failed to parse query[%s]: %v", data, err)
	}

	count := 0
	for i := 0; ; i++ {
		for _, path := range req.paths {
			select {
			case <-stop:
				glog.Infof("workder[%d] is done with %d requests", index, count)
				return count
			default:
				url := fmt.Sprintf("%s%s", req.host, path)
				if count == 0 {
					glog.Infof("url: %v", url)
				}

				result := do_send_request(url, qdata)
				if count%1000 == 0 {
					glog.Infof("count=%d: %v", count, result)
				}
				count += 1
			}

		}

		if i%1000 == 0 {
			glog.V(3).Infof("worker[%d] sends %d requests", index, i)
		}
	}
}

func wait(du int) {
	if du < 1 {
		//wait infinitely
		select {}
	}

	tick := time.Tick(time.Second * time.Duration(du))
	<-tick
}

func parallel_send(req *requestConfig) {
	var wg sync.WaitGroup
	stop := make(chan struct{})
	results := make(chan int, req.threadNum)

	//1. start the threads
	for i := 0; i < req.threadNum; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			result := send_request(index, req, stop)
			results <- result
		}(i)
	}

	//2. wait for timer
	wait(req.time)
	close(stop)
	wg.Wait()

	//3. statistic the result
	total := 0
	for i := 0; i < req.threadNum; i++ {
		total += <-results
	}

	glog.Infof("test sends %d reqeusts in total.", total)
}

func parseFlag(req *requestConfig) {
	flag.StringVar(&(req.host), "host", "http://127.0.0.1:8080", "the host:port of webserver")
	flag.IntVar(&(req.cpu), "cpu", 100, "the computation amount: calculate md5 for #cpu * 1000 times")
	flag.IntVar(&(req.memory), "memory", 64, "the memory used for each request in MB")
	flag.IntVar(&(req.latency), "delay", 30, "the latency of the memory-type request")
	flag.IntVar(&(req.threadNum), "threadNum", defaultThreadNum, "number of threads to send requests")
	flag.StringVar(&(req.kind), "kind", "cpu", "three kinds of workloads: cpu/memory/mix")
	flag.IntVar(&(req.time), "time", 0, "how long to do this tests in seconds; 0 means unlimited")
	flag.Set("logtostderr", "true")

	flag.Parse()

	paths := []string{cpuPath, memPath}
	if req.kind == "cpu" {
		paths = []string{cpuPath}
	} else if req.kind == "memory" {
		paths = []string{memPath}
	}

	req.paths = paths
	fmt.Println(paths)
	glog.V(2).Infof("kind=%s, paths=%++v", req.kind, req.paths)
}

func main() {
	req := &requestConfig{}
	parseFlag(req)
	parallel_send(req)
}
