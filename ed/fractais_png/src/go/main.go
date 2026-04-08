package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func arvere(pen *Pen, dist float64) {
	if dist < 10 {
		/*if ri(0, 50) == 0 {
			pen.SetRGB(255, 0, 0)
			pen.FillCircle(10)
		}*/
		return
	}
	
	//ang_dir := 90
	//ang_esq := 180

    /*pen.SetHeading(float64(90 ))       
	pen.Walk(dist)
	pen.SetHeading(float64(180))
	pen.Walk(dist)
	 dist= dist/1.05
	pen.SetHeading(float64(270 ))       
	pen.Walk(dist)
	pen.SetHeading(float64(0))
	pen.Walk(dist)
	arvere(pen, dist/1.05)*/
	pen.SetRGB(ri(0 , 255),ri(0 , 255),ri(0 , 255))
	pen.SetLineWidth(dist / 30)
	pen.Left(90)
    pen.Walk(dist)
	arvere(pen, dist/1.05)
	//pen.SetLineWidth(dist / 5)
	//pen.SetRGB(0, 0, 0)
	//pen.Walk(dist)
	//pen.Right(ang_dir)
	//arvere(pen, dist*(ri(80, 85)/100))
	//pen.Left(ang_dir + ang_esq)
	//arvere(pen, dist*(ri(80, 85)/100))
	//pen.Right(ang_esq)
	//pen.SetRGB(0, 0, 0)
	//pen.Walk(-dist)
}

func main() {
	pen := NewPen(500, 500)
	pen.SetHeading(90)
	pen.SetPosition(450, 50)
	arvere(pen, 400)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}

/* func main() {
	pen := NewPen(500, 500)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(255, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(250, 500) // move o pincel para x 250, y 500
	pen.SetHeading(90)        // coloca o pincel apontando para cima
	pen.Walk(100)             // anda 100 pixels
	pen.Left(30)              // dobra 30 graus para esquerda
	pen.Walk(100)             // anda 100 pixels
	pen.DrawCircle(50)        // desenha um círculo de raio 50
	pen.Right(60)             // gira para direita 60 graus
	pen.Walk(150)
	for range 10 {
		pen.Up()
		pen.Walk(30) // anda sem riscar
		pen.Down()

		pen.DrawCircle(10) //desenha um circulo pequeno

		pen.Up()
		pen.Walk(-30) // volta sem riscar
		pen.Down()

		pen.Left(36) // gira
	} */
