package main

import "testing"

//func TestMain(m *testing.M) {
//	main()
//}

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
