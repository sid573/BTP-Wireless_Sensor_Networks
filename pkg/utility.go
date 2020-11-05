package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandomNumber will generate a slice
// with a random permutation on 1 .... n
func GenerateRandomNumber(size int) []int {
	rand.Seed(time.Now().UnixNano())
	res := rand.Perm(size)
	return res
}

// StatsGenerator generates all the stats
func (env *SDSController) StatsGenerator() {
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
