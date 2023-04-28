package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	ID   int
	Name string
	Type string
}

type Notification struct {
	UserID  int
	Message string
	Type    string
	SentAt  time.Time
}

type NotificationService struct {
	// Notification channels for each user type
	SMS     chan Notification
	Email   chan Notification
	Push    chan Notification
	Users   []User
	Workers int
}

func main() {
	users := []User{
		{ID: 1, Name: "Delivery Man 1", Type: "Delivery Man"},
		{ID: 2, Name: "Swiggy Employee 1", Type: "Swiggy Employee"},
		{ID: 3, Name: "End User 1", Type: "End User"},
	}

	message := "Your order is on the way!"

	notificationService := NotificationService{
		SMS:     make(chan Notification),
		Email:   make(chan Notification),
		Push:    make(chan Notification),
		Users:   users,
		Workers: 3,
	}

	go notificationService.SendNotifications(message)

	notificationService.Start()
}

func (s *NotificationService) SendNotification(notification Notification) {
	switch notification.Type {
	case "Delivery Man":
		s.SMS <- notification

	case "Swiggy Employee":
		s.Email <- notification

	case "End User":
		s.Push <- notification
	}
}

func (s *NotificationService) SendNotifications(message string) {
	wg := sync.WaitGroup{}
	wg.Add(len(s.Users))

	for _, user := range s.Users {
		go func(user User) {
			defer wg.Done()
			notification := Notification{
				UserID:  user.ID,
				Message: message,
				Type:    user.Type,
				SentAt:  time.Now(),
			}
			s.SendNotification(notification)
		}(user)
	}

	wg.Wait()

	close(s.SMS)
	close(s.Email)
	close(s.Push)
}

func (s *NotificationService) Start() {
	wg := sync.WaitGroup{}
	wg.Add(s.Workers)

	worker := func(c chan Notification) {
		defer wg.Done()
		for notification := range c {
			fmt.Printf("Sending notification to User %d - %s: %s\n", notification.UserID, notification.Type,
				notification.Message)
			// Implement notification logic here based on the notification type
			time.Sleep(time.Second)
		}
	}

	for i := 0; i < s.Workers; i++ {
		go worker(s.SMS)
		go worker(s.Email)
		go worker(s.Push)
	}

	wg.Wait()
}
