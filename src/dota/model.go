package dota

import (
	//"log"
	"time"
)

func (s *skill) attack(h *hero) {
	nowTime := time.Now().Unix() //当前时间戳

	if nowTime-s.lastTime < s.harm {
		return //技能处于冷却状态
	}

	s.lastTime = nowTime

	// begin miss的概率
	randArr := RandPerm(100)

	min := 10
	max := 20

	if randArr[0] >= min && randArr[0] <= max {
		//log.Println("miss")
		return //miss
	}
	// end miss的概率
	//ori_hp := h.hp

	h.hp = h.hp - s.harm

	if h.hp < 0 {
		h.hp = 0
	}

	//log.Println(ori_hp, "=>", h.hp)
}

func (s *skill) aoeAttack(heros []*hero) {
	nowTime := time.Now().Unix()

	if nowTime-s.lastTime < 2*s.harm {
		return // aoe技能处在冷却时间
	}

	s.lastTime = nowTime

	for k, h := range heros {

		// begin miss的概率
		randArr := RandPerm(100)
		min := 15
		max := 30
		if randArr[0] >= min && randArr[0] <= max {
			//log.Println("miss")
			return //miss
		}
		// end miss的概率

		//ori_hp := heros[k].hp

		heros[k].hp = h.hp - s.harm

		if heros[k].hp < 0 {
			heros[k].hp = 0
		}
		//log.Println(ori_hp, "=>", heros[k].hp)
	}
}

func (*ob) over(a, b *Team) (r bool, t string) {
	i, j := 0, 0
	for _, h := range a.heros {
		if h.hp <= int64(0) {
			i++
		}
	}
	for _, h := range b.heros {
		if h.hp <= int64(0) {
			j++
		}
	}
	if i == 5 && j == 5 {
		return true, "both"
	} else if i == 5 {
		return true, b.name
	} else if j == 5 {
		return true, a.name
	} else {
		return false, "none"
	}
}

type Team struct {
	name  string
	heros [5]hero
}

type hero struct {
	name   string
	hp     int64
	skills [4]skill
	aoe    skill
}

type ob struct {
}

type skill struct {
	name     string
	harm     int64
	lastTime int64
}
