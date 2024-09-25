package routes

import (
	"github.com/Owariq/em-music/internal/db"
	lg "github.com/Owariq/em-music/internal/logger"
	"github.com/Owariq/em-music/internal/models"
	"github.com/gin-gonic/gin"
)

// UpdateSong - update song by id
//	@Summary		Update song by id
//	@Description	Update song by id
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param			updateSong	body		models.Song	true	"song"
//	@Success		200			{object}	models.Song
//	@Failure		400			{object}	map[string]interface{}
//	@Failure		500			{object}	map[string]interface{}
func UpdateSong(c *gin.Context) {
	lg.Log.Info("update song")

	var updateSong models.Song
	var song models.Song

	err := c.ShouldBindJSON(&updateSong)
	
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		lg.Log.Error(err.Error())
		return
	}


	lg.Log.Debug("song", updateSong)
	if updateSong.ID == 0 {
		c.JSON(400, gin.H{"error": "id is required"})
		lg.Log.Error("id is required")
		return
	}

	if updateSong.Group == "" || updateSong.Song == "" {
		c.JSON(400, gin.H{"error": "group and song are required"})
		lg.Log.Error("group and song are required")
		return
	}

	err = db.DB.First(&song, updateSong.ID).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		lg.Log.Error(err.Error())
		return
	}
	lg.Log.Debug("bd response", song)


	song = updateSong

	err = db.DB.Save(&song).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		lg.Log.Error(err.Error())
		return
	}

	
	c.JSON(200, song.ID)
}