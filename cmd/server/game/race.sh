#!/bin/sh

go run -race main.go \
--stderrthreshold INFO \
--color true \
--server-name GameExample \
--server-host 0.0.0.0 \
--server-port 3011 \
--server-ws-protocol ws \
--server-ws-route-path ws \
--server-host-for-client 0.0.0.0 \
--server-port-for-client 3011 \
--server-ws-protocol-for-client ws \
--server-ws-route-path-for-client ws \
--run-mode dev \
--gin-mode debug \
--ip-address 0.0.0.0 \
--db-host localhost \
--db-port 3306 \
--db-schema dev_g2g_server \
--db-schema-game dev_g2g_game_example \
--db-user dev_g2gserverbase \
--db-password !QAZ2wsx \
--db-max-connection 32 \
--use-database-server-info true \
--game-type 1001 \
--v 2