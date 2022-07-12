package interview

// 2022-07-12

// https://leetcode.cn/problems/generate-parentheses/
// 22. 括号生成

//数字 n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//示例 1：
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//示例 2：
//输入：n = 1
//输出：["()"]

func generateParenthesis918(n int) []string {
	if n == 0 {
		return []string{""}
	} else if n == 1 {
		return []string{"()"}
	}
	var ans []string
	// (A)B
	for k := 0; k <= n-1; k++ { // 此处要特别小心！！！0~n-1，包含n-1！！！
		leftAns, rigthAns := generateParenthesis918(k), generateParenthesis918(n-k-1)
		// fmt.Println(k,leftAns,"  ",n-k-1,rigthAns)
		for _, l := range leftAns {
			for _, r := range rigthAns {
				ans = append(ans, "("+l+")"+r)
			}
		}
	}
	return ans
}

// dfs的做法
func generateParenthesis9182(n int) []string {
	var ans []string
	var bs []byte
	var dfs func(index int)
	dfs = func(index int) {
		if index > 2*n {
			// fmt.Println(string(bs),valid918(bs))
			if valid918(bs) {
				ans = append(ans, string(bs))
			}
			return
		}
		for _, b := range []byte{'(', ')'} {
			bs = append(bs, b)
			dfs(index + 1)
			bs = bs[:len(bs)-1]
		}
	}
	dfs(1)
	return ans
}

func valid918(bs []byte) bool {
	var q []byte
	for i := 0; i < len(bs); i++ {
		if bs[i] == '(' {
			q = append(q, bs[i])
		} else {
			if len(q) == 0 || q[len(q)-1] == ')' {
				return false
			}
			q = q[:len(q)-1]
		}
	}
	return len(q) == 0
}
