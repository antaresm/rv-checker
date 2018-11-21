package handlers

import (
	"net/http"
	"fmt"
	"strings"
	"rv-check/dataProvider"
)

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var payload = "common_status 200\n"

	orm := dataProvider.OrmProvider
	checkpoints := orm.GetChecks()
	for _, checkpoint := range checkpoints {
		payload = payload + fmt.Sprintf("check_%s %d\n", strings.ToLower(checkpoint.Name), checkpoint.LastStatus)
	}

	payload = payload + fmt.Sprintf("%s %d\n", "check_finish", 200)


	w.Write([]byte(payload))
})
