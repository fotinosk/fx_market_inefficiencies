## Some notes on graph traversal


### Problem specification
- This is a weighted directed grap
- We are interested in find the max difference between the longest and shorted path (no in terms of nodes visited but distance)
- Distance is not addition but multiplication
- We do not allow for nodes to be revisited to avoid infinite loops
- The graph is dense, ie most nodes are connected to most other nodes
- No starting and ending point is 


### Notes
- If max distance A->B is x and distance B->C is y and the nodes visited in the path B->C are a disjoint set of the nodes visited in the best path of A->B then A->C is the best path with distance x*y -- **Is this true?**
- Treat complements 

- They have to be cycles but not contain cycles? Ie A->B->A