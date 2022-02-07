package beater

// import (
// 	"testing"
// )

// func BenchmarkScheduler(b *testing.B) {
// 	s := SynchronousScheduler{}
// 	h := TestResourceHandler{
// 		called: []interface{}{},
// 	}
// 	list := map[string][]interface{}{
// 		"provider": {},
// 	}
// 	for n := 0; n < b.N; n++ {
// 		list["provider"] = append(list["provider"], 1)
// 	}

// 	s.ScheduleResources(list, h.func1)
// }
