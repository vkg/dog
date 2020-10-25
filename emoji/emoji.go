package emoji

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	emojiWidth  = 200
	emojiHeight = 200

	lineSpacing float64 = 1
)

var Colors = map[string]image.Image{
	"white":       White,
	"black":       Black,
	"transparent": Transparent,
	"opaque":      Opaque,
	"olive":       Olive,
	"yellow":      Yellow,
	"fuchsia":     Fuchsia,
	"silver":      Silver,
	"aqua":        Aqua,
	"lime":        Lime,
	"red":         Red,
	"gray":        Gray,
	"blue":        Blue,
	"green":       Green,
	"purple":      Purple,
	"navy":        Navy,
	"teal":        Teal,
	"maroon":      Maroon,
}

var (
	White       = image.NewUniform(color.White)
	Black       = image.NewUniform(color.Black)
	Transparent = image.NewUniform(color.Transparent)
	Opaque      = image.NewUniform(color.Opaque)

	Olive   = image.NewUniform(color.NRGBA{128, 128, 0, 0xff})
	Yellow  = image.NewUniform(color.NRGBA{255, 255, 0, 0xff})
	Fuchsia = image.NewUniform(color.NRGBA{255, 0, 255, 0xff})
	Silver  = image.NewUniform(color.NRGBA{192, 192, 192, 0xff})
	Aqua    = image.NewUniform(color.NRGBA{0, 255, 255, 0xff})
	Lime    = image.NewUniform(color.NRGBA{0, 255, 0, 0xff})
	Red     = image.NewUniform(color.NRGBA{255, 0, 0, 0xff})
	Gray    = image.NewUniform(color.NRGBA{128, 128, 128, 0xff})
	Blue    = image.NewUniform(color.NRGBA{0, 0, 255, 0xff})
	Green   = image.NewUniform(color.NRGBA{0, 128, 0, 0xff})
	Purple  = image.NewUniform(color.NRGBA{128, 0, 128, 0xff})
	Navy    = image.NewUniform(color.NRGBA{0, 0, 128, 0xff})
	Teal    = image.NewUniform(color.NRGBA{0, 128, 128, 0xff})
	Maroon  = image.NewUniform(color.NRGBA{128, 0, 0, 0xff})
)

type Config struct {
	Fg, Bg   image.Image
	Fontpath string
}

func Gen(text string, conf Config) (string, error) {
	if text == "" {
		return "", fmt.Errorf("empty text is passed")
	}

	if utf8.RuneCountInString(text) > 16 {
		return "", fmt.Errorf("emoji length must be less than 16")
	}

	texts := splitText(text)

	if conf.Fontpath == "" {
		conf.Fontpath = "./font/Koruri-Regular.ttf"
	}

	if conf.Fg == nil {
		conf.Fg = White
	}

	if conf.Bg == nil {
		conf.Bg = Transparent
	}

	fontBytes, err := ioutil.ReadFile(conf.Fontpath)
	if err != nil {
		return "", err
	}

	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return "", err
	}

	fontSize := calcFontSize(texts)

	rgba := image.NewRGBA(image.Rect(0, 0, emojiWidth, emojiHeight))
	draw.Draw(rgba, rgba.Bounds(), conf.Bg, image.Point{}, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(conf.Fg)
	c.SetHinting(font.HintingNone)

	pt := freetype.Pt(0, int(c.PointToFixed(fontSize)>>6))
	for _, s := range texts {
		_, err = c.DrawString(s, pt)
		if err != nil {
			return "", err
		}
		pt.Y += c.PointToFixed(fontSize * lineSpacing)
	}

	filepath := "out/" + strings.Join(texts, "") + ".png"
	outFile, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	b := bufio.NewWriter(outFile)

	if err = png.Encode(b, rgba); err != nil {
		return "", err
	}

	if err = b.Flush(); err != nil {
		return "", err
	}

	return filepath, nil
}

func splitText(text string) []string {
	length := utf8.RuneCountInString(text)

	split := func(nums ...int) []string {
		ret := make([]string, len(nums))
		for i := 0; i < len(nums); i++ {
			ret[i] = string([]rune(text)[0:nums[i]]) // head n chars
			text = string([]rune(text)[nums[i]:])    // cut head n chars
		}

		return ret
	}
	switch length {
	case 1:
		return split(1)
	case 2:
		return split(1, 1)
	case 3:
		return split(2, 1)
	case 4:
		return split(2, 2)
	case 5:
		return split(3, 2)
	case 6:
		return split(3, 2, 1)
	case 7:
		return split(3, 2, 2)
	case 8:
		return split(3, 3, 2)
	case 9:
		return split(3, 3, 3)
	case 10:
		return split(4, 2, 2, 2)
	case 11:
		return split(4, 3, 2, 2)
	case 12:
		return split(4, 3, 3, 2)
	case 13:
		return split(4, 3, 3, 3)
	case 14:
		return split(4, 4, 3, 3)
	case 15:
		return split(4, 4, 4, 3)
	case 16:
		return split(4, 4, 4, 4)
	default:
		return []string{}
	}
}

func calcFontSize(texts []string) float64 {
	length := len(texts)
	for _, text := range texts {
		if length < utf8.RuneCountInString(text) {
			length = utf8.RuneCountInString(text)
		}
	}

	return float64(emojiHeight/length) * 0.95
}
