package api

import (
	"net/http"
	"unico/internal/unico/models"
	"unico/internal/unico/services"
	"unico/pkg/rest"

	"github.com/gin-gonic/gin"
)

// GetMarket godoc
// @Summary Get existing market
// @Description Get all the existing market
// @Accept  json
// @Produce  json
// @Param  record path string true "Market Record"
// @Success 200  {object} entities.Market
// @Failure 404 {object} rest.Response
// @Failure 400 {object} rest.Response
// @Failure 500 {object} rest.Response
// @Router /market/:id [get]
func (ch *controlerHandler) GetMarket(c *gin.Context) {
	market, err := ch.db.GetMarketByRecord(c.Param("record"))
	if err != nil {
		log.Error(err)
		rest.ResponseBadRequest(c, err)
		return
	}
	if market == nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, market)
}

// GetAllMarket godoc
// @Summary List existing market
// @Description Get all the existing market
// @Accept  json
// @Produce  json
// @Param name query string false "name"
// @Param zone query string false "zone"
// @Param district query string false "district"
// @Param neighborhood query string false "neighborhood"
// @Success 200 {array} entities.Market
// @Failure 500 {object} rest.Response
// @Router /market [get]
func (ch *controlerHandler) GetSearchMarket(c *gin.Context) {
	search := new(models.MarketSearchRequest)
	err := c.BindQuery(&search)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}
	markets, err := ch.db.GetAllMarket(search)
	if err != nil {
		log.Error(err)
		rest.ResponseBadRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, markets)
}

// CreateMarket godoc
// @Summary Create Market
// @Description Create a new Market
// @Accept  json
// @Produce  json
// @Param market body models.MarketCreateRequest true "Create Market"
// @Success 200=1 {object} entities.Market
// @Failure 400 {object} rest.Response
// @Failure 500 {object} rest.Response
// @Router /market [post]
func (ch *controlerHandler) CreateMarket(c *gin.Context) {
	marketreq := new(models.MarketCreateRequest)
	err := c.BindJSON(&marketreq)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}
	// Validade Request
	err = ch.validate.Struct(marketreq)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}

	// Create New Market
	market, err := services.CreateMarket(marketreq, ch.db)
	if err != nil {
		log.WithField("Name", market.Name).
			WithField("ID", market.ID).
			WithField("Message", err).
			Error("Erro trye Market Create")
		rest.ResponseBadRequest(c, err)
		return
	}

	log.WithField("Name", market.Name).
		WithField("ID", market.ID).
		Info("Market Created")
	c.JSON(http.StatusCreated, market)
}

// UpdateMarket godoc
// @Summary Create Market
// @Description Create a new Market
// @Accept  json
// @Produce  json
// @Param  record path string true "Market Record"
// @Param market body models.MarketUpdateRequest true "Update Market"
// @Success 200
// @Failure 500 {object} rest.Response
// @Router /market/{id} [put]
func (ch *controlerHandler) UpdateMarket(c *gin.Context) {
	// Check Request
	marketreq := new(models.MarketUpdateRequest)
	err := c.BindJSON(&marketreq)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}
	// Validade Request
	err = ch.validate.Struct(marketreq)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}

	// Get Market
	market, err := ch.db.GetMarketByRecord(c.Param("record"))
	if err != nil {
		log.Error(err)
		rest.ResponseBadRequest(c, err)
		return
	}
	if market == nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// UpdateMarket
	err = services.UpdateMarket(market, marketreq, ch.db)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}

	log.WithField("Name", market.Name).
		WithField("ID", market.ID).
		Info("Market Updated")
	c.JSON(http.StatusOK, gin.H{})
}

// RemoveMarket godoc
// @Summary Remove Market
// @Description Remove Market by ID
// @Accept  json
// @Produce  json
// @Param  record path string true "Market Record"
// @Success 200
// @Failure 500 {object} rest.Response
// @Router /market/{id} [delete]
func (ch *controlerHandler) RemoveMarket(c *gin.Context) {
	market, err := ch.db.GetMarketByRecord(c.Param("record"))
	if err != nil {
		log.Error(err)
		rest.ResponseBadRequest(c, err)
		return
	}
	if market == nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// Remove Record
	err = ch.db.RemoveMarket(market.Record)
	if err != nil {
		log.Error(err)
		rest.ResponseBadRequest(c, err)
		return
	}
	log.WithField("ID", market.ID).
		Info("Market Removed")
	c.JSON(http.StatusOK, gin.H{})
}
