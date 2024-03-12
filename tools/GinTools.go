package tools

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
)

//v 1.2

func CheckAcceptHeader(c *gin.Context) {
	if c.GetHeader("restype") == "xml" || c.GetHeader("Accept") == "application/xml" {
		c.Set("isXML", true)
	}
	if c.GetHeader("restype") == "xml2" {
		c.Set("isXMLGW", true)
	}

	c.Next()
}
func ConvertBooleanStrings() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonBody map[string]interface{}

		if err := c.ShouldBindJSON(&jsonBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		// recursive function to handle nested maps
		var convertBooleanStringsRecursive func(map[string]interface{})
		convertBooleanStringsRecursive = func(m map[string]interface{}) {
			for k, v := range m {
				switch value := v.(type) {
				case string:
					if value == "true" {
						m[k] = true
					} else if value == "false" {
						m[k] = false
					}
				case map[string]interface{}:
					convertBooleanStringsRecursive(value)
				}
			}
		}

		convertBooleanStringsRecursive(jsonBody)

		c.Set("jsonBody", jsonBody)

		c.Next()
	}
}

func Respond(c *gin.Context, status int, data interface{}) {
	if isXML, exists := c.Get("isXMLGW"); exists && isXML.(bool) {
		xmlData, err := xml.MarshalIndent(data, "", "  ")
		if err != nil {
			c.String(http.StatusInternalServerError, "XML generation failed")
			return
		}
		c.Data(status, "application/xml", xmlData)
	} else if isXML2, exists2 := c.Get("isXML"); exists2 && isXML2.(bool) {
		c.XML(status, data)
	} else {
		c.JSON(status, data)
	}
}

func GetIP(c *gin.Context) string {
	i1 := c.Request.Header.Get("X-Forwarded-For")
	i2 := c.Request.RemoteAddr
	if i1 == "" {
		return i2
	} else {
		return i1
	}
}
