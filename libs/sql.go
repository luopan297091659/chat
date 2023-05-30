package libs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type LBVerify struct {
	Id int32  `gorm:"column:id"`
	LbVmZoneId int `gorm:"column:lb_vm_zone_id"`
	LbVmZoneName string `gorm:"column:lb_vm_zone_name"`
	LbVmRs1 string `gorm:"column:lb_vm_rs1"`
	LbVmRs2 string `gorm:"column:lb_vm_rs2"`
	LbUuid string `gorm:"column:lb_uuid"`
	LbVmSecgrouId string `gorm:"column:lb_vm_secgroupid"`
	PubNetworkId string `gorm:"column:pub_network_id"`
	PubSubNetworkId string `gorm:"column:pub_subnetwork_id"`
	PriNetworkId string `gorm:"column:pri_network_id"`
	PriSubNetworkId string `gorm:"column:pri_subnetwork_id"`
	Status  int `gorm:"column:status"`
	HistoryStatus   int `gorm:"column:history_status"`
	CreateTime time.Time `gorm:"create_time"`
	EndTime time.Time `gorm:"end_time"`
}


// 初始化数据库连接，并赋值给全局变量 db
func InitDB() (*gorm.DB, error) {

	// 打开一个数据库连接并连接到数据库
	dsn := "root:123456@tcp(45.43.33.78:4306)/es_verify?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 将数据库连接赋值给全局变量 db
	db.Table("lb_zone_verify").AutoMigrate(&LBVerify{})
	return db, nil
}

func StrToInt(ss string) int {
	s, err :=  strconv.Atoi(ss)
	if err != nil {
		// 处理转换错误
		log.Print("string to int failed：", err)
	}
	return s
}


func InsetSql(db *gorm.DB, lb_uuid string) int32 {
	//lb_zone_id, err :=  strconv.Atoi(zone_id)
	//if err != nil {
	//	// 处理转换错误
	//	log.Print("string to int failed：", err)
	//}
	lb_zone_id := StrToInt(zone_id)
	lb_zone_name := GetZoneName(zone_id)
	//db.Table("lb_zone_verify").AutoMigrate(&LBVerify{})
	lbverify :=&LBVerify{LbUuid: lb_uuid, LbVmZoneId: lb_zone_id, LbVmZoneName: lb_zone_name, LbVmRs1: rs1_ip,LbVmSecgrouId: seg_id,
		PubNetworkId: pub_network_id,PubSubNetworkId: pub_subnetwork_id,
		PriNetworkId: pri_network_id,PriSubNetworkId: pri_subnetwork_id,
		CreateTime:	time.Now(),EndTime: time.Now()}
	log.Print(lbverify)
	db.Table("lb_zone_verify").Create(lbverify)
	return lbverify.Id
}

func UpdateSql(db *gorm.DB, lb_uuid string,status_num int)  {
	if err := db.Table("lb_zone_verify").Where("lb_uuid = ?", lb_uuid).Updates(LBVerify{Status:status_num,
		EndTime: time.Now(),
		HistoryStatus: 1}).Error; err != nil {
		log.Print("update colimn faild ,id is ",&LBVerify{})
	}
}

func UpdateSqlHistory(db *gorm.DB)  {
	lb_zone_id := StrToInt(zone_id)
	if err := db.Table("lb_zone_verify").Where("lb_vm_zone_id = ? AND history_status=1", lb_zone_id).Updates(LBVerify{
		HistoryStatus: -1}).Error; err != nil {
		log.Print("update colimn faild")
	}
}

func QuerySql(db *gorm.DB) (bool) {
	//var lbverify LBVerify
	lb_zone_id := StrToInt(zone_id)
	var count int64
	//rows,err := db.Table("lb_zone_verify").Where("lb_vm_zone_id = ?", lb_zone_id).Find(&lbverify).Rows()
	db.Table("lb_zone_verify").Where("lb_vm_zone_id = ?", lb_zone_id).Count(&count)

	log.Print(count)
	if count == 0 {
		// 处理查询错误
		return false
	}  else {
		// 处理查询结果不为空的情况
		return true
	}
}

func QueryNewZoneVerify(db *gorm.DB) []LBVerify {
	lb_zone_id := StrToInt(zone_id)
	log.Print(lb_zone_id)
	var lbverify []LBVerify
	db.Table("lb_zone_verify").Where("lb_vm_zone_id = ? and history_status=1", lb_zone_id).Find(&lbverify)
	log.Print(lbverify)
	return lbverify
}

func QueryZoneNewVerify(db *gorm.DB) []LBVerify {
	var lbverify []LBVerify
	db.Table("lb_zone_verify").Where("history_status=1").Find(&lbverify)
	log.Print(lbverify)
	return lbverify
}

func QueryZoneVerify(db *gorm.DB) []LBVerify {
	lb_zone_id := StrToInt(zone_id)
	var lbverify []LBVerify
	db.Table("lb_zone_verify").Where("lb_vm_zone_id = ?", lb_zone_id).Find(&lbverify)
	log.Print(lbverify)
	return lbverify
}

func QueryAllZoneVerify(db *gorm.DB) []LBVerify {
	var lbverify []LBVerify
	db.Table("lb_zone_verify").Find(&lbverify)
	log.Print(lbverify)
	return lbverify
}