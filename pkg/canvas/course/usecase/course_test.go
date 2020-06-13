package usecase

import (
	"context"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/abmid/icanvas-analytics/pkg/canvas/course/repository"
	mock_course "github.com/abmid/icanvas-analytics/pkg/canvas/course/repository/mock"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func SetupTest(ctrl *gomock.Controller) *mock_course.MockCourseRepository {
	ListCoursePage1 := []entity.Course{} // 50
	ListCoursePage2 := []entity.Course{} // 50
	ListCoursePage3 := []entity.Course{} // 50
	ListCoursePage4 := []entity.Course{} // 50
	ListCoursePage5 := []entity.Course{}
	ListCoursePage6 := []entity.Course{}
	for i := 0; i < 200; i++ {
		if i < 50 {
			ListCoursePage1 = append(ListCoursePage1, entity.Course{
				ID: uint32(i) + 1,
			})
		} else if i < 100 {
			ListCoursePage2 = append(ListCoursePage2, entity.Course{
				ID: uint32(i) + 1,
			})
		} else if i < 150 {
			ListCoursePage3 = append(ListCoursePage3, entity.Course{
				ID: uint32(i) + 1,
			})
		} else {
			ListCoursePage4 = append(ListCoursePage4, entity.Course{
				ID: uint32(i) + 1,
			})
		}
	}
	mockRepoCourse := mock_course.NewMockCourseRepository(ctrl)
	mockRepoCourse.EXPECT().Courses(uint32(1), uint32(1)).Return(ListCoursePage1, nil).AnyTimes()
	mockRepoCourse.EXPECT().Courses(uint32(1), uint32(2)).Return(ListCoursePage2, nil).AnyTimes()
	mockRepoCourse.EXPECT().Courses(uint32(1), uint32(3)).Return(ListCoursePage3, nil).AnyTimes()
	mockRepoCourse.EXPECT().Courses(uint32(1), uint32(4)).Return(ListCoursePage4, nil).AnyTimes()
	mockRepoCourse.EXPECT().Courses(uint32(1), uint32(5)).Return(ListCoursePage5, nil).AnyTimes()
	mockRepoCourse.EXPECT().Courses(uint32(1), uint32(6)).Return(ListCoursePage6, nil).AnyTimes()

	return mockRepoCourse
}

func RealRepository() *repository.APIRepository {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   0,
			KeepAlive: 0,
		}).Dial,
		TLSHandshakeTimeout: 60 * time.Second,
	}

	client := http.Client{Transport: transport}
	repo := repository.NewRepositoryAPI(&client, "https://lms.umm.ac.id/", "2Q8LJIJs7gCo8XsftFOtq53UT3cUlBIHsTQS7WAi6Le0TTjT2sL7bNtkm5ERT7cb")
	return repo
}

func TestListUserInCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ListUser := []entity.User{
		{ID: 1, Name: "Test User Teacher"},
	}
	mockRepoCourse := mock_course.NewMockCourseRepository(ctrl)
	mockRepoCourse.EXPECT().ListUserInCourse(uint32(1), "TeacherEnrollment").Return(ListUser, nil)
	courseUC := NewCourseUseCase(mockRepoCourse)
	res, err := courseUC.ListUserInCourse(uint32(1), "TeacherEnrollment")
	t.Log(res)
	assert.NilError(t, err, "Have error")
	assert.Equal(t, len(ListUser), len(res))
}

func TestCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ListCourse := []entity.Course{
		{ID: 1},
	}
	mockRepoCourse := mock_course.NewMockCourseRepository(ctrl)
	mockRepoCourse.EXPECT().Courses(uint32(1), uint32(1)).Return(ListCourse, nil)
	courseUseCase := NewCourseUseCase(mockRepoCourse)
	result, err := courseUseCase.Courses(uint32(1), uint32(1))
	assert.NilError(t, err, "Err Result")
	assert.Equal(t, result[0].ID, uint32(1))
	assert.Equal(t, len(result), 1)
}

func TestWorkerCourse(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	setupTest := SetupTest(ctrl)
	courseUseCase := NewCourseUseCase(setupTest)
	var pool, countPage uint32 = 2, 1
	ch := make(chan []entity.Course)
	go workerCourses(pool, &countPage, 1, ch, courseUseCase)
	result := <-ch
	assert.Equal(t, 100, len(result))
}

func TestAllCourse(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	setupTest := SetupTest(ctrl)
	courseUseCase := NewCourseUseCase(setupTest)
	result, err := courseUseCase.AllCourse(uint32(1), uint32(2))
	t.Log(len(result))
	for _, res := range result {
		t.Log(res)
	}
	assert.NilError(t, err, "Error All Course")
	assert.Equal(t, len(result), 200, "Length Result not same")
}

func TestGoAllCourse(t *testing.T) {
	ctrl, _ := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()
	setupTest := SetupTest(ctrl)
	courseUseCase := NewCourseUseCase(setupTest)
	ch := make(chan []entity.Course)
	wg := new(sync.WaitGroup)
	// Init Worker Receive data from channel GoAllCourse
	go func(ch <-chan []entity.Course) {
		res := []entity.Course{}
		for result := range ch {
			res = append(res, result...)
			wg.Done()
		}
		assert.Equal(t, len(res), 200)
	}(ch)
	courseUseCase.GoAllCourse(uint32(1), ch, wg)
	wg.Wait()
}
