package routes

import (
	"github.com/Owariq/em-music/internal/db"
	lg "github.com/Owariq/em-music/internal/logger"
	"github.com/Owariq/em-music/internal/models"
	"github.com/gin-gonic/gin"
)

// DeleteSong - Delete song
//	@Summary		Delete song
//	@Description	Delete a song
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			song	body		models.Song	true	"song info"
//	@Success		200		{object}	gin.H
//	@Failure		400		{object}	gin.H
//	@Failure		500		{object}	gin.H
//	@Router			/lib/song [delete]
func DeleteSong(c *gin.Context) {
	lg.Log.Info("delete song")

	var song models.Song
	err := c.ShouldBindJSON(&song)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		lg.Log.Error(err.Error())
		return
	}
	lg.Log.Debug("song", song)

	if song.Group == "" || song.Song == "" {
		c.JSON(400, gin.H{"error": "group and song are required"})
		lg.Log.Error("group and song are required")
		return
	}

	err = db.DB.Where(&models.Song{Group: song.Group, Song: song.Song}).First(&song).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		lg.Log.Error(err.Error())
		return
	}
	lg.Log.Debug("bd response", song)

	err = db.DB.Where(&models.Song{Group: song.Group, Song: song.Song}).Delete(&song).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		lg.Log.Error(err.Error())
		return
	}
	c.JSON(200, song.ID)
}