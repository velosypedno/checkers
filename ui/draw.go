package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/velosypedno/checkers/backend"
	"golang.org/x/image/font/basicfont"
)

const (
	Title                            = "Checkers"
	boardSizePX                      = 640
	offsetXY                         = 20
	frameSizePX                      = 5
	cellSizePX                       = boardSizePX / boardSize
	checkerOffsetInsideCellPX        = 10
	checkerSizePX                    = cellSizePX - 2*checkerOffsetInsideCellPX
	highlightOutlineThicknessPX      = 5
	possibleSelectOutlineThicknessPX = 5
	circleRadiusPX                   = 10
	boardSize                        = 8
	startButtonWidthPX               = 200
	startButtonHeightPX              = 60
)

func DrawFrame(frameImg *ebiten.Image) {
	frameImg.Fill(FrameOfTheBoardColor)
}

func DrawBoard(boardImg *ebiten.Image) {
	boardImg.Fill(FrameOfTheBoardColor)
	isBlack := true
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			cellImg := ebiten.NewImage(cellSizePX, cellSizePX)
			if isBlack = !isBlack; isBlack {
				cellImg.Fill(color.Black)
			} else {
				cellImg.Fill(color.White)
			}
			op := &ebiten.DrawImageOptions{}
			xOffset := j * cellSizePX
			yOffset := i * cellSizePX
			op.GeoM.Translate(float64(xOffset), float64(yOffset))
			boardImg.DrawImage(cellImg, op)
		}
		isBlack = !isBlack
	}

}

func DrawHighlightOutline(outlineImg *ebiten.Image) {
	bounds := outlineImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y < highlightOutlineThicknessPX || y >= height-highlightOutlineThicknessPX ||
				x < highlightOutlineThicknessPX || x >= width-highlightOutlineThicknessPX {
				outlineImg.Set(x, y, HighlightOutlineColor)
			} else {
				outlineImg.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}
}

func DrawPossibleSelectOutline(outlineImg *ebiten.Image) {
	bounds := outlineImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y < possibleSelectOutlineThicknessPX || y >= height-possibleSelectOutlineThicknessPX ||
				x < possibleSelectOutlineThicknessPX || x >= width-possibleSelectOutlineThicknessPX {
				outlineImg.Set(x, y, PossibleSelectOutlineColor)
			} else {
				outlineImg.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}
}
func DrawCenterCircle(circleImg *ebiten.Image) {
	bounds := circleImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	cx, cy := float64(width)/2, float64(height)/2
	r := float64(circleRadiusPX)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dx := float64(x) - cx
			dy := float64(y) - cy
			dist := math.Hypot(dx, dy)
			if dist <= r {
				circleImg.Set(x, y, HighlightOutlineColor)
			} else {
				circleImg.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}
}

func DrawStartButton(buttonImg *ebiten.Image) {
	// Fill button background
	buttonImg.Fill(StartButtonColor)

	// Draw button border
	bounds := buttonImg.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Draw border
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y < 2 || y >= height-2 || x < 2 || x >= width-2 {
				buttonImg.Set(x, y, color.RGBA{0, 0, 0, 255}) // Black border
			}
		}
	}

	// Draw text "START"
	face := basicfont.Face7x13
	textStr := "START"
	textBounds := text.BoundString(face, textStr)
	textX := (width - textBounds.Dx()) / 2
	textY := (height + textBounds.Dy()) / 2

	text.Draw(buttonImg, textStr, face, textX, textY, StartButtonTextColor)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.Fill(GameBackgroundColor)

	// If in StartScreen state, draw start button
	if g.state == StartScreen {
		// Create start button
		buttonImg := ebiten.NewImage(startButtonWidthPX, startButtonHeightPX)
		DrawStartButton(buttonImg)

		// Position button in center of screen
		buttonX := (WindowWithPX - startButtonWidthPX) / 2
		buttonY := (WindowHighPX - startButtonHeightPX) / 2

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(buttonX), float64(buttonY))
		screen.DrawImage(buttonImg, op)
		return
	}

	// Draw frame
	frameImg := ebiten.NewImage(boardSizePX+frameSizePX*2, boardSizePX+frameSizePX*2)
	DrawFrame(frameImg)

	// Draw board
	boardImg := ebiten.NewImage(boardSizePX, boardSizePX)
	DrawBoard(boardImg)

	// Draw checkers on the board
	stateOfGame := g.gameBackend.GetState()
	stateOfBoard := stateOfGame.Board
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			var checker *ebiten.Image
			c := stateOfBoard[i][j]
			if c.Side == backend.Red {
				checker = GetRedChecker(checkerSizePX, c.IsQueen)
			} else if c.Side == backend.Blue {
				checker = GetBlueChecker(checkerSizePX, c.IsQueen)
			} else {
				continue
			}
			checkerOffsetInsideBoardByX := float64(cellSizePX*j + checkerOffsetInsideCellPX)
			checkerOffsetInsideBoardByY := float64(cellSizePX*i + checkerOffsetInsideCellPX)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(checkerOffsetInsideBoardByX, checkerOffsetInsideBoardByY)
			boardImg.DrawImage(checker, op)
		}
	}

	var op *ebiten.DrawImageOptions
	var xOffset float64
	var yOffset float64

	// Draw highlighted outline
	if g.selected != nil {
		outlineImg := ebiten.NewImage(cellSizePX, cellSizePX)
		DrawHighlightOutline(outlineImg)
		xOffset = float64(cellSizePX * g.selected.X)
		yOffset = float64(cellSizePX * g.selected.Y)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(xOffset, yOffset)
		boardImg.DrawImage(outlineImg, op)
	}

	// Draw possible moves
	for _, point := range g.possibleMovesOfSelected {
		pointImage := ebiten.NewImage(cellSizePX, cellSizePX)
		DrawCenterCircle(pointImage)
		xOffset = float64(cellSizePX * point.X)
		yOffset = float64(cellSizePX * point.Y)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(xOffset, yOffset)
		boardImg.DrawImage(pointImage, op)
	}

	for _, attack := range g.possibleAttacksSelected {
		pointImage := ebiten.NewImage(cellSizePX, cellSizePX)
		DrawCenterCircle(pointImage)
		xOffset = float64(cellSizePX * attack.Move.X)
		yOffset = float64(cellSizePX * attack.Move.Y)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(xOffset, yOffset)
		boardImg.DrawImage(pointImage, op)
	}

	for _, candidateToAttack := range g.candidatesToAttack {
		outlineImg := ebiten.NewImage(cellSizePX, cellSizePX)
		DrawPossibleSelectOutline(outlineImg)
		xOffset = float64(cellSizePX * candidateToAttack.X)
		yOffset = float64(cellSizePX * candidateToAttack.Y)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(xOffset, yOffset)
		boardImg.DrawImage(outlineImg, op)
	}

	// Draw board on the frame
	op = &ebiten.DrawImageOptions{}
	xOffset = frameSizePX
	yOffset = frameSizePX
	op.GeoM.Translate(xOffset, yOffset)
	frameImg.DrawImage(boardImg, op)

	// Draw frame on the screen
	op = &ebiten.DrawImageOptions{}
	xOffset = offsetXY
	yOffset = offsetXY
	op.GeoM.Translate(xOffset, yOffset)
	screen.DrawImage(frameImg, op)
}

func (g *Game) Layout(outsideWith, outsideHeight int) (int, int) {
	return WindowWithPX, WindowHighPX
}
