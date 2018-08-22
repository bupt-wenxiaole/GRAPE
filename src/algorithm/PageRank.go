package algorithm
/*
import (
	"graph"
	"math"
	"log"
)

const eps = 0.01

type PRPair struct {
	PRValue float64
	ID graph.ID
}

func PageRank_PEVal(g graph.Graph, prVal map[int64]float64, workerNum int) (int64, map[int64]float64) {
	nodeNum := len(g.GetNodes())
	initVal := 1.0 / float64(nodeNum * workerNum)
	for id := range g.GetNodes() {
		prVal[id.IntVal()] = initVal
	}

	tempPr := make(map[int64]float64)
	loopTime := 0
	for {
		if loopTime == 0 {
			log.Println("finish peval")
			break
		}

		updated := false
		still := 0.0
		loopTime++

		log.Printf("loop time:%v\n", loopTime)

		for id := range g.GetNodes() {
			targets, _ := g.GetTargets(id)
			if len(targets) == 0 {
				still += prVal[id.IntVal()]
			} else {
				num := float64(len(targets))
				for dstId := range targets {
					tempPr[dstId.IntVal()] += 0.85 * prVal[id.IntVal()] / num
				}
			}
		}
		still = 0.85 * still / float64(nodeNum) + 0.15 * initVal
		for id := range g.GetNodes() {
			tempPr[id.IntVal()] += still
			if math.Abs(tempPr[id.IntVal()] - prVal[id.IntVal()]) > eps * initVal {
				updated = true
			}
			//maxerr = math.Max(maxerr, math.Abs(tempPr[id.IntVal()] - prVal[id.IntVal()]))
		}
		//log.Printf("max error:%v\n", maxerr)

		if !updated {
			prVal = tempPr
			break
		}

		prVal = tempPr
		tempPr = make(map[int64]float64)
	}

	log.Printf("loop time:%v\n", loopTime)
	return int64(nodeNum), prVal
}

func PageRank_IncEval(g graph.Graph, prVal map[int64]float64, oldPr map[int64]float64, workerNum int, partitionId int, outerMsg map[int64][]int64, messages map[int64]float64, totalVertexNum int64) (bool, map[int][]*PRMessage, map[int64]float64, map[int64]float64) {
	maxerr := 0.0

	still := 0.0
	initVal := 1.0 / float64(totalVertexNum)
	updated := false

	//var receiveSum float64 = 0
	//var sendSum float64 = 0

	for id, msg := range messages {
		//log.Printf("id:%v, val:%v\n", id, msg)
		if id != -1 {
			prVal[id] += msg
			//receiveSum += msg
		} else {
			still += msg
		}
	}

	log.Printf("threshold:%v\n", eps * initVal)

	var sum float64 = 0
	for id := range g.GetNodes() {
		prVal[id.IntVal()] += still * 0.85
		if math.Abs(prVal[id.IntVal()] - oldPr[id.IntVal()]) > eps * initVal {
			updated = true
		}
		maxerr = math.Max(maxerr, math.Abs(prVal[id.IntVal()] - oldPr[id.IntVal()]))
		sum += prVal[id.IntVal()]
	}
	//log.Printf("total vertex num:%v\n", totalVertexNum)
	//log.Printf("still receive:%v\n", still)
	log.Printf("max error:%v\n", maxerr)

	tempPr := make(map[int64]float64)
	messagePr := make(map[int64]float64)
	still = 0

	for id := range g.GetNodes() {
		targets, _ := g.GetTargets(id)
		sonNum := len(targets) + len(outerMsg[id.IntVal()])
		tempPr[id.IntVal()] += 0.15 * initVal
		if sonNum == 0 {
			still += prVal[id.IntVal()] / float64(totalVertexNum)
		} else {
			val := prVal[id.IntVal()] / float64(sonNum)
			//log.Printf("val: %v\n", val)
			for target := range targets {
				tempPr[target.IntVal()] += 0.85 * val
			}
			for _, outer := range outerMsg[id.IntVal()] {
				//log.Printf("out node:%v\n", outer)

				messagePr[outer] += 0.85 * val
				//sendSum += 0.85 * val
			}
		}
	}
	for id := range g.GetNodes() {
		tempPr[id.IntVal()] += 0.85 * still
	}

	reduceMsg := make(map[int][]*PRMessage)

	for i := 0; i < workerNum; i++ {
		if i == partitionId {
			continue
		}
		reduceMsg[i] = make([]*PRMessage, 0)
		reduceMsg[i] = append(reduceMsg[i], &PRMessage{PRValue:still,ID:graph.ID(-1)})
	}

	//log.Printf("still send:%v\n", still)

	//log.Printf("receive sum:%v\n", receiveSum)
	//log.Printf("send sum:%v\n", sendSum)
	return updated, reduceMsg, prVal, tempPr
}
*/