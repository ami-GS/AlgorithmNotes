package main

import (
	"fmt"

	"github.com/ami-GS/AlgorithmNotes/map/cuckoo"
)

func main() {
	// m := closedhash.NewMap(3000)
	// m.Insert(9, 123)
	// m.Insert(19, 1234)
	// m.Insert(29, 12345)
	// m.Insert(39, 123456)
	// fmt.Println(m.Get(29))
	// m.Remove(19)
	// m.Remove(9)
	// m.Insert(19, 4321)
	// fmt.Println(m.Get(29))

	m := cuckoo.NewMap(5)
	m.Insert(1, 1231)
	m.Insert(6, 12311)
	m.Insert(2, 1232)
	m.Insert(7, 12322)
	m.Insert(3, 1233)
	m.Insert(8, 12333)
	m.Insert(4, 1234)
	m.Insert(9, 12344)
	m.Insert(5, 1235)
	m.Insert(10, 12355)
	m.Insert(11, 12355)
	fmt.Println(m.Get(3))

	// m := openhash.NewMap(10)
	// for i := 0; i < 100; i++ {
	// 	m.Insert(i, i*10)
	// }
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(m.Get(i))
	// }
	// for i := 99; i >= 0; i-- {
	// 	fmt.Println("remove", i)
	// 	m.Remove(i)
	// }
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(m.Get(i))
	// }
}
