package utils

import (
	"strconv"
	"sync"
	"time"
)

var (
	offerCurrentNumber     int = 1
	productCurrentNumber   int = 1
	affiliateCurrentNumber int = 1
	CategoryCurrentNumber  int = 1
	mu                     sync.Mutex
)

type IdGenerator interface {
	GenerateUniqueId() int
}

// GenerateOfferUniqueId generates a unique id for an offer
func GenerateOfferUniqueId() string {
	mu.Lock()
	defer mu.Unlock()
	n := offerCurrentNumber
	offerCurrentNumber++
	return strconv.Itoa(n)
}

func GenerateProductUniqueId() string {
	mu.Lock()
	defer mu.Unlock()
	n := productCurrentNumber
	productCurrentNumber++
	return strconv.Itoa(n)
}

func GenerateAffiliateUniqueId() string {
	mu.Lock()
	defer mu.Unlock()
	n := affiliateCurrentNumber
	affiliateCurrentNumber++
	return strconv.Itoa(n)
}

func GenerateCategoryUniqueId() string {
	mu.Lock()
	defer mu.Unlock()
	n := CategoryCurrentNumber
	CategoryCurrentNumber++
	return strconv.Itoa(n)
}

func StringInsideSlice(s string, slice []string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

// /check if time is before current time
func IsTimeCrossed(timestamp string) bool {

	parsedTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		// handle error
		return false
	}
	return parsedTime.Before(time.Now())

}

// /check if cancelWindow is crossed or not
func IsCancelWindowCrossed(timestamp string, cancelWindow int) bool {
	parsedTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		// handle error
		return false
	}
	return parsedTime.Add(time.Duration(cancelWindow) * time.Minute).Before(time.Now())
}
