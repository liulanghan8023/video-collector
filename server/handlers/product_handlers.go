package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"video-collector/server/models"
)

// GetProducts godoc
// @Summary Get all products
// @Description Get all products with optional pagination
// @Tags products
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} models.Product
// @Router /products [get]
func (h *Handler) GetProducts(c *gin.Context) {
	if c.Query("page") == "" && c.Query("pageSize") == "" {
		var products []models.Product
		if err := h.DB.Order("created_at DESC").Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
			return
		}
		c.JSON(http.StatusOK, products)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	status := c.Query("status")

	var products []models.Product
	var total int64

	db := h.DB.Model(&models.Product{})

	if status != "" {
		db = db.Where("status = ?", status)
	}

	db.Count(&total)

	offset := (page - 1) * pageSize
	if err := db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	for i := range products {
		var count int64
		h.DB.Model(&models.Video{}).Where("product_id = ?", products[i].ID).Count(&count)
		products[i].VideoCount = count
	}

	c.JSON(http.StatusOK, gin.H{"total": total, "data": products})
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the given name and URL
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product object"
// @Success 200 {object} models.Product
// @Router /products [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.Status = "video_collecting"
	if err := h.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// UpdateProductStatus godoc
// @Summary Update a product's status
// @Description Update a product's status by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param status body object{status=string} true "Status object"
// @Success 200 {object} models.Product
// @Router /products/{id}/status [put]
func (h *Handler) UpdateProductStatus(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var body struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Model(&product).Update("status", body.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product status"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product object"
// @Success 200 {object} models.Product
// @Router /products/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Router /products/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// AddVideosToProduct godoc
// @Summary Add videos to a product
// @Description Add multiple videos to a product by product ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param urls body []string true "List of video URLs"
// @Success 200 {object} gin.H{"message": "Videos added successfully"}
// @Router /products/{id}/videos [post]
func (h *Handler) AddVideosToProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var body struct {
		URLs []string `json:"urls"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var videos []models.Video
	for _, url := range body.URLs {
		videos = append(videos, models.Video{ProductID: uint(productID), URL: url, Status: "video_collecting"})
	}

	if err := h.DB.Create(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create videos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Videos added successfully"})
}
