// https://leetcode-cn.com/problems/nim-game/

package main

func main() {

}

// Nim游戏，只要数字总数n，能取得的最小值a和能取得的最大值b加起来的和c。有如下关系。
// n % c == 0 先手必输，前提是可选的数字要是连续的。
// 否则，先手只要抢到余数就必胜。
//
// 为什么呢？
// 其实也很简单，你想啊，c 相当于每次游戏先手可以控制的两人走的最近的一步。
// 什么意思呢，就是说，每次两个人报数，是两个数字，不管对方选择了范围内哪个数字。
// 都可以在第二次报数的时候，把两个人报的数字和凑成c。
//
// 这就是上面我说的前提。
// 比如如果是1，2，3，不管对面报几，你都可以选另一个数字，把总和凑成4。
// 比如如果是1，2，3，4，不管对面报几，你都可以选另一个数字，把总和凑成5。
//
// 所以只要先手抢到了余数，那么剩下的每一轮，先手都可以保证整体前进c。所以必胜。
//
// PS：
// 如果是不连续的，就比较麻烦。
// 比如如果是1，2，4，对面报了2，就凑不到5了。就要分成几个情况讨论了。
func canWinNim(n int) bool {
	return n%4 != 0
}
