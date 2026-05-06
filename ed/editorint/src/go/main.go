package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {

    if e.cursor != e.line.Value.Front() {
        e.cursor = e.cursor.Prev()
        return
    }

    if e.line.Prev() != nil {
        e.line = e.line.Prev()
        e.cursor = e.line.Value.Back()
    }
}

func (e *Editor) KeyEnter() {
	nova := NewList[rune]()
	e.lines.Insert(e.line.Next(), nova)
	e.line = e.line.Next()
	e.cursor = e.line.Value.Front()
}

func (e *Editor) KeyRight() {

    if e.cursor != e.line.Value.Back() {
        e.cursor = e.cursor.Next()
        return
    }

    if e.line.Next() != nil {
        e.line = e.line.Next()
        e.cursor = e.line.Value.Front()
    }
}

func (e *Editor) KeyUp() {
	if e.line != e.lines.root.Next()  {
	
		if   e.line.Prev().Value.Size() == 0 {
			e.line = e.line.Prev()
			e.cursor = e.line.Value.Back()
			
		}else{
			
			e.line = e.line.Prev()
			e.cursor = e.line.Value.Back()
		}
		return
	}
}

func (e *Editor) KeyDown() {
	if e.line != e.lines.root.Prev()  {
		
		if   e.line.Next().Value.Size() == 0 {
			e.line = e.line.Next()
			e.cursor = e.line.Value.Back()
			
		}else{
			
			e.line = e.line.Next()
			e.cursor = e.line.Value.Back()
		}
		return
	}
}

func (e *Editor) juntarLinhas() {

    atual := e.line
    anterior := atual.Prev()

    if anterior == nil {
        return
    }

    // move todos os caracteres da linha atual para a anterior
    for n := atual.Value.Front(); n != atual.Value.End(); n = n.Next() {
        anterior.Value.PushBack(n.Value)
    }

    // remove linha atual
    e.lines.Erase(atual)

    // atualiza estado
    e.line = anterior
    e.cursor = anterior.Value.Back()
}

func (e *Editor) KeyBackspace() {

    chars := e.line.Value

    // CASO 1: não está no início da linha → apaga caractere anterior
    if e.cursor != chars.Front() {

        prev := e.cursor.Prev()

        chars.Erase(e.cursor)

        e.cursor = prev
        return
    }

    // CASO 2: início da linha → junta linhas
    if e.line.Prev() != nil {
        e.juntarLinhas()
    }
}



func (e *Editor) KeyDelete() {

    chars := e.line.Value

    // CASO 1: tem próximo caractere
    if e.cursor != chars.Back() {

        next := e.cursor.Next()

        chars.Erase(e.cursor)

        e.cursor = next
        return
    }

    // CASO 2: fim da linha → junta com próxima
    if e.line.Next() != nil {

        linhaAtual := e.line.Next()

        for n := linhaAtual.Value.Front(); n != linhaAtual.Value.End(); n = n.Next() {
            e.line.Value.PushBack(n.Value)
        }

        e.lines.Erase(linhaAtual)
    }
}



func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
