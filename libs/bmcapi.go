package libs

import (
	"encoding/json"
	"log"
	"time"
)

var zone_id string = getvalue("conf/app.conf", "lb_vm_zone_id")
var seg_id string = getvalue("conf/app.conf", "lb_vm_secgroupid")
var rs1_ip string = getvalue("conf/app.conf", "lb_vm_rs1")
//var rs2_ip string = getvalue("conf/app.conf", "lb_vm_rs2")
var pub_subnetwork_id string = getvalue("conf/app.conf", "pub_subnetwork_id")
var pub_network_id string = getvalue("conf/app.conf", "pub_network_id")
var pri_subnetwork_id string = getvalue("conf/app.conf", "pri_subnetwork_id")
var pri_network_id string = getvalue("conf/app.conf", "pri_network_id")

type Responseinfo struct {
	Resultcode int    `json:"code"`
	Resultmsg  string `json:"msg"`
	Success bool `json:"success"`
}

type ListenerRES struct {
	ListenerUUID string `json:"listenerUuid"`
	Msg string `json:"msg"`
	Success bool `json:"success"`
}

type BackendRES struct {
	BackendsUUID string `json:"backendsUuid"`
	Msg string `json:"msg"`
	Success bool `json:"success"`
}


type ListLB struct {
	Data []Data `json:"data"`
	Success bool `json:"success"`
	TotalCount int `json:"totalCount"`
}
type Data struct {
	ID int `json:"Id"`
	Name string `json:"Name"`
	UserID int `json:"UserId"`
	ZoneID int `json:"ZoneId"`
	Type string `json:"Type"`
	Vip []string `json:"Vip"`
	RateLimit int `json:"RateLimit"`
	UUID string `json:"Uuid"`
	Vxnet string `json:"Vxnet"`
	MasterPrivateIP string `json:"MasterPrivateIp"`
	SlavePrivateIP string `json:"SlavePrivateIp"`
	IPNum int `json:"IpNum"`
	MaxConnect int `json:"MaxConnect"`
	Cps int `json:"Cps"`
	QPS int `json:"Qps"`
	UserLanID int `json:"UserLanId"`
	IsWorking int `json:"IsWorking"`
	CreateTime string `json:"CreateTime"`
	UpdateTime string `json:"UpdateTime"`
	AddIPNum int `json:"AddIpNum"`
	UserName string `json:"UserName"`
	VMLBStatus bool `json:"VmLBStatus"`
	Platform string `json:"Platform,omitempty"`
	Vendor string `json:"Vendor,omitempty"`
}

type ListLBZoneRES struct {
	Data []LBZoneData `json:"data"`
	Success bool `json:"success"`
}

type LBZoneData struct {
	Id int `json:"Id"`
	Name string `json:"Name"`
}



func CreateLB() string {
	body :="{\"Name\":\"test-lb-NAT-111\",\"ZoneId\":"+ zone_id +",\"Type\":\"NLB\",\"IpNum\":1,\"RateLimit\":1,\"Vlan\":100,\"Vxnet\":\"10.10.10.0/24\",\"WanLanId\":3988,\"UserLanId\":6539,\"MaxConnect\":200000,\"Cps\":200000,\"Qps\":200000,\"Vendor\":\"estack\",\"Platform\":\"vm\"," +
		"\"SecGroupId\":\""+ seg_id +"\"," +
		"\"EstackNetworkPorts\":[{\"Bandwidth\":11," +
		"\"SecurityGroupId\":\""+ seg_id +"\"," +
		"\"SubnetId\":\""+ pub_subnetwork_id +"\"," +
		"\"NetworkId\":\""+ pub_network_id +"\"," +
		"\"Wan\":true}," +
		"{\"SecurityGroupId\":\""+ seg_id +"\"," +
		"\"NetworkId\":\""+ pri_network_id +"\"," +
		"\"SubnetId\":\""+ pri_subnetwork_id +"\"," +
		"\"Wan\":false}]}"
	var res string = Post("/v2/api/pub/add_lb",body)
	var responseinfo Responseinfo
	err := json.Unmarshal([]byte(res), &responseinfo)
	if err != nil {
		log.Print(err.Error())
	}
	return responseinfo.Resultmsg
}

