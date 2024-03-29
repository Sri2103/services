package dependency

import (
	"fmt"

	"github.com/Sri2103/services/internal/api-gateway/config"
	"github.com/Sri2103/services/pkg/logger"
	"github.com/eko/gocache/store/redis/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Dependency struct {
	Logger      *zap.Logger
	Redis       *redis.RedisStore
	Config      *config.AppConfig
	Server      *echo.Echo
	ProductConn *grpc.ClientConn
	CartConn    *grpc.ClientConn
	OrderConn   *grpc.ClientConn
	UserConn    *grpc.ClientConn
}

func NewDependency() (*Dependency, error) {
	var d Dependency
	d.setUpLogger()
	d.setUpConfig()
	d.Server = echo.New()
	d.SetUpProductClient()
	d.SetUpCartClient()
	d.SetUpOrderClient()
	d.SetUpUserClient()
	return &d, nil
}

func (d *Dependency) setUpLogger() {
	l, err := logger.NewLogger()
	d.Logger = l
	if err != nil {
		panic(err)
	}
}

func (d *Dependency) setUpConfig() {
	cfg := config.ReadConfig()
	d.Config = cfg
}

func (d *Dependency) SetUpProductClient() {
	connStr := fmt.Sprintf("%s:%d", d.Config.ProductConfig.Host, d.Config.ProductConfig.Port)
	conn, err := grpc.Dial(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	d.ProductConn = conn

}

func (d *Dependency) SetUpCartClient() {
	connStr := fmt.Sprintf("%s:%d", d.Config.CartConfig.Host, d.Config.CartConfig.Port)
	fmt.Println(connStr, "Connection string for  cart server")
	conn, err := grpc.Dial(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	d.CartConn = conn
}
func (d *Dependency) SetUpOrderClient() {
	connStr := fmt.Sprintf("%s:%d", d.Config.OrderConfig.Host, d.Config.OrderConfig.Port)
	fmt.Println(connStr, "Connection string for  order server")
	conn, err := grpc.Dial(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	d.OrderConn = conn
}
func (d *Dependency) SetUpUserClient() {
	connStr := fmt.Sprintf("%s:%d", d.Config.UserConfig.Host, d.Config.UserConfig.Port)
	fmt.Println(connStr, "Connection string for  user server")
	conn, err := grpc.Dial(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	d.UserConn = conn
}
