package integrations

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func CleanDatabase()error{
	resp, err := http.Post("https://portal.vicuesoft.com/QNSvyxD[DELETED]Ruaa043VG","",nil)
	if err != nil {
		return err
	}
	if resp.StatusCode!=200{
		return errors.New("Сервер не принял запрос")
	}
	return nil
}
func DeleteKey(arg string)error{
	r := struct {
		Key string `json:"key"`
	}{
		Key: arg,
	}
	j, _ := json.Marshal(r)
	reader := bytes.NewReader(j)
	resp, _ := http.Post("https://portal.vicuesoft.com/CToO[DELETED]]UENfe", "application/json", reader)
	if resp.StatusCode!=200{
		return errors.New("Сервер не принял запрос")
	}
	return nil
}

func RenewDroplet(arg string){
	r := struct{
		IP string
	}{
		IP: arg,
	}
	j,_ := json.Marshal(r)
	reader := bytes.NewBuffer(j)
	http.Post("https://portal.vicuesoft.com/PjLeq[DELETED]Df7c5SREugaAdzF", "application/json", reader)

}