// https://www.acwing.com/problem/content/description/284/
// 区间dp
// 想象这里的区间时，要从离散的角度去想，想象一堆一堆的石头
// 比如 s[1]，也可以单独取
//
// 所有的区间dp问题，第一维都是枚举区间长度，一般 len = 1 用来初始化，
// 枚举从 len = 2 开始，第二维枚举起点 i （右端点 j 自动获得，j = i + len - 1）
//
//```
//for (int i = 1; i <= n; i++) {
//    dp[i][i] = 初始值
//}
//for (int len = 2; len <= n; len++)           //区间长度
//    for (int i = 1; i + len - 1 <= n; i++) { //枚举起点
//        int j = i + len - 1;                 //区间终点
//        for (int k = i; k < j; k++) {        //枚举分割点，构造状态转移方程
//            dp[i][j] = max(dp[i][j], dp[i][k] + dp[k + 1][j] + w[i][j]);
//        }
//    }
//```

package main

import (
	"fmt"
	"math"
)

const N = 301

var s [N]int
var dp [N][N]int

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ { // 从 1 开始
		fmt.Scan(&s[i])
		s[i] += s[i-1]
	}
	for space := 2; space <= n; space++ { // 取石子的范围，从连续的几堆石子中取，从小到大，因为至少两堆石子合并，所以2开始，比如s[1]和s[2]两个点，他们的间距是2，
		for l := 1; l+space-1 <= n; l++ { // 范围是 1 到 n - space - 1，因为要留空间给间距
			r := l + space - 1       // 右端点
			dp[l][r] = math.MaxInt64 // 下面有 min 比较，所以先把值置为最大
			for k := l; k < r; k++ { // 注意这里的 k 从 l 开始，要这么想象，s[k] 一堆，s[k+1:]一堆
				dp[l][r] = min(dp[l][r], dp[l][k]+dp[k+1][r]+s[r]-s[l-1]) // s[r]-s[l-1] 是最后两堆石子合并的代价，按照题意，正好是所有石子的总重量
			}
		}
	}
	fmt.Println(dp[1][n]) // [1, n]，最小值
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
