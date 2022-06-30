#!/bin/sh

go run -race main.go \
--stderrthreshold WARNING \
--color true \
--server-name Gateway001 \
--server-host 0.0.0.0 \
--server-port 3000 \
--run-mode dev \
--gin-mode debug \
--ip-address 0.0.0.0 \
--db-host localhost \
--db-port 3306 \
--db-schema dev_g2g_server_base \
--db-user dev_g2gserverbase \
--db-password !QAZ2wsx \
--db-max-connection 32 \
--use-database-server-info true