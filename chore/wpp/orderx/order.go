package orderx

import "strings"

const (
	StatusOrdered   = 0
	Ordered         = "ordered"
	StatusUnpaid    = 1
	Unpaied         = "unpaid"
	StatusPaid      = 2
	Paid            = "paid"
	StatusUnshipped = 3
	Unshipped       = "unshipped"
	StatusShipping  = 4
	Shipping        = "shipping"
	StatusShipped   = 5
	Shipped         = "shipped"
	StatusConfirmed = 6
	Confirmed       = "confirmed"
	StatusUnrated   = 7
	Unrated         = "unrated"
	StatusRated     = 8
	Rated           = "rated"
	StatusCompleted = 9
	Completed       = "completed"
	StatusCanceled  = 10
	Canceled        = "canceled"
	StatusClosed    = 11
	Closed          = "closed"
	StatusDeleted   = 12
	Deleted         = "deleted"
	StatusReturning = 13
	Returning       = "returning"
	StatusReturned  = 14
	Returned        = "returned"
)

func OrderStatus2Pb(orderStatus string) int64 {
	if strings.EqualFold(orderStatus, Ordered) {
		return StatusOrdered
	}
	// Unpaied
	if strings.EqualFold(orderStatus, Unpaied) {
		return StatusUnpaid
	}
	// Paid
	if strings.EqualFold(orderStatus, Paid) {
		return StatusPaid
	}
	// Unshipped
	if strings.EqualFold(orderStatus, Unshipped) {
		return StatusUnshipped
	}
	// Shipping
	if strings.EqualFold(orderStatus, Shipping) {
		return StatusShipping
	}
	// Shipped
	if strings.EqualFold(orderStatus, Shipped) {
		return StatusShipped
	}
	// Confirmed
	if strings.EqualFold(orderStatus, Confirmed) {
		return StatusConfirmed
	}
	// Unrated
	if strings.EqualFold(orderStatus, Unrated) {
		return StatusUnrated
	}
	// Rated
	if strings.EqualFold(orderStatus, Rated) {
		return StatusRated
	}
	// Completed
	if strings.EqualFold(orderStatus, Completed) {
		return StatusCompleted
	}
	// Canceled
	if strings.EqualFold(orderStatus, Canceled) {
		return StatusCanceled
	}
	// Closed
	if strings.EqualFold(orderStatus, Closed) {
		return StatusClosed
	}
	// Deleted
	if strings.EqualFold(orderStatus, Deleted) {
		return StatusDeleted
	}
	// Returning
	if strings.EqualFold(orderStatus, Returning) {
		return StatusReturning
	}
	// Returned
	if strings.EqualFold(orderStatus, Returned) {
		return StatusReturned
	}

	return -1
}

func OrderStatus2Web(orderStatus int64) string {
	if orderStatus == StatusOrdered {
		return Ordered
	}

	if orderStatus == StatusUnpaid {
		return Unpaied
	}
	// Paid
	if orderStatus == StatusPaid {
		return Paid
	}
	// Unshipped
	if orderStatus == StatusUnshipped {
		return Unshipped
	}
	// Shipping
	if orderStatus == StatusShipping {
		return Shipping
	}
	// Shipped
	if orderStatus == StatusShipped {
		return Shipped
	}
	// Confirmed
	if orderStatus == StatusConfirmed {
		return Confirmed
	}
	// Unrated
	if orderStatus == StatusUnrated {
		return Unrated
	}
	// Rated
	if orderStatus == StatusRated {
		return Rated
	}
	// Completed
	if orderStatus == StatusCompleted {
		return Completed
	}
	// Canceled
	if orderStatus == StatusCanceled {
		return Canceled
	}
	// Closed
	if orderStatus == StatusClosed {
		return Closed
	}
	// Deleted
	if orderStatus == StatusDeleted {
		return Deleted
	}
	// Returning
	if orderStatus == StatusReturning {
		return Returning
	}
	// Returned
	if orderStatus == StatusReturned {
		return Returned
	}

	return ""
}
