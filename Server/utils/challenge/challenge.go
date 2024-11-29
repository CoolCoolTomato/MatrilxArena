package challenge

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"math"
)

// DynamicScore 动态计算分数
// s 初始分数
// r 最低分数比率
// d 衰减速度
// x 解题率
func DynamicScore(s float64, r float64, d float64, x float64) int {
	exponentialPart := (1 - math.Exp((1-x)*d)) / (1 - math.Exp(d))
	result := s * (r + (1-r)*exponentialPart)
	return int(math.Floor(result))
}

func CalculateCTFChallengeScore(ctfChallenge model.CTFChallenge) model.CTFChallenge {
	s := float64(ctfChallenge.Score)
	r := 0.1
	d := float64(1)
	var x float64
	if ctfChallenge.CTF.TeamSignIn {
		x = float64(len(ctfChallenge.CTFTeams)) / float64(len(ctfChallenge.CTF.CTFTeams))
	} else {
		x = float64(len(ctfChallenge.Users)) / float64(len(ctfChallenge.CTF.Users))
	}
	ctfChallenge.Score = DynamicScore(s, r, d, x)
	return ctfChallenge
}

func CalculateCTFChallengeListScore(ctfChallengeList []model.CTFChallenge) []model.CTFChallenge {
	for index, ctfChallenge := range ctfChallengeList {
		ctfChallengeList[index] = CalculateCTFChallengeScore(ctfChallenge)
	}
	return ctfChallengeList
}
