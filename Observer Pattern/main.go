package main

import "fmt"

// User represents a user of the Swiggy app
type User struct {
	ID   int
	Name string
	Type string // user type: "delivery-man", "swiggy-employee", or "end-user"
}

// Observer is an interface for users who can receive notifications
type Observer interface {
	Notify(string)
}

// DeliveryMan is an observer that represents a delivery man
type DeliveryMan struct {
	User
}

// Notify sends a notification to the delivery man
func (d *DeliveryMan) Notify(message string) {
	fmt.Printf("Delivery man %s received notification: %s\n", d.Name, message)
}

// SwiggyEmployee is an observer that represents a Swiggy employee
type SwiggyEmployee struct {
	User
}

// Notify sends a notification to the Swiggy employee
func (s *SwiggyEmployee) Notify(message string) {
	fmt.Printf("Swiggy employee %s received notification: %s\n", s.Name, message)
}

// EndUser is an observer that represents an end user of the Swiggy app
type EndUser struct {
	User
}

// Notify sends a notification to the end user
func (e *EndUser) Notify(message string) {
	fmt.Printf("End user %s received notification: %s\n", e.Name, message)
}

// Subject is a struct that manages a list of observers
type Subject struct {
	observers []Observer
}

// RegisterObserver adds an observer to the list of observers
func (s *Subject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

// RemoveObserver removes an observer from the list of observers
func (s *Subject) RemoveObserver(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

// NotifyObservers sends a notification to all observers
func (s *Subject) NotifyObservers(message string) {
	for _, o := range s.observers {
		o.Notify(message)
	}
}

// NotificationService is a struct that represents the notification service for Swiggy
type NotificationService struct {
	Subject
}

// NewNotificationService creates a new NotificationService object
func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

// RegisterUsers registers a list of users with the notification service
func (s *NotificationService) RegisterUsers(users []User) {
	for _, user := range users {
		switch user.Type {
		case "delivery-man":
			s.RegisterObserver(&DeliveryMan{User: user})
		case "swiggy-employee":
			s.RegisterObserver(&SwiggyEmployee{User: user})
		case "end-user":
			s.RegisterObserver(&EndUser{User: user})
		}
	}
}

// SendNotification sends a notification to all registered users
func (s *NotificationService) SendNotification(message string) {
	s.NotifyObservers(message)
}

func main() {
	service := NewNotificationService()

	// Register delivery men
	service.RegisterUsers([]User{
		{ID: 1, Name: "John", Type: "delivery-man"},
		{ID: 2, Name: "Jane", Type: "delivery-man"},
	})

	// Register Swiggy employees
	service.RegisterUsers([]User{
		{ID: 3, Name: "Bob", Type: "swiggy-employee"},
	})

	// Register end users
	service.RegisterUsers([]User{
		{ID: 4, Name: "Alice", Type: "end-user"},
		{ID: 5, Name: "Charlie", Type: "end-user"},
	})

	// Send notification to all users
	service.SendNotification("New offers available on Swiggy!")
}
