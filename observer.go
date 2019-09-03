// Package main
package main

import (
	"fmt"
	"sync"
)

// Param ...
type Param string

// Observer ...
type Observer interface {
	Notify(p Param)
}

// Subject ...
type Subject struct {
	mu        sync.RWMutex
	observers []Observer
}

// AddObserver ...
func (nb *Subject) AddObserver(o Observer) {
	nb.mu.Lock()
	defer nb.mu.Unlock()
	nb.observers = append(nb.observers, o)
}

// DeleteObserver ...
func (nb *Subject) DeleteObserver(o Observer) {
	nb.mu.Lock()
	defer nb.mu.Unlock()
	obs := nb.observers[0:0]
	for _, v := range nb.observers {
		if v != o {
			obs = append(obs, v)
		}
	}
	nb.observers = obs
}

// Notify ...
func (nb *Subject) Notify(p Param) {
	nb.mu.RLock()
	defer nb.mu.RUnlock()
	for _, v := range nb.observers {
		v.Notify(p)
	}
}

// Something ...
type Something struct {
	Subject
}

// Observer1 ...
type Observer1 struct{}

// Notify ...
func (o1 *Observer1) Notify(p Param) {
	fmt.Println("obs1:", p)
}

// Observer2 ...
type Observer2 struct{}

// Notify ...
func (o1 *Observer2) Notify(p Param) {
	fmt.Println("obs2:", p)
}

func main() {
	s := &Something{}
	ob1 := &Observer1{}
	ob2 := &Observer2{}
	s.AddObserver(ob1)
	s.AddObserver(ob2)
	s.Notify("hello!")
	s.DeleteObserver(ob1)
	s.Notify("World!")
}
