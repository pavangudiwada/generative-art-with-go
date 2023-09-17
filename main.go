package main

import (
	"math"
	"math/rand"

	"github.com/fogleman/gg"
)

const W = 1920.0
const H = 1080.0
const step = 20.0

var colorPalettes = map[string][]string{
	"Palette1": {"2B2E4A", "E84545", "903749", "53354A"},
	"Palette2": {"2e4252", "738e85", "bebba1", "e5dab9"},
	"Palette3": {"0f0f1b", "565a75", "c6b7be", "fafbf6"},
	"Palette4": {"46425e", "5b768d", "d17c7c", "f6c6a8"},
	"Palette5": {"090110", "1a1e25", "3d3e32", "6f653d"},
	"Palette6": {"99B898", "FECEAB", "FF847C", "E84A5F"},
}

func draw_line() {

	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)

	for i := 0.0; i < W; i += step {
		for j := 0.0; j < H; j += step {

			dc.DrawLine(i, j, i+step, j+step)
			dc.SetLineWidth(5.0)
			dc.Stroke()

		}

	}

	dc.SavePNG("draw_lines.png")

}

func draw_line_color() {

	palette := "Palette1"
	dc := gg.NewContext(W, H)
	dc.SetHexColor(colorPalettes[palette][0])
	dc.Clear()

	for i := 0.0; i < W; i += step {
		for j := 0.0; j < H; j += step {

			dc.DrawLine(i, j, i+step, j+step)
			dc.SetLineWidth(5.0)
			color_index := rand.Intn(3) + 1
			dc.SetHexColor(colorPalettes[palette][color_index])
			dc.Stroke()

		}

	}
	dc.SavePNG("draw_line_color.png")

}

func draw_line_palette() {

	dc := gg.NewContext(W, H)

	for palette, colors := range colorPalettes {
		dc.SetHexColor(colors[0])
		dc.Clear()
		for i := 0.0; i < W; i += step {
			for j := 0.0; j < H; j += step {

				dc.DrawLine(i, j, i+step, j+step)
				dc.SetLineWidth(5.0)
				color_index := rand.Intn(3) + 1
				dc.SetHexColor(colors[color_index])
				dc.Stroke()

			}

		}

		dc.SavePNG("draw_line_palette_" + palette + ".png")
	}

}

func draw_line_rotate() {

	dc := gg.NewContext(W, H)

	for palette, colors := range colorPalettes {

		dc.SetHexColor(colors[0])
		dc.Clear()
		for i := 0.0; i < W; i += step {

			for j := 0.0; j < H; j += step {
				dc.Push()
				dc.RotateAbout(gg.Radians(90.0*float64(rand.Intn(4))), (i+i+step)*0.5, (j+j+step)*0.5)
				// dc.DrawLine(i, j, i+i+step/2, i+j+step/2) // Weird
				// dc.DrawLine(i, j, i+step, j+step) //Minions
				dc.DrawLine((i+i+step)*0.5, (j+j+step)*0.5, i+step, j+step) //straight

				// fmt.Printf("%v, %v, %v, %v", i, j, i+step, j+step)
				dc.SetLineWidth(5.0)
				color_index := rand.Intn(3) + 1
				dc.SetHexColor(colors[color_index])
				dc.Stroke()
				dc.Pop()

			}

		}

		dc.SavePNG("draw_line_rotate_" + palette + ".png")

	}
}

func draw_line_stroke() {
	dc := gg.NewContext(W, H)
	for palette, colors := range colorPalettes {

		dc.SetHexColor(colors[0])
		dc.Clear()
		for i := 0.0; i < W; i += step {

			for j := 0.0; j < H; j += step {
				dc.Push()
				dc.RotateAbout(gg.Radians(90.0*float64(rand.Intn(4))), (i+i+step)*0.5, (j+j+step)*0.5)
				// dc.DrawLine(i, j, i+i+step/2, i+j+step/2) // Weird
				dc.DrawLine(i, j, i+step, j+step) //Minions
				// dc.DrawLine((i+i+step)*0.5, (j+j+step)*0.5, i+step, j+step) //straight

				dc.SetLineWidth(rand.Float64()*10 + 1.0)
				color_index := rand.Intn(3) + 1
				dc.SetHexColor(colors[color_index])
				dc.Stroke()
				dc.Pop()

			}

		}

		dc.SavePNG("draw_line_stroke_" + palette + ".png")

	}
}

func draw_line_mixed() {

	dc := gg.NewContext(W, H)

	for palette, colors := range colorPalettes {
		dc.SetHexColor(colors[0])
		dc.Clear()

		for i := 0.0; i < W; i += step {
			for j := 0.0; j < H; j += step {
				dc.Push()
				dc.RotateAbout(gg.Radians(90.0*float64(rand.Intn(4))), (i+i+step)*0.5, (j+j+step)*0.5)
				// dc.DrawLine(i, j, i+i+step/2, i+j+step/2) // Weird
				dc.DrawLine(i, j, i+step, j+step) //Minions
				// dc.DrawLine((i+i+step)*0.5, (j+j+step)*0.5, i+step, j+step) //straight

				if i < W/4 {
					dc.SetLineWidth(rand.Float64()*4 + 0.4)

				} else if i < W/3 {
					dc.SetLineWidth(rand.Float64()*6 + 0.6)

				} else if i < W/2 {
					dc.SetLineWidth(rand.Float64()*8 + 0.8)
				} else {
					dc.SetLineWidth(rand.Float64()*10 + 1.0)
				}

				color_index := rand.Intn(3) + 1
				dc.SetHexColor(colors[color_index])
				dc.Stroke()
				dc.Pop()

			}

		}

		dc.SavePNG("draw_line_mixed_" + palette + ".png")

	}
}

func draw_lines_circles() {
	dc := gg.NewContext(W, H)
	step := 40.0 //Updated step for just this function because smaller value is making it chaotic

	count := 1.0
	for palette, colors := range colorPalettes {

		dc.SetHexColor(colors[0])
		dc.Clear()
		for i := 0.0; i < W; i += step {

			for j := 0.0; j < H; j += step {
				dc.Push()
				dc.RotateAbout(gg.Radians(90.0*float64(rand.Intn(4))), (i+i+step)*0.5, (j+j+step)*0.5)
				color_index := rand.Intn(3) + 1

				if math.Remainder(count, 2) == 0.0 {

					// dc.DrawLine(i, j, i+i+step/2, i+j+step/2) // Weird
					// dc.DrawLine(i, j, i+step, j+step) //Minions
					dc.DrawLine((i+i+step)*0.5, (j+j+step)*0.5, i+step, j+step) //straight

				} else {
					dc.DrawCircle(i, j, 10.0)
				}

				dc.SetLineWidth(5.0)
				dc.SetHexColor(colors[color_index])
				dc.Stroke()
				dc.Pop()
				count += 1.0

			}
		}
		dc.SavePNG("draw_lines_circles_" + palette + ".png")
	}
}

func main() {

	//Uncomment the name of the you want to generate wallpapers of

	// draw_line()
	// draw_line_color()
	// draw_line_palette()
	// draw_line_rotate()
	// draw_line_stroke()
	// draw_line_mixed()
	draw_lines_circles()
}
