package test

import (
	"TheLabSystem/Dao/MentorRecordDao"
	"TheLabSystem/Types/ServiceType/MentorRecord"
	"fmt"
	"sync"
)

func testInsertMentorRecord(wt *sync.WaitGroup) {
	var mr = MentorRecord.MentorRecord{
		TeacherID: 123,
		StudentID: 321,
	}
	err := MentorRecordDao.InsertMentorRecord(mr)
	if err != nil {
		fmt.Println(err)
	}
	wt.Done()
}
func testDeleteMentorRecord() {
	var mr = MentorRecord.MentorRecord{
		TeacherID: 123,
		StudentID: 321,
	}
	err := MentorRecordDao.DeleteMentorRecord(mr)
	if err != nil {
		fmt.Println(err)
	}
}
func testFindMentorRecord() {
	fmt.Println(MentorRecordDao.FindMentorRecordByStudentID(321))
	fmt.Println(MentorRecordDao.FindMentorRecordByTeacherID(123))
}
func main() {
	//waiter := &sync.WaitGroup{}
	//waiter.Add(90)
	//for i := 1; i <= 90; i++ {
	//	go testInsertMentorRecord(waiter)
	//}
	//waiter.Wait()
	//testDeleteMentorRecord()
	//testFindMentorRecord()
}
