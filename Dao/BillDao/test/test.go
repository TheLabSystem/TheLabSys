package main

import (
	"TheLabSystem/Dao/BillDao"
	"TheLabSystem/Types/ServiceType/Bill"
	"fmt"
)

func testInsertBill() {
	bill := Bill.Bill{
		PayerID:    1,
		Money:      100000,
		BillStatus: 2,
	}
	err := BillDao.InsertBill(bill)
	if err != nil {
		fmt.Println("failed to insert bill!")
	}
}
func testFindBill() {
	fmt.Println(BillDao.FindBillsByPayerID(123))
}
func testUpdateBill() {
	err := BillDao.UpdateBillStatus(123, 1)
	if err != nil {
		fmt.Println("failed to update bill!")
	}
}
func main() {
	testInsertBill()
	//testFindBill()
	//testUpdateBill()
}
