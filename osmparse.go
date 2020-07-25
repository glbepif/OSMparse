package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
	"os"
)

func main() {
	var f, err = os.Open("C:/Users/Глеб/go/src/OSM/RU.highways.osm.pbf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := osmpbf.New(context.Background(), f, 3)
	defer scanner.Close()

	outFile, err := os.Create("out.txt")
	defer outFile.Close()

	i := 0
	for scanner.Scan() {
		data := scanner.Object()
		if err != nil {
			fmt.Println(err)
		}

		if data.ObjectID().Type() != osm.TypeWay {
			if (data.ObjectID().Type() == osm.TypeNode) {
				node := data.(*osm.Node)
				//1881001201,1881001200,365464328,1845336102,1449582594,947067468,1429798792,2054601347,1827277136,1429801834
				d, _ := json.Marshal(node)
				if (node.ID == 2694776285 || node.ID == 2694776287) {
					_, _ = outFile.WriteString(string(d))
					_, _ = outFile.WriteString("\n")
					i++
				}
			}
			continue
		}
		w := data.(*osm.Way)

		name := w.Tags.Find("addr:city")
		if name != "Челябинск" {
			continue
		}

		d, _ := json.Marshal(w)

		_, _ = outFile.WriteString(string(d))
		_, _ = outFile.WriteString("\n")

		if i == 1 {
			return
		}
	}
	scanErr := scanner.Err()
	if scanErr != nil {
		panic(scanErr)
	}
}
