package lesson7

import (
	"fmt"
)

func Entry() {
	fmt.Println("Arrays and slices")

	// Theory
	// Arrays vs slices
	// Arrays: numbered list of a single type, fixed length
	// Slices: Can be resized  (built on top of arrays) (slices are references)

	// slice (type, length, capacity - max size of the slice)
	myCourses := make([]string, 10, 20)
	myCourses2 := []string{"test1", "test2", "test3", "test4", "test5"}
	fmt.Printf("Length: %d. Capacity: %d\n", len(myCourses), cap(myCourses))
	fmt.Printf("Length: %d. Capacity: %d\n", len(myCourses2), cap(myCourses2))

	sliceOfSlice := myCourses2[0:2]
	fmt.Println(sliceOfSlice)

	mySlice := make([]int, 1, 4)
	for i := 1; i < 10; i++ {
		mySlice = append(mySlice, i) // mySlice = append(mySlice, mySlice...)
		fmt.Printf("\nCapacity: %d", cap(mySlice))
	}

}
