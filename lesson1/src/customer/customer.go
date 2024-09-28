/*
Package customer implements the Customer type and the actions a customer can
do.
*/
package customer

// Customer represents a customer visiting the coffee cart.
type Customer struct{}

// JoinQueue adds the customer to the ordering queue, to be asyncrhonously
// served by the barista. An error is returned if the queue is full, or the
// cart is closed.
func (c *Customer) JoinQueue() error { return nil }

// CreateOrder allows a customer to make an order with a barista. A customer
// can always make an order if they are in the queue, so this function does
// not return an error.
func (c *Customer) CreateOrder() {}
