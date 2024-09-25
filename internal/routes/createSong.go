package routes

import (
	"github.com/Owariq/em-music/internal/db"
	lg "github.com/Owariq/em-music/internal/logger"
	"github.com/Owariq/em-music/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateSong - Create a new song
//	@Summary		Create a new song
//	@Description	Create a new song
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			song	body		models.Song	true	"Song"
//	@Success		201		{object}	models.Song
//	@Failure		400		{object}	gin.H
//	@Failure		500		{object}	gin.H
//	@Router			/lib/song [post]
func CreateSong(c *gin.Context) {
	lg.Log.Info("create song")

	var song models.Song
	err := c.ShouldBindJSON(&song)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		lg.Log.Error(err.Error())
		return
	}

	if song.Group == "" || song.Song == "" {
		c.JSON(400, gin.H{"error": "group and song are required"})
		lg.Log.Error("group and song are required")
		return
	}

	lg.Log.Debug("song", song)

	err = db.DB.Create(&song).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	lg.Log.Error(err.Error())
		return
	}
	c.JSON(201, song)
}