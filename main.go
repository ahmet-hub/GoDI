package main

import "fmt"

type RockClimber struct {
	rocksClimbed int
	sp           SafetyPlacer
}

func newRockClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{
		sp: sp,
	}
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

type SafetyPlacer interface {
	placeSafeties()
}

type IceSafetyPlacer struct{}

func (rc IceSafetyPlacer) placeSafeties() {
	fmt.Println("placing my Ice safeties...")
}

type NOPSafetyPlacer struct{}

func (sp NOPSafetyPlacer) placeSafeties() {
	fmt.Println("placing NO safeties...")
}

func main() {
	rc := newRockClimber(NOPSafetyPlacer{})

	for i := 0; i < 11; i++ {
		rc.climbRock()
	}
}
