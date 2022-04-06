package test

import (
	"TheLabSystem/Dao/ReservationRecordDao"
	"TheLabSystem/Types/ReservationRecord"
	"fmt"
	"sync"
)

func testInsertReservationRecord(wt *sync.WaitGroup) {
	var rr = ReservationRecord.ReservationRecord{
		ReservationID: 1,
		OperatorID:    233,
		OperationType: "permit",
		OperatingDay:  "2022-4-6",
	}
	err := ReservationRecordDao.InsertReservationRecord(rr)
	if err != nil {
		fmt.Println(err)
	}
	wt.Done()
}
func testFindReservationRecord() {
	fmt.Println(ReservationRecordDao.FindReservationRecordByReservationID(1))
	fmt.Println(ReservationRecordDao.FindReservationRecordByOperatorID(233))
	fmt.Println(ReservationRecordDao.FindReservationRecordByOperatingDay("2022-4-6"))
}
func main() {
	//waiter := &sync.WaitGroup{}
	//waiter.Add(90)
	//for i := 1; i <= 90; i++ {
	//	go testInsertReservationRecord(waiter)
	//}
	//waiter.Wait()
	//testFindReservationRecord()
}
