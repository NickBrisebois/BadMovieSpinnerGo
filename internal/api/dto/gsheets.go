package dto

const (
	ColSuggestedByIndex = iota
	ColMovieTitleIndex
	ColTMDBLinkIndex
	ColSheetsMoviePoster // ColSheetsMoviePoster is GSheets-specific and not useable here
	ColWatched
	ColBingoCard // ColBingoCard is GSheets-specific and not useable here
)

type GSheetsMoviesEntry struct {
	SuggestedBy string
	MovieTitle  string
	TMDBLink    string
	Watched     bool
}

func (g *GSheetsMoviesEntry) FromRowData(rowData []any) {
	g.SuggestedBy = rowData[ColSuggestedByIndex].(string)
	g.MovieTitle = rowData[ColMovieTitleIndex].(string)
	g.TMDBLink = rowData[ColTMDBLinkIndex].(string)
	g.Watched = strToBool(rowData[ColWatched].(string))
}

func strToBool(val string) bool {
	return val == "TRUE"
}
