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
//TODO: fix with interface
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

func validateIsInsideBoard(p [2]float64) bool {
	x := p[0]
	y := p[1]
	if x >= 0 && x < 8 && y >= 0 && y < 8 {
		return true
	}else
	{
		return false
	}
}

func (n Knight) AvailableMoves(x, y float64) [][2]float64 {
	var nextMoves [][2]float64
	//var posibleNextMoves [2]float64
	// array of slices
	posibleNextMoves := [][2]float64{
		{x-2.0,y-1.0},
		{x-2.0,y+1.0},
		{x-1.0,y-2.0},
		{x-1.0,y+2.0},
		{x+1,y-2},
		{x+1,y+2},
		{x+2,y-1},
		{x+2,y+1},
	}
	 for _, p := range posibleNextMoves {
		 if validateIsInsideBoard(p) {
			nextMoves = append(nextMoves, p)
		 }
	 }

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
//TODO: use interface
var board [8][8]Pawn


func (g Game) PopulateInitialBoard(c bool) {
	//place pawns
	//TODO this test dont work
	/*if ( c == WHITE ) {
		strcolor := "w"
		strcolor2 := "b"
	} else {
		strcolor := "b"
		strcolor2 := "w"
	} */
	strcolor := "w"
	strcolor2 := "b"

	log.Print(strcolor)
	for x:= 0; x < 8; x++ {
			board[x][6] = Pawn{Piece{ImageFile:"images/" + strcolor +"pawn.png",Color:true}}
			board[x][1] = Pawn{Piece{ImageFile:"images/" + strcolor2 +"pawn.png",Color:true}}
	}
	// place pieces
	//TODO: apply interface make it work
	//board[0][7] = Rook{Piece{ImageFile:"images/wrook.png",Color:true}}
	board[0][7] = Pawn{Piece{ImageFile:"images/"+strcolor+"rook.png",Color:true}}
	//board[1][7] = Knight{Piece{ImageFile:"images/wknight.png",Color:true}}
	board[1][7] = Pawn{Piece{ImageFile:"images/" + strcolor +"knight.png",Color:true}}
	//board[2][7] = Bishop{Piece{ImageFile:"images/wbishop.png",Color:true}}
	board[2][7] = Pawn{Piece{ImageFile:"images/" + strcolor +"bishop.png",Color:true}}
	//board[3][7] = Queen{Piece{ImageFile:"images/wqueen.png",Color:true}}
	board[3][7] = Pawn{Piece{ImageFile:"images/" + strcolor +"queen.png",Color:true}}
	//board[4][7] = King{Piece{ImageFile:"images/wking.png",Color:true}}
	board[4][7] = Pawn{Piece{ImageFile:"images/" + strcolor +"king.png",Color:true}}
	//board[5][7] = Bishop{Piece{ImageFile:"images/wbishop.png",Color:true}}
	board[5][7] = Pawn{Piece{ImageFile:"images/" + strcolor +"bishop.png",Color:true}}
	//board[6][7] = Knight{Piece{ImageFile:"images/wknight.png",Color:true}}
	board[6][7] = Pawn{Piece{ImageFile:"images/" + strcolor +"knight.png",Color:true}}
	//board[7][7] = Rook{Piece{ImageFile:"images/wrook.png",Color:true}}
	board[7][7] = Pawn{Piece{ImageFile:"images/" + strcolor +"rook.png",Color:true}}

	board[0][0] = Pawn{Piece{ImageFile:"images/" + strcolor2 +"rook.png",Color:false}}
	//board[0][0] = Rook{Piece{ImageFile:"images/brook.png",Color:false}}
	board[1][0] = Pawn{Piece{ImageFile:"images/" + strcolor2 +"knight.png",Color:false}}
	//board[1][0] = Knight{Piece{ImageFile:"images/bknight.png",Color:false}}
	board[2][0] = Pawn{Piece{ImageFile:"images/" + strcolor2 +"bishop.png",Color:false}}
	//board[2][0] = Bishop{Piece{ImageFile:"images/bbishop.png",Color:false}}
	board[3][0] = Pawn{Piece{ImageFile:"images/" + strcolor2 +"queen.png",Color:false}}
	//board[3][0] = Queen{Piece{ImageFile:"images/bqueen.png",Color:false}}
	board[4][0] = Pawn{Piece{ImageFile:"images/" + strcolor2 +"king.png",Color:false}}
	//board[4][0] = King{Piece{ImageFile:"images/bking.png",Color:false}}
	board[5][0] = Pawn{Piece{ImageFile:"images/" + strcolor2 +"bishop.png",Color:false}}
	//board[5][0] = Bishop{Piece{ImageFile:"images/bbishop.png",Color:false}}
	board[6][0] = Pawn{Piece{ImageFile:"images/" + strcolor2 +"knight.png",Color:false}}
	//board[6][0] = Knight{Piece{ImageFile:"images/bknight.png",Color:false}}
	board[7][0] = Pawn{Piece{ImageFile:"images/" + strcolor2 +"rook.png",Color:false}}
	//board[7][0] = Rook{Piece{ImageFile:"images/brook.png",Color:false}}
}

func (g Game) DrawBoard(da *gtk.DrawingArea, cr *cairo.Context) {
	i := 0;
	for x := range board {
		i++
		for y := range board[x] {
			i++
			ipiece := board[x][y]
			log.Print("x: ")
			log.Print(x)
			log.Print("y: ")
			log.Print(y)
			log.Print("Piece: ")
			log.Print(ipiece)
			log.Print("+++++")
			log.Print(i);
			log.Print("+++++")
			/*if (ipiece == nil) {
				continue
			}*/
			if i % 2 == 0 {
				cr.SetSourceRGBA(0, 0, 2, 0.1)
			} else
			{
				cr.SetSourceRGBA(0, 0, 2, 0.3)
			}
			//TODO apply interface
			//if (ipiece.(Piece).ImageFile == "") {
			if (ipiece.ImageFile == "") {
		 	    cr.Rectangle(float64(x)*UNITSIZE, float64(y)*UNITSIZE, UNITSIZE, UNITSIZE)
				cr.Fill()
				continue
			}
		    cr.Rectangle(float64(x)*UNITSIZE, float64(y)*UNITSIZE, UNITSIZE, UNITSIZE)
			cr.Fill()
			//p, _ := gdk.PixbufNewFromFile(ipiece.(Piece).ImageFile)
			p, _ := gdk.PixbufNewFromFile(ipiece.ImageFile)
			gtk.GdkCairoSetSourcePixBuf(cr, p, float64(x)*UNITSIZE, float64(y)*UNITSIZE)
		    cr.Rectangle(float64(x)*UNITSIZE, float64(y)*UNITSIZE, UNITSIZE, UNITSIZE)
			cr.Fill()
		}
	}
}


func main() {
	piecen := Pawn{Piece{"images/wpawn.png",true}}
	board[0][6] = piecen

	gtk.Init(nil)
	GameI := Game{"raul","pedro", true}
	GameI.PopulateInitialBoard(true)
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
