/*
Package barista implements the Barista type and the actions a barista can
do.

The functions contain pseudocode referenced in lesson 2.
*/
package barista

// Barista represents a staff member at the coffee cart.
type Barista struct {
	id int // each barista is identified by a staff id number
}

// MakeOrder takes ingredients from stock, waits for the time-to-make of
// the ordered item, and then signals the customer that the order is ready.
// No error is returned because the cart has infinite storeroom stock.
func MakeOrder() {
	/*
		order = pop(orderQueue)
		pause(order.TimeToMakeThis)
		signalCustomer(order)
	*/

}

// AcceptOrder accepts a valid order from a customer, and puts it in a queue
func AcceptOrder() {
	/*
		validateOrder(order)
		push(orderQueue, order)
	*/
}

// Restock replenishes the cart stock from the store room
func Restock() {
	/*
		decrementStoreStock(sizeOfCartStock)
		pause(timeToRestock)
	*/
}
