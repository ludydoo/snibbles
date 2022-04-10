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

package bstgame

import (
	"dsa/pkg/bst2"
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var node *bst2.Node
var selected *bst2.Node
var isInserting bool
var insertText string

func Run() error {
	ints := uniqueRandomInts(30)
	node = bst2.New(ints...)
	selected = node

	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	draw()
loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyEsc {
					if isInserting {
						isInserting = false
					} else {
						break loop
					}
				} else {

					if selected == nil {
						selected = node
					}

					if isInserting {
						if ev.Key == termbox.KeyEnter {
							val, err := strconv.Atoi(insertText)
							if err != nil {
								continue
							}
							bst2.Insert(node, val)
							insertText = ""
							isInserting = false
						} else if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
							if len(insertText) > 0 {
								insertText = insertText[:len(insertText)-1]
							}
						} else {
							r := ev.Ch
							if r >= '0' && r <= '9' {
								insertText += string(r)
							}
						}
						continue
					}

					if ev.Key == termbox.KeyArrowLeft {
						if selected.Left != nil {
							selected = selected.Left
						}
					} else if ev.Key == termbox.KeyArrowRight {
						selected = selected.Right
					} else if ev.Key == termbox.KeyArrowUp {
						if selected != nil && selected != node {
							selected = bst2.GetParent(node, selected.Key)
						}
					} else if ev.Ch == 'n' {
						successor := bst2.FindInOrderSuccessor(node, selected.Key)
						if successor != nil {
							selected = successor
						}
					} else if ev.Ch == 'p' {
						successor := bst2.FindInOrderPredecessor(node, selected.Key)
						if successor != nil {
							selected = successor
						}
					} else if ev.Ch == 'd' {
						if selected != nil {
							toDelete := selected
							if toDelete.Right != nil && toDelete.Left != nil {
								selected = toDelete
							} else if toDelete.Right != nil {
								selected = toDelete.Right
							} else if toDelete.Left != nil {
								selected = toDelete.Left
							} else if selected == node {
								continue
							} else {
								selected = bst2.GetParent(node, toDelete.Key)
							}
							node = bst2.DeleteKey(node, toDelete.Key)
						}
					} else if ev.Ch == 'i' {
						isInserting = true
					}
				}
			}
		default:
			draw()
			time.Sleep(10 * time.Millisecond)
		}
	}
	return nil
}

func uniqueRandomInts(n int) []int {
	var m = map[int]struct{}{}
	var ints []int
	var i int
	for {
		r := rand.Intn(100)
		if _, ok := m[r]; !ok {
			ints = append(ints, r)
			m[r] = struct{}{}
			i++
		}
		if i == n {
			break
		}
	}
	return ints
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	bstStr := bst2.Print(node, bst2.PrintCallback(func(node *bst2.Node) string {
		if selected == node {
			return fmt.Sprintf("-> %d <-", node.Key)
		}
		return ""
	}))

	upCmd := "Up Arrow to select parent"
	leftCmd := "Left Arrow to select left child"
	rightCmd := "Right Arrow to select right child"
	nextCmd := "n to select successor"
	prevCmd := "p to select predecessor"
	insertCmd := "i to insert"
	deleteCmd := "d to delete the selected node"
	cmds := []string{upCmd, leftCmd, rightCmd, nextCmd, prevCmd, insertCmd, deleteCmd}
	cmdCount := len(cmds)

	if !isInserting {
		for l, cmd := range cmds {
			cmdRunes := []rune(cmd)
			for i := 0; i < len(cmdRunes); i++ {
				termbox.SetCell(i, l, cmdRunes[i], termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	} else {
		insertRunes := []rune(insertText)
		for i, insertRune := range insertRunes {
			termbox.SetCell(i, 0, insertRune, termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	rows := strings.Split(bstStr, "\n")
	for i, row := range rows {
		runes := []rune(row)
		for j, c := range runes {
			termbox.SetCell(j, i+cmdCount, c, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	termbox.Flush()
}
