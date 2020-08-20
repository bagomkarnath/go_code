package err

import "fmt"

var (
	// ErrMapStoreNotInitialized MapStore not initialized
	ErrMapStoreNotInitialized = fmt.Errorf("MapStore not initialized")

	// ErrCustomerDoesNotExists Customer doesn't exists
	ErrCustomerDoesNotExists = fmt.Errorf("Customer doesn't exists")

	// ErrNoDataInMapStore No data in Map Store
	ErrNoDataInMapStore = fmt.Errorf("No data in Map Store")
)
