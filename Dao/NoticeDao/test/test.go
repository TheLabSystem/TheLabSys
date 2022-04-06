package test

import (
	"TheLabSystem/Dao/NoticeDao"
	"TheLabSystem/Types/ServiceType/Notice"
	"fmt"
	"sync"
)

func testInsertNotice(wt *sync.WaitGroup) {
	var notice = Notice.Notice{
		NoticeInfo: "123",
		IssuerID:   1,
	}
	err := NoticeDao.InsertNotice(notice)
	if err != nil {
		fmt.Println(err)
	}
	wt.Done()
}
func testDeleteNotice() {
	var notice = Notice.Notice{
		NoticeID: 1,
	}
	err := NoticeDao.DeleteNotice(notice)
	if err != nil {
		fmt.Println(err)
	}
}
func testUpdateNotice() {
	notice := Notice.Notice{
		NoticeID:   4,
		NoticeInfo: "12345",
		IssuerID:   1,
	}
	fmt.Println(NoticeDao.UpdateNotice(notice))
}
func testFindNotice() {
	fmt.Println(NoticeDao.FindNoticeByID(2))
	fmt.Println(NoticeDao.FindNoticeByIssuerID(1))
}

func main() {
	//waiter := &sync.WaitGroup{}
	//waiter.Add(90)
	//for i := 1; i <= 90; i++ {
	//	go testInsertNotice(waiter)
	//}
	//waiter.Wait()
	//testDeleteNotice()
	//testFindNotice()
	//testUpdateNotice()
	notices, err := NoticeDao.FindNoticeByOffset(1, 100)
	if err != nil {
		fmt.Println(err)
	} else {
		for key := range notices {
			fmt.Println(notices[key])
		}
	}
}
