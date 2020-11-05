package pkg

import (
	"math/rand"
)

// RandomAllocation for those bots who are
// in passive state at initial stage and for
// those bot who are still in passive state after
// communication
func (env *SDSController) RandomAllocation(bot *Bot) {
	regionCount := len(env.Regions)
	subRegionCount := len(env.SubRegions) / regionCount
	// Randomly allocating region and subregion
	regionID := rand.Intn(regionCount) + 1
	subRegionID := rand.Intn(subRegionCount) + 1
	bot.RegionID = regionID
	bot.SubRegionID = subRegionID
}

// DetectionProcess is the step in which all bots
// will compute at the allocated subregion inside
// the region
func (env *SDSController) DetectionProcess(bot *Bot) {
	for _, subRegion := range env.SubRegions {
		if bot.RegionID == subRegion.RegionID && bot.SubRegionID == subRegion.ID {
			if env.Limit >= subRegion.Available {
				bot.Detection = true
			}
		}
	}
}

// SubRegionAllocation after communication in Region
// is allocated then inside that a subregion is selected
// at random
func (env *SDSController) SubRegionAllocation(bot *Bot) {
	regionCount := len(env.Regions)
	subRegionCount := len(env.SubRegions) / regionCount
	subRegionID := rand.Intn(subRegionCount) + 1
	bot.SubRegionID = subRegionID
}

// Communication ...
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
	res := rand.Perm(size)
	return res
}
