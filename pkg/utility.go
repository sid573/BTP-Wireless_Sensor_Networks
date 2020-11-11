package pkg

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Arafatk/glot"
)

// GenerateRandomNumber will generate a slice
// with a random permutation on 1 .... n
func GenerateRandomNumber(size int) []int {
	rand.Seed(time.Now().UnixNano())
	res := rand.Perm(size)
	return res
}

// StatsGenerator generates all the stats
func (env *SDSController) StatsGenerator(disp int) float64 {
	botCountActive := 0
	actualRegion := 0
	region := make(map[int]int)
	for _, bot := range env.Bots {
		if bot.Detection == true {
			botCountActive++
			region[bot.RegionID]++
		}
	}

	for _, reg := range env.SubRegions {
		if env.Limit < reg.Count {
			actualRegion++
		}
	}
	percentageBotActive := float64(botCountActive) / float64(len(env.Bots))
	if disp == 1 {
		fmt.Println("---------------------------")
		fmt.Println("Active Bots in the System")
		fmt.Println(percentageBotActive)
		fmt.Println("---------------------------")
		fmt.Println("Region Details")
		fmt.Println(region)
		fmt.Println("---------------------------")
		fmt.Println("Total Subregion with excess count")
		fmt.Println(actualRegion)
		fmt.Println("---------------------------")
	}
	return percentageBotActive

}

// DisplayEnv will show all details of bots
// and region of the environment
func (env *SDSController) DisplayEnv() {
	fmt.Println("Details of all bots in the environment")
	for _, bot := range env.Bots {
		fmt.Println(*bot)
	}
	fmt.Println("Details of all regions in the environment")
	for _, region := range env.Regions {
		fmt.Println(*region)
	}
	fmt.Println("Details of all sub-regions in the environment")
	for _, subRegion := range env.SubRegions {
		fmt.Println(*subRegion)
	}
}

// PlotFunction ...
func PlotFunction(file string, points [][]float64) {
	dimensions := 2
	// The dimensions supported by the plot
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	pointGroupName := "line"
	style := "lines"
	// points := [][]float64{{7, 3, 13, 5.6, 11.1}, {12, 13, 11, 1, 7}}
	// Adding a point group
	plot.AddPointGroup(pointGroupName, style, points)
	// A plot type used to make points/ curves and customize and save them as an image.
	plot.SetTitle("Active Bots in system")
	// Optional: Setting the title of the plot
	plot.SetXLabel("Iterations")
	plot.SetYLabel("Ratio of active bots")
	// Optional: Setting label for X and Y axis
	plot.SetXrange(0, 100)
	plot.SetYrange(0, 1)
	// Optional: Setting axis ranges
	plot.SavePlot(file)
}

// RandFloat ...
func RandFloat(min float64, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*max
}
