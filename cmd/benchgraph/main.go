package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/tools/benchmark/parse"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	bs := parseFile(os.Args[1])
	var groups map[string]map[string]*parse.Benchmark
	groups = make(map[string]map[string]*parse.Benchmark)
	for name, bb := range bs {
		parent, sub := extractParent(name)
		// group by parent
		//fmt.Print(parent, " ", sub, len(bb), " ")
		if len(bb) != 1 {
			fmt.Printf("expect only one benchmark, but got %d", len(bb))
		}
		b := bb[0]
		//fmt.Print(b.N, b.NsPerOp, b.AllocedBytesPerOp, b.AllocsPerOp, b.MBPerS, b.Measured, b.Ord, "\n")
		group, ok := groups[parent]
		if !ok {
			group = make(map[string]*parse.Benchmark)
		}
		group[sub] = b
		groups[parent] = group
	}
	fmt.Println("groups", len(groups))
	charts := ECharts{Title: os.Args[1]}
	// one graph for each group
	for groupName, group := range groups {
		// nanoseconds per iteration
		nsChart := EChartOption{
			// TODO: groupName might not be a valid javascript variable name, sanitize?
			Name:   groupName + "Ns",
			Title:  groupName + " NsPerOp",
			Legend: []string{"nanosecond per iteration"},
		}
		bChart := EChartOption{
			Name:   groupName + "Bytes",
			Title:  groupName + " BytesPerOp",
			Legend: []string{"bytes allocated per iteration"},
		}
		nsSeries := Series{
			Name: "nanosecond per iteration",
			Type: "bar",
		}
		bSeries := Series{
			Name: "bytes allocated per iteration",
			Type: "bar",
		}
		// TODO: sort by sub ...
		for subName, b := range group {
			nsChart.XAxis = append(nsChart.XAxis, subName)
			nsSeries.Data = append(nsSeries.Data, float64(b.NsPerOp))
			bChart.XAxis = append(bChart.XAxis, subName)
			bSeries.Data = append(bSeries.Data, float64(b.AllocedBytesPerOp))
		}
		nsChart.Series = append(nsChart.Series, nsSeries)
		bChart.Series = append(bChart.Series, bSeries)
		charts.Charts = append(charts.Charts, nsChart, bChart)
	}
	b, err := charts.Render()
	if err != nil {
		fatal(err)
	}
	if err := ioutil.WriteFile("zap.html", b, os.ModePerm); err != nil {
		fatal(err)
	}
	fmt.Println("graph zap.html generated")
}

func usage() {
	fmt.Printf("usage: %s bench.txt\n\n", os.Args[0])
}

func fatal(msg interface{}) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseFile(path string) parse.Set {
	f, err := os.Open(path)
	if err != nil {
		fatal(err)
	}
	defer f.Close()
	b, err := parse.ParseSet(f)
	if err != nil {
		fatal(err)
	}
	return b
}

func extractParent(bechmarkName string) (string, string) {
	slash := strings.Index(bechmarkName, "/")
	if slash == -1 {
		return bechmarkName, ""
	}
	return bechmarkName[:slash], bechmarkName[slash+1:]
}
