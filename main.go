package main

import (
	"package-mvc/app/controller"
	"package-mvc/app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	cfg.AllowHeaders = []string{"Authorization", "Origin", "Accept", "X-Requested-With", " Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	router.Use(cors.New(cfg))
	router.POST("/api/v1/account/add", controller.CreateAccount)
	router.POST("/api/v1/login", controller.Login)
	router.GET("/api/v1/account", middleware.Auth, controller.GetAccount)
	router.POST("/api/v1/transfer", middleware.Auth, controller.Transfer)
	router.POST("/api/v1/withdraw", middleware.Auth, controller.Withdraw)
	router.POST("/api/v1/deposit", middleware.Auth, controller.Deposit)
	router.POST("/api/v1/interest", middleware.Auth, controller.Interest)
	// router.POST("/test", testing)
	router.Run(":8000")
}

// func testing(c *gin.Context) {

// 	var auth model.Auth
// 	if err := c.Bind(&auth); err != nil {
// 		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	pass, err := utils.HashGenerator(auth.Password)
// 	if err != nil {
// 		utils.WrapAPIError(c, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	auth.Password = pass
// 	// account.AccountNumber = utils.RangeIn(111111, 999999)
// 	// account.Saldo = 0
// 	// account.IdAccount = fmt.Sprintf("id-%d", utils.RangeIn(111, 999))
// 	if err := model.DB.Create(&auth).Error; err != nil {
// 		utils.WrapAPIError(c, "error", http.StatusBadRequest)
// 		return
// 	}
// 	utils.WrapAPISuccess(c, "success", http.StatusOK)
// 	return
// 	// e, err := json.Marshal(account)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println(string(e))
// 	// body, _ := ioutil.ReadAll(c.Request.Body)
// 	// println(string(body))
// }
