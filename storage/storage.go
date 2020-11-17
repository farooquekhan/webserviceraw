package storage

// TodoItem represent a single todo item
type TodoItem struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

// TodoList represents the list of todo items
type TodoList []TodoItem

// Our storage
var store TodoList
var currentID = 1

// GetAll returns the current TodoList
func GetAll() TodoList {
	return store
}

// GetItem returns a single item for given id
func GetItem(id int) (bool, TodoItem) {

	index := -1

	for i, items := range store {
		if items.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		return true, store[index]
	}

	return false, TodoItem{}
}

// AddUpdate function adds a new item to the list
func AddUpdate(newItem TodoItem) int {

	index := -1

	for i, items := range store {
		if items.ID == newItem.ID {
			index = i
			break
		}
	}

	if index != -1 {
		// Update
		store[index] = newItem
	} else {
		newItem.ID = currentID
		currentID++
		store = append(store, newItem)
	}

	return newItem.ID
}

// Remove removes the item with given id from the list
func Remove(id int) bool {
	index := -1

	for i, item := range store {
		if item.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		store = append(store[:index], store[index+1:]...)
	}

	// Returns true if item was found & removed
	return index != -1
}
