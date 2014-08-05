package gocamp

type WorldRepresenter interface {
	GetTerrain() *Terrain
	GetEntityManager() *EntityManager

	// TODO: Implement TaskManager
	//	GetTaskManager() *TaskManager

	// TODO: Implement RegionManager
	// GetRegionManager()
}
