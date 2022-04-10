/*
 * MIT License
 *
 * Copyright (c) 2022 Ludovic Cleroux
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package mazegen

import (
	"github.com/nsf/termbox-go"
	"time"
)

var createMaze func()

func Run(width, height int) error {
	createMaze = func() {
		maze = NewMaze(width, height)
	}
	createMaze()
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()
	evQueue := make(chan termbox.Event)
	go func() {
		for {
			evQueue <- termbox.PollEvent()
		}
	}()
	draw()
loop:
	for {
		select {
		case ev := <-evQueue:
			if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyEsc {
					break loop
				} else {
					createMaze()
				}
			}
		default:
			hasNext := maze.Next()
			draw()
			if hasNext {
				// time.Sleep(1 * time.Nanosecond)
			} else {
				time.Sleep(50 * time.Millisecond)
			}

		}
	}

	return nil
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < len(maze.Maze); y++ {
		for x := 0; x < len(maze.Maze[y]); x++ {
			if maze.Maze[y][x] {
				termbox.SetBg(x*2, y, termbox.ColorWhite)
				termbox.SetBg(x*2+1, y, termbox.ColorWhite)
			} else {
				termbox.SetBg(x*2, y, termbox.ColorBlack)
				termbox.SetBg(x*2+1, y, termbox.ColorBlack)
			}
		}
	}
	termbox.Flush()
}
