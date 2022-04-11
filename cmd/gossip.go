/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"dsa/cmd/gossip"
	"github.com/spf13/cobra"
)

var gossipAddr string
var gossipSeed []string

// gossipCmd represents the gossip command
var gossipCmd = &cobra.Command{
	Use:   "gossip",
	Short: "Simple implementation of gossip protocol",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		srv := gossip.NewServer(gossipSeed)
		go func() {
			if err := srv.Start(ctx, gossipAddr); err != nil {
				println(err.Error())
			}
		}()
		<-ctx.Done()
	},
}

func init() {
	rootCmd.AddCommand(gossipCmd)
	gossipCmd.Flags().StringVar(&gossipAddr, "addr", ":8080", "gossip address")
	gossipCmd.Flags().StringSliceVar(&gossipSeed, "seed", []string{}, "gossip seed")
}
