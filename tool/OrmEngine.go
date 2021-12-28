package tool

import (
	"CloudRestaurant/model"

	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

type Orm struct {
	*xorm.Engine
}

var Dbengine *Orm //进行全局赋值

func OrmEngine(cfg *Config) (*Orm, error) {
	dbconfig := cfg.Database
	conn := dbconfig.User + ":" + dbconfig.Password + "@tcp(" + dbconfig.Host + ":" + dbconfig.Port + ")/" + dbconfig.DbName + "?charset=" + dbconfig.CharSet
	engine, err := xorm.NewEngine(dbconfig.Driver, conn)
	if err != nil {
		return nil, err
	}

	//是否提示操作状态
	engine.ShowSQL(dbconfig.IsShowsql)
	// 将结构体映射为数据库的表
	err = engine.Sync2(new(model.Smscode), new(model.Member))
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	Dbengine = orm //进行全局赋值
	return orm, nil
}
