from typing import Callable


class Edge:
    """
    Edge constructer class til en graf
    
    Attributes
    ----------
    to_node : int
        den node som edgen går til
    from_node : int
        den node som edgen går fra
    distance_between : float
        den 'weight' som edgen har
    
    """
    to_vertex: int
    from_vertex: int
    weight: float = 1

    def __init__(self, to_node, from_node) -> None:
        """
        Parameters
        ----------
        to_node
            the node an edge goes to
        from_node
            the node an edge comes from
        """
        self.to_vertex = to_node
        self.from_vertex = from_node

    def __str__(self) -> str:
        return f"Edge(to_node={self.to_vertex}, from_node={self.from_vertex}, distance_between={self.weight})"

    def __repr__(self) -> str:
        return self.__str__()

class Vertex:
    def __init__(self, key: int, name: str | None) -> None:
        self.key = key
        self.name = name
    
    def __str__(self) -> str:
        return f"Node(key={self.key}, name='{self.name}')"

class Graph:
    """
    Graph constructer class 
    
    Attributes
    ----------
    edges : dict
        key/value pair, key er en int og value er en liste af edgesclass som noden har
    nodes : dict
        key/value pair, key er en int og value er node class

    Functions
    ---------
    addNode(key, int, node)

    addEdge(from_node, edge, distance)
        adds an edge from a given node to another with a given distance
    printGraph
        Prints a graph
    """
    edges: dict[int, list[Edge]]
    vertices: dict[int, Vertex]

    def __init__(self):
        self.edges = {}
        self.vertices = {}

    def addNode(self, v: int | Vertex) -> None:
        'add or update a node to the graph'
        if type(v) is Vertex:
            self.vertices[v.key] = v
            key = v.key
        elif type(v) is int:
            key = v
            self.vertices[v] = Vertex(v, None)
        
        if key not in self.edges:
            self.edges[key] = []

    def addEdge(self, from_vertex: int, edge: Edge | int, distance: float = 0) -> None:
        'add an edge from a given node to another with a given distance, add node if it doesn\'t exists'
        if type(edge) is int:
            edge = Edge(edge, from_vertex)
            # set distance for edge
            edge.weight = distance
        else:
            edge.from_vertex = from_vertex

        if from_vertex not in self.vertices:
            self.addNode(from_vertex)
        if edge.to_vertex not in self.vertices:
            self.addNode(edge.to_vertex)

        if from_vertex not in self.edges:
            self.edges[from_vertex] = [edge]
        else:
            self.edges[from_vertex].append(edge)

    def printGraph(self):
        for source, edges in self.edges.items():
            destinations = [str(e.to_vertex) for e in edges]
            
            print(f"{source}-->{', '.join(destinations)}")

    def getKey(self, predicate: Callable[[Vertex], bool]) -> Vertex | None:
        for k, v in self.vertices.items():
            if(predicate(v)):
                return k