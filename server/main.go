package main

import (
	"context"
	"net/http"
	_ "net/http/pprof"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
)

// User model ----------------------------------------------------------------------------------------------------------

type GroupID int

const (
	GroupGuest = 1
	GroupUser  = 2
)

type User struct {
	ID     int
	Groups []GroupID
}

func NewUser() *User {
	return &User{}
}

// User repository -----------------------------------------------------------------------------------------------------

type UserRepository struct {
	lastID int
}

func (r *UserRepository) NextID() int {
	r.lastID++
	return r.lastID
}

//go:noinline
func (r *UserRepository) persists(user *User) {
	runtime.KeepAlive(user)
}

var userRepo = &UserRepository{}

// Make users function -------------------------------------------------------------------------------------------------

func MakeUsers(count int) {
	for i := 0; i < count; i++ {
		user := &User{ID: userRepo.NextID(), Groups: []GroupID{GroupGuest, GroupUser}}
		userRepo.persists(user)
	}
}

// Http server ---------------------------------------------------------------------------------------------------------

func main() {
	go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()

	go func() {
		err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userCount, _ := strconv.Atoi(r.URL.Query().Get("cnt"))
			MakeUsers(userCount)
			w.WriteHeader(http.StatusOK)
		}))
		if err != nil {
			panic(err)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	<-ctx.Done()
}
