/*
Package barista implements the Barista type and the actions a barista can
do.
*/
package barista

// Barista represents a staff member at the coffee cart
type Barista struct{}

// MakeOrder takes ingredients from stock, waits for the time-to-make of
// the ordered item, and then signals the customer that the order is ready.
// No error is returned because the cart has infinite storeroom stock.
func (b *Barista) MakeOrder() {}
