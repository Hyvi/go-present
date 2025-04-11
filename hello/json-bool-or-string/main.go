package main

// 解决bool 类型兼容“true” “false”的情况
// 参考 https://stackoverflow.com/questions/30856454/how-to-unmarshall-both-0-and-false-as-bool-from-json
import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ConvertibleBoolean 变体bool
type ConvertibleBoolean bool

// UnmarshalJSON 兼容“true”和1
func (bit *ConvertibleBoolean) UnmarshalJSON(data []byte) error {
	asString := strings.ToLower(strings.Trim(string(data), `"`))
	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" {
		*bit = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}

// T 测试
type T struct {
	Name string
	Eng  *ConvertibleBoolean
}

func main() {
	r := gin.Default()
	r.POST("/jsonbool", func(c *gin.Context) {
		var t T
		if err := c.BindJSON(&t); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	if err := r.Run(); err != nil {
		return
	}
}
