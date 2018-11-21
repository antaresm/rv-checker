package config

import (
	"time"
	"os"
)

var StartDelay time.Duration
var CheckTimer time.Duration
var Timeout time.Duration

func init() {
	s := os.Getenv("START_DELAY")
	StartDelay, _ = time.ParseDuration(s)

	s = os.Getenv("CHECK_TIMER")
	CheckTimer, _ = time.ParseDuration(s)
}