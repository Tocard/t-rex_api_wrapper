package trex_wrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"t-rex_wrapper/utils"
)

func SetFanPercent(f float64) {
	for i, s := range utils.Cfg.HostTarget {
		fmt.Println(i, s, f)
		values := map[string]string{"fan": fmt.Sprint(f)}
		jsonData, err := json.Marshal(values)
		utils.Check(err)
		_, err = http.Post("http://"+s, "application/json",
			bytes.NewBuffer(jsonData))
		utils.Warn(err.Error())
		//if err == nil {
		//	var res map[string]interface{}
		//	json.NewDecoder(resp.Body).Decode(&res)
		//	fmt.Println(res["json"])
		//}
	}
}
