// 小A是某公司的员工，在一次公司年会上，
// 主持人宣布进行一项游戏来活跃年会气氛，游戏规则如下：
// n个人随机站成一排，按照他们所站的顺序依次给他们编号从1到n，
// 接下来就从编号为1的人开始，按从左到右的顺序每隔一人选出一个人，
// 选出的这些被淘汰，剩下的需要重新站成一排，其中首尾是接龙的，
// 即如果倒数第二个被淘汰，则隔一人即第一个人被淘汰，
// 如此循环一直到最后剩下两个人为止，那么这剩下的最后两个人就是本场晚会的幸运儿，
// 得到神秘大奖，小A想成为这个幸运儿，
// 请你帮小A算出来开始时他应该站在什么位置才最终可以成为幸运儿。（3<=n<=50）
//
// 输入描述
// 开始的进行游戏的总人数n
// 输出描述
// 第一行是选出顺序，第二行是两名幸运儿的初始位置（按升序排列），要求位置编号之间用一个空格空开。
//
// 样例输入
// 5
// 样例输出
// 2 4 1
// 5 3

package main

import "fmt"
import "strings"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]bool, n)
	i := 1
	count := n
	var strs1, strs2 []string
	for count > 0 {
		if count > 2 {
			a[i] = true
			strs1 = append(strs1, fmt.Sprintf("%d", i+1))
		} else {
			strs2 = append(strs2, fmt.Sprintf("%d", i+1))
		}
		for j := 0; j < 2; j++ {
			for a[i] {
				i = (i + 1) % n
			}
			if j == 0 {
				i = (i + 1) % n
			}
		}
		count--
	}
	fmt.Println(strings.Join(strs1, " "))
	fmt.Println(strings.Join(strs2, " "))
}
