package dota

import (
	"log"
	"runtime"
	"strconv"
)

var good = Team{name: "good"}
var evil = Team{name: "evil"}

// 初始化两支队伍
func init() {
	for i, _ := range good.heros {
		good.heros[i].name = "good hero " + strconv.Itoa(i)
		good.heros[i].hp = int64(100)
		good.heros[i].aoe.harm = RandInt64(5, 10)
		good.heros[i].aoe.lastTime = int64(0)
		for j, _ := range good.heros[i].skills {
			good.heros[i].skills[j].harm = RandInt64(10, 15)
			good.heros[i].skills[j].lastTime = int64(0)
			good.heros[i].skills[j].name = "skill " + strconv.Itoa(j)
		}
	}
	for i, _ := range evil.heros {
		evil.heros[i].name = "evil hero " + strconv.Itoa(i)
		evil.heros[i].hp = int64(100)
		evil.heros[i].aoe.harm = RandInt64(5, 10)
		evil.heros[i].aoe.lastTime = int64(0)
		for j, _ := range evil.heros[i].skills {
			evil.heros[i].skills[j].harm = RandInt64(10, 15)
			evil.heros[i].skills[j].lastTime = int64(0)
			evil.heros[i].skills[j].name = "skill " + strconv.Itoa(j)
		}
	}
}

func Run() {
	var obOne, obTwo ob
	runtime.GOMAXPROCS(runtime.NumCPU())
	k := 1
	for {
		r, t := obOne.over(&good, &evil)
		r1, _ := obTwo.over(&good, &evil)
		if r == true && r1 == true {
			log.Println(t + " win!")
			break
		}
		log.Println("第" + strconv.Itoa(k) + "次团战开始---------------------------------------------------")

		log.Println("---------------good-------------")
		for _, h := range good.heros {
			log.Println("name: " + h.name + "  hp: " + strconv.FormatInt(h.hp, 10))
		}
		log.Println("--------------------------------")
		log.Println("---------------evil-------------")
		for _, h := range evil.heros {
			log.Println("name: " + h.name + "  hp: " + strconv.FormatInt(h.hp, 10))
		}
		log.Println("--------------------------------")

		var goodHero = []int{} // good 实际可以参战的英雄下标
		var evilHero = []int{} // evil 实际可以参战的英雄下标

		for i, h := range good.heros {
			if h.hp > int64(0) {
				goodHero = append(goodHero, i)
			}
		}
		for i, h := range evil.heros {
			if h.hp > int64(0) {
				evilHero = append(evilHero, i)
			}
		}

		g := len(goodHero) // good 可以参战的英雄数量
		e := len(evilHero) // evil 可以参战的英雄数量

		log.Println("good 还有" + strconv.Itoa(g) + "位活着的英雄")
		log.Println("evil 还有" + strconv.Itoa(e) + "位活着的英雄")

		randGoodHero := RandSlice(goodHero, g) // shuff后的可以参战的英雄索引
		randEvilHero := RandSlice(evilHero, e) // shuff后的可以参战的英雄索引

		var m, n int

		m = RandIntn(g) + 1 //good 实际参战的英雄数量
		n = RandIntn(e) + 1 //evil 实际参战的英雄数量

		var goodWarrier = []*hero{} // good 实际参战的英雄
		var evilWarrier = []*hero{} // evil 实际参战的英雄

		for _, i := range randGoodHero[:m] {
			goodWarrier = append(goodWarrier, &good.heros[i])
		}

		for _, i := range randEvilHero[:n] {
			evilWarrier = append(evilWarrier, &evil.heros[i])
		}

		log.Println("good 参战" + strconv.Itoa(m) + "位英雄")
		for _, h := range goodWarrier {
			log.Println("name: " + h.name)
		}
		log.Println("evil 参战" + strconv.Itoa(n) + "位英雄")
		for _, h := range evilWarrier {
			log.Println("name: " + h.name)
		}

		c := make(chan bool)

		// 团战开始
		for _, h := range goodWarrier {
			go func(h *hero, c chan bool) {
				defer func(c chan bool) {
					c <- true
				}(c)
				if h.hp <= 0 {
					return
				}
				for _, s := range h.skills {
					var indexArr []int
					k := 0
					for i, w := range evilWarrier {
						if w.hp > 0 {
							indexArr = append(indexArr, i)
							k++
						}
					}
					if k == 0 {
						break
					}
					randArr := RandSlice(indexArr, k)
					ori_hp := evilWarrier[randArr[0]].hp
					s.attack(evilWarrier[randArr[0]])
					log.Println(h.name, "attack", evilWarrier[randArr[0]].name, "技能:", s.name, "伤害:", s.harm, "hp:", ori_hp, "=>", evilWarrier[randArr[0]].hp)
				}
				h.aoe.aoeAttack(evilWarrier)
			}(h, c)
		}

		for _, h := range evilWarrier {
			go func(h *hero, c chan bool) {
				defer func(c chan bool) {
					c <- true
				}(c)
				if h.hp <= 0 {
					return
				}
				for _, s := range h.skills {
					var indexArr []int
					k := 0
					for i, w := range goodWarrier {
						if w.hp > 0 {
							indexArr = append(indexArr, i)
							k++
						}
					}
					if k == 0 {
						break
					}
					randArr := RandSlice(indexArr, k)
					ori_hp := goodWarrier[randArr[0]].hp
					s.attack(goodWarrier[randArr[0]])
					log.Println(h.name, "attack", goodWarrier[randArr[0]].name, "技能:", s.name, "伤害:", s.harm, "hp:", ori_hp, "=>", goodWarrier[randArr[0]].hp)
				}
				h.aoe.aoeAttack(goodWarrier)
			}(h, c)
		}

		for i := 0; i < m+n; i++ {
			<-c
		}

		// 团战结束

		log.Println("---------------good-------------")
		for _, h := range good.heros {
			log.Println("name: " + h.name + "  hp: " + strconv.FormatInt(h.hp, 10))
		}
		log.Println("--------------------------------")
		log.Println("---------------evil-------------")
		for _, h := range evil.heros {
			log.Println("name: " + h.name + "  hp: " + strconv.FormatInt(h.hp, 10))
		}
		log.Println("--------------------------------")
		//time.Sleep(1 * time.Second)
		k++
	}
}
