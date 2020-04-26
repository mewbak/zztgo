// ZZT ported to Go

package main

import (
	"math/rand"
	"time"
)

func main() {
	WorldFileDescCount = 7
	WorldFileDescKeys[0] = "TOWN"
	WorldFileDescValues[0] = "TOWN       The Town of ZZT"
	WorldFileDescKeys[1] = "DEMO"
	WorldFileDescValues[1] = "DEMO       Demo of the ZZT World Editor"
	WorldFileDescKeys[2] = "CAVES"
	WorldFileDescValues[2] = "CAVES      The Caves of ZZT"
	WorldFileDescKeys[3] = "DUNGEONS"
	WorldFileDescValues[3] = "DUNGEONS   The Dungeons of ZZT"
	WorldFileDescKeys[4] = "CITY"
	WorldFileDescValues[4] = "CITY       Underground City of ZZT"
	WorldFileDescKeys[5] = "BEST"
	WorldFileDescValues[5] = "BEST       The Best of ZZT"
	WorldFileDescKeys[6] = "TOUR"
	WorldFileDescValues[6] = "TOUR       Guided Tour ZZT's Other Worlds"

	rand.Seed(time.Now().UnixNano())

	StartupWorldFileName = "TOWN"
	ResourceDataFileName = "ZZT.DAT"
	GameTitleExitRequested = false
	EditorEnabled = false

	VideoInstall()
	TextWindowInit(5, 3, 50, 18)
	New(IoTmpBuf)
	VideoHideCursor()
	VideoClrScr()
	TickSpeed = 4
	DebugEnabled = false
	SavedGameFileName = "SAVED"
	SavedBoardFileName = "TEMP"
	GenerateTransitionTable()
	WorldCreate()

	GameTitleLoop()

	SoundUninstall()
	SoundClearQueue()
	VideoUninstall()
}
