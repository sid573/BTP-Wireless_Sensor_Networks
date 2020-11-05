package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomAllocation for those bots who are
// in passive state at initial stage and for
// those bot who are still in passive state after
// communication
func (env *SDSController) RandomAllocation(bot *Bot) {
	rand.Seed(time.Now().UnixNano())
	regionCount := len(env.Regions)
	subRegionCount := len(env.SubRegions) / regionCount
	// Randomly allocating region and subregion
	regionID := rand.Intn(regionCount) + 1
	subRegionID := rand.Intn(subRegionCount) + 1
	bot.RegionID = regionID
	bot.SubRegionID = subRegionID
	env.SubRegionOverlap(bot, regionID, subRegionID)
}

// DetectionProcess is the step in which all bots
// will compute at the allocated subregion inside
// the region
func (env *SDSController) DetectionProcess(bot *Bot) {
	for _, subRegion := range env.SubRegions {
		if bot.RegionID == subRegion.RegionID && bot.SubRegionID == subRegion.ID {
			if env.Limit < subRegion.Count {
				bot.Detection = true
				subRegion.Available = false
				break
			}
			bot.RegionID = 0
			bot.SubRegionID = 0
			bot.Detection = false
			break
		}
	}
}

// SubRegionAllocation after communication in Region
// is allocated then inside that a subregion is selected
// at random
func (env *SDSController) SubRegionAllocation(bot *Bot) {
	rand.Seed(time.Now().UnixNano())
	regionCount := len(env.Regions)
	subRegionCount := len(env.SubRegions) / regionCount
	subRegionID := rand.Intn(subRegionCount) + 1
	bot.SubRegionID = subRegionID
	env.SubRegionOverlap(bot, bot.RegionID, subRegionID)
}

// Communication it is the step at which
// one bot can only communicate with other
// and if his detection is true then in same region
// the other bot is given a random subregion
// to detect
func (env *SDSController) Communication() {
	orderCommunication := GenerateRandomNumber(len(env.Bots))
	for i := 0; i < len(orderCommunication); i = i + 2 {
		if i+1 == len(orderCommunication) {
			break
		}
		if env.Bots[orderCommunication[i]].Detection == true && env.Bots[orderCommunication[i+1]].Detection == false {
			env.Bots[orderCommunication[i+1]].RegionID = env.Bots[orderCommunication[i]].RegionID
			env.SubRegionAllocation(env.Bots[orderCommunication[i+1]])
		}
		if env.Bots[orderCommunication[i+1]].Detection == true && env.Bots[orderCommunication[i]].Detection == false {
			env.Bots[orderCommunication[i]].RegionID = env.Bots[orderCommunication[i+1]].RegionID
			env.SubRegionAllocation(env.Bots[orderCommunication[i]])
		}
	}

}

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

// SubRegionOverlap ...
func (env *SDSController) SubRegionOverlap(bot *Bot, regionID int, subRegionID int) {
	for _, subRegion := range env.SubRegions {
		if subRegion.RegionID == regionID && subRegion.ID == subRegionID {
			if subRegion.Available == false {
				bot.RegionID = 0
				bot.SubRegionID = 0
				bot.Detection = false
			}
			if subRegion.Available == true {
				subRegion.Available = false
			}
			break
		}
	}
}
