package main

import (
	"flag"
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var VERSION = "0.1"
var PROJECTNAME = "collatzConjecture"

func main() {
	fmt.Printf("%s %s ", PROJECTNAME, VERSION)

	flag.Parse()
	numberCLIArg := flag.Arg(0)
	var value int
	fmt.Println(value)

	_, err := fmt.Sscan(numberCLIArg, &value)

	if err != nil {
		fmt.Println(err)
	}
	var newValue int = value
	values := buildArray(newValue)
	printGraph(value, values, numberCLIArg)

}

func buildArray(newValue int) []int {
	values := []int{}
	values = append(values, newValue)
	for {
		if newValue != 1 {
			if newValue%2 == 0 {
				newValue = isEven(newValue)
			} else {
				newValue = isOdd(newValue)
			}
			fmt.Println(newValue)
			values = append(values, newValue)
		} else {
			break
		}
	}
	return values
}

func isOdd(number int) int {
	calcresult := 3*number + 1

	return calcresult
}

func isEven(number int) int {
	calcresult := number / 2

	return calcresult
}

func printGraph(value int, values []int, numberCLIArg string) {
	graph := plot.New()
	title := "Collatz Conjecture - " + numberCLIArg
	graph.Title.Text = title
	graph.X.Label.Text = "X"
	graph.Y.Label.Text = "Y"

	coords := make(plotter.XYs, len(values))
	for i := range coords {
		coords[i].X = float64(i)
		coords[i].Y = float64(values[i])

	}
	fmt.Println(coords)

	err := plotutil.AddLinePoints(graph, "", coords)
	if err != nil {
		panic(err)
	}
	//save as PNG
	if err := graph.Save(8*vg.Inch, 4*vg.Inch, title+".png"); err != nil {
		panic(err)
	}
}
