package go_sprite_split

import (
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
	_ "image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Process(inputImgPath, outputFolderPath, mode string) bool {
	file, err := os.Open(inputImgPath)
	if err != nil {
		println("file can not be found! abort.")
		return false
	}
	defer file.Close()
	img, msg, err := image.Decode(file)
	if err != nil {
		println("file can not be decode, abort.", msg)
		return false
	}
	sprite := make(map[string]interface{})

	bounds := img.Bounds()
	mask := make([][]bool, bounds.Dx())
	for i := range mask {
		mask[i] = make([]bool, bounds.Dy())
	}

	counter := 1
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			if _, _, _, a := img.At(x, y).RGBA(); a > 0 && !mask[x][y] {
				minX, minY, maxX, maxY := dfs(x, y, img, mask)

				if mode == "split" {
					rect := image.Rect(minX, minY, maxX+1, maxY+1)
					subImg := image.NewRGBA(rect)
					draw.Draw(subImg, rect, img, rect.Min, draw.Src)
					savaPath := fmt.Sprintf("%s/%d.png", outputFolderPath, counter)
					fmt.Printf("saving %s \n", savaPath)
					imaging.Save(subImg, savaPath)
				} else if mode == "sprite" {
					sprite[strconv.Itoa(counter)] = map[string]int{
						"x1": minX,
						"y1": minY,
						"x2": maxX,
						"y2": maxY,
						"w":  maxX - minX + 1,
						"h":  maxY - minY + 1,
					}
				}
				counter++
			}
		}
	}
	fmt.Printf("number of sprite: %d\n", counter)
	filename := filepath.Base(inputImgPath)
	extension := filepath.Ext(filename)
	newFilename := strings.TrimSuffix(filename, extension) + ".json"
	outPath := fmt.Sprintf("%s/%s", outputFolderPath, newFilename)
	outFile, _ := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(sprite)
	if err != nil {
		fmt.Println("encode and write output: fail", err)
		return false
	}
	fmt.Println("encode and write output: success")
	return true
}

func dfs(x, y int, img image.Image, mask [][]bool) (minX, minY, maxX, maxY int) {
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	stack := [][]int{{x, y}}
	mask[x][y] = true

	minX, minY = x, y
	maxX, maxY = x, y

	for len(stack) > 0 {
		x, y = stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]

			_, _, _, a := img.At(x, y).RGBA()
			if nx >= 0 && nx < len(mask) && ny >= 0 && ny < len(mask[0]) &&
				!mask[nx][ny] && a > 0 {

				mask[nx][ny] = true
				stack = append(stack, []int{nx, ny})

				if nx < minX {
					minX = nx
				}
				if nx > maxX {
					maxX = nx
				}
				if ny < minY {
					minY = ny
				}
				if ny > maxY {
					maxY = ny
				}
			}
		}
	}

	return
}
