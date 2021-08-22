package trex_wrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"t-rex_wrapper/logger"
	"t-rex_wrapper/utils"
)

// ManageFan interact with all t rex client listed in config file
func ManageFan(f string) {
	for i, s := range utils.Cfg.HostTarget {
		fmt.Println(i, s, f)
		values := map[string]string{"fan": f}
		jsonData, err := json.Marshal(values)
		logger.ErrorMessage(err)
		_, err = http.Post("http://"+s, "application/json",
			bytes.NewBuffer(jsonData))
		logger.WarnErrorMessage(err)
		//if err == nil {
		//	var res map[string]interface{}
		//	json.NewDecoder(resp.Body).Decode(&res)
		//	fmt.Println(res["json"])
		//}
	}
}
