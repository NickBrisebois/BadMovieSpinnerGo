package spinner

import (
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/data"
	"NickBrisebois/BadMovieSpinnerGo/internal/spinner/render"
	"NickBrisebois/BadMovieSpinnerGo/pkg/models"
	"image/color"
)

var (
	outlineColour = color.RGBA{0, 0, 0, 255}
	sliceColours  = [...]color.RGBA{
		{242, 145, 153, 255},
		{107, 167, 191, 255},
		{242, 242, 56, 255},
		{242, 90, 56, 255},
		{64, 1, 1, 255},
		{42, 45, 53, 255},
		{7, 67, 91, 255},
		{42, 242, 56, 255},
		{42, 90, 56, 255},
		{4, 1, 1, 255},
	}
)

type Spinner struct {
	DrawHandler *render.DrawHandler
	Slices      *[]data.Slice
}

func NewSpinner() *Spinner {
	return &Spinner{nil, nil}
}

func (s *Spinner) Init(centreX, centreY, radiusX, radiusY float32) {
	mockMovies := s.pullMovieList()
	sliceAngle := render.GetAnglePerSlice(len(mockMovies))

	s.Slices = s.genSlices(
		sliceAngle,
		centreX,
		centreY,
		radiusX,
		radiusY,
	)
	s.DrawHandler = render.NewDrawHandler(s.Slices, sliceAngle, centreX, centreY, radiusX, radiusY)
}

func (s *Spinner) pullMovieList() []models.MovieMeta {
	return []models.MovieMeta{
		{Title: "Movie 1", Link: "http://example.com/movie1", Watched: false, SuggestedBy: "user1", PosterURL: "https://http.cat/images/100.jpg"},
		{Title: "Movie 2", Link: "http://example.com/movie2", Watched: true, SuggestedBy: "user2", PosterURL: "https://http.cat/images/101.jpg"},
		{Title: "Movie 3", Link: "http://example.com/movie3", Watched: false, SuggestedBy: "user3", PosterURL: "https://http.cat/images/102.jpg"},
	}
}

func (s *Spinner) genSlices(sliceAngle, centreX, centreY, radiusX, radiusY float32) *[]data.Slice {
	movies := s.pullMovieList()
	slices := make([]data.Slice, 0, len(movies))

	for i, movie := range movies {
		slices = append(slices, data.Slice{
			ID:         i,
			Movie:      movie,
			Label:      movie.Title,
			FillColour: sliceColours[i%len(sliceColours)],
			DrawProperties: &data.DrawProperties{
				Step:  i,
				Start: float32(i) * sliceAngle,
				End:   (float32(i) + 1) * sliceAngle,
			},
		})
	}

	return &slices
}

func (s *Spinner) Update() {
	if s.Slices == nil {
		return
	}

	for i := range *s.Slices {
		drawProperties := (*s.Slices)[i].DrawProperties
		if drawProperties == nil {
			(*s.Slices)[i].DrawProperties = data.GetNextStateSliceDrawProperties(drawProperties)
		}
	}
}
