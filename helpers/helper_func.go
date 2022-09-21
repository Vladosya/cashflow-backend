package helpers

import (
	"math/rand"
	"strings"
	"time"
)

func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func RandomStrGeneration(needLength int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := needLength
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return strings.ToLower(b.String())
}

func ReturnJSONB(needJson string) []byte {
	s := []byte(needJson)
	return s
}

func testTableNums(participantCount int) int { // посчитать кол-во столов от кол-ва участников
	if participantCount == 0 {
		return 0
	}
	if participantCount < 8 {
		return 1
	} else {
		var numOfTab = participantCount % 7
		if numOfTab == 0 {
			return participantCount / 7
		} else {
			return (participantCount / 7) + 1
		}
	}
}
