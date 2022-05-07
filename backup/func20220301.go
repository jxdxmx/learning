package backup

func f20220301() {
	//var arr = new([]int)
	//(*arr)[0] = 5
	var v = new(map[int]int)
	(*v)[0] = 1

	//var arr []int
	//arr = append(arr, 1, 2, 3)
	//arr = append(arr, 4) // 此处将扩容，cap从3->6
	//var arr2 = arr
	//arr2[0] = 10           // 10 2 3 4
	//arr2 = append(arr2, 5) // 10 20 3 4 5
	//arr2[1] = 20
	//var arr3 = arr2[:5]
	//arr3 = append(arr3, 11, 12, 13, 14, 15, 16, 17, 18)
	//arr3[0] = 999
	//fmt.Println(arr, arr2, arr3) // 10 20 3 4			// 10 20 3 4 5			// 999 20 3 4 5 11, 12, 13, 14, 15, 16, 17, 18
}

//func f20220301() { //array := [4]int{10, 20, 30, 40}
//	array := make([]int, 4, 10)
//	array[0] = 10
//	array[1] = 20
//	array[2] = 30
//	array[3] = 40
//	slice := array[0:2]                   // 此处slice、array使用的是同一底层数组。
//	newSlice := append(slice, 50, 44, 45) // append，是否使用原数组还是新生成数组，看是否超过cap！ 10 20 50 44 45
//	//newSlice := append(slice, 50, )
//	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice)) // 10 20
//	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
//	newSlice[1] += 10 //  10 30 50 44 45
//	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
//	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
//	fmt.Printf("After array = %v\n", array) // 10 30 50 44
//}
