module ohdada/g2gserver

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-resty/resty/v2 v2.3.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	google.golang.org/protobuf v1.24.0
)

replace github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b => github.com/KairosAerospace/glog v0.0.0-20170802202929-7e85df754651
