package routes

import (
	"github.com/Owariq/em-music/internal/db"
	lg "github.com/Owariq/em-music/internal/logger"
	"github.com/Owariq/em-music/internal/models"
	"github.com/Owariq/em-music/internal/utils"
	"github.com/gin-gonic/gin"
)

// GetText - Get song text
//	@Summary		Get song text
//	@Description	Get song text by group and song names
//	@Tags			text
//	@Accept			json
//	@Produce		json
//	@Param			group	path		string	true	"Group name"
//	@Param			song	path		string	true	"Song name"
//	@Param			verse	query		int		false	"Verse number"
//	@Success		200		{object}	models.Song
//	@Failure		400		{object}	gin.H
//	@Failure		404		{object}	gin.H
//	@Failure		500		{object}	gin.H
//	@Router			/lib/song/text [get]
func GetText(c *gin.Context) {
	lg.Log.Info("get text")

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

	if song.Text == "" {
		c.JSON(404, gin.H{"error": "text not found"})
		lg.Log.Error("text not found")
		return
	}

	if queryParams.Verse != 0 {
		splitedVerse := utils.VerseSplit(song.Text)
		if queryParams.Verse > len(splitedVerse) {
			c.JSON(404, gin.H{"error": "verse not found"})
			lg.Log.Error("verse not found")
			return
		}
		c.JSON(200, splitedVerse[queryParams.Verse-1])
		return
	}
	lg.Log.Debug("bd response", song)
	c.JSON(200, song.Text)
}