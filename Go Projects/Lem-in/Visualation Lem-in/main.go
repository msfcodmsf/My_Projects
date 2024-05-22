package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

type farm struct {
	antAmount int
	rooms     map[string][2]int
	links     []string
	start     string
	end       string
}

type Node struct {
	Name        string
	Coordinates [2]int
}

type Edge struct {
	Start  *Node
	End    *Node
	Weight int
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . example.txt")
		os.Exit(1)
	}

	filename := os.Args[1]
	farminfo, err := readfile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	graph := createGraphFromRoomsAndLinks(farminfo.rooms, farminfo.links)

	paths := findAllPaths(*graph, farminfo.start, farminfo.end)

	// Sort paths by length
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	// Initialize variables to track the best antPaths and its length
	var bestAntPaths [][]string
	minLength := -1

	for i := 0; i < len(paths); i++ {
		filterPaths := filterUniquePaths(paths, i)
		antPaths := moveAnts(filterPaths, farminfo.antAmount, farminfo.start, farminfo.end)
		fmt.Println(len(antPaths))
		// Update the best antPaths if the current one is shorter or if it's the first one
		if minLength == -1 || len(antPaths) < minLength {
			minLength = len(antPaths)
			bestAntPaths = antPaths
		}
	}

	// Çıktıyı dosyaya yaz
	file, err := os.Create("ant_output.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    for _, step := range bestAntPaths {
        fmt.Fprintln(file, strings.Join(step, " "))
    }

    // Create web server
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html")
    })
    http.HandleFunc("/output", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "ant_output.txt")
    })

    fmt.Println("Server started at http://localhost:8065...")
    http.ListenAndServe(":8065", nil)

    printAntPaths(bestAntPaths)
}

func filterUniquePaths(paths [][]string, index int) [][]string {
	firstPath := paths[index]
	midElements := make([]string, len(firstPath)-2)
	copy(midElements, firstPath[1:len(firstPath)-1])

	var uniquePaths [][]string
	uniquePaths = append(uniquePaths, firstPath)

	for i := 0; i < len(paths); i++ {
		if i == index {
			continue
		}
		path := paths[i]
		containsAny := false
		for _, elem := range midElements {
			if pathContainsElement(path, elem) {
				containsAny = true
				break
			}
		}

		if !containsAny {
			uniquePaths = append(uniquePaths, path)
			newMidElements := path[1 : len(path)-1]
			for _, newElem := range newMidElements {
				if !pathContainsElement(midElements, newElem) {
					midElements = append(midElements, newElem)
				}
			}
		}
	}
	return uniquePaths
}

func pathContainsElement(path []string, elem string) bool {
	for _, item := range path {
		if item == elem {
			return true
		}
	}
	return false
}

func readfile(filename string) (farminfo farm, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return farminfo, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var filelines []string
	for scanner.Scan() {
		filelines = append(filelines, scanner.Text())
	}

	var links []string
	rooms := make(map[string][2]int)
	for i, line := range filelines {
		fields := strings.Split(line, " ")
		if len(fields) > 3 {
			return farminfo, fmt.Errorf("invalid line format: %s", line)
		}
		if line == "##start" {
			start := strings.Fields(filelines[i+1])
			farminfo.start = start[0]
		} else if line == "##end" {
			end := strings.Fields(filelines[i+1])
			farminfo.end = end[0]
		} else if strings.Contains(line, "-") {
			links = append(links, line)
		} else if !(strings.Contains(line, "#")) && len(fields) == 3 {
			x, err := strconv.Atoi(fields[1])
			if err != nil {
				return farminfo, fmt.Errorf("invalid coordinate: %s", fields[1])
			}
			y, err := strconv.Atoi(fields[2])
			if err != nil {
				return farminfo, fmt.Errorf("invalid coordinate: %s", fields[2])
			}
			rooms[fields[0]] = [2]int{x, y}
		}
	}
	antamount, err := strconv.Atoi(filelines[0])
	if err != nil {
		return farminfo, fmt.Errorf("invalid ant amount: %s", filelines[0])
	}
	farminfo.antAmount = antamount
	farminfo.rooms = rooms
	farminfo.links = links
	return farminfo, nil
}

