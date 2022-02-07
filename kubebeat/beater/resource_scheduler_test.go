package beater

// import (
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/suite"
// )

// type SchedulerTestSuite struct {
// 	suite.Suite
// 	scheduler ResourceScheduler
// 	handler   *TestResourceHandler
// }

// func TestSchedulerTestSuite(t *testing.T) {
// 	suite.Run(t, &SchedulerTestSuite{
// 		scheduler: new(SynchronousScheduler),
// 		handler:   new(TestResourceHandler),
// 	})
// }

// func (s *SchedulerTestSuite) TestScheduler() {
// 	tests := []struct {
// 		resources map[string][]interface{}
// 		expected  []interface{}
// 	}{
// 		{
// 			map[string][]interface{}{
// 				"provider":  {1, 2, 3},
// 				"provider2": {4, 5, 6},
// 			},
// 			[]interface{}{
// 				1, 2, 3, 4, 5, 6,
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		s.scheduler.ScheduleResources(test.resources, s.handler.func1)
// 		s.Equal(test.expected, s.handler.called)
// 	}

// }

// type TestResourceHandler struct {
// 	called []interface{}
// }

// func (h *TestResourceHandler) func1(r interface{}) {
// 	h.called = append(h.called, r)
// 	time.Sleep(100 * time.Millisecond)
// }
