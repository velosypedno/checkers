package ui

import (
	"image/color"
	"math"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/velosypedno/checkers/backend"
)

const (
	WindowWithPX                = 700
	WindowHighPX                = 700
	Title                       = "Checkers"
	boardSizePX                 = 640
	offsetXY                    = 20
	frameSizePX                 = 5
	cellSizePX                  = boardSizePX / boardSize
	checkerOffsetInsideCellPX   = 10
	checkerSizePX               = cellSizePX - 2*checkerOffsetInsideCellPX
	highlightOutlineThicknessPX = 5
	circleRadiusPX              = 10
	boardSize                   = 8
)

var (
	backgroundColor       = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	frameColor            = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
	highlightOutlineColor = color.RGBA{255, 200, 100, 180}
)

type point struct {
	x, y int
}

type Game struct {
	gameBackend   *backend.GameBackend
	selected      *point
	possibleMoves []point
}

func NewGame() *Game {
	return &Game{
		gameBackend:   backend.NewGameBackend(),
		possibleMoves: []point{},
	}
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		xCell := (x - offsetXY - frameSizePX) / cellSizePX
		yCell := (y - offsetXY - frameSizePX) / cellSizePX
		if g.gameBackend.IsMoveable(xCell, yCell) {
			g.selected = &point{xCell, yCell}
			possibleMoves := []point{}
			for _, p := range g.gameBackend.PossibleMoves(xCell, yCell) {
				possibleMoves = append(possibleMoves, point{p.X, p.Y})
			}
			g.possibleMoves = possibleMoves
		} else {
			g.selected = nil
			g.possibleMoves = []point{}
		}

	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.selected = nil
		g.possibleMoves = []point{}
	}
	return nil
}

func DrawFrame(frameImg *ebiten.Image) {
	frameImg.Fill(frameColor)
}

func DrawBoard(boardImg *ebiten.Image) {
	boardImg.Fill(frameColor)
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
				outlineImg.Set(x, y, highlightOutlineColor)
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
				circleImg.Set(x, y, highlightOutlineColor)
			} else {
				circleImg.Set(x, y, color.RGBA{0, 0, 0, 0})
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.Fill(backgroundColor)

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
			side := stateOfBoard[i][j]
			if side == backend.Red {
				checker = GetRedChecker(checkerSizePX)
			} else if side == backend.Blue {
				checker = GetBlueChecker(checkerSizePX)
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
		xOffset = float64(cellSizePX * g.selected.x)
		yOffset = float64(cellSizePX * g.selected.y)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(xOffset, yOffset)
		boardImg.DrawImage(outlineImg, op)
	}

	// Draw possible moves
	for _, point := range g.possibleMoves {
		pointImage := ebiten.NewImage(cellSizePX, cellSizePX)
		DrawCenterCircle(pointImage)
		xOffset = float64(cellSizePX * point.x)
		yOffset = float64(cellSizePX * point.y)
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(xOffset, yOffset)
		boardImg.DrawImage(pointImage, op)
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
