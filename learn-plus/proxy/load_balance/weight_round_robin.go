package load_balance

import (
	"errors"
	"strconv"
)

type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
	rsw      []int
}


type WeightNode struct {
	addr string
	weight int
	currentWeight int
	effectiveWeight int
}

func (r *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}
	parInt, err := strconv.ParseInt(params[1],10,64)
	if err != nil {
		return err
	}

	node := &WeightNode{
		addr:             params[0],
		weight:          int(parInt),
	}
	node.effectiveWeight = node.weight
	r.rss = append(r.rss, node)
	return nil
}


func (r *WeightRoundRobinBalance) Next() string {
	total := 0
	var best *WeightNode
	for i:=0; i<len(r.rss);i++ {
		w := r.rss[i]
		// step1 统计所有有效权重之和
		total += w.effectiveWeight

		// step2 变更节点临时权重的节点临时权重+节点有效权重
		w.currentWeight += w.effectiveWeight

		// step3 有效权重默认与权重相同，异常时-1，成功时+1，知道恢复到weight大小
		if w.effectiveWeight<w.weight{
			w.effectiveWeight++
		}

		// step4 选最大的临时权重节点
		if best==nil || w.currentWeight > best.currentWeight {
			best = w
		}

	}
	if best == nil {
		return ""
	}
	//step 5 变更临时权重为 临时权重-有效权重之和
	best.currentWeight -= total
	return best.addr
}

func (r *WeightRoundRobinBalance) Get(key string) (string, error) {
	return r.Next(), nil
}