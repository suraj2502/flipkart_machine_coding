package main

import (
	"fmt"
	"suraj_projects/flipkart_machine_coding/affiliate"
	"suraj_projects/flipkart_machine_coding/order"
	"suraj_projects/flipkart_machine_coding/product"
)

// Commission App

// Platform Overview:
// Flipkart has developed an Affiliate feature where an affiliate influences someone to place an order on Flipkart. For each order placed through an affiliate, the affiliate gets a certain commission. The commission needs to be paid out to the affiliate.

// Here is how the flow works:
// An affiliate (example: a YouTube channel with a lot of subscribers) shares a FK product link on their channel.
// When the affiliate’s subscribers land on Flipkart using that affiliate link and place an order, the affiliate gets a commission from Flipkart.
// Internal Capabilities:
// Order lifecycle:
// 	The order consists of:
// OrderId: unique id that identifies an order.
// Price: Price paid by the user for the order.
// ProductId: unique productId that identifies what product was purchased (even though Flipkart orders can contain multiple products, but here we can assume each order contains exactly one productId).
// CategoryId: which category this product belongs to. There can be four categories: [“Clothing”, “Mobiles”, “MobileCovers”, “Furniture”].
// AffiliateId: Affiliate who should get the commission.
// Timestamp: time when the order was placed.

// Commission calculation:
// There are multiple commission rules defined in the system (one for each category). Commission is calculated in the following way:
// Each category has a percentage/flat-rate and a maxCap. Each order is paid based on these two.
// Ex:
// {
// “Category”: “Mobiles”,
// “Percentage”: “10”,
// “MaxCap”: “50” 		//we pay min(10% of orderPrice, maxCap)
// }

// {
// “Category”: “MobileCovers”,
// “FlatRate”: “5”,			//we pay a flat rate of Rs. 5
// }
// When an order is placed, we need to figure out which rule is applicable to the order, and calculate the commission accordingly.
// We can assume that the commission rules will never change.

// Order Transition and Payout:
// An order goes through multiple stages. A typical order stage looks like this:

// Happy case: where we need to pay commissions
// CREATED -> DISPATCHED ->DELIVERED -> RETURN_PERIOD_EXPIRED

// Other cases: where we don’t need to pay any commissions
// CREATED -> CANCELED
// CREATED -> DISPATCHED -> CANCELED
// CREATED -> DISPATCHED -> DELIVERED -> RETURNED

// We need to expose an API to listen to these state change events at the orderId level.
// If the order reaches a CANCELED/RETURNED stage, we don’t need to pay out anything.
// Once the order’s return period has expired, the commission can be paid. Each payment requires a transaction fee. So, it is better if we wait till the total pending commission reaches Rs. 100 and only then initiate the payouts. Each payout should have a unique transactionId for accounting purposes. We just need to create a transaction entity, and assume that the payout will happen via another batch job (No need to take transaction fee into account while calculating commission).
// Note- The commission calculation needs to happen at the CREATED state, but the actual payout needs to happen at the RETURN_PERIOD_EXPIRY state. If we reach CANCELED/RETURNED state instead of the RETURN_PERIOD_EXPIRY state, we simply cancel the commission for that order.

// Example:
// Let’s say we have three orders (o1, o2 and o3) in the DELIVERED state for an affiliate a1, each with commission Rs. 60. (Total potential payment = 180).
// When we receive the RETURN_PERIOD_EXPIRED event for o1, we do nothing (as total commission to payout is < 100). (Total potential payment = 120, total earning = 60).
// When we receive the RETURNED event for o3, we simply cancel the pending commission of Rs. 60 for o3. (Total potential payment = 60, total earning = 60).
// When we receive the RETURN_PERIOD_EXPIRED event for o2, the total commission to pay out becomes 120 (> 100), so we create a new transaction and add o1 and o2.

// External Capabilities:
// An affiliate can come and fetch all his/her transactions along with other details like:
// transactionId
// Total amount paid.
// OrderIds (and associated commission) that were paid as part of this transactionId
// Get all orders by state and affiliateId.

// Bonus Question (to be attempted only when everything else is done!):
// Assume that the commission rules can change anytime. Also, the order events can arrive with a delay. So, it is possible that the rules might have changed by the time an Order Event reaches our system. So, now we need to calculate the commissions in a “time-sensitive” way, where we apply the rules that were applicable during the order placement time.
// Handle concurrent Order Events

// Guidelines:
// Input can be read from a file or STDIN or coded in a driver method.
// Output can be written to a file or STDOUT.
// Feel free to store all interim/output data in-memory.
// Restrict internet usage to looking up syntax.
// Bonus Question is OPTIONAL, and should only be attempted once everything else is COMPLETE.
// You are free to use the language of your choice.
// Save your code/project by your name and email it. Your program will be executed on another machine. So, explicitly specify dependencies, if any, in your email.

// Expectations:
// Code should be demo-able (very important). Code should be functionally correct and complete. (Have a Main/Driver class to demo the code)
// At the end of this interview round, an interviewer will provide multiple inputs to your program for which it is expected to work
// Code should handle edge cases properly and fail gracefully. Add suitable exception handling, wherever applicable.
// Code should have good object-oriented design.
// Code should be readable, modular, testable and extensible. Use intuitive names for your variables, methods and classes.
// It should be easy to add/remove functionality without rewriting a lot of code.
// Do not write monolithic code.
// No need to use any real databases, please use appropriate data structures to maintain the data in-memory.
// Only public methods need to be exposed (tested via Driver/Main class), REST apis are not needed.

func main() {
	// Driver code
	OrderS := order.NewOrderService()
	ProductServive := product.NewProductMgr()
	AffiliateService := affiliate.NewAffiliateMgr()
	// Category1 := category.NewCategory("Mobiles")
	// Category2 := category.NewCategory("Clothing")

	AffiliateService.CreateAffiliate(affiliate.Affiliate{
		AffiliateName: "suraj",
	})

	ProductServive.CreateProduct(product.Product{
		ProductId:   "1",
		CategoryId:  "1",
		Price:       100,
		ProductName: "toothpaste",
	})
	ordererr := OrderS.CreateOrder(order.Order{
		Price:         100,
		ProductId:     "1",
		CategoryId:    "1",
		AffiliateId:   "1",
		CommisionType: "Percentage",
		Timestamp:     "1",
		State:         "Created",
	})
	if ordererr != nil {
		fmt.Println("orderErr", ordererr)
	} else {
		o, _ := OrderS.GetOrder("1")
		fmt.Println("order created successfully", o)

	}

	OrderS.UpdateOrderState("1", "Dispatched")
	OrderS.UpdateOrderState("1", "Delivered")
	onewState, _ := OrderS.GetOrder("1")
	fmt.Println("Order State", onewState)
	value, e := OrderS.GetTotalCommisionForAffiliate("1")
	if e != nil {
		fmt.Println("error", e)
	} else {
		fmt.Println("Total Commission", value)

	}

	// Create an affiliate

	// Create an order
	// Update the order state
	// Get the order
	// Get the orders by state and affiliateId
	// Create a commission rule
	// Calculate the commission
	// Create a commission
	// Create a transaction
	// Get the transactions by affiliateId

}
