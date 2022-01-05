package tool

import (
	"CloudRestaurant/model"
	"fmt"

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

	//初始化时将结构体映射为数据库的表
	err = engine.Sync2(
		new(model.Smscode),
		new(model.Member),
		new(model.FoodCategory),
		new(model.Shop),
		new(model.Service),
		new(model.ShopService))
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	Dbengine = orm //进行全局赋值

	//初始化shop数据
	//InitShopData()
	return orm, nil
}

//向shop表中初始化数据
func InitShopData() {
	shops := []model.Shop{
		{Id: 1, Name: "嘉禾一品（温都水城）", Address: "北京市昌平区宏福苑温都水城F1", Longitude: 116.36868, Latitude: 40.10039,
			Phone: "13437850035", Status: 1, RecentOrderNum: 106, RatingCount: 961, Rating: 4.7, PromotionInfo: "欢迎光临，用餐高峰请提前下单，谢谢",
			OpeningHours: "8:30/20:30", IsNew: true, IsPremium: true, Imagepath: "", MinimumOrderAmount: 20, DeliveryFee: 5},

		{Id: 479, Name: "杨国福麻辣烫", Address: "北京市市蜀山区南二环路天鹅湖万达广场8号楼1705室", Longitude: 117.22124, Latitude: 31.81948, Phone: "13167583411",
			Status: 1, RecentOrderNum: 755, RatingCount: 167, Rating: 4.2, PromotionInfo: "欢迎光临，用餐高峰请提前下单，谢谢", OpeningHours: "8:30/20:30",
			IsNew: true, IsPremium: true, Imagepath: "", MinimumOrderAmount: 20, DeliveryFee: 5},

		{Id: 485, Name: "好适口", Address: "北京市海淀区西二旗大街58号", Longitude: 120.65355, Latitude: 31.26578, Phone: "12345678901",
			Status: 1, RecentOrderNum: 58, RatingCount: 576, Rating: 4.6, PromotionInfo: "欢迎光临，用餐高峰请提前下单，谢谢", OpeningHours: "8:30/20:30",
			IsNew: true, IsPremium: true, Imagepath: "", MinimumOrderAmount: 20, DeliveryFee: 5},

		{Id: 486, Name: "东来顺旗舰店", Address: "北京市天河区东圃镇汇彩路38号1领汇创展商务中心401", Longitude: 113.41724, Latitude: 23.1127, Status: 1,
			Phone: "13544323775", RecentOrderNum: 542, RatingCount: 372, Rating: 4.2, PromotionInfo: "老北京正宗涮羊肉,非物质文化遗产",
			OpeningHours: "09:00/21:30", IsNew: true, IsPremium: true, Imagepath: "", MinimumOrderAmount: 20, DeliveryFee: 5},

		{Id: 487, Name: "北京酒家", Address: "北京市海淀区上下九商业步行街内", Longitude: 113.24826, Latitude: 23.11488, Phone: "13257482341", Status: 0,
			RecentOrderNum: 923, RatingCount: 871, Rating: 4.2, PromotionInfo: "北京第一家传承300年酒家", OpeningHours: "8:30/20:30", IsNew: true, IsPremium: true, Imagepath: "",
			MinimumOrderAmount: 20, DeliveryFee: 5},

		{Id: 488, Name: "和平鸽饺子馆", Address: "北京市越秀区德政中路171", Longitude: 113.27521, Latitude: 23.12092,
			Phone: "17098764762", Status: 1, RecentOrderNum: 483, RatingCount: 273, Rating: 4.2, PromotionInfo: "吃饺子就来和平鸽饺子馆", OpeningHours: "8:30/20:30",
			IsNew: true, IsPremium: true, Imagepath: "", MinimumOrderAmount: 20, DeliveryFee: 5}}

	//多条数据-启用事务
	//事务
	session := Dbengine.NewSession()
	defer session.Close()

	//事务操作:开始-执行-回滚-提交
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, shop := range shops {
		_, err := session.Insert(&shop)
		if err != nil {
			session.Rollback() //回滚
			return
		}
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}

}
