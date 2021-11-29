package main

import (
	_middleware "investaBackend/app/middlewares"
	"investaBackend/app/routes"
	bankUsecase "investaBackend/business/bank"
	investUseCase "investaBackend/business/investasi"
	proyekMitraUseCase "investaBackend/business/proyek_mitra"

	userUsecase "investaBackend/business/users"
	bankController "investaBackend/controllers/bank"
	investController "investaBackend/controllers/investasi"
	proyekMitraController "investaBackend/controllers/proyek_mitra"

	userController "investaBackend/controllers/users"
	bankRepo "investaBackend/drivers/databases/bank"
	proyekMitraRepo "investaBackend/drivers/databases/proyek_mitra"

	investRepo "investaBackend/drivers/databases/investasi"
	userRepo "investaBackend/drivers/databases/users"

	userInvestasiUsecase "investaBackend/business/user_investasi"
	userInvestasiController "investaBackend/controllers/user_investasi"
	proyekMitra "investaBackend/drivers/databases/proyek_mitra"
	sektorRepo "investaBackend/drivers/databases/sektor"
	userInvestorRepo "investaBackend/drivers/databases/user_investasi"

	"investaBackend/drivers/mysql"

	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("app/config/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&bankRepo.Bank{})
	db.AutoMigrate(&sektorRepo.Sektor{})
	db.AutoMigrate(&userRepo.User{})
	db.AutoMigrate(&userInvestorRepo.UserInvestasi{})
	db.AutoMigrate(&proyekMitra.ProyekMitra{})
	db.AutoMigrate(&investRepo.Investasi{})

}

func main() {
	configDb := mysql.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDb.InitialDb()
	dbMigrate(db)

	jwt := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()
	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUsecase.NewUsecase(userRepoInterface, timeoutContext, &jwt)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	userInvestorRepoInterface := userInvestorRepo.NewUserInvestasiRepository(db)
	userInvestorUseCaseInterface := userInvestasiUsecase.NewUsecase(userInvestorRepoInterface, timeoutContext, &jwt)
	userInvestorControllerInterface := userInvestasiController.NewUserInvestasiController(userInvestorUseCaseInterface)

	bankRepoInterface := bankRepo.NewBankRepository(db)
	bankUseCaseInterface := bankUsecase.NewBankUseCase(bankRepoInterface, timeoutContext)
	bankControllerInterface := bankController.NewBankController(bankUseCaseInterface)

	investRepoInterface := investRepo.NewInvestasiRepository(db)
	investUseCaseInterface := investUseCase.NewInvestasiUseCase(investRepoInterface, timeoutContext)
	investController := investController.NewInvestasiController(investUseCaseInterface)

	proyekMitraInterface := proyekMitraRepo.NewProyekMitraRepository(db)
	proyekMitraUseCaseInterface := proyekMitraUseCase.NewProyekMitraUseCase(proyekMitraInterface, timeoutContext)
	proyekMitraControllerInterface := proyekMitraController.NewProyekMitraController(proyekMitraUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController:         *userControllerInterface,
		BankController:         *bankControllerInterface,
		UserInvestorController: *userInvestorControllerInterface,
		InvestasiController:    *investController,
		ProyekMitraController:  *proyekMitraControllerInterface,
		JWTConfig:              jwt.Init(),
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}
