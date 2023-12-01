package repo

import "testing"

func TestInMemoryRepo(t *testing.T) {

	r := &InMemoryRepo{}

	sampleData := []*User{
		{1, "John", "john@email.com", "1234"},
		{2, "Jane", "jane@email.com", "5678"},
	}

	for _, user := range sampleData {
		r.Create(user)
	}

	// Get all
	users, err := r.GetAll()
	if err != nil {
		t.Errorf("Error fetching data: %v", err)
	}
	testCount(t, len(users), 2)

	// Get by id
	user, err := r.GetByID(1)
	if err != nil {
		t.Errorf("Error fetching data: %v", err)
	}
	testValue(t, user.Name, "John")

	// Create
	newUser := &User{3, "Jim", "jim@email.com", "abcd"}
	err = r.Create(newUser)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}
	testCount(t, len(r.data), 3)

	// Update
	updatedUser := &User{3, "James", "james@email.com", "abcd"}
	err = r.Update(updatedUser)
	if err != nil {
		t.Errorf("Error updating user: %v", err)
	}
	testValue(t, r.data[2].Name, "James")

	// Delete
	err = r.Delete(2)
	if err != nil {
		t.Errorf("Error deleting user: %v", err)
	}
	testCount(t, len(r.data), 2)
}

func testCount(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func testValue(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
