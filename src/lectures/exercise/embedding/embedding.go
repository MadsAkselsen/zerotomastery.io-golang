//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

func (s BandwidthUsage) String() string {
	return "BandwidthUsage"
}

type CpuTemp struct {
	temp []Celcius
}

func (s CpuTemp) String() string {
	return "CpuTemp"
}

type MemoryUsage struct {
	amount []Bytes
}

func (s MemoryUsage) String() string {
	return "MemoryUsage"
}

func (b BandwidthUsage) CalcAvg() int {
	var sum Bytes
	for i := 0; i < len(b.amount); i++ {
		amount := b.amount[i]	
		sum += amount
	}

	return int(sum)/len(b.amount)
}

func (b CpuTemp) CalcAvg() int {
	var sum Celcius
	for i := 0; i < len(b.temp); i++ {
		amount := b.temp[i]	
		sum += amount
	}
	return int(sum)/len(b.temp)
}

func (b MemoryUsage) CalcAvg() int {
	var sum Bytes
	for i := 0; i < len(b.amount); i++ {
		amount := b.amount[i]	
		sum += amount
	}
	return int(sum)/len(b.amount)
}

type Calculator interface {
	CalcAvg() int
}

func Calculate(item Calculator) {
	// fmt.Printf("Avg for v% is %v", item, item.CalcAvg())
	fmt.Println(item.CalcAvg())
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{bandwidth, temp, memory}
	Calculate(dashboard.CpuTemp)
}
