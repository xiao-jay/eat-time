package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var food = "馄饨 拉面 烩面 热干面 刀削面 油泼面 炸酱面 炒面 重庆小面 米线 酸辣粉 土豆粉 螺狮粉 凉皮儿 麻辣烫 肉夹馍 羊肉汤 炒饭 盖浇饭 卤肉饭 烤肉饭 黄焖鸡米饭 驴肉火烧 川菜 麻辣香锅 火锅 酸菜鱼 烤串 披萨 烤鸭 汉堡 炸鸡 寿司 蟹黄包 煎饼果子 生煎 炒年糕"

func main() {
	fmt.Println()
	foods := strings.Split(food, " ")
	for {
		hour := time.Now().Hour()
		min := time.Now().Minute()
		time.Now().Weekday()
		if (hour == 18 && min <= 0) || (hour == 12 && min <= 0) && !(time.Now().Weekday().String() == "Saturday") && !(time.Now().Weekday().String() == "Sunday") {
			idx := rand.Intn(len(foods))
			SendMessage("吃饭啦！！\n 今天吃：" + foods[idx])
		}
		time.Sleep(time.Minute)
	}
}
