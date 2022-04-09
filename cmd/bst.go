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
	"strings"
	"time"
)

var backbuf []termbox.Cell
var bbw, bbh int
var node *bst2.Node
var selected *bst2.Node

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

		ints := uniqueRandomInts()
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
						break loop
					} else {

						if selected == nil {
							selected = node
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

func uniqueRandomInts() []int {
	var m = map[int]struct{}{}
	var ints []int
	for i := 0; i < 100; i++ {
		r := rand.Intn(100)
		if _, ok := m[r]; !ok {
			ints = append(ints, r)
			m[r] = struct{}{}
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
	cmds := []string{upCmd, leftCmd, rightCmd, nextCmd, prevCmd}
	cmdCount := len(cmds)

	for l, cmd := range cmds {
		cmdRunes := []rune(cmd)
		for i := 0; i < len(cmdRunes); i++ {
			termbox.SetCell(i, l, cmdRunes[i], termbox.ColorDefault, termbox.ColorDefault)
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
