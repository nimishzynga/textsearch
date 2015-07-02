package main

import "strings"
import "math"
import "fmt"

type wf struct {
	count int
	w     map[string]int
}

type ind struct {
	totalDoc  int
	wordToDoc map[string]*wf
}

func (x *ind) index(id string, data string) {
	x.totalDoc++
	newMap := make(map[string]int)
	splitted := strings.Split(data, " ")
	for _, v := range splitted {
		newMap[v]++
	}
	for k, v := range newMap {
		if x.wordToDoc[k] == nil {
			t := make(map[string]int)
			t[id] = v
			x.wordToDoc[k] = &wf{w: t, count: 1}
		} else {
			m := x.wordToDoc[k]
			m.w[id] = v
			m.count++
			x.wordToDoc[k] = m
		}
	}
}

func (x *ind) queryMatch(str string) {
	strs := strings.Split(str, " ")
    for _,j := range x.wordToDoc {
        fmt.Println(*j)
    }
	scoreMap := make(map[string]float64)
	for _, str := range strs {
		val := x.wordToDoc[str]
		h := math.Log(float64(x.totalDoc / len(val.w)))
        fmt.Println("h is ", h);
		for k, v := range val.w {
			scoreMap[k] += float64(v) * (1+h)
		}
	}
    vs := NewValSorter(scoreMap)
    vs.Sort()
    fmt.Println(vs.Keys)
}

func main() {
    i := &ind{wordToDoc:make(map[string]*wf)}
    i.index("nimish1", "is my name")
    i.index("nimish", "nimish gupta is my name")
    i.index("nimish2", "gupta gupta gupta my name")
    i.queryMatch("gupta is")
}
