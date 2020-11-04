package pkg

import (
	"fmt"
	"math/rand"
)

// Bot for creating each bot
type Bot struct {
	ID          int
	Detection   bool
	RegionID    int
	SubRegionID int
}

// Region for creating different region
type Region struct {
	ID int
	// More information regarding particular region can be stored
}

// SubRegion ...
type SubRegion struct {
	ID        int
	RegionID  int
	Available float64
}

// SDSController to be used as reciever for functions implemented
type SDSController struct {
	Bots       []*Bot
	Regions    []*Region
	SubRegions []*SubRegion
}

// SDSEnv to create environment for simulation
func SDSEnv(botCount int, regionCount int, subRegionCount int) *SDSController {
	bots := make([]*Bot, 0)
	regions := make([]*Region, 0)
	subRegions := make([]*SubRegion, 0)
	// Creating bots for environment
	for i := 0; i < botCount; i++ {
		bot := &Bot{
			ID: i,
		}
		bots = append(bots, bot)
	}
	// Creating Region and SubRegions
	for i := 0; i < regionCount; i++ {
		region := &Region{
			ID: i,
		}
		for j := 0; j < subRegionCount; j++ {
			available := rand.Float64()
			subRegion := &SubRegion{
				ID:        j,
				RegionID:  i,
				Available: available,
			}
			subRegions = append(subRegions, subRegion)
		}
		regions = append(regions, region)
	}
	controller := &SDSController{
		Bots:       bots,
		Regions:    regions,
		SubRegions: subRegions,
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
	fmt.Println("Details of all sub-regions in the environment")
	for _, subRegion := range env.SubRegions {
		fmt.Println(*subRegion)
	}
}
