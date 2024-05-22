package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Farm yapısı, karınca sayısı, odalar, bağlantılar, başlangıç ve bitiş noktalarını tutar
type Farm struct {
	AntAmount int
	Rooms     map[string][2]int
	Links     []string
	Start     string
	End       string
}

// Node yapısı, oda ismi ve koordinatlarını tutar
type Node struct {
	Name        string
	Coordinates [2]int
}

// Edge yapısı, başlangıç ve bitiş noktalarını ve ağırlığını tutar
type Edge struct {
	Start  *Node
	End    *Node
	Weight int
}

// Graph yapısı, düğümler ve kenarları tutar
type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func main() {
	// Komut satırı argüman kontrolü
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . example.txt")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Dosyayı okuma ve farm bilgilerini alma
	farminfo, err := ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Dosya içeriğini okuma ve ekrana yazdırma
	error := ReadFileAndPrintContent(filename)
	if error != nil {
		fmt.Println("Error reading file:", error)
		return
	}

	// Odalar ve bağlantılardan grafik oluşturma
	graph := CreateGraphFromRoomsAndLinks(farminfo.Rooms, farminfo.Links)

	// Başlangıç ve bitiş noktaları arasındaki tüm yolları bulma
	paths := FindAllPaths(*graph, farminfo.Start, farminfo.End)

	// Yolları uzunluklarına göre sıralama
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	var bestAntPaths [][]string
	minLength := -1

	// En iyi karınca yollarını bulma
	for i := 0; i < len(paths); i++ {
		filterPaths := FilterUniquePaths(paths, i)
		antPaths := MoveAnts(filterPaths, farminfo.AntAmount, farminfo.Start, farminfo.End)
		if minLength == -1 || len(antPaths) < minLength {
			minLength = len(antPaths)
			bestAntPaths = antPaths
		}
	}

	// En iyi yolları ekrana yazdırma
	PrintAntPaths(bestAntPaths)
}

// Dosyayı okur ve Farm yapısına dönüştürür
func ReadFile(filename string) (farminfo Farm, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return farminfo, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var links []string
	rooms := make(map[string][2]int)
	coordinates := make(map[[2]int]string)
	edgeSet := make(map[string]bool) // Bağlantıların benzersizliğini kontrol etmek için bir küme

	for i, line := range lines {
		fields := strings.Split(line, " ")
		if len(fields) > 3 {
			return farminfo, fmt.Errorf("invalid line format: %s", line)
		}

		if line == "##start" {
			start := strings.Fields(lines[i+1])
			if len(start) != 3 {
				return farminfo, fmt.Errorf("invalid start line format: %s", lines[i+1])
			}
			farminfo.Start = start[0]
		} else if line == "##end" {
			end := strings.Fields(lines[i+1])
			if len(end) != 3 {
				return farminfo, fmt.Errorf("invalid end line format: %s", lines[i+1])
			}
			farminfo.End = end[0]
		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if parts[0] == parts[1] {
				return farminfo, fmt.Errorf("self link found: %s", line)
			}
			edge := fmt.Sprintf("%s-%s", parts[0], parts[1])
			reverseEdge := fmt.Sprintf("%s-%s", parts[1], parts[0])

			// Bağlantıların benzersiz olup olmadığını kontrol et
			if edgeSet[edge] || edgeSet[reverseEdge] {
				return farminfo, fmt.Errorf("duplicate link found: %s", line)
			}

			// Bağlantıları ekle
			links = append(links, line)
			edgeSet[edge] = true
			edgeSet[reverseEdge] = true
		} else if !(strings.Contains(line, "#")) && len(fields) == 3 {
			x, err := strconv.Atoi(fields[1])
			if err != nil {
				return farminfo, fmt.Errorf("invalid coordinate: %s", fields)
			}
			y, err := strconv.Atoi(fields[2])
			if err != nil {
				return farminfo, fmt.Errorf("invalid coordinate: %s", fields)
			}
			coords := [2]int{x, y}
			if existingRoom, exists := coordinates[coords]; exists {
				return farminfo, fmt.Errorf("duplicate coordinates: %s and %s have the same coordinates %v", fields[0], existingRoom, coords)
			}
			coordinates[coords] = fields[0]
			rooms[fields[0]] = coords
		}

		// Oda isimlerinin geçerliliğini kontrol et
		if len(fields) == 3 && RuneVarMi(fields[0]) {
			return farminfo, fmt.Errorf("invalid room name: %s", fields)
		}
	}

	// Karınca miktarını kontrol et
	if len(lines) == 0 {
		return farminfo, fmt.Errorf("no data found in the file")
	}

	antamount, err := strconv.Atoi(lines[0])
	if err != nil {
		return farminfo, fmt.Errorf("invalid ant amount: %s", lines[0])
	} else if antamount <= 0 {
		return farminfo, fmt.Errorf("invalid ant amount: %d", antamount)
	}

	farminfo.AntAmount = antamount
	farminfo.Rooms = rooms
	farminfo.Links = links
	return farminfo, nil
}

