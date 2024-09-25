package routes

import (
	"github.com/Owariq/em-music/internal/db"
	lg "github.com/Owariq/em-music/internal/logger"
	"github.com/Owariq/em-music/internal/models"
	"github.com/gin-gonic/gin"
)

// GetSong - Get song by group and song
//
//	@Summary	Get song by group and song
//	@Tags		song
//	@Accept		json
//	@Produce	json
//	@Param		group	query		string	true	"Group"
//	@Param		song	query		string	true	"Song"
//	@Success	200		{object}	models.Song
//	@Failure	400		{object}	gin.H
//	@Failure	500		{object}	gin.H
//	@Router		/lib/song [get]
func GetSong(c *gin.Context) {
	lg.Log.Info("get song")

	var song models.Song
	var queryParams models.QueryParams

	err := c.ShouldBindQuery(&queryParams)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		lg.Log.Error(err.Error())
		return
	}

	if queryParams.Group == "" || queryParams.Song == "" {
		c.JSON(400, gin.H{"error": "group and song are required"})
		lg.Log.Error("group and song are required")
		return
	}
	
	lg.Log.Debug("query params", queryParams)

	err = db.DB.Where(&models.Song{Group: queryParams.Group, Song: queryParams.Song}).First(&song).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		lg.Log.Error(err.Error())
		return
	}
	lg.Log.Debug("bd response", song)

	c.JSON(200, song)
}