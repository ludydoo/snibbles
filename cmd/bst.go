/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dsa/pkg/bst2"
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
	"strconv"
	"strings"
	"time"
)

var node *bst2.Node
var selected *bst2.Node
var isInserting bool
var insertText string

// bstCmd represents the bst command
var bstCmd = &cobra.Command{
	Use:   "bst",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

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
	},
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

func init() {
	rootCmd.AddCommand(bstCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bstCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bstCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
