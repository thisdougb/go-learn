/*
Package business implements tasks and proceses of the coffee cart.
do.
*/
package business

// Business represents the business state, such as cash flow and stock levels.
type Business struct{}

// Open sets the state of the coffee cart and enables customers to join the
// ordering queue.
func (b *Business) Open() {}

// Close sets the state of the coffee cart and prevents customers joining the
// ordering queue.
func (b *Business) Close() {}
