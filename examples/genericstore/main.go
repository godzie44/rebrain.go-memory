package main

import (
	"fmt"
	"runtime"
	"time"
)

func ManualBenchmark(f func(), benchName string) {
	iterCnt := 10_000

	startMs := &runtime.MemStats{}
	runtime.ReadMemStats(startMs)
	startTime := time.Now().UnixNano()

	for i := 0; i < iterCnt; i++ {
		f()
	}

	endTime := time.Now().UnixNano()
	endMs := &runtime.MemStats{}
	runtime.ReadMemStats(endMs)

	fmt.Printf(
		"%s       %d ns/op    %d b/op    %d allocs/op \n",
		benchName,
		(endTime-startTime)/int64(iterCnt),
		(endMs.HeapAlloc-startMs.HeapAlloc)/uint64(iterCnt),
		(endMs.Mallocs-startMs.Mallocs)/uint64(iterCnt),
	)
}

func main() {
	var nums = []int{3, 2, 20, 5}

	ManualBenchmark(func() {
		intStorage := &Storage{}
		for _, v := range nums {
			intStorage.Add(v)
		}
	}, "DefaultStorageBenchmark")

	ManualBenchmark(func() {
		intStorage := &IntStorage{}
		for _, v := range nums {
			intStorage.Add(v)
		}
	}, "IntStorageBenchmark")

	ManualBenchmark(func() {
		intStorage := &genericStorage[int]{}
		for _, v := range nums {
			intStorage.Add(v)
		}
	}, "GenericStorageBenchmark")
}

type Storage struct {
	data []interface{}
}

func (s *Storage) Add(el int) {
	s.data = append(s.data, el)
}

func (s *Storage) Find(el interface{}) (bool, interface{}) {
	for i := range s.data {
		if s.data[i] == el {
			return true, s.data[i]
		}
	}
	return false, nil
}

type IntStorage struct {
	data []int
}

func (s *IntStorage) Add(el int) {
	s.data = append(s.data, el)
}

func (s *IntStorage) Find(el int) (bool, int) {
	for i := range s.data {
		if s.data[i] == el {
			return true, s.data[i]
		}
	}
	return false, 0
}

type genericStorage[T comparable] struct {
	data []T
}

func (s *genericStorage[T]) Add(el T) {
	s.data = append(s.data, el)
}

func (s *genericStorage[T]) Find(el T) (bool, T) {
	for i := range s.data {
		if s.data[i] == el {
			return true, s.data[i]
		}
	}
	var zero T
	return false, zero
}