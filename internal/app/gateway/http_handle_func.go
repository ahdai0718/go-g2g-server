package gateway

import (
	"net/http"
	"ohdada/g2gserver/internal/pkg/auth"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/static/data"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	authenticator = auth.DefaultAuthenticatorSimpleFactory.Create(&pb.Auth{
		Type:   pb.AuthType_AT_JWT,
		Secret: "12345678",
	})
)

// AllowIntranetOnly .
func AllowIntranetOnly(context *gin.Context) {
	isAllow := false

	if strings.Contains(context.Request.RemoteAddr, "192.168.") {
		isAllow = true
	}

	if strings.Contains(context.Request.RemoteAddr, "127.0.0.1") {
		isAllow = true
	}

	if strings.Contains(context.Request.RemoteAddr, "0.0.0.0") {
		isAllow = true
	}

	for _, serverInfo := range data.DefaultManager().GetServerInfoMapByType(pb.ServerType_ST_GAME) {
		if strings.Contains(context.Request.RemoteAddr, serverInfo.Host) {
			isAllow = true
		}
	}

	if !isAllow {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}

// HandleRequestServerInfo .
func HandleRequestServerInfo(context *gin.Context) {
	context.ProtoBuf(http.StatusOK, ServerInfo())
}

// HandleGetUserTransactions .
func HandleGetUserTransactions(context *gin.Context) {

}

// CheckOAuthToken .
func CheckOAuthToken(context *gin.Context) {
	if authenticator.Type() == pb.AuthType_AT_OAUTH {
		authorization := context.GetHeader("Authorization")
		strList := strings.Split(authorization, "Bearer")
		if len(strList) != 2 {
			context.JSON(400, gin.H{
				"error": gin.H{
					"code":    -1,
					"message": "Authorization: not bearer.",
				},
			})
			context.Abort()
		} else {
			accessToken := strings.Trim(strList[1], " ")
			if accessToken != "12345678" {
				context.JSON(400, gin.H{
					"error": gin.H{
						"code":    -1,
						"message": "Authorization: access token invalid.",
					},
				})
				context.Abort()
			}
		}
	}
}

// CheckAccessToken .
func CheckAccessToken(context *gin.Context) {

	accessToken := context.GetHeader("X-Access-Token")

	if len(strings.Trim(accessToken, " ")) == 0 {
		context.JSON(400, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": "Access token is required.",
			},
		})
	}
}

// HandleOAuthAccessToken .
func HandleOAuthAccessToken(context *gin.Context) {
	clientID, _ := context.GetPostForm("client_id")
	clientSecret, _ := context.GetPostForm("client_secret")
	grantType, _ := context.GetPostForm("grant_type")
	scope, _ := context.GetPostForm("scope")

	if clientID == "default" &&
		clientSecret == "12345678" &&
		grantType == "client_credentials" &&
		scope == "bet" {
		context.JSON(200, gin.H{
			"access_token": "12345678",
			"expires_in":   time.Now().Add(time.Hour).Unix(),
			"token_type":   "Bearer",
			"scope":        "bet",
		})
	} else {
		context.JSON(401, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": "OAuth failed.",
			},
		})
	}
}

// HandlePlayerTokenValidate .
func HandlePlayerTokenValidate(context *gin.Context) {

	accessToken := context.GetHeader("X-Access-Token")

	authResult, err := authenticator.Verify(accessToken)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	context.JSON(200, gin.H{
		"id":            authResult.PlayerID,
		"name":          authResult.PlayerID,
		"language":      "zh-CN",
		"balance":       1e6,
		"currency_code": "TWD",
		"user_sn":       0,
		"icon":          "01",
	})
}

// HandlePlayerBalance .
func HandlePlayerBalance(context *gin.Context) {

	accessToken := context.GetHeader("X-Access-Token")

	_, err := authenticator.Verify(accessToken)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	context.JSON(200, gin.H{
		"balance":       1e10,
		"currency_code": "TWD",
		"userSN":        0,
	})
}

// HandlePlayerTransactionLock .
func HandlePlayerTransactionLock(context *gin.Context) {
	accessToken := context.GetHeader("X-Access-Token")

	_, err := authenticator.Verify(accessToken)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	body := struct {
		Amount        float64 `json:"amount"`
		CurrencyCode  string  `json:"currency_code"`
		TransactionID string  `json:"transaction_id"`
		GameID        string  `json:"game_id"`
		// Timestamp     int64   `json:"timestamp"`
		Timestamp string `json:"timestamp"`
	}{}

	err = context.BindJSON(&body)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	context.JSON(200, gin.H{
		"balance":       1e6,
		"currency_code": "TWD",
	})

}

// HandlePlayerTransactionCancel .
func HandlePlayerTransactionCancel(context *gin.Context) {
	accessToken := context.GetHeader("X-Access-Token")

	_, err := authenticator.Verify(accessToken)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	body := struct {
		CurrencyCode        string `json:"currency_code"`
		TransactionIDCancel string `json:"transaction_id_cancel"`
		TransactionIDOrigin string `json:"transaction_id_origin"`
		// Timestamp           int64  `json:"timestamp"`
		Timestamp string `json:"timestamp"`
	}{}

	err = context.BindJSON(&body)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	context.JSON(200, gin.H{
		"balance":       1e6,
		"currency_code": "TWD",
	})
}

// HandlePlayerTransactionUnlock .
func HandlePlayerTransactionUnlock(context *gin.Context) {
	accessToken := context.GetHeader("X-Access-Token")

	_, err := authenticator.Verify(accessToken)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	body := struct {
		CurrencyCode        string `json:"currency_code"`
		TransactionIDCancel string `json:"transaction_id_cancel"`
		TransactionIDOrigin string `json:"transaction_id_origin"`
		Timestamp           int64  `json:"timestamp"`
		// Timestamp string `json:"timestamp"`
	}{}

	err = context.BindJSON(&body)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	context.JSON(200, gin.H{
		"balance":       1e6,
		"currency_code": "TWD",
	})
}

// HandlePlayerBetPlace .
func HandlePlayerBetPlace(context *gin.Context) {
	accessToken := context.GetHeader("X-Access-Token")

	_, err := authenticator.Verify(accessToken)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	body := struct {
		Amount        float64 `json:"amount"`
		CurrencyCode  string  `json:"currency_code"`
		TransactionID string  `json:"transaction_id"`
		GameID        string  `json:"game_id"`
		// Timestamp     int64   `json:"timestamp"`
		Timestamp string `json:"timestamp"`
	}{}

	err = context.BindJSON(&body)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	context.JSON(200, gin.H{
		"balance":       1e6,
		"currency_code": "TWD",
	})

}

// HandlePlayerBetSettle .
func HandlePlayerBetSettle(context *gin.Context) {
	accessToken := context.GetHeader("X-Access-Token")

	_, err := authenticator.Verify(accessToken)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	body := struct {
		Amount              float64 `json:"amount"`
		CurrencyCode        string  `json:"currency_code"`
		TransactionIDSettle string  `json:"transaction_id_settle"`
		TransactionIDBet    string  `json:"transaction_id_bet"`
		GameID              string  `json:"game_id"`
		// Timestamp           int64   `json:"timestamp"`
		Timestamp string `json:"timestamp"`
	}{}

	err = context.BindJSON(&body)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	context.JSON(200, gin.H{
		"balance":       1e6,
		"currency_code": "TWD",
	})

}

// HandlePlayerBetCancel .
func HandlePlayerBetCancel(context *gin.Context) {
	accessToken := context.GetHeader("X-Access-Token")

	_, err := authenticator.Verify(accessToken)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	body := struct {
		CurrencyCode        string `json:"currency_code"`
		TransactionIDCancel string `json:"transaction_id_cancel"`
		TransactionIDOrigin string `json:"transaction_id_origin"`
		Timestamp           int64  `json:"timestamp"`
		// Timestamp string `json:"timestamp"`
	}{}

	err = context.BindJSON(&body)

	if err != nil {
		context.JSON(200, gin.H{
			"error": gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
	}

	context.JSON(200, gin.H{
		"balance":       1e6,
		"currency_code": "TWD",
	})
}
