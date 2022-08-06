package array

import "math/rand"

type RandomizedSet struct {
	set      []int
	location map[int]int
}

//func Constructor() RandomizedSet {
//	return RandomizedSet{
//		set:      make([]int, 0),
//		location: make(map[int]int),
//	}
//}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exi := this.location[val]; exi {
		return false
	} else {
		this.location[val] = len(this.set)
		this.set = append(this.set, val)
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if loc, exi := this.location[val]; exi {
		this.set[loc] = this.set[len(this.set)-1]
		this.location[this.set[len(this.set)-1]] = loc
		this.set = this.set[:len(this.set)-1]
		delete(this.location, val)
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.set[rand.Intn(len(this.set))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
