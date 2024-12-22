package main

import (
	"fmt"
	"strings"

	"github.com/trippingcoin/Pugashbek_Torekhan/Assaignment_1/Assaignment_1/Shapes/shapes"
)

func main() {
	shapesList := []shapes.Shape{}

	for {
		fmt.Println("Choose an option: Add (Shape), List (Shapes), Exit")
		var command string
		fmt.Print("> ")
		fmt.Scanln(&command)

		switch strings.ToLower(command) {
		case "add":
			fmt.Println("Select Shape: Rectangle, Circle, Square, Triangle")
			var shapeType string
			fmt.Print("> ")
			fmt.Scanln(&shapeType)

			switch strings.ToLower(shapeType) {
			case "rectangle":
				var length, width float64
				fmt.Print("Enter Length: ")
				fmt.Scanln(&length)
				fmt.Print("Enter Width: ")
				fmt.Scanln(&width)
				rectangle := shapes.Rectangle{Length: length, Width: width}
				shapesList = append(shapesList, rectangle)

			case "circle":
				var radius float64
				fmt.Print("Enter Radius: ")
				fmt.Scanln(&radius)
				circle := shapes.Circle{Radius: radius}
				shapesList = append(shapesList, circle)

			case "square":
				var length float64
				fmt.Print("Enter Side Length: ")
				fmt.Scanln(&length)
				square := shapes.Square{Length: length}
				shapesList = append(shapesList, square)

			case "triangle":
				var sideA, sideB, sideC float64
				fmt.Print("Enter Side A: ")
				fmt.Scanln(&sideA)
				fmt.Print("Enter Side B: ")
				fmt.Scanln(&sideB)
				fmt.Print("Enter Side C: ")
				fmt.Scanln(&sideC)
				triangle := shapes.Triangle{SideA: sideA, SideB: sideB, SideC: sideC}
				shapesList = append(shapesList, triangle)

			default:
				fmt.Println("Invalid shape type.")
			}

		case "list":
			for i, shape := range shapesList {
				fmt.Printf("\nShape %d:\n", i+1)
				shapes.PrintShapeDetails(shape)
			}

		case "exit":
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
