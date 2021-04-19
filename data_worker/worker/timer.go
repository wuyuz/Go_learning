package worker

import (
	. "data_worker/app"
	"data_worker/common"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type timeConfig struct {
	GET_DATA int
	GET_URL  string
}

var timerCost *timeConfig

func loadConfig() {
	common.ViperConfig.SetDefault("GET_DATA", "1")
	common.ViperConfig.SetDefault("GET_URL", "")
	common.ViperConfig.Unmarshal(&timerCost)
}

// Timer Worker
func TimerWork() {
	loadConfig()
	t := time.NewTicker(time.Duration(timerCost.GET_DATA) * time.Second)
	TimerChan = make(chan bool)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			Log.Info().Msg("[+] Get data worker running")
			get_data()
		case stop := <-TimerChan:
			if stop {
				Log.Error().Msg("[+] Stop data worker ...")
				return
			}
		}
	}
}

func get_data() {
	var (
		client *http.Client
		err    error
		resp   *http.Response
	)
	client = &http.Client{Timeout: 5 * time.Second}
	resp, err = client.Get(timerCost.GET_URL)
	if err != nil {
		Log.Error().Msg(err.Error())
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}
