package main

import "fmt"

// Set class
type Set struct {
	anyMap map[any]bool
}

// New creates the map of any and boolean
func (set *Set) New() {
	set.anyMap = make(map[any]bool)
}

// Contains checks wether or not the given element exists in map
func (set *Set) Contains(element any) (exists bool) {
	exists = set.anyMap[element]
	return
}

// Add adds the given element to a set
func (set *Set) Add(element any) {
	if !set.Contains(element) {
		set.anyMap[element] = true
	}
}

// Delete deletes the given element from map using delete method
func (set *Set) Delete(element any) {
	delete(set.anyMap, element)
}

// Intersect returns an intersect set that consists of the intersection of set and the given other set
func (set *Set) Intersect(anotherSet *Set) *Set {
	intersectSet := &Set{}
	intersectSet.New()

	for value := range set.anyMap {
		if anotherSet.Contains(value) {
			intersectSet.Add(value)
		}
	}

	return intersectSet
}

// Union returns a union set that consists of a union of set and the given other set
func (set *Set) Union(anotherSet *Set) *Set {
	unionSet := &Set{}
	unionSet.New()

	for value := range set.anyMap {
		unionSet.Add(value)
	}

	for value := range anotherSet.anyMap {
		unionSet.Add(value)
	}

	return unionSet
}

func main() {
	set := &Set{}
	set.New()
	set.Add(1)        // 1:true
	set.Add('A')      // 65:true
	set.Add(6.9)      // 6.9:true
	set.Add("String") // String:true
	fmt.Println(set)  // &{map[6.9:true 1:true 65:true String:true]}

	fmt.Println(set.Contains('A')) // true
	fmt.Println(set.Contains(3))   // false

	anotherSet := &Set{}
	anotherSet.New()
	anotherSet.Add(2)    // 2:true
	anotherSet.Add('A')  // 65:true
	anotherSet.Add(4.20) // 4.20:true

	fmt.Println(set.Intersect(anotherSet)) // &{map[65:true]}
	fmt.Println(set.Union(anotherSet))     // &{map[4.2:true 6.9:true 1:true 2:true 65:true String:true]}

}
