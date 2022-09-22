package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/Vladosya/our_project/appl_row"
	"github.com/bradfitz/iter"
	"math/rand"
	"strings"
	"time"
)

func getCurrentDate() time.Time {
	return time.Now()
}

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

func CalculateByTableNums(participantCount int) int { // посчитать кол-во столов от кол-ва участников
	if participantCount == 0 {
		return 0
	}
	if participantCount < 9 {
		return 1
	} else {
		var numOfTab = participantCount % 8
		if numOfTab == 0 {
			return participantCount / 8
		} else {
			return (participantCount / 8) + 1
		}
	}
}

func CalculateByTableLimit(limitTable int, currentParticipant int) bool { // проверка на ограничение по количеству столов (если true, то разрешена регистрация на мероприятие, иначе false)
	if (limitTable * 8) > currentParticipant {
		return true
	}
	if (limitTable * 8) <= currentParticipant {
		return false
	}
	return false
}

func GenSeatAtTableByTableLen(needTable int) []byte { // ф-ция, которая генерирует json форму для созданного мероприятия, где администратор будет рассаживать игроков в зависимости от того, сколько столов мы передали
	if needTable == 0 {
		return []byte{}
	} else {
		var candidates []appl_row.SeatAtTables
		for i := range iter.N(needTable) {
			candidates = append(candidates, appl_row.SeatAtTables{
				Id:   i + 1,
				Name: fmt.Sprintf("Стол %d", i+1),
				SeatAtTable: []appl_row.SeatAtTable{
					{
						Id:     1,
						Name:   nil,
						IdUser: nil,
						Rank:   nil,
					},
					{
						Id:     2,
						Name:   nil,
						IdUser: nil,
						Rank:   nil,
					},
					{
						Id:     3,
						Name:   nil,
						IdUser: nil,
						Rank:   nil,
					},
					{
						Id:     4,
						Name:   nil,
						IdUser: nil,
						Rank:   nil,
					},
					{
						Id:     5,
						Name:   nil,
						IdUser: nil,
						Rank:   nil,
					},
					{
						Id:     6,
						Name:   nil,
						IdUser: nil,
						Rank:   nil,
					},
					{
						Id:     7,
						Name:   nil,
						IdUser: nil,
						Rank:   nil,
					},
					{
						Id:     8,
						Name:   nil,
						IdUser: nil,
						Rank:   nil,
					},
				},
			})
		}
		jsonData, err := json.Marshal(candidates)
		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return []byte{}
		}
		fmt.Println("jsonData -->", jsonData)
		return jsonData
	}
}
