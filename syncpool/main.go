package main

import "sync"

type Person struct {
	Name string
	Age  int
}

var personPool = sync.Pool{
	New: func() interface{} { return &Person{} },
}

//BenchmarkWithoutPool-12                1462            804840 ns/op          238159 B/op      19744 allocs/op
//BenchmarkWithPool-12                1462               834840 ns/op          0 B/op      0 allocs/op
func GetUser() {
	var p *Person
	for i := 0; i < 1000; i++ {
		p = personPool.Get().(*Person)
		p.Name = "fsdsdf"
		p.Age = 132434
		personPool.Put(p)
	}
}
