/*
 * File Created: Thursday, 4th June 2020 10:37:11 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2020 Author
 */

package usecase

import (
	"net"
	"runtime"
	"sync"

	"github.com/abmid/icanvas-analytics/pkg/canvas/course/repository"
	"github.com/abmid/icanvas-analytics/pkg/canvas/entity"
)

type courseUseCase struct {
	CourseRepo repository.CourseRepository
}

func NewCourseUseCase(courseRepo repository.CourseRepository) *courseUseCase {
	return &courseUseCase{
		CourseRepo: courseRepo,
	}
}

func (CUC *courseUseCase) Courses(accountId, page uint32) (res []entity.Course, err error) {
	courses, err := CUC.CourseRepo.Courses(accountId, page)
	if err != nil {
		return nil, err
	}
	res = courses
	return res, nil
}

/**
* This method for worker get All Course
* @param pool int
* @param page *int
* @param accountID int
* @param chCourse chan []entity.course
* @param *courseUseCase
* @param *sync.WaitGroup
 */
func workerCourses(pool uint32, page *uint32, accountId uint32, chCourse chan<- []entity.Course, CUC *courseUseCase) {
	workerChCourse := make(chan []entity.Course, 50)
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var i uint32
	for i = 0; i < pool; i++ {
		wg.Add(1)
		go func(page *uint32) {
			mtx.Lock()
			course, err := CUC.CourseRepo.Courses(accountId, *page)
			*page++
			mtx.Unlock()
			if err != nil {
				panic(err)
			}
			workerChCourse <- course
			defer wg.Done()
		}(page)
	}
	go func() {
		wg.Wait()
		close(workerChCourse)
	}()
	workerCourse := []entity.Course{}
	for course := range workerChCourse {
		if len(course) != 0 {
			workerCourse = append(workerCourse, course...)
		}
	}
	chCourse <- workerCourse
}

func (CUC *courseUseCase) AllCourse(accountId, pool uint32) (res []entity.Course, err error) {
	/*
		TODO : Create Worker
		? 1. Buat Worker yang memiliki pool dinamis (ex. 100 Concurrent)
		? 2. Loop secara infinity, pada statement loop menjalankan worker dan mengembalikan nilai ke channel.
		? 3. Pada saat itu juga hasil dari channel worker di simpan ke variable.
		? 4. Jika hasil len 50, maka infinite loop jalan lagi dan memanggil worker
	*/
	runtime.GOMAXPROCS(runtime.NumCPU())
	var countPage uint32 = 1
	for {
		ch := make(chan []entity.Course)
		go workerCourses(pool, &countPage, accountId, ch, CUC)
		if chCourse := <-ch; len(chCourse) == 0 {
			break
		} else {
			res = append(res, chCourse...)
		}
	}

	return res, nil
}

/*
This method for get all course
with distribution result per 50 course per loop
*/
func (CUC *courseUseCase) GoAllCourse(accountID uint32, ch chan<- []entity.Course, wg *sync.WaitGroup) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var countPage uint32 = 1
	countTry := 0
outerLoop:
	for {
		course, err := CUC.Courses(accountID, countPage)
		if err != nil {
			if countTry > 2 {
				panic(err)
			}
			if err, ok := err.(net.Error); ok && err.Timeout() {
				countTry++
				break outerLoop
			} else {
				panic(err)
			}
		}
		if len(course) < 1 {
			break
		}
		wg.Add(1)
		ch <- course
		countPage++
	}
	close(ch)
}

/**
* This method for get user in course
* @param courseID int
* @param enrollmentRole string
 */
func (CUC *courseUseCase) ListUserInCourse(courseID uint32, enrollmentRole string) (res []entity.User, err error) {
	res, err = CUC.CourseRepo.ListUserInCourse(courseID, enrollmentRole)
	if err != nil {
		return nil, err
	}
	return res, nil
}
