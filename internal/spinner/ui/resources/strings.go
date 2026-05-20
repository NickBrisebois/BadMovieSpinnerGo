package res

import "fmt"

func strikethrough(s string) string {
	// returns a unicode strikethrough version of the given string
	result := ""
	for _, r := range s {
		result += string(r) + "\u0336"
	}
	return result
}

var (
	// ====== GLOBAL HEADER ======
	StrHeaderTitle    = "Astounding, Phenomenal, Jaw-Dropping and truly Stupendous Movie Spinner"
	StrHeaderSubtitle = ""

	// ====== SIDEBAR ======
	StrSidebarHeaderTitle = fmt.Sprintf("Spin to %s Lose", strikethrough("Win"))
	StrSidebar            = "Click the button to spin the wheel and get a random bad movie to watch :)"

	StrSidebarNamesSubheaderTitle         = "Respected Movie Connoisseurs"
	StrSidebarAllMoviesSubheaderTitle     = "All Movies"
	StrSidebarWatchedMoviesSubheaderTitle = "Watched Movies"
)
