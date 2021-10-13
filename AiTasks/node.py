import time
from collections import deque

class State():

    def __init__(self, value, move):
        self.value = value
        self.move = move

    def add_move(self, move):
        self.move.append(move)

    def is_equal_to(self, state):
        for i in range(len(self.value)):
            for j in range(len(self.value[i])):
                if self.value[i][j] != state.value[i][j]:
                    return False
        return True


    def find_space(self):
        # print("err",self.value.value)
        for i in range(len(self.value)):
            for j in range(len(self.value[i])):
                if self.value[i][j] == 0:
                    return i, j

    def print(self):
        for i in range(len(self.value)):
            print("|", end="")
            for j in range(len(self.value[i])):
                print(f"{self.value[i][j]}|", end="")
            print()
        print()

    def is_goal(self):
        goal = [[1, 2, 3], [4, 5, 6], [7, 8, 0]]
        return goal == self.value


from copy import deepcopy
def move(state, direction):
    i, j = state.find_space()
    if direction == "left" and j != 0:
        value = deepcopy(state.value)
        value[i][j], value[i][j - 1] = value[i][j - 1], value[i][j]
        moves = deepcopy(state.move)
        moves.append("left")
        s = State(value, moves)
        return s

    if direction == "right" and j != 2:
        value = deepcopy(state.value)
        value[i][j], value[i][j + 1] = value[i][j + 1], value[i][j]
        moves = deepcopy(state.move)
        moves.append("right")
        return State(value, moves)

    if direction == "up" and i != 0:
        value = deepcopy(state.value)
        value[i][j], value[i - 1][j] = value[i - 1][j], value[i][j]
        moves = deepcopy(state.move)
        moves.append("up")
        s = State(value, moves)
        return s

    if direction == "down" and i != 2:
        value = deepcopy(state.value)
        value[i][j], value[i + 1][j] = value[i + 1][j], value[i][j]
        moves = deepcopy(state.move)
        moves.append("down")
        s = State(value, moves)
        return s
    return None


def contains(visitedStates, state):
    for s in visitedStates:
        if s.is_equal_to(state):
            return True
    return False

def dfs(root):
    visitedStates = set()
    stack = [root]
    while len(stack) > 0:
        currentState = stack.pop()
        visitedStates.add(currentState)
        if currentState.is_goal():
            print("Goal Found")
            return currentState

        currentState.print()
        ng = [state for state in [move(currentState, "up"),
                                  move(currentState, "down"),
                                  move(currentState, "left"),
                                  move(currentState, "right")] if state is not None]

        for state in ng:
            if not contains(visitedStates, state) and not contains(stack, state):
                stack.append(state)

        print("depth: ", len(currentState.move))
        # time.sleep(1)


def bfs(root):
    visitedStates = set()
    queue = deque([root])
    while len(queue) > 0:
        currentState = queue.popleft()
        visitedStates.add(currentState)
        if currentState.is_goal():
            print("Goal Found")
            return currentState

        # currentState.print()
        ng = [state for state in [move(currentState, "up"),
                                  move(currentState, "down"),
                                  move(currentState, "left"),
                                  move(currentState, "right")] if state is not None]

        for state in ng:
            if not contains(visitedStates, state) and not contains(queue, state):
                queue.append(state)
        # time.sleep(1)
        print("depth: ",len(currentState.move))

if __name__ == '__main__':
    # root_value = [[1,2,0],[4,5,3],[7,8,6]]
    root_value = [[2,8,3],[1,6,4],[7,0,5]]
    # root_value = [[0,8,7],[6,5,4],[3,2,1]]
    root = State(root_value, [])
    # goal = dfs(root)
    goal = bfs(root)
    print(goal.move)
    goal.print()
