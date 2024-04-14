import random

class CellType:
    """Enum for different types of cells in the maze."""
    WALL = 0
    ROAD = 1
    ENTRANCE = 2
    EXIT = 3
    TRAP = 4
    TREASURE = 5


class MazeGenerator:
    """Class to generate a maze."""
    def __init__(self, size):
        """Initialize the maze generator with a specified size."""
        self.size = size
        self.maze = [[CellType.WALL for _ in range(size)] for _ in range(size)]
        self.directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]

    def generate_maze(self):
        """Generate the maze using recursive backtracking algorithm."""
        self._recursive_backtracker(1, 1)

        # Place entrance and exit
        self.maze[0][random.randint(1, self.size - 2)] = CellType.ENTRANCE
        self.maze[self.size - 1][random.randint(1, self.size - 2)] = CellType.EXIT

        # Place treasure and traps
        treasure_x, treasure_y = random.randint(1, self.size - 2), random.randint(1, self.size - 2)
        self.maze[treasure_x][treasure_y] = CellType.TREASURE
        num_traps = random.randint(0, 5)
        for _ in range(num_traps):
            trap_x, trap_y = random.randint(1, self.size - 2), random.randint(1, self.size - 2)
            self.maze[trap_x][trap_y] = CellType.TRAP

        return self.maze

    def _recursive_backtracker(self, x, y):
        """Recursive function to carve paths in the maze."""
        self.maze[x][y] = CellType.ROAD

        random.shuffle(self.directions)
        for dx, dy in self.directions:
            nx, ny = x + 2 * dx, y + 2 * dy
            if 0 <= nx < self.size and 0 <= ny < self.size and self.maze[nx][ny] == CellType.WALL:
                self.maze[x + dx][y + dy] = CellType.ROAD
                self._recursive_backtracker(nx, ny)

def display_maze(maze):
    """Display the maze."""
    for row in maze:
        for cell in row:
            if cell == CellType.WALL:
                print("█", end="")
            elif cell == CellType.ROAD:
                print(" ", end="")
            elif cell == CellType.ENTRANCE:
                print("E", end="")
            elif cell == CellType.EXIT:
                print("X", end="")
            elif cell == CellType.TREASURE:
                print("O", end="")
            elif cell == CellType.TRAP:
                print("T", end="")
        print()

if __name__ == "__main__":
    size = int(input("Enter the size of the maze (odd number): "))
    if size % 2 == 0:
        print("Size must be an odd number!")
    else:
        maze_generator = MazeGenerator(size)
        maze = maze_generator.generate_maze()
        display_maze(maze)
