package constants

import (
	"github.com/maria-mz/bash-battle/config"
	"github.com/maria-mz/bash-battle/status"
)

/* TEXT MAPS */
var DifficultyTextMap = map[config.GameDifficulty]string{
	config.DiffVaried: "Varied",
	config.DiffEasy:   "Easy",
	config.DiffMedium: "Medium",
	config.DiffHard:   "Hard",
}

var FileSizeTextMap = map[config.GameFileSize]string{
	config.FileVaried: "Varied",
	config.FileSmall:  "Small",
	config.FileMedium: "Medium",
	config.FileBig:    "Large",
}

var StatusTextMap = map[status.GameStatus]string{
	status.Initializing:      "Initializing...",
	status.WaitingForPlayers: "Waiting for players...",
}

var ConnectionTextMap = map[status.ConnStatus]string{
	status.Connecting:   "Connecting...",
	status.Connected:    "Connected",
	status.Disconnected: "Disconnected",
}

/* TITLE */
const (
	BashCubeASCII = `
                  +++++++         
		       +++*     *+++      
            +++             *++*  
		 *+*                   *++
		 +                    ++++
		 +                *+++++++
		 +             *++++++++++
		 +           +++++++++++++
		 +          ++++++++++++++
		 +          ++++++++++++++
		 +          +++++*++++++++
		 ++*        +++++*++*+++++
		    +++     +++++++++++*
			   ++++ ++++++++
			      +++++++ 
	`
	BashBattleTitle = `
        __               __       __          __  __  __   
       / /_  ____ ______/ /_     / /_  ____ _/ /_/ /_/ /__ 
      / __ \/ __  / ___/ __ \   / __ \/ __  / __/ __/ / _ \
     / /_/ / /_/ (__  ) / / /  / /_/ / /_/ / /_/ /_/ /  __/
    /_.___/\__,_/____/_/ /_/  /_.___/\__,_/\__/\__/_/\___/ 
	`
	BashBattleWelcome = "echo \"Welcome to Bash Battle!\""

	WindowTitle = "Bash Battle"
)

/* PLAYERS TABLE */
const PlayersTableTitle = "Players"

/* CONFIG TABLE */
const (
	ConfigLabelServerAddr    = "Server address"
	ConfigLabelRounds        = "Number of rounds"
	ConfigLabelRoundDuration = "Round duration (seconds)"
	ConfigLabelDifficulty    = "Difficulty"
	ConfigLabelFileSize      = "File size"

	ConfigTableTitle = "Config"
)
