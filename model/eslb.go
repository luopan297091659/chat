package model

import (
	"chat/libs"
	"gorm.io/gorm"
	"log"
	//"os"
	"errors"
)

func EsLBCheck(db *gorm.DB) (string,bool,error) {
	//db,_ := libs.InitDB()
	var lb_uuid string = libs.CreateLB()
	bol := libs.QuerySql(db)
	if bol {
		libs.UpdateSqlHistory(db)
	}
	//var lb_uuid string = "ad3b8e12-b7b7-40fe-abda-939ce309d8fb"
	libs.InsetSql(db, lb_uuid)
	log.Print("LB "+ lb_uuid +" have created" )
	if libs.WaitLBNoramal(lb_uuid){
		log.Print("LB "+ lb_uuid +" have got ready" )
	} else {
		libs.UpdateSql(db, lb_uuid, -1)
		log.Print("LB "+ lb_uuid +" haven't got ready, you must check the environment" )
		return lb_uuid,false,errors.New("vms or vip haven't got ready")
		//os.Exit(1)
	}

	var listener_uuid = libs.CreateListener(lb_uuid)
	libs.CreateBackend(listener_uuid)
	if libs.ApplyLB(lb_uuid) {
		log.Print("LB "+ lb_uuid +" apply successful")
		if libs.JudgeLBWorkStatus(lb_uuid) {
			log.Print("LB "+ lb_uuid +" work normal")
		} else {
			libs.UpdateSql(db, lb_uuid, -1)
			log.Print("LB "+ lb_uuid +" work unnormal")
			return lb_uuid,false,errors.New("LB "+ lb_uuid +" work unnormal")
		}
	} else {
		libs.UpdateSql(db, lb_uuid, -1)
		log.Print("LB "+ lb_uuid +" apply failed")
		return lb_uuid,false,errors.New("LB "+ lb_uuid +" apply failed")
		//os.Exit(1)
	}
	if DeleteLb(db, lb_uuid)==false {
		return lb_uuid,false,errors.New("LB "+ lb_uuid +" delete failed")
	}
	return lb_uuid,true,nil
}

func DeleteLb(db *gorm.DB,lb_uuid string) bool {
	//db,_ := libs.InitDB()
	if libs.JudgeLBWorkStatus(lb_uuid){
		if libs.RemoveLB(lb_uuid) {
			log.Print("LB instance "+ lb_uuid +" have remove successful" )
			if libs.DeleteLB(lb_uuid) {
				libs.UpdateSql(db, lb_uuid, 1)
				log.Print("LB "+ lb_uuid +" delete successful")
			} else {
				log.Print("LB "+ lb_uuid +" delete failed")
				return false
			}

		} else {
			log.Print("LB "+ lb_uuid +" removed failed")
			return false
		}
	} else {
		if libs.DeleteLB(lb_uuid) {
			libs.UpdateSql(db, lb_uuid, -1)
			log.Print("LB "+ lb_uuid +" delete successful")
		} else {
			log.Print("LB "+ lb_uuid +" delete failed")
			return false
		}
	}
	return true
}
