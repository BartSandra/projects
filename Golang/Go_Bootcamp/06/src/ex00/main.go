package main

import (
	"github.com/fogleman/gg"
	"log"
	"math"
)

func main() {
	const W = 300
	const H = 300
	dc := gg.NewContext(W, H)

	dc.SetRGB(0.9, 0.9, 0.9) // Цвет фона
	dc.Clear()

	dc.SetRGB(0.2, 0.3, 0.4)     // Цвет формы
	dc.DrawCircle(W/2, H/2, 100) // Форма
	dc.Fill()

	// Рисование глаз
	dc.SetRGB(1, 1, 1)                // Белый цвет глаз
	dc.DrawCircle(W/2-35, H/2-10, 30) // Глаза чуть ниже
	dc.DrawCircle(W/2+35, H/2-10, 30)
	dc.Fill()

	dc.SetRGB(0, 0, 0) // Черные зрачки
	dc.DrawCircle(W/2-35, H/2-10, 15)
	dc.DrawCircle(W/2+35, H/2-10, 15)
	dc.Fill()

	// Рисование ушей
	dc.SetRGB(0.2, 0.3, 0.4)
	dc.DrawCircle(W/2-70, H/2-70, 30) // Уши
	dc.DrawCircle(W/2+70, H/2-70, 30)
	dc.Fill()

	// Рисование носика
	dc.SetRGB(0, 0, 0)             // Черный цвет носика
	dc.DrawCircle(W/2, H/2+10, 10) // Носик чуть ниже
	dc.Fill()

	// Рисование улыбки
	dc.SetRGB(1, 1, 1)                                  // Белый цвет улыбки
	dc.SetLineWidth(3)                                  // Толщина линии дуги
	dc.DrawArc(W/2, H/2+20, 40, math.Pi/4, 3*math.Pi/4) // Дуга
	dc.Stroke()

	err := dc.SavePNG("amazing_logo.png")
	if err != nil {
		log.Fatalf("error while creating logo: %v", err)
	}
}
