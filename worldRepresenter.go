package gocamp

type WorldRepresenter interface {
	GetTerrain() *Terrain
	EntityManager() EntityManager

	// TODO: Implement TaskManager
	//	GetTaskManager() *TaskManager

	// TODO: Implement RegionManager
	// GetRegionManager()
}
