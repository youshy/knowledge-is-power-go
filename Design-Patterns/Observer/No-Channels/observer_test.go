package observer

import "testing"

func Test_Observer(t *testing.T) {
	watchTower := WatchTower{}
	soldier1 := Soldier{id: 1, zone: "B"}
	soldier2 := Soldier{id: 2, zone: "A"}

	watchTower.Add(soldier1)
	watchTower.Add(soldier2)

	// Notify Zone A
	watchTower.Notify("A")

	// Notify Zone B
	watchTower.Notify("B")

	// Remove soldier 1 (No soldier in zone B)
	watchTower.Remove(soldier1)

	// Notify zone B (as it's free)
	watchTower.Notify("B")
}
