package main

type pair struct {
	key int
	val string
}

type arrayHashMap struct {
	buckets []*pair
}

func newArrayHashMap() *arrayHashMap {
	return &arrayHashMap{
		buckets: make([]*pair, 100),
	}
}

func (a *arrayHashMap) put(key int, val string) {
	pair := &pair{key: key, val: val}
	index := a.hashFunc(key)
	a.buckets[index] = pair
}

func (a *arrayHashMap) remove(key int) {
	index := a.hashFunc(key)
	a.buckets[index] = nil
}

func (a *arrayHashMap) pairSet() []*pair {
	var pairs []*pair
	for _, pair := range a.buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func (a *arrayHashMap) keySet() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}

	}
	return keys
}
func (a *arrayHashMap) valueSet() []string {
	var values []string
	for _, pair := range a.buckets {
		if pair != nil {
			values = append(values, pair.val)
		}

	}
	return values
}
func (a *arrayHashMap) hashFunc(key int) int {
	return key % 100
}

func (a *arrayHashMap) get(key int) string {
	index := a.hashFunc(key)
	pair := a.buckets[index]
	if pair == nil {
		return ""
	}
	return pair.val
}

func (a *arrayHashMap) print() {
	for _, pair := range a.pairSet() {
		if pair != nil {
			println(pair.key, pair.val)
		}

	}
}

func main() {

}
