package load_balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	currIndex int
	rss []string
	// 观察主体

}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}

	add := params[0]
	r.rss=append(r.rss,add)
	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.rss) == 0{
		return ""
	}
	r.currIndex = rand.Intn(len(r.rss))
	return r.rss[r.currIndex]
}

func (r *RandomBalance)  Get( key string) (string, error){
	return r.Next(),nil
}
