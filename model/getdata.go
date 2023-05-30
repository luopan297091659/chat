package model

import (
	"chat/libs"
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
)

func LbVerfyres(c *gin.Context) { //不同于net/http库的路由函数，gin.Context封装了request和response
	db,_ := libs.InitDB()
	datas := libs.QueryAllZoneVerify(db)
	datasJSON, _ := json.Marshal(datas)
	c.String(http.StatusOK, string(datasJSON))
}

func SpecialZoneLbVerfy(c *gin.Context) { //不同于net/http库的路由函数，gin.Context封装了request和response
	db,_ := libs.InitDB()
	datas := libs.QueryZoneVerify(db)
	datasJSON, _ := json.Marshal(datas)
	c.String(http.StatusOK, string(datasJSON))
}

func GetAllZoneLbNewVerfy(c *gin.Context) { //不同于net/http库的路由函数，gin.Context封装了request和response
	db,_ := libs.InitDB()
	datas := libs.QueryNewZoneVerify(db)
	datasJSON, _ := json.Marshal(datas)
	c.String(http.StatusOK, string(datasJSON))
}
