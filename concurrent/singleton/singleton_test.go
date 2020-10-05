package singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5000

	fmt.Printf("Before loop1, current count is %d\n", singleton.GetCount())
	fmt.Printf("Before loop2, current count is %d\n", singleton2.GetCount())

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Middle loop1, current count is %d\n", singleton.GetCount())
	fmt.Printf("Middle loop2, current count is %d\n", singleton2.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(time.Millisecond * 10)
	}

	fmt.Printf("After loop1, current count is %d\n", singleton.GetCount())
	fmt.Printf("After loop2, current count is %d\n", singleton2.GetCount())

	singleton.Stop()

}

func TestSingletonLock(t *testing.T) {
	singletonLock := GetInstanceLock()
	singletonLock2 := GetInstanceLock()

	n := 5000

	fmt.Printf("Before loop1, current count is %d\n", singletonLock.GetCountLock())
	fmt.Printf("Before loop2, current count is %d\n", singletonLock2.GetCountLock())

	for i := 0; i < n; i++ {
		go singletonLock.AddOneLock()
		go singletonLock2.AddOneLock()
	}

	fmt.Printf("Middle loop1, current count is %d\n", singletonLock.GetCountLock())
	fmt.Printf("Middle loop2, current count is %d\n", singletonLock2.GetCountLock())

	var val int
	for val != n*2 {
		val = singletonLock.GetCountLock()
		time.Sleep(time.Millisecond * 10)
	}

	fmt.Printf("Before loop1, current count is %d\n", singletonLock.GetCountLock())
	fmt.Printf("Before loop2, current count is %d\n", singletonLock2.GetCountLock())

}
