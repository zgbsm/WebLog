package manage

import (
	"WebLog/data"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/cel-go/cel"
	"github.com/google/uuid"
	"time"
)

func Create(c *gin.Context) {
	cel_code := c.PostForm("rule")
	u := uuid.NewString()
	code := md5.Sum([]byte(u))
	codeStr := fmt.Sprintf("%x", code)
	env, err := cel.NewEnv(
		cel.Variable("method", cel.StringType),
		cel.Variable("url", cel.StringType),
		cel.Variable("headers", cel.MapType(cel.StringType, cel.StringType)),
		cel.Variable("body", cel.StringType),
	)
	data.ErrHandle(err)
	ast, issue := env.Compile(cel_code)
	data.ErrHandle(issue.Err())
	prg, err := env.Program(ast)
	info := data.Info{Start: time.Now(), Rule: prg, Requested: false}
	data.DataLock.Lock()
	data.Data[codeStr] = info
	data.DataLock.Unlock()
	c.JSON(200, gin.H{
		"success": true,
		"message": codeStr,
	})
}

func Get(c *gin.Context) {
	code := c.Query("c")
	data.DataLock.Lock()
	info, ok := data.Data[code]
	data.DataLock.Unlock()
	if ok {
		c.JSON(200, gin.H{
			"success": true,
			"message": info,
		})
	} else {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Invalid code.",
		})
	}
}

func Clean() {
	for {
		time.Sleep(time.Minute * 5)
		for k, v := range data.Data {
			if v.Start.Add(time.Minute * 5).Before(time.Now()) {
				delete(data.Data, k)
			}
		}
	}
}
