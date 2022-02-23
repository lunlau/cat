package mysqldao

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"log/syslog"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

type mysqlDAOConfig struct {
	DriverName string
	DataSourceName string
}

type UserDao struct {
	engine *xorm.Engine
	cfg *mysqlDAOConfig
}

const (
	driveName = "mysql"
	dataSourceName = "root:123@/test?charset=utf8"
)
func InitDao()  {

}

func NewUserDao(ctx context.Context, )  {

}
//func main() {
//	var err error
//	engine, err = xorm.NewEngine(driveName, dataSourceName)
//	fmt.Println(err)
//}

type TUser struct {
	Id   int64    `xorm:"not null pk autoincr INT"`
	Name string `xorm:"not null default '' VARCHAR(128)"`
}

type UserDO struct {}

func (s * UserDao)AddUser(ctx context.Context, userDo TUser) (int64, error) {
	uniqueId, err := s.engine.Insert(&userDo)
	if err != nil {
		return 0, fmt.Errorf("db insert failed, err : %+v", err)
	}
	return uniqueId, err
}

func (s * UserDao)UpdateUser(ctx context.Context, userDo TUser) (int64, error) {
	uniqueId, err := s.engine.Update(&userDo)
	if err != nil {
		return 0, fmt.Errorf("db update failed, err : %+v", err)
	}
	return uniqueId, err
}


func (s * UserDao)QueryOne(ctx context.Context, id int64) (*TUser, error) {
	userDo := &TUser{
		Id:id,
	}
	userInfoList, err := s.engine.QueryString(&userDo)
	if err != nil {
		return &TUser{}, fmt.Errorf("db query  failed, err : %+v", err)
	}
	fmt.Println("query one record %+v", userInfoList)
	return &TUser{}, err
}

func (s * UserDao)QueryList(ctx context.Context, id int64) ([]*TUser, error) {
	userDo := &TUser{
		Id:id,
	}
	userInfoList, err := s.engine.QueryString(&userDo)
	if err != nil {
		return nil, fmt.Errorf("db query  failed, err : %+v", err)
	}
	fmt.Println("query one record %+v", userInfoList)
	resList := make([]*TUser, 0)
	return resList, err
}

func (s * UserDao)createLoger(ctx context.Context) error {
	logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	if err != nil {
		log.Fatalf("Fail to create xorm system logger: %v\n", err)
		return err
	}
	logger := xlog.NewSimpleLogger(logWriter)
	logger.ShowSQL(true)
	s.engine.SetLogger(logger)
	return nil
}