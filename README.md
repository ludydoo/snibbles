# About

Just a few bobs and snibbles. Implementing
some data structures, challenges and games in Go and stuff.

## Data Structures

- [Binary Seach Tree (Generic)](pkg/bst)
- [Binary Seach Tree (Iterative)](pkg/bst2)
- [Hashmap](pkg/hashmap)
- [Heap](pkg/heap)
- [Linked List](pkg/linkedlist)
- [Queue](pkg/queue)
- [Stack](pkg/stack)
- [Disjoint Set](pkg/disjointset)

## Algorithms

- [Dijkstra's Algorithm](pkg/dijkstra)
- [Kruskal's Algorithm](pkg/kruskal)
- [Prim's Algorithm](pkg/prim)
- [Bellman-Ford Algorithm](pkg/bellmanford)

# Challenges

- [Remove Islands](pkg/challenges/islands)

## Games

### Binary Search Tree

`go run . bst`

#### Commands

```
up    - move up
left  - select left child
right - select right child
n     - in order successor
p     - in order predecessor
i     - insert a node
d     - delete the selected node
```

![BST](images/bst.gif)

### Maze Generator

`go run . maze --width 10 --height 10`

#### Commands

```
press any key to generate a new maze
press escape to exit
```

![BST](images/maze.gif)