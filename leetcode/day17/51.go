package main

import "strings"

// 提交一份 Golang的写法
func solveNQueens(n int) [][]string {
	return NBoard(n)
}

func Board(n int, str string) [][]string {
	res := make([][]string, n)
	for i, _ := range res {
		// 第二维数组数据填
		v2 := make([]string, n)
		for j, _ := range v2 {
			v2[j] = str
		}
		res[i] = v2
	}
	return res
}

func NBoard(n int) [][]string {
	res := make([][]string, 0, n)
	// 结果集初始化
	for i, _ := range res {
		res[i] = make([]string, 0, n)
	}
	var dfs func(board [][]string, row int)
	dfs = func(board [][]string, row int) {
		// 结束条件,row的长度等于board长度
		if row == len(board) {
			tmp := make([]string, 0, n)
			for i := 0; i < n; i++ {
				tmp = append(tmp, strings.Join(board[i], ""))
			}
			res = append(res, tmp)
			return
		}

		// 获取当前棋盘的列数量
		n = len(board[row])
		// 对该列进行遍历
		for col := 0; col < n; col++ {
			// 排除不合法的选择
			if !IsValid(board, row, col) {
				continue
			}
			// 做选择
			board[row][col] = "Q"
			// 进行下一项决策
			dfs(board, row+1)
			// 撤销选择
			board[row][col] = "."
		}
	}
	board := Board(n, ".")
	dfs(board, 0)
	return res
}

// IsValid 判断当前位置是否可以放置皇后
func IsValid(board [][]string, row, col int) bool {
	// 获取棋盘的宽度
	n := len(board)
	// 检查该列是否有皇后冲突
	for i := 0; i < n; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	// 检查左上角是否有冲突，遍历条件行减列减
	i := row - 1
	j := col - 1
	for i >= 0 && j >= 0 {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j--
	}
	// 检查右上角是否有冲突，遍历条件行减列加
	i1 := row - 1
	j1 := col + 1
	for i1 >= 0 && j1 < n {
		if board[i1][j1] == "Q" {
			return false
		}
		i1--
		j1++
	}
	return true
}
