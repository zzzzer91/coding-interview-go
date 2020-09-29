# 刷题总结

- 算法题就是总结规律（找性质），然后归类（搜索、二分、动归等，也可能不属于某类），然后考虑边界条件。

- 时刻考虑时间复杂度。

- 时刻考虑考虑溢出问题，感觉有溢出可能就用 `long`。

- 使用二分法需要满足性质：1、区间能被分成两部分；2、一半区间满足某一性质，另一半区间不满足同一性质。

- 看见有序数组，首先想到二分法，或者左右指针扫描。

- 二分法解题模板（来自 https://www.acwing.com/blog/content/31/）：

  ```java
  // 类型 1：
  // 当我们将区间[l, r]划分成[l, mid]和[mid + 1, r]时，
  // 其更新操作是r = mid或者l = mid + 1;，计算mid时不需要加1。
  int bsearch1(int l, int r) {
      while (l < r) {
          // 防止溢出
          int mid = l + (r - l) / 2;
          if (check(mid)) {
              // 答案在区间 [l, mid] 中
              r = mid;
          } else {
              // 否则答案在区间 [mid+1, r] 中
              l = mid + 1;
          }
      }
      return l;
  }
  
  // 类型 2：
  // 当我们将区间[l, r]划分成[l, mid - 1]和[mid, r]时，
  // 其更新操作是r = mid - 1或者l = mid;，此时为了防止死循环，计算mid时需要加1。
  int bsearch2(int l, int r) {
      while (l < r) {
          int mid = l + (r - l + 1) / 2;
          if (check(mid)) {
              l = mid;
          } else {
              r = mid - 1;
          }
      }
      return l;
  }
  
  // 与上面对比，有序数组搜索插入，二分查找
  int searchInsert(int[] nums, int target) {
      int l = 0, r = nums.length - 1;
      while (l <= r) { // 注意等于号
          int mid = l + (r - l) / 2;
          int val = nums[mid];
          if (val < target) {
              l = mid + 1;
          } else if (val > target) {
              r = mid - 1;
          } else {
              return mid;
          }
      }
      return l;
  }
  ```

- 看见有重叠子计算的，首先想到动态规划。动态规划中使用一维数组，一般可以把数组优化掉；使用二维数组，一般可以优化成一维数组。 

- 动态规划解题三大步骤：1、定义 `dp` 数组含义，这个数组用来保存历史数据；2、找出数组元素之间的关系式，如计算 `dp[n]` 时，可以利用 `dp[n-1]`；3、找出初始值。 

- 位运算常用：

  ```java
  // 1、判断一个数字 x 二进制下第 i 位是不是等于1
  if ((x & (1 << (i - 1)) > 0) {
      // ...
  }
  
  // 2、把一个数字 x 二进制下第 i 位更改成 1
  x = x | (1 << (i - 1))
  
  // 3、把一个数字 x 二进制下最右边的第一个 1 去掉
  x = x & (x - 1)
      
  // 4、判断数字 x 中只有 1 位为 1
  if ((i & (i - 1)) == 0) {
      // ...
  }
  ```
  
- 链表有关问题中，快慢指针是个好东西。

- 区间问题，首先想到排序，根据左端点或者右断点排序。

- DFS 寻找所有方案（但搜的时候也可以加条件搜某一种），BFS 寻找最优方案。