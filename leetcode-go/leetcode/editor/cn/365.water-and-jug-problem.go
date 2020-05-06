package main

/*
[365]水壶问题
//有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？
//
// 如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
//
// 你允许：
//
//
// 装满任意一个水壶
// 清空任意一个水壶
// 从一个水壶向另外一个水壶倒水，直到装满或者倒空
//
//
// 示例 1: (From the famous "Die Hard" example)
//
// 输入: x = 3, y = 5, z = 4
//输出: True
//
//
// 示例 2:
//
// 输入: x = 2, y = 6, z = 5
//输出: False
//
// Related Topics 数学
*/

//leetcode submit region begin(Prohibit modification and deletion)
//记录当前x,y壶盛水量
type State struct {
	X int
	Y int
}

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	if x+y < z {
		return false
	}
	return bfs(x, y, z)
}

/*

装满任意一个水壶，
（1）装满 A，包括 A 为空和 A 非空的时候把 A 倒满的情况；
（2）装满 B，包括 B 为空和 B 非空的时候把 B 倒满的情况。
清空任意一个水壶，
（1）清空 A，B不变
（2）清空 B，A不变
从一个水壶向另外一个水壶倒水，直到装满或者倒空，（四选二）
（1）从 A 到 B，使得 B 满，A 还有剩；
（2）从 A 到 B，此时 A 的水太少，A 倒尽，B 没有满；
（3）从 B 到 A，使得 A 满，B 还有剩余；
（4）从 B 到 A，此时 B 的水太少，B 倒尽，A 没有满。
当前状态可能延伸八个状态，且状态存在循环，即有向有环图，类似于java类加载排序
循环{从初始状态开始，遍历所有状态，加入队列，逐个取出，查看结果是否为z}，另需要额外的map记录遍历过的节点，避免环。
*/
func bfs(x, y, z int) bool {
	var queue []State
	initState := State{
		X: 0,
		Y: 0,
	}
	queue = append(queue, initState)
	visitMap := make(map[State]bool)

	for len(queue) != 0 {
		now := queue[0]
		queue = queue[1:]
		if visitMap[now] {
			continue
		}
		visitMap[now] = true
		if now.X == z || now.Y == z || now.X+now.Y == z {
			return true
		}

		next := State{
			X: x,
			Y: now.Y,
		}
		queue = append(queue, next)
		next = State{
			X: now.X,
			Y: y,
		}
		queue = append(queue, next)
		next = State{
			X: now.X,
			Y: 0,
		}
		queue = append(queue, next)
		next = State{
			X: 0,
			Y: now.Y,
		}
		queue = append(queue, next)
		//x的水倒入y， y存在剩余可装下x的水和装不完的情况
		if now.X > y-now.Y {
			next = State{
				X: now.Y + now.X - y,
				Y: y,
			}
			queue = append(queue, next)
		} else {
			next = State{
				X: 0,
				Y: now.Y + now.X,
			}
			queue = append(queue, next)
		}

		if now.Y > x-now.X {
			next = State{
				X: x,
				Y: now.Y + now.X - x,
			}
			queue = append(queue, next)
		} else {
			next = State{
				X: now.X + now.Y,
				Y: 0,
			}
			queue = append(queue, next)
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
