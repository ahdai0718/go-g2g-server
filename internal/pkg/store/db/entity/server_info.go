package entity

import "database/sql"

// ServerInfo .
type ServerInfo struct {
	Name                        sql.NullString
	Group                       sql.NullString
	Type                        sql.NullInt32
	Host                        sql.NullString
	Port                        sql.NullInt64
	Protocol                    sql.NullString
	ServerInfoRoutePath         sql.NullString
	WebsocketProtocol           sql.NullString
	WebsocketRoutePath          sql.NullString
	HostForClient               sql.NullString
	PortForClient               sql.NullInt64
	WebsocketProtocolForClient  sql.NullString
	WebsocketRoutePathForClient sql.NullString
	PublicIPAddress             sql.NullString
	IsOffline                   sql.NullBool
	GameType                    sql.NullInt32
}
