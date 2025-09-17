package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"video-collector/server/models"
)

// GetVideos godoc
// @Summary Get all videos
// @Description Get all videos with optional pagination and product ID filter
// @Tags videos
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Param product_id query int false "Product ID"
// @Success 200 {array} models.Video
// @Router /videos [get]
func (h *Handler) GetVideos(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	productId := c.Query("product_id")

	var videos []models.Video
	var total int64

	query := h.DB.Model(&models.Video{})
	if productId != "" {
		query = query.Where("product_id = ?", productId)
	}

	query.Count(&total)

	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Preload("Product").Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve videos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total, "data": videos})
}

// CreateVideo godoc
// @Summary Create a new video
// @Description Create a new video with the given URL and product ID
// @Tags videos
// @Accept json
// @Produce json
// @Param video body models.Video true "Video object"
// @Success 200 {object} models.Video
// @Router /videos [post]
func (h *Handler) CreateVideo(c *gin.Context) {
	var video models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create video"})
		return
	}
	c.JSON(http.StatusOK, video)
}

// UpdateVideo godoc
// @Summary Update a video
// @Description Update a video by ID
// @Tags videos
// @Accept json
// @Produce json
// @Param id path int true "Video ID"
// @Param video body models.Video true "Video object"
// @Success 200 {object} models.Video
// @Router /videos/{id} [put]
func (h *Handler) UpdateVideo(c *gin.Context) {
	id := c.Param("id")
	var video models.Video
	if err := h.DB.First(&video, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Save(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update video"})
		return
	}
	c.JSON(http.StatusOK, video)
}

// DeleteVideo godoc
// @Summary Delete a video
// @Description Delete a video by ID
// @Tags videos
// @Param id path int true "Video ID"
// @Success 204 "No Content"
// @Router /videos/{id} [delete]
func (h *Handler) DeleteVideo(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Video{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete video"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
