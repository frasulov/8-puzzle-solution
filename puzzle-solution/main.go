package main

import (
	"fmt"
	"time"
)

const (
	N = 3
)

type Stack []State

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str State) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (State, bool) {
	if s.IsEmpty() {
		return State{}, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}


type State struct {
	value [N][N]int
	moves []string
}

func NewState(value [N][N]int, move ...string) State {
	return State{
		value: value,
		moves: move,
	}
}



func (s State) AddMove(move string){
	s.moves = append(s.moves, move)
}

func (s State) GetNeighbours() []State {
	var states = make([]State, 0)
	i, j := s.findSpace()
	// left
	if j != 0 {
		value := s.value
		value[i][j], value[i][j-1] = value[i][j-1], value[i][j]
		states = append(states, NewState(value, append(s.moves, "left")...))
	}

	// right
	if j != 2 {
		value := s.value
		value[i][j], value[i][j+1] = value[i][j+1], value[i][j]
		states = append(states, NewState(value, append(s.moves, "right")...))
	}

	if i != 2 {
		value := s.value
		value[i][j], value[i+1][j] = value[i+1][j], value[i][j]
		states = append(states, NewState(value, append(s.moves, "down")...))
	}

	if i !=0 {
		value := s.value
		value[i][j], value[i-1][j] = value[i-1][j], value[i][j]
		states = append(states, NewState(value, append(s.moves, "up")...))
	}
	return states
}

func (s State) toString() string {
	str := ""
	for _, value := range s.value {
		for _, value2 := range value {
			str += fmt.Sprintf("%v", value2)
		}
	}
	return str
}

func (s State) findSpace() (int, int){
	for i, value := range s.value {
		for j, value2 := range value {
			if value2 == 0{
				return i,j
			}
		}
	}
	return -1,-1
}

func (s State) print() {
	for _, val := range s.value {
		fmt.Print("|")
		for _, value := range val {
			fmt.Printf("%v|", value)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (s State) isGoalState() bool {
	var goal = [3][3]int{
		{1,2,3},{4,5,6},{7,8,0},
	}
	for i, val := range s.value {
		for j, value := range val {
			if goal[i][j] != value {
				return false
			}
		}
	}
	return true
}

func (s State) isEqualTo(s2 State) bool{
	for i, val := range s.value {
		for j, value := range val {
			if s2.value[i][j] != value {
				return false
			}
		}
	}
	return true
}

func Contains(nodes []State, s State) bool{
	for _, n := range nodes {
		if n.isEqualTo(s){
			return true
		}
	}
	return false
}

func dfs(root State) State {
	visitedNodes := []State{}
	var stack Stack
	stack.Push(root)

	for !stack.IsEmpty() {
		currentNode, _ := stack.Pop()
		visitedNodes = append(visitedNodes, currentNode)

		if currentNode.isGoalState() {
			fmt.Print("Found\n")
			return currentNode
		}

		//neighbors := currentNode.GetNeighbours()


	}

	return State{}
}

func bfs(root State) State {
	visitedNodes := []State{}
	queue := []State{}
	queue = append(queue, root)
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		visitedNodes = append(visitedNodes, currentNode)

		if currentNode.isGoalState() {
			fmt.Print("Found\n")
			return currentNode
		}

		neighbourStates := currentNode.GetNeighbours()

		for _, state := range neighbourStates {
			if !Contains(visitedNodes, state) {
					queue = append(queue, state)
					visitedNodes = append(visitedNodes, state)
				}
		}
		currentNode.print()
		time.Sleep(time.Second*1)
	}
	return State{}
}

func main() {
	initial_value := [3][3]int{{0,8,7},{6,5,4},{3,2,1}}

	initial_state := State{
		value: initial_value,
	}

	resultState := bfs(initial_state)
	resultState.print()
}
