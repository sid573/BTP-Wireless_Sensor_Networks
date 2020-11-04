package pkg

import (
	"fmt"
	"math/rand"
)

// Bot for creating each bot
type Bot struct {
	ID        int
	Detection bool
}

// Region for creating different region
type Region struct {
	ID          int
	SubRegionID int
	Available   float64
	BotCount    int
}

// SDSController to be used as reciever for functions implemented
type SDSController struct {
	Bots    []*Bot
	Regions []*Region
}

// SDSEnv to create environment for simulation
func SDSEnv(botCount int, regionCount int, subRegionCount int) *SDSController {
	bots := make([]*Bot, 0)
	regions := make([]*Region, 0)
	// Creating bots for environment
	for i := 0; i < botCount; i++ {
		bot := &Bot{
			ID: i,
		}
		bots = append(bots, bot)
	}
	// Creating Region and SubRegions
	for i := 0; i < regionCount; i++ {
		for j := 0; j < subRegionCount; j++ {
			available := rand.Float64()
			region := &Region{
				ID:          i,
				SubRegionID: j,
				Available:   available,
			}
			regions = append(regions, region)
		}
	}
	controller := &SDSController{
		Bots:    bots,
		Regions: regions,
	}
	return controller
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
}
