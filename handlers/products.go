package handlers

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"go-search-db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PopulateProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		for i := 0; i < 50; i++ {
			db.Create(&models.Product{
				Title:       fmt.Sprintf("Title %d", i),
				Description: fmt.Sprintf("Description %d", i),
				Image:       fmt.Sprintf("http://google.com/200/200?%s", fmt.Sprintf("UUIDDigit-%d", i)),
				Price:       rand.Intn(90) + 10,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}

func GetFrontendProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product

		db.Find(&products)

		c.JSON(http.StatusOK, products)
	}
}

func GetBackendProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product

		sql := "SELECT * FROM products"

		if s := c.Query("s"); s != "" {
			sql = fmt.Sprintf("%s WHERE title LIKE '%%%s%%' OR description LIKE '%%%s%%'", sql, s, s)
		}

		if sort := c.Query("sort"); sort != "" {
			sql = fmt.Sprintf("%s ORDER BY price %s", sql, sort)
		}

		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		perPage := 5
		var total int64

		db.Raw(sql).Count(&total)

		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)

		db.Raw(sql).Scan(&products)

		c.JSON(http.StatusOK, gin.H{
			"data":      products,
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total / int64(perPage))),
		})
	}
}