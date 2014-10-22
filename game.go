package main

import (
	"github.com/conformal/gotk3/cairo"
	"github.com/conformal/gotk3/gdk"
	"github.com/conformal/gotk3/gtk"
	"log"
)

const(
	WHITE bool = true
	BLACK bool = false
	UNITSIZE float64 = 45.0
)

type Mover interface {
	AvailableMoves(x, y float64) [][2]float64
}

type Piece struct {
	ImageFile string
	Color bool
}

func (pp Piece) AvailableMoves(x, y float64) [][2]float64 {
	var nextMoves [][2]float64
	//TODO
	currentPlace := [2]float64{x,y}
	nextMoves = append(nextMoves, currentPlace)
	return nextMoves
}

type Pawn struct {
	Piece
}

func (p Pawn) AvailableMoves(x, y float64) [][2]float64 {
	var nextMoves [][2]float64
	//TODO
	currentPlace := [2]float64{x,y}
	nextMoves = append(nextMoves, currentPlace)
	return nextMoves
}

type Bishop struct {
	Piece
}

func (b Bishop) AvailableMoves(x, y float64) [][2]float64 {
	var nextMoves [][2]float64
	//TODO
	currentPlace := [2]float64{x,y}
	nextMoves = append(nextMoves, currentPlace)
	return nextMoves
}

type Knight struct {
	Piece
}

func (n Knight) AvailableMoves(x, y float64) [][2]float64 {
	var nextMoves [][2]float64
	//TODO
	currentPlace := [2]float64{x,y}
	nextMoves = append(nextMoves, currentPlace)
	return nextMoves
}
type Rook struct {
	Piece
}

func (r Rook) AvailableMoves(x, y float64) [][2]float64 {
	var nextMoves [][2]float64
	//TODO
	currentPlace := [2]float64{x,y}
	nextMoves = append(nextMoves, currentPlace)
	return nextMoves
}

type Queen struct {
	Piece
}

func (q Queen) AvailableMoves(x, y float64) [][2]float64 {
	var nextMoves [][2]float64
	//TODO
	currentPlace := [2]float64{x,y}
	nextMoves = append(nextMoves, currentPlace)
	return nextMoves
}
type King struct {
	Piece
}

func (k King) AvailableMoves(x, y float64) [][2]float64 {
	var nextMoves [][2]float64
	//TODO
	currentPlace := [2]float64{x,y}
	nextMoves = append(nextMoves, currentPlace)
	return nextMoves
}

type Game struct {
	whitetime string
	blacktime string
	playingwhite bool
}

type Player struct {
	name string
}

var board [8][8]Mover