func CreateListener(lbuuid string) string {
	body := "{\"Uuid\":\""+ lbuuid +"\"," +
		"\"Name\":\"NLB-NAT-test\"," +
		"\"ListenerProto\":\"TCP\",\"" +
		"ListenerPort\":\"80\"," +
		"\"BackendProto\":\"TCP\",\"Algorithm\":\"rr\",\"Kind\":\"NAT\",\"Cert_select\":0,\"HealthyCheck\":{\"Enable\":true,\"CheckType\":\"HTTP\",\"ConnectTimeout\":7,\"Retry\":5,\"DelayBeforeRetry\":3,\"DelayLoop\":6,\"ConnectPort\":80,\"HTTPProtoVersion\":\"HTTP/1.0\",\"Path\":\"/\",\"StatusCode\":200},\"Notify\":{\"Enable\":true,\"ApiAddress\":\"https://open.feishu.cn/open-apis/bot/v2/hook/b4c94-bcdc-d1ebcc43d157\"}}"
	var res string = Post("/v2/api/pub/add_listener",body)
	var responseinfo ListenerRES
	err := json.Unmarshal([]byte(res), &responseinfo)
	if err != nil {
		log.Print(err.Error())
	}
	return responseinfo.ListenerUUID
}

func CreateBackend(listeneruuid string) string {
	body := "{\"Name\":\"RS1\",\"IpAddr\":\""+ rs1_ip +"\",\"Port\":\"80,3000\",\"Weight\":80," +
		"\"ListenerUuid\":\""+ listeneruuid +"\"," +
		"\"Platform\":\"vm\"}"
	var res string = Post("/v2/api/pub/add_backends",body)
	var responseinfo BackendRES
	err := json.Unmarshal([]byte(res), &responseinfo)
	if err != nil {
		log.Print(err.Error())
	}
	return responseinfo.BackendsUUID
}

func ApplyLB(lbuuid string) bool {
	body := "{\"Uuid\":[\""+ lbuuid +"\"],\"Platform\":\"vm\"}\n"
	var res string = Post("/v2/api/pub/apply_lb",body)
	var responseinfo Responseinfo
	err := json.Unmarshal([]byte(res), &responseinfo)
	if err != nil {
		log.Print(err.Error())
	}
	return responseinfo.Success
}

func RemoveLB(lbuuid string) bool {
	body := "{\"Uuid\":[\""+ lbuuid +"\"],\"Platform\":\"vm\"}\n"
	var res string = Post("/v2/api/pub/remove_lb",body)
	var responseinfo Responseinfo
	err := json.Unmarshal([]byte(res), &responseinfo)
	if err != nil {
		log.Print(err.Error())
	}
	return responseinfo.Success
}

func DeleteLB(lbuuid string) bool {
	body := "{\"Uuid\":[\""+ lbuuid +"\"],\"Platform\":\"vm\"}\n"
	var res string = Post("/v2/api/pub/del_lb",body)
	var responseinfo Responseinfo
	err := json.Unmarshal([]byte(res), &responseinfo)
	if err != nil {
		log.Print(err.Error())
	}
	return responseinfo.Success
}

func GetLBList() []Data {
	var res string = Post("/v2/api/pub/list_lb?pageSize=200","")
	var responseinfo ListLB
	err := json.Unmarshal([]byte(res), &responseinfo)
	if err != nil {
		log.Print(err.Error())
	}
	//msg,_ := json.Marshal(responseinfo.Data)
	return responseinfo.Data
}

func JudgeLBStatus(lbuuid string) bool {
	var datas []Data = GetLBList()
	for i := range datas {
		if lbuuid == datas[i].UUID {
			log.Print("LB ",lbuuid," vip list is ",datas[i].Vip)
			if len(datas[i].Vip) > 0 && datas[i].VMLBStatus == true {
				return true
			}

		}
	}
	return false
}

func JudgeLBWorkStatus(lbuuid string) bool {
	var datas []Data = GetLBList()
	for i := range datas {
		if lbuuid == datas[i].UUID {
			if datas[i].IsWorking == 1 {
				return true
			}

		}
	}
	return false
}
func WaitLBNoramal(lbuuid string) bool {
	for i := 0; i <= 10; i++ {
		if JudgeLBStatus(lbuuid) {
			return true
		}
		time.Sleep(60 * time.Second)
	}
	return false
}


func GetListLbZone() []LBZoneData {
	var res string = Post("/v2/api/pub/list_lb_zone","")
	var responseinfo ListLBZoneRES
	err := json.Unmarshal([]byte(res), &responseinfo)
	if err != nil {
		log.Print(err.Error())
	}
	//msg,_ := json.Marshal(responseinfo.Data)
	return responseinfo.Data
}

func GetZoneName(zoneid string) string {
	var datas []LBZoneData = GetListLbZone()
	log.Print(datas)
	lb_zone_id := StrToInt(zone_id)
	log.Print(lb_zone_id)
	for i := range datas {
		if lb_zone_id == datas[i].Id {
			log.Print(datas[i])
			{
				return datas[i].Name
			}
		}
	}
	return ""
}