package main

import "fmt"

type Transaction struct {
	Amount          float64
	TransactionsNum int
	AccountAge      int
	LocationChange  bool
	FailedLogins    int
}

type Node struct {
	Question string
	YesNode  *Node
	NoNode   *Node
}

const (
	LARGE_TXN_AMOUNT                  = "Is the transaction amount greater than $10,000?"
	MULTIPLE_TXN_LAST_HRS             = "Does the customer have more than 3 transactions in the last hour?"
	SUSPICIOUS_TRANSFER_LOCATION      = "Is the transaction from a different country than the customer's home address?"
	MULTIPLE_FAILED_ATTEMPTS_LAST_HRS = "Have there been multiple failed login attempts in the past 24 hours?"
	ACCOUNT_IS_NEW                    = "Is the account less than 1 month old?"
)

func main() {
	// Define the decision tree
	root := Node{
		Question: LARGE_TXN_AMOUNT,
		YesNode: &Node{
			Question: MULTIPLE_TXN_LAST_HRS,
			YesNode: &Node{
				Question: SUSPICIOUS_TRANSFER_LOCATION,
				YesNode: &Node{
					Question: MULTIPLE_FAILED_ATTEMPTS_LAST_HRS,
					YesNode:  &Node{Question: "Fraud"},
					NoNode:   &Node{Question: "Not Fraud"},
				},
				NoNode: &Node{Question: "Not Fraud"},
			},
			NoNode: &Node{Question: "Not Fraud"},
		},
		NoNode: &Node{
			Question: MULTIPLE_FAILED_ATTEMPTS_LAST_HRS,
			YesNode:  &Node{Question: "Fraud"},
			NoNode:   &Node{Question: "Not Fraud"},
		},
	}

	// Define a sample transaction
	txn := Transaction{
		Amount:          15000,
		TransactionsNum: 4,
		AccountAge:      2,
		LocationChange:  true,
		FailedLogins:    5,
	}

	// Evaluate the transaction
	result := evaluate(&root, txn)
	fmt.Println(result)
}

func evaluate(n *Node, t Transaction) string {
	switch n.Question {
	case LARGE_TXN_AMOUNT:
		if t.Amount > 10000 {
			return evaluate(n.YesNode, t)
		} else {
			return evaluate(n.NoNode, t)
		}
	case MULTIPLE_TXN_LAST_HRS:
		if t.TransactionsNum > 3 {
			return evaluate(n.YesNode, t)
		} else {
			return evaluate(n.NoNode, t)
		}
	case SUSPICIOUS_TRANSFER_LOCATION:
		if t.LocationChange {
			return evaluate(n.YesNode, t)
		} else {
			return evaluate(n.NoNode, t)
		}
	case MULTIPLE_FAILED_ATTEMPTS_LAST_HRS:
		if t.FailedLogins > 0 {
			return evaluate(n.YesNode, t)
		} else {
			return evaluate(n.NoNode, t)
		}
	case ACCOUNT_IS_NEW:
		if t.AccountAge < 1 {
			return evaluate(n.YesNode, t)
		} else {
			return evaluate(n.NoNode, t)
		}
	default:
		return n.Question
	}
}
