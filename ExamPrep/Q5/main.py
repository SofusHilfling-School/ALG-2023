from graph import Edge, Graph
from pq import IndexMinPrioityQueue

class Dijkstra:
    distTo: dict[int, float]
    edgeTo: dict[int, Edge]
    
    def __init__(self, g: Graph, start: int) -> None:
        self.edgeTo = {v: None for v in g.vertices.keys()}
        self.distTo = {v:float('inf') for v in g.vertices.keys()}
        self.distTo[start] = 0

        pq = IndexMinPrioityQueue()
        pq.push(start, 0)
        while not pq.empty():
            v = pq.pop()
            for e in g.edges[v]:
                w = e.to_vertex
                new_dist = self.distTo[v] + e.weight
                if self.distTo[w] > new_dist:
                    self.distTo[w] = new_dist
                    self.edgeTo[w] = e
                    pq.push(w, new_dist)


    ### Helper methods for displaying the result ###

    def hasPathTo(self, v: int) -> bool:
        return self.distTo[v] < float('inf')
    
    def pathTo(self, v: int) -> list[Edge] | None:
        if not self.hasPathTo(v):
            return None

        path: list[Edge] = []
        e = self.edgeTo[v]
        while e != None:
            path.append(e)
            e = self.edgeTo[e.from_vertex]

        return path

g = Graph()
g.addEdge(0, 1, 4)
g.addEdge(0, 2, 2)
g.addEdge(1, 2, 3)
g.addEdge(1, 3, 2)
g.addEdge(2, 3, 1)

g.addNode(4)

# print keys 
# print([k for k in g.vertices.keys()])


sp = Dijkstra(g, 0)
pathTo3 = sp.pathTo(3)
pathTo3.reverse()
print(f"path from 0 to 3: {[ x.to_vertex for x in pathTo3 ]}")
print(f"unconnected node: {sp.pathTo(4)}")