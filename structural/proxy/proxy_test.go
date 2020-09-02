package proxy

import (
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	someDatabase := UserList{}

	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCache:    UserList{},
		StackCapacity: 2,
	}

	knownIDs := [3]int32{someDatabase[3].ID, someDatabase[4].ID, someDatabase[5].ID}

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err.Error())
		}
		if user.ID != knownIDs[0] {
			t.Errorf("Expected ID: %d\nActual ID: %d\n", knownIDs, user.ID)
		}
		if len(proxy.StackCache) != 1 {
			t.Errorf("Expected lenght: 1\nActual lenght: %d", len(proxy.StackCache))
		}
		if proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - One user on cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err.Error())
		}
		if user.ID != knownIDs[0] {
			t.Errorf("Expected ID: %d\nActual ID: %d\n", knownIDs, user.ID)
		}
		if len(proxy.StackCache) != 1 {
			t.Errorf("Expected lenght: 1\nActual lenght: %d", len(proxy.StackCache))
		}
		if !proxy.DidLastSearchUsedCache {
			t.Error("Expected user returned from cache")
		}
	})

	t.Run("FindUser - Full cache & user not on cache", func(t *testing.T) {
		user1, err1 := proxy.FindUser(knownIDs[0])
		if err1 != nil {
			t.Fatal(err1.Error())
		}
		user2, err2 := proxy.FindUser(knownIDs[1])
		if err2 != nil {
			t.Fatal(err2.Error())
		}
		user3, err3 := proxy.FindUser(knownIDs[2])
		if err3 != nil {
			t.Fatal(err3.Error())
		}

		for _, v := range proxy.StackCache {
			if v.ID == user1.ID {
				t.Errorf("UserID: %d shouldn't be on cache", user1.ID)
			}
		}

		for _, v := range proxy.StackCache {
			if v.ID != user2.ID && v.ID != user3.ID {
				t.Errorf("Expected userIDs: %d & %d on cache", user2.ID, user3.ID)
			}
		}

		if len(proxy.StackCache) != 2 {
			t.Errorf("Expected lenght: 2\nActual lenght: %d", len(proxy.StackCache))
		}
	})

	t.Run("FindUser - NonExistent user", func(t *testing.T) {
		_, err := proxy.FindUser(-1)
		if err == nil {
			t.Fatal("Expected error user not found")
		}
	})
}
