package routes

import (
	"github.com/Owariq/em-music/internal/db"
	lg "github.com/Owariq/em-music/internal/logger"
	"github.com/Owariq/em-music/internal/models"
	"github.com/gin-gonic/gin"
)

// GetLib - Get songs from library
//	@Summary		Get songs from library
//	@Description	Get songs from library by group and song name with pagination
//	@Tags			library
//	@Accept			json
//	@Produce		json
//	@Param			group	query		string	false	"Group name"
//	@Param			song	query		string	false	"Song name"
//	@Param			limit	query		int		false	"Limit of songs"
//	@Param			offset	query		int		false	"Offset of songs"
//	@Success		200		{array}		models.Song
//	@Failure		400		{object}	gin.H
//	@Failure		500		{object}	gin.H
//	@Router			/lib [get]
func GetLib(c *gin.Context) {
	lg.Log.Info("get lib")
	var queryParams models.QueryParams
	c.ShouldBindQuery(&queryParams)
	lg.Log.Debug("query params", queryParams)

	var lib []models.Song

	if queryParams.Limit == 0 {
		queryParams.Limit = -1
	}
	if queryParams.Offset == 0 {
		queryParams.Offset = -1
	}

	err := db.DB.Limit(queryParams.Limit).Offset(queryParams.Offset).Where(&models.Song{Group: queryParams.Group, Song: queryParams.Song}).Find(&lib)
	lg.Log.Debug("bd response", )
	if err.Error != nil {
		c.JSON(500, gin.H{"error": err.Error.Error()})
		lg.Log.Error(err.Error.Error())
		return
	}
	lg.Log.Debug("bd response", lib)

	c.JSON(200, lib)
}