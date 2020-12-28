package model

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/util"

	"database/sql"
)

const (
	_ = iota
	stmtTransactionAddLog
)

// NewTransaction .
func NewTransaction(db *sql.DB, connection *pb.StoreConnection) *Transaction {
	model := &Transaction{}
	model.db = db
	model.connection = connection
	model.stmtMap = make(map[int]*sql.Stmt)

	var err error

	if _, err = model.genStmt(stmtTransactionAddLog, model.genSPStrWithSchema("Transaction_SP_AddLog(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")); err != nil {
		glog.Error(err)
	}

	if err != nil {
		glog.Error(err)
	}

	return model
}

// Transaction .
type Transaction struct {
	base
}

// AddLog .
// IN player_id varchar(64),
// IN game_type varchar(64),
// IN game_id varchar(64),
// IN round_id varchar(64),
// IN timestamp bigint unsigned,
// IN id_lock varchar(64),
// IN id_cancel_lock varchar(64),
// IN id_bet varchar(64),
// IN id_cancel_bet varchar(64),
// IN id_settle varchar(64),
// IN id_unlock varchar(64),
// IN amount decimal(10,6) unsigned,
// IN credit bigint unsigned,
// IN jackpot bigint unsigned,
// IN currency decimal(10,6) unsigned,
// IN currency_code varchar(8),
// IN system_fee bigint unsigned,
// IN system_fee_jackpot bigint unsigned,
// IN percentage_win_to_system_fee int unsigned,
// IN percentage_system_fee_to_jackpot int unsigned,
// IN percentage_win_jackpot int unsigned,
// IN success tinyint,
// IN platform varchar(64),
// IN platform_player_id varchar(64),
// IN platform_player_name varchar(64),
// IN platform_player_display_name varchar(64),
// IN platform_player_balance decimal(16,6) unsigned,
// IN raw_data text,
func (model *Transaction) AddLog(transaction *pb.Transaction) (err error) {
	stmt := model.getStmt(stmtTransactionAddLog)

	jsonString, err := util.ConvertProtoMessageToJSONString(transaction)

	_, err = stmt.Exec(
		transaction.Type,
		transaction.PlayerId,
		transaction.GameType,
		transaction.GameId,
		transaction.RoundId,
		transaction.Timestamp,
		transaction.TransactionIdLock,
		transaction.TransactionIdCancel,
		transaction.TransactionIdBet,
		transaction.TransactionIdCancelBet,
		transaction.TransactionIdSettle,
		transaction.TransactionIdUnlock,
		transaction.Amount,
		transaction.Credit,
		transaction.Jackpot,
		transaction.Currency,
		transaction.CurrencyCode,
		transaction.SystemFee,
		transaction.SystemFeeJackpot,
		transaction.PercentageWinToSystemFee,
		transaction.PercentageSystemFeeToJackpot,
		transaction.PercentageWinJackpot,
		transaction.PlatformPlayer.Platform,
		transaction.PlatformPlayer.Id,
		transaction.PlatformPlayer.Name,
		transaction.PlatformPlayer.DisplayName,
		transaction.PlatformPlayer.Balance,
		jsonString,
	)

	if err != nil {
		glog.Error(err)
	}

	return
}
