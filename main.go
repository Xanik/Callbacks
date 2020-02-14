package main

import "fmt"

func main() {
	po := new(PurchaseOrder)

	po.Value = 45.92

	ch := make(chan *PurchaseOrder)

	go SavePo(po, ch)

	newPo := <- ch

	fmt.Printf("PO: %v", newPo)

}

type PurchaseOrder struct {
	Number int
	Value  float64
}

func SavePo(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 1234

	callback <- po
}