// Dosyayı okur ve içeriğini ekrana yazdırır
func ReadFileAndPrintContent(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	lines := string(content)
	fmt.Println(lines)
	fmt.Println()

	return nil
}

// Oda isimlerinin geçerliliğini kontrol eden yardımcı fonksiyon
func RuneVarMi(s string) bool {
	return s[0] == '#' || s[0] == 'L'
}

// Odalar ve bağlantılardan graf oluşturur
func CreateGraphFromRoomsAndLinks(rooms map[string][2]int, links []string) *Graph {
	nodes := make(map[string]*Node)
	var edges []*Edge

	// Odalardan düğümler oluşturma
	for room, coords := range rooms {
		node := &Node{Name: room, Coordinates: coords}
		nodes[room] = node
	}

	// Bağlantılardan kenarlar oluşturma
	for _, conn := range links {
		parts := strings.Split(conn, "-")
		if len(parts) != 2 {
			continue
		}
		startNode := nodes[parts[0]]
		endNode := nodes[parts[1]]
		if startNode != nil && endNode != nil {
			edge1 := &Edge{Start: startNode, End: endNode, Weight: 1}
			edge2 := &Edge{Start: endNode, End: startNode, Weight: 1}
			edges = append(edges, edge1, edge2)
		}
	}

	// Graf yapısını oluşturma
	graph := &Graph{}
	for _, node := range nodes {
		graph.Nodes = append(graph.Nodes, node)
	}
	graph.Edges = edges

	return graph
}

// Başlangıç ve bitiş noktaları arasındaki tüm yolları bulur
func FindAllPaths(graph Graph, start string, end string) [][]string {
	visited := make(map[string]bool)
	paths := [][]string{}
	findPathsRecursive(graph, start, end, []string{start}, visited, &paths)
	return paths
}

// Rekürsif olarak yolları bulan yardımcı fonksiyon
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

// Benzersiz yolları filtreler
func FilterUniquePaths(paths [][]string, index int) [][]string {
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
			if PathContainsElement(path, elem) {
				containsAny = true
				break
			}
		}

		if !containsAny {
			uniquePaths = append(uniquePaths, path)
			newMidElements := path[1 : len(path)-1]
			for _, newElem := range newMidElements {
				if !PathContainsElement(midElements, newElem) {
					midElements = append(midElements, newElem)
				}
			}
		}
	}
	return uniquePaths
}

// Yolun belirli bir elemanı içerip içermediğini kontrol eder
func PathContainsElement(path []string, elem string) bool {
	for _, item := range path {
		if item == elem {
			return true
		}
	}
	return false
}

// Karıncaları odalar arasında hareket ettirir
func MoveAnts(paths [][]string, antAmount int, startRoom string, endRoom string) [][]string {
	moves := [][]string{}
	antPositions := make([]int, antAmount)
	antPaths := make([][]string, antAmount)
	occupiedRooms := make(map[int]map[string]bool)
	occupiedEdges := make(map[int]map[string]bool)

	totalPathLength := 0
	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path)
		totalPathLength += len(path)
	}

	pathWeights := make([]float64, len(paths))
	antsPerPath := make([]int, len(paths))
	for i, length := range pathLengths {
		pathWeights[i] = float64(totalPathLength) / float64(length)
	}

	totalWeight := 0.0
	for _, weight := range pathWeights {
		totalWeight += weight
	}

	remainingAnts := antAmount
	for i := 0; i < len(paths); i++ {
		antsPerPath[i] = 1
		remainingAnts--
	}

	for remainingAnts > 0 {
		minIndex := -1
		minValue := int(^uint(0) >> 1)
		for i := 0; i < len(paths); i++ {
			value := (len(paths[i]) - 2) + antsPerPath[i]
			if value < minValue {
				minValue = value
				minIndex = i
			}
		}
		antsPerPath[minIndex]++
		remainingAnts--
	}

	antIndex := 0
	for i, ants := range antsPerPath {
		for j := 0; j < ants; j++ {
			antPaths[antIndex] = paths[i]
			antIndex++
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

// Karınca hareketlerini ekrana yazdırır
func PrintAntPaths(antPaths [][]string) {
	for _, step := range antPaths {
		fmt.Println(strings.Join(step, " "))
	}
}
