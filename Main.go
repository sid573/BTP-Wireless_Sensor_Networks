package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"pkg/pkg"
	"strconv"
)

func main() {
	env := pkg.SDSEnv(1000, 1000, 100, 0.9, 100)
	// for ploting purpose
	activeBot := make([]float64, 0)
	iteration := make([]float64, 0)
	plotPoints := make([][]float64, 0)
	// to save communication data
	file, err := os.Create("sim3-communication1")
	if err != nil {
		log.Fatal(err)
	}
	file2, err2 := os.Create("sim3-activeBotStats")
	if err2 != nil {
		log.Fatal(err)
	}
	for _, bot := range env.Bots {
		env.RandomAllocation(bot)
	}
	for i := 0; i < 100; i++ {
		for _, bot := range env.Bots {
			env.DetectionProcess(bot)
		}
		env.Communication(file, i)
		for _, bot := range env.Bots {
			if bot.Detection == false {
				env.RandomAllocation(bot)
			}
		}
		act := env.StatsGenerator(0)
		activeBot = append(activeBot, act)
		iteration = append(iteration, float64(i))
		mw := io.MultiWriter(os.Stdout, file2)
		fmt.Fprintln(mw, "-----------------------------------------")
		fmt.Fprintln(mw, "Ratio of active bots in system after iteration "+strconv.Itoa(i))
		fmt.Fprintln(mw, act)
		fmt.Fprintln(mw, "-----------------------------------------")
	}
	plotPoints = append(plotPoints, iteration)
	plotPoints = append(plotPoints, activeBot)
	pkg.PlotFunction("sim3-fig1.png", plotPoints)
	env.StatsGenerator(1)
}
