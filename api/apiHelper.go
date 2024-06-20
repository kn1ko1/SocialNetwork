package api

import (
	"math/rand"
	"socialnetwork/repo"
	"strconv"
	"time"
)

var CurrentTime = time.Now()

var Timestamp = CurrentTime.Unix()

var R = repo.NewDummyRepository()

var RandomNumberInt = rand.Intn(1000000)

var RandomNumberStr = strconv.Itoa(rand.Intn(1000000))
