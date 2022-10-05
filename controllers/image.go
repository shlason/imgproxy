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

func GetImagesByQS(c *gin.Context) {
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

	if errW != nil || errH != nil || errB != nil {
		fmt.Println(errW, errH, errB)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errW,
		})
		return
	}

	// https://www.nasa.gov/sites/default/files/thumbnails/image/pia22228.jpg
	response, err := http.Get(urlQs)

	fmt.Println("Get image")

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

	fmt.Println("Read Response")

	newImage := bimg.NewImage(bodyBytes)

	imageSize, err := newImage.Size()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	fmt.Println(imageSize, resizeQs, blur)

	options := bimg.Options{
		Width:       width,
		Height:      height,
		Crop:        true,
		Gravity:     bimg.GravitySmart,
		Quality:     100,
		Lossless:    true,
		Compression: 0,
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
