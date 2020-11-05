package main

import "pkg/pkg"

func main() {
	env := pkg.SDSEnv(10, 20, 10, 0.5)
	for _, bot := range env.Bots {
		env.RandomAllocation(bot)
	}
	for _, bot := range env.Bots {
		env.DetectionProcess(bot)
	}
	env.Communication()
	env.DisplayEnv()
}