func (g Game) PopulateInitialBoard(c bool) {
	if c == WHITE {
				//place pawns
				for x:= 0; x < 8; x++ {
					board[x][6] = Pawn{Piece{ImageFile:"images/wpawn.png",Color:true}}
					board[x][1] = Pawn{Piece{ImageFile:"images/bpawn.png",Color:true}}
				}

				// place pieces
				board[0][7] = Rook{Piece{ImageFile:"images/wrook.png",Color:true}}
				//board[0][7].AvailableMoves(1.0, 2.0);
				board[1][7] = Knight{Piece{ImageFile:"images/wknight.png",Color:true}}
			//	board[1][7].AvailableMoves(1.0, 2.0);
				board[2][7] = Bishop{Piece{ImageFile:"images/wbishop.png",Color:true}}
			//	board[2][7].AvailableMoves(1.0, 2.0);
				board[3][7] = Queen{Piece{ImageFile:"images/wqueen.png",Color:true}}
		//		board[3][7].AvailableMoves(1.0, 2.0);
				board[4][7] = King{Piece{ImageFile:"images/wking.png",Color:true}}
	//			board[4][7].AvailableMoves(1.0, 2.0);
				board[5][7] = Bishop{Piece{ImageFile:"images/wbishop.png",Color:true}}
	//			board[5][7].AvailableMoves(1.0, 2.0);
				board[6][7] = Knight{Piece{ImageFile:"images/wknight.png",Color:true}}
	//			board[6][7].AvailableMoves(1.0, 2.0);
				board[7][7] = Rook{Piece{ImageFile:"images/wrook.png",Color:true}}
		//		board[7][7].AvailableMoves(1.0, 2.0);

				board[0][0] = Rook{Piece{ImageFile:"images/brook.png",Color:false}}
		//		board[0][0].AvailableMoves(1.0, 2.0)
				board[1][0] = Knight{Piece{ImageFile:"images/bknight.png",Color:false}}
		//		board[1][0].AvailableMoves(1.0, 2.0)
				board[2][0] = Bishop{Piece{ImageFile:"images/bbishop.png",Color:false}}
		//		board[2][0].AvailableMoves(1.0, 2.0)
				board[3][0] = Queen{Piece{ImageFile:"images/bqueen.png",Color:false}}
		//		board[3][0].AvailableMoves(1.0, 2.0)
				board[4][0] = King{Piece{ImageFile:"images/bking.png",Color:false}}
	//			board[4][0].AvailableMoves(1.0, 2.0)
				board[5][0] = Bishop{Piece{ImageFile:"images/bbishop.png",Color:false}}
	//			board[5][0].AvailableMoves(1.0, 2.0)
				board[6][0] = Knight{Piece{ImageFile:"images/bknight.png",Color:false}}
	//			board[6][0].AvailableMoves(1.0, 2.0)
				board[7][0] = Rook{Piece{ImageFile:"images/brook.png",Color:false}}
	//			board[7][0].AvailableMoves(1.0, 2.0)
	}
}

func (g Game) DrawBoard(da *gtk.DrawingArea, cr *cairo.Context) {
	for x := range board {
		for y := range board[x] {
			ipiece := board[x][y]
			if (ipiece == nil) {
				continue
			}
			//TODO better way to check if no piece object in the array
			if y %2 == 0 {
				cr.SetSourceRGBA(0, 0, 2, 0.1)
			} else
			{
				cr.SetSourceRGBA(0, 0, 2, 0.5)
			}
			if (ipiece.(Piece).ImageFile == "") {
		 	    cr.Rectangle(float64(x)*UNITSIZE, float64(y)*UNITSIZE, UNITSIZE, UNITSIZE)
				cr.Fill()
				continue
			}
		    cr.Rectangle(float64(x)*UNITSIZE, float64(y)*UNITSIZE, UNITSIZE, UNITSIZE)
			cr.Fill()
			p, _ := gdk.PixbufNewFromFile(ipiece.(Piece).ImageFile)
			gtk.GdkCairoSetSourcePixBuf(cr, p, float64(x)*UNITSIZE, float64(y)*UNITSIZE)
		    cr.Rectangle(float64(x)*UNITSIZE, float64(y)*UNITSIZE, UNITSIZE, UNITSIZE)
			cr.Fill()
		}
	}
}


func main() {

	gtk.Init(nil)
	GameI := Game{"raul","pedro", true}
	GameI.PopulateInitialBoard(WHITE)
	// gui boilerplate
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	grid, err := gtk.GridNew()
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	da, err := gtk.DrawingAreaNew()
	if err != nil {
		log.Fatal("Unable to create drawing area:", err)
	}
	//win.Add(da)
	win.SetTitle("Chess GoNuts")
	win.Connect("destroy", gtk.MainQuit)
	//win.ShowAll()
	// Create a new grid widget to arrange child widgets
	// Create some widgets to put in the grid.
	lab, err := gtk.LabelNew("Here to place the clock")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// Data

	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		GameI.DrawBoard(da, cr)
	})

	grid.Add(lab)
	grid.Attach(da, 1, 1, 1, 2)
	// has to be set for drawing area to work inside grid
	da.SetHExpand(true)
	da.SetVExpand(true)

	win.Resize(800, 600)
	win.Add(grid)
	win.ShowAll()
	gtk.Main()
}
