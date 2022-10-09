package manage

import (
	"WebLog/data"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func Listener() {
	http.Handle("/", http.HandlerFunc(handler))
	err := http.ListenAndServe(fmt.Sprintf(":%d", data.Config.Listener), nil)
	data.ErrHandle(err)
}

func handler(res http.ResponseWriter, req *http.Request) {
	method := req.Method
	url := req.URL.String()
	headers := make(map[string]string)
	bodyData, _ := io.ReadAll(req.Body)
	body := string(bodyData)
	for k, v := range req.Header {
		headers[k] = v[0]
	}
	data.DataLock.Lock()
	for k, v := range data.Data {
		out, _, _ := v.Rule.Eval(map[string]interface{}{
			"method":  method,
			"url":     url,
			"headers": headers,
			"body":    body,
		})
		res, _ := out.ConvertToNative(reflect.TypeOf(true))
		if res.(bool) == true {
			v.Requested = true
			data.Data[k] = v
			break
		}
	}
	data.DataLock.Unlock()
	res.Write([]byte("ok"))
}
