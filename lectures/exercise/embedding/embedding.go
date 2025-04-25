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

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (c *CpuTemp) AverageCpu() int {
	sumParam := 0
	for i := 0; i < len(c.temp); i++ {
		sumParam += int(c.temp[i])
	}

	return sumParam / len(c.temp)
}

func (m *MemoryUsage) AverageMem() int {
	sumParam := 0
	for i := 0; i < len(m.amount); i++ {
		sumParam += int(m.amount[i])
	}

	return sumParam / len(m.amount)
}

func (b *BandwidthUsage) AverageBand() int {
	sumParam := 0
	for i := 0; i < len(b.amount); i++ {
		sumParam += int(b.amount[i])
	}

	return sumParam / len(b.amount)
}

type Dashboard struct {
	CpuTemp
	MemoryUsage
	BandwidthUsage
}

func (d Dashboard) String() string {
	return fmt.Sprintf("Dashboard: CPU temp: %v, Memory Usage: %v, Bandwidth Usage: %v", d.temp, d.MemoryUsage.amount, d.BandwidthUsage.amount)
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{
		CpuTemp:        temp,
		MemoryUsage:    memory,
		BandwidthUsage: bandwidth,
	}

	fmt.Printf("Average usage: CPU %v, Memory %v, Bandwidth %v", dash.AverageCpu(), dash.AverageMem(), dash.AverageBand())

}