func createGraphFromRoomsAndLinks(rooms map[string][2]int, links []string) *Graph {
	nodes := make(map[string]*Node)
	var edges []*Edge

	for room, coords := range rooms {
		node := &Node{Name: room, Coordinates: coords}
		nodes[room] = node
	}

	for _, conn := range links {
		parts := strings.Split(conn, "-")
		startNode := nodes[parts[0]]
		endNode := nodes[parts[1]]
		if startNode != nil && endNode != nil {
			edge1 := &Edge{Start: startNode, End: endNode, Weight: 1}
			edge2 := &Edge{Start: endNode, End: startNode, Weight: 1}
			edges = append(edges, edge1, edge2)
		}
	}

	graph := &Graph{}
	for _, node := range nodes {
		graph.Nodes = append(graph.Nodes, node)
	}
	graph.Edges = edges

	return graph
}

func findAllPaths(graph Graph, start string, end string) [][]string {
	visited := make(map[string]bool)
	paths := [][]string{}
	findPathsRecursive(graph, start, end, []string{start}, visited, &paths)
	return paths
}

func findPathsRecursive(graph Graph, current string, end string, path []string, visited map[string]bool, paths *[][]string) {
	visited[current] = true
	if current == end {
		newPath := make([]string, len(path))
		copy(newPath, path)
		*paths = append(*paths, newPath)
	} else {
		for _, edge := range graph.Edges {
			if edge.Start.Name == current && !visited[edge.End.Name] {
				findPathsRecursive(graph, edge.End.Name, end, append(path, edge.End.Name), visited, paths)
			}
		}
	}
	visited[current] = false
}

func moveAnts(paths [][]string, antAmount int, startRoom string, endRoom string) [][]string {
	moves := [][]string{}
	antPositions := make([]int, antAmount)
	antPaths := make([][]string, antAmount)
	occupiedRooms := make(map[int]map[string]bool)
	occupiedEdges := make(map[int]map[string]bool)

	// Calculate total path lengths and weights
	totalPathLength := 0
	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path)
		totalPathLength += len(path)
	}

	// Determine the weight for each path
	pathWeights := make([]float64, len(paths))
	for i, length := range pathLengths {
		pathWeights[i] = float64(totalPathLength) / float64(length)
	}

	// Normalize the weights to allocate ants proportionally
	totalWeight := 0.0
	for _, weight := range pathWeights {
		totalWeight += weight
	}

	// Calculate number of ants for each path based on weights
	antsPerPath := make([]int, len(paths))
	for i, weight := range pathWeights {
		antsPerPath[i] = int(weight / totalWeight * float64(antAmount))
	}

	// Ensure total ants are assigned
	assignedAnts := 0
	for _, ants := range antsPerPath {
		assignedAnts += ants
	}
	for i := 0; assignedAnts < antAmount; i = (i + 1) % len(paths) {
		antsPerPath[i]++
		assignedAnts++
	}

	// Assign ants to paths
	antIndex := 0
	for i, ants := range antsPerPath {
		for j := 0; j < ants; j++ {
			antPaths[antIndex] = paths[i]
			antIndex += 1
		}
	}

	step := 0
	for {
		moveStep := []string{}
		allFinished := true

		if _, exists := occupiedRooms[step]; !exists {
			occupiedRooms[step] = make(map[string]bool)
		}
		if _, exists := occupiedEdges[step]; !exists {
			occupiedEdges[step] = make(map[string]bool)
		}

		for i := 0; i < antAmount; i++ {
			if antPositions[i] < len(antPaths[i])-1 {
				currentPosition := antPositions[i]
				nextPosition := currentPosition + 1
				currentRoom := antPaths[i][currentPosition]
				nextRoom := antPaths[i][nextPosition]
				edge := fmt.Sprintf("%s-%s", currentRoom, nextRoom)
				reverseEdge := fmt.Sprintf("%s-%s", nextRoom, currentRoom)

				if nextRoom != startRoom && nextRoom != endRoom && occupiedRooms[step][nextRoom] {
					continue
				}
				if occupiedEdges[step][edge] || occupiedEdges[step][reverseEdge] {
					continue
				}

				moveStep = append(moveStep, fmt.Sprintf("L%d-%s", i+1, nextRoom))
				occupiedRooms[step][nextRoom] = true
				occupiedEdges[step][edge] = true
				occupiedEdges[step][reverseEdge] = true
				antPositions[i]++
				allFinished = false
			}
		}

		if allFinished {
			break
		}

		moves = append(moves, moveStep)
		step++
	}

	return moves
}

// en kısa tur sayısını bulan fonksiyondan bilgileri alır ve yazdırır.
func printAntPaths(antPaths [][]string) {
	for _, step := range antPaths {
		fmt.Println(strings.Join(step, " "))
	}
}
