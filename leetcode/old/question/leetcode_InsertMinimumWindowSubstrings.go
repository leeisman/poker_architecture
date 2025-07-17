// 380. Insert Delete GetRandom O(1)
// 學習觀念slice 要快速刪除用swap搭配尾部[:n-1]
// rand用slice 0(1)

type RandomizedSet struct {
	nums []int
	idx  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: []int{},
		idx:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.idx[val]; exists {
		return false
	}
	this.nums = append(this.nums, val)
	this.idx[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, exists := this.idx[val]
	if !exists {
		return false
	}

	lastVal := this.nums[len(this.nums)-1]
	this.nums[i] = lastVal
	this.idx[lastVal] = i

	this.nums = this.nums[:len(this.nums)-1]
	delete(this.idx, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	r := rand.Intn(len(this.nums))
	return this.nums[r]
}