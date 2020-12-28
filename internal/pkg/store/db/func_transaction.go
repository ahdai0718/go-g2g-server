package db

import "ohdada/g2gserver/internal/pkg/pb"

// TransactionAddLog .
func TransactionAddLog(transaction *pb.Transaction) error {
	return transactionModel.AddLog(transaction)
}
