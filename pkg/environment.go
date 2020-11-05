package pkg

import (
	"math/rand"
	"time"
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
	Count     float64
	Available bool
}

// SDSController to be used as reciever for functions implemented
type SDSController struct {
	Bots       []*Bot
	Regions    []*Region
	SubRegions []*SubRegion
	Limit      float64
}

// SDSEnv to create environment for simulation
func SDSEnv(botCount int, regionCount int, subRegionCount int, limit float64) *SDSController {
	bots := make([]*Bot, 0)
	regions := make([]*Region, 0)
	subRegions := make([]*SubRegion, 0)
	// Creating bots for environment
	for i := 1; i <= botCount; i++ {
		bot := &Bot{
			ID: i,
		}
		bots = append(bots, bot)
	}
	// Creating Region and SubRegions
	for i := 1; i <= regionCount; i++ {
		region := &Region{
			ID: i,
		}
		for j := 1; j <= subRegionCount; j++ {
			rand.Seed(time.Now().UnixNano())
			available := rand.Float64()
			subRegion := &SubRegion{
				ID:        j,
				RegionID:  i,
				Count:     available,
				Available: true,
			}
			subRegions = append(subRegions, subRegion)
		}
		regions = append(regions, region)
	}
	controller := &SDSController{
		Bots:       bots,
		Regions:    regions,
		SubRegions: subRegions,
		Limit:      limit,
	}
	return controller
}
