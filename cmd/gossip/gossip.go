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

package gossip

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	keyAddress = "gossip_address"
)

// Server is the main struct of the gossip package.
type Server struct {
	// metadata is the metadata of the server.
	metadata ClusterMetadata
	// listener is the listener of the server.
	listener net.Listener
	// id is the id of the gossip server.
	id string
	// seedNodes are the initial seed nodes of the gossip server.
	seedNodes []string
	lock      sync.RWMutex
}

// NewServer creates a new gossip server.
func NewServer(seedNodes []string) *Server {
	return &Server{
		metadata:  make(ClusterMetadata),
		seedNodes: seedNodes,
	}
}

// addLocalState adds a key value pair to the local node state
func (s *Server) addLocalState(key string, value string) {
	if s.metadata[s.id] == nil {
		s.metadata[s.id] = make(map[string]*VersionedStr)
	}
	if s.metadata[s.id][key] == nil {
		s.metadata[s.id][key] = newVersionedStr(value)
	} else {
		s.metadata[s.id][key].set(value)
	}
}

// doGossip performs the gossip protocol
func (s *Server) doGossip() {
	liveNodes := s.liveNodes()
	if len(liveNodes) == 0 {
		// seed
		s.gossip(randomNode(s.seedNodes))
	} else {
		// gossip
		s.gossip(randomNode(liveNodes))
	}
}

// randomNode selects a random node from the list of given nodes
func randomNode(nodes []string) string {
	if len(nodes) == 0 {
		return ""
	}
	i := rand.Intn(len(nodes))
	return nodes[i]
}

// gossip performs gossip with the given node
func (s *Server) gossip(node string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if node == "" {
		return
	}
	jsonBytes, err := json.Marshal(s.metadata)
	if err != nil {
		fmt.Println("error marshalling metadata", err)
		return
	}
	resp, err := http.Post(fmt.Sprintf("http://%s/gossip", node), "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("error sending gossip to", node, err)
		return
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body", err)
		return
	}

	var newMetadata ClusterMetadata
	err = json.Unmarshal(respBytes, &newMetadata)
	if err != nil {
		fmt.Println("error unmarshalling response", err)
		return
	}

	merge(s.metadata, newMetadata)
	printMetadata(s.metadata)

}

// liveNodes returns a list of nodes that are alive
// except the local node
func (s *Server) liveNodes() []string {
	var result []string
	for nId, nState := range s.metadata {
		if nId == s.id {
			continue
		}
		if nAddr, ok := nState[keyAddress]; ok {
			result = append(result, nAddr.Value)
		}
	}
	return result
}

// Start starts the gossip server
func (s *Server) Start(ctx context.Context, addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.listener = l
	s.id = l.Addr().String()
	s.addLocalState(keyAddress, s.id)
	go func() {
		if err := http.Serve(l, s); err != http.ErrServerClosed {
			return
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				s.doGossip()
				time.Sleep(time.Second * 1)
			}
		}
	}()
	<-ctx.Done()
	if err := s.listener.Close(); err != nil {
		return err
	}
	return nil
}

// ServeHTTP handles http requests
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "invalid content type", http.StatusBadRequest)
		return
	}

	if r.URL.Path == "/gossip" && r.Method == http.MethodPost {
		var gossipRequest ClusterMetadata
		if err := json.NewDecoder(r.Body).Decode(&gossipRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.lock.Lock()
		defer s.lock.Unlock()
		merge(s.metadata, gossipRequest)
		printMetadata(s.metadata)

		if err := json.NewEncoder(w).Encode(s.metadata); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else if r.URL.Path == "/state" && r.Method == http.MethodGet {
		s.lock.RLock()
		defer s.lock.RUnlock()
		if err := json.NewEncoder(w).Encode(s.metadata); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.URL.Path == "/" && r.Method == http.MethodPost {
		var req map[string]string
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s.lock.Lock()
		defer s.lock.Unlock()
		for key, value := range req {
			println("Setting key", key, "to", value)
			s.addLocalState(key, value)
		}
	}

}

// printMetadata prints the node metadata
func printMetadata(metadata ClusterMetadata) {
	jsonBytes, _ := json.Marshal(metadata)
	fmt.Println(string(jsonBytes))
}

// Address returns the address of the server
func (s *Server) Address() string {
	return s.listener.Addr().String()
}
