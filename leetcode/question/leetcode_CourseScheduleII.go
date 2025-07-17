// 210. Course Schedule II
// 👉 適用於 有向無環圖（DAG）
// 👉 目的是：找出所有節點（任務、課程、建置步驟）的一種合法執行順序，使得所有的先決條件都被滿足。
// •	每門課（0 ~ numCourses-1）是一個節點
// •	prerequisites[i] = [a, b] 表示要修 a，必須先修 b ⇒ 邊 b → a
// •	要完成所有課程，就要找出一種不違反先修順序的排列路徑
// Example 2:

// Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
// Output: [0,2,1,3]
// Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
// So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3]

func findOrder(numCourses int, prerequisites [][]int) []int {
    // 建立入度與圖
    inDegree := make([]int, numCourses)
    graph := make([][]int, numCourses)

    for _, pre := range prerequisites {
        course, prereq := pre[0], pre[1]
        graph[prereq] = append(graph[prereq], course)
        inDegree[course]++
    }

    // 將入度為 0 的課程加入 queue
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    var result []int
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        result = append(result, cur)

        for _, next := range graph[cur] {
            inDegree[next]--
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }

    // 有些課修不到 → 圖中有環
    if len(result) != numCourses {
        return []int{}
    }

    return result
}