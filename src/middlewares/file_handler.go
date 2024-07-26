package middleware

import "github.com/gin-gonic/gin"

func (m middleware) LogoHandler(reference string, storePath string) func(c *gin.Context) {
	return func(c *gin.Context) {
		file, err := c.FormFile(reference)
		if err == nil {
			dst := storePath + file.Filename

			c.Set("company_logo", "/"+dst)
			c.SaveUploadedFile(file, dst)
		}
		c.Next()
	}
}
