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

package cmd

import (
	"dsa/cmd/mazegen"
	"github.com/spf13/cobra"
)

var mazeWidth int
var mazeHeight int

// mazeCmd represents the maze command
var mazeCmd = &cobra.Command{
	Use:   "maze",
	Short: "Simple maze generator",
	Long:  `Generates random mazes using the Kruskal's algorithm.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if mazeWidth <= 0 {
			mazeWidth = 20
		}
		if mazeHeight <= 0 {
			mazeHeight = 20
		}
		return mazegen.Run(mazeWidth, mazeHeight)
	},
}

func init() {
	rootCmd.AddCommand(mazeCmd)
	mazeCmd.Flags().IntVar(&mazeWidth, "width", 30, "Width of the maze")
	mazeCmd.Flags().IntVar(&mazeHeight, "height", 30, "Height of the maze")
}
