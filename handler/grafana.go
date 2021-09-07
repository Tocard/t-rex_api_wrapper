package handler

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"t-rex_wrapper/auth"
	"t-rex_wrapper/utils"
)

type UserStruct struct {
	User string `form:"user" json:"user" xml:"user"  binding:"required"`
}

func StartPage(c *gin.Context) {
	user := UserStruct{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password := auth.StartPage()
	values := map[string]interface{}{
		"name":     user.User,
		"email":    "user@graf.com",
		"login":    user.User,
		"password": password,
		"OrgId":    5}
	jsonData, err := json.Marshal(values)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	url := fmt.Sprintf("https://%s:%s@grafana.ether-source.fr/api/admin/users", utils.CfgGrafana.Username, utils.CfgGrafana.Password)
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode, resp.Status, string(body))

	values = map[string]interface{}{
		"password":  password,
		"roles":     "['role1', 'role2']",
		"full_name": user.User,
		"email":     "none@mail.com",
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	respEs, err := client.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	bodyEs, _ := ioutil.ReadAll(respEs.Body)
	fmt.Println(respEs.StatusCode, respEs.Status, string(bodyEs))

	if resp.StatusCode == 200 {
		message := fmt.Sprintf("Created user %s with password %s on datasaource %s-data", user.User, password, user.User)
		c.JSON(http.StatusOK, gin.H{"message": message})
	} else {
		c.JSON(http.StatusInternalServerError, string(body))
	}
}
