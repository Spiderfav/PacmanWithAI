package characters

// This method means type NPC implements the interface Character,
// but I don't need to explicitly declare that it does so.
type Player struct {
	Attributes NPC
}
