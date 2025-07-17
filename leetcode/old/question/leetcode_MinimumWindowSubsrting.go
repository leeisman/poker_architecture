// 76. Minimum Window Substring
// 經典的滑動視窗（Sliding Window）題目
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	// 最終結果
	res := ""

	// 雙指標與符合條件數量
	left, right := 0, 0
	matchCount := 0

	// 記錄 t 中每個字元的需求量
	need := make(map[byte]int)

	// 記錄目前視窗內的字元統計
	window := make(map[byte]int)

	// 初始化需求表
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// 右指標開始滑動
	for right < len(s) {
		ch := s[right]

		// 如果是目標字元，加入視窗內並更新匹配數
		if count, exists := need[ch]; exists {
			window[ch]++
			if window[ch] == count {
				matchCount++
			}
		}

		// 當所有字元匹配時，嘗試收縮視窗
		for matchCount == len(need) {
			// 更新結果
			if res == "" || len(res) > right-left+1 {
				res = s[left : right+1]
			}

			// 移除左邊界字元
			leftCh := s[left]
			if count, exists := need[leftCh]; exists {
				if window[leftCh] == count {
					matchCount--
				}
				window[leftCh]--
			}
			left++
		}

		right++
	}

	return res
}