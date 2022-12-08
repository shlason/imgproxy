package controllers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/h2non/bimg"
)

var acceptResize = map[string]string{
	"fit":  "fit",
	"fill": "fill",
}

// GetImagesByQS godoc
// @Summary      抓取圖片
// @Description  抓取由指定參數所處理過後的圖片
// @Tags         image
// @Accept       json
// @Produce      image/*
// @Param        url     query      string  true  "Image URL"
// @Param        width   query      string  true  "Desire Width"
// @Param        height   query      string  true  "Desire height"
// @Param        resize   query      string  true  "fit or fill"
// @Param        blur   query      string  true  "Desire blur"
// @Router       /image [get]
// TODO: 需要整理一下
func GetImagesByQS(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=604800, immutable")

	urlQs := c.Query("url")
	widthQs := c.Query("width")
	heightQs := c.Query("height")
	resizeQs := c.Query("resize")
	blurQs := c.Query("blur")

	_, err := url.ParseRequestURI(urlQs)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid image URL",
		})
		return
	}

	width, errW := strconv.Atoi(widthQs)
	height, errH := strconv.Atoi(heightQs)
	blur, errB := strconv.Atoi(blurQs)

	fmt.Println("Convert")

	if errW != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errW,
		})
		return
	}
	if errH != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errH,
		})
		return
	}
	if errB != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errB,
		})
		return
	}

	response, err := http.Get(urlQs)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	defer response.Body.Close()

	bodyBytes, e := io.ReadAll(response.Body)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	newImage := bimg.NewImage(bodyBytes)

	imageSize, err := newImage.Size()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	options := bimg.Options{
		GaussianBlur: bimg.GaussianBlur{
			Sigma:   float64(blur),
			MinAmpl: float64(blur),
		},
		Gravity:     bimg.GravitySmart,
		Quality:     100,
		Lossless:    true,
		Compression: 0,
	}

	if resizeQs == acceptResize["fill"] {

		options.Width = width
		options.Height = height
		options.Crop = true
	}

	if resizeQs == acceptResize["fit"] {
		if width-imageSize.Width < height-imageSize.Height {
			options.Width = width
		} else {
			options.Height = height
		}
	}

	result, err := newImage.Process(options)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.Data(http.StatusOK, fmt.Sprintf("image/%s", newImage.Type()), result)
}
