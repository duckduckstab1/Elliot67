package main

import (
	"hash/fnv"
	"image/color"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
)

func main() {
	stats := getStats()
	hash := GetMD5Hash(stats)
	rand.Seed(int64(hash - 9223372036854775808))
	generateImage()
}

func getStats() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/users/elliot67/events/public", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("accept", "application/vnd.github.v3+json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func GetMD5Hash(text string) uint64 {
	hasher := fnv.New64()
	hasher.Write([]byte(text))
	return hasher.Sum64()
}

func generateImage() {
	c := generativeart.NewCanva(800, 350)
	setCanvaOptions(c, OptionsOverride{
		Background: common.Black,
		LineColor:  color.RGBA{165, 174, 217, 255},
	})
	c.FillBackground()
	c.Draw(arts.NewBlackHole(rand.Intn(800-100)+100, 900, 0.005+rand.Float64()*(.04-0.005)))
	c.ToPNG("../.repo/profilehash.png")
}

type OptionsOverride struct {
	Background color.RGBA
	LineColor  color.RGBA
}

func setCanvaOptions(c *generativeart.Canva, options OptionsOverride) {
	c.SetBackground(options.Background)
	c.SetLineColor(options.LineColor)
}
