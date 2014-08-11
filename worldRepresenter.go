package gocamp

type WorldRepresenter interface {
	GetTerrain() *Terrain
	EntityManager() EntityManager

	// TODO: Implement TaskManager
	//	GetTaskManager() *TaskManager

	// TODO: Implement RegionManager
	// GetRegionManager()

	// Get a map of the world as a representation of passable terrain
	// Primarily used for pathfinding
	GetTraversableMap(int) [][]int
}
