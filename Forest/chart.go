package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func generateChart1(data []map[int]float64) {
	p := plot.New()

	p.Title.Text = "Wykres zależności procentu zalesienia lasu od procentu spalonego lasu"
	p.X.Label.Text = "Procent zalesienia lasu"
	p.Y.Label.Text = "Procent spalonego lasu"

	pts := make(plotter.XYs, 0)
	for _, v := range data {
		for k, val := range v {
			if math.IsNaN(val) {
				continue
			}
			pt := plotter.XY{
				X: float64(k),
				Y: val,
			}
			pts = append(pts, pt)
		}
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}

	p.Add(line)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "burned.png"); err != nil {
		panic(err)
	}
}

func generateChart2(data []map[int]float64) {
	p := plot.New()

	p.Title.Text = "Ile drzew przetrwało po pożarze w zależności od procentu zalesienia lasu"
	p.X.Label.Text = "Procent zalesienia lasu"
	p.Y.Label.Text = "Ilość drzew przetrwałych po pożarze"

	pts := make(plotter.XYs, 0)
	for _, v := range data {
		for k, val := range v {
			if math.IsNaN(val) {
				continue
			}
			pt := plotter.XY{
				X: float64(k),
				Y: val,
			}
			pts = append(pts, pt)
		}
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}

	p.Add(line)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "howManyAlive.png"); err != nil {
		panic(err)
	}
}
