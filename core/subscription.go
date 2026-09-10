package core

import (
	"fmt"
)

type Subscription struct {
	id       int64
	Name     string
	Expense  float64
	Currency Currency
}

var Subscriptions map[int64]Subscription = make(map[int64]Subscription, 0)

func newSubscription(id int64, name string, expense float64, currency Currency) Subscription {
	return Subscription{id: id, Name: name, Expense: expense, Currency: currency}
}

func SubscriptiontoString(entry Subscription) string {
	return fmt.Sprintf("Subscription(id=%d, name=%s, expense=%f, currency=%s)", entry.id, entry.Name, entry.Expense, entry.Currency.Name)
}

func AddSubscription(id int64, name string, expense float64, currency Currency) Subscription {
	ret := newSubscription(id, name, expense, currency)
	Subscriptions[id] = ret
	return ret
}

func GetSubscription(id int64) (Subscription, bool) {
	val, exists := Subscriptions[id]
	return val, exists
}

func RemoveSubscription(id int64) {
	delete(Subscriptions, id)
}

func UpdateSubscriptionName(id int64, newName string) {
	if val, exists := Subscriptions[id]; exists {
		val.Name = newName
		Subscriptions[id] = val
	}
}

func UpdateSubscriptionExpense(id int64, newExpense float64) {
	if val, exists := Subscriptions[id]; exists {
		val.Expense = newExpense
		Subscriptions[id] = val
	}
}

func UpdateSubscriptionCurrency(id int64, newCurrency Currency) {
	if val, exists := Subscriptions[id]; exists {
		val.Currency = newCurrency
		Subscriptions[id] = val
	}
}
