package main

import "pkg/pkg"

func main() {
	env := pkg.SDSEnv(10, 20, 10)
	env.DisplayEnv()
}
