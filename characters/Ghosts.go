package characters

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This method means type NPC implements the interface CharacterFunctions,
// but I don't need to explicitly declare that it does so.
type NPC struct {
	Attributes *Character
}

func (npc *NPC) Init(pos mazegrid.Position) {
	npc.Attributes.Init(pos)
}

func (npc *NPC) GetPosition() mazegrid.Position {
	return npc.Attributes.GetPosition()
}

func (npc *NPC) SetPosition(pos mazegrid.Position) {
	npc.Attributes.SetPosition(pos)
}

func (npc *NPC) GetAlgo() int {
	return npc.Attributes.GetAlgo()
}

func (npc *NPC) GetFrameProperties() FrameProperties {
	return npc.Attributes.GetFrameProperties()
}

func (npc *NPC) SetFrameProperties(fp FrameProperties) {
	npc.Attributes.SetFrameProperties(fp)
}

func (npc *NPC) UpdateCount() {
	npc.Attributes.Count += 1
}

func (npc *NPC) GetCount() int {
	return npc.Attributes.GetCount()
}

func (npc *NPC) GetSprite() *ebiten.Image {
	return npc.Attributes.GetSprite()
}
