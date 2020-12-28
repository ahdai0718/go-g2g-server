package platform

import (
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/store"
)

// LogTransaction .
func LogTransaction(transaction *pb.Transaction, platformPlayer *pb.PlatformPlayer) (err error) {

	if platformPlayer != nil {
		transaction.PlatformPlayer = platformPlayer
	}

	store.DefaultManager().AddTransactionLog(transaction)

	return
}
