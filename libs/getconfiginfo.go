package libs

import (
	"github.com/unknwon/goconfig"
	"log"
)

func getvalue(path string, key string) string {
	cfg, _ := goconfig.LoadConfigFile(path)
	value, _ := cfg.GetValue("info", key)
	return value
}
func main() {
	log.Print(getvalue("conf/app.conf", "lb_vm_zone_id"))
}
