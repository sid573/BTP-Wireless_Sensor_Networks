package main

import "pkg/pkg"

func main() {
	env := pkg.SDSEnv(10000, 100, 100, 0.81)
	for _, bot := range env.Bots {
		env.RandomAllocation(bot)
	}
	for i := 0; i < 100; i++ {
		for _, bot := range env.Bots {
			env.DetectionProcess(bot)
		}
		env.Communication()
		for _, bot := range env.Bots {
			if bot.Detection == false {
				env.RandomAllocation(bot)
			}
		}
	}
	env.StatsGenerator()
}
