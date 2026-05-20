package sidebar

import (
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"

	"github.com/ebitenui/ebitenui/widget"
)

type SuggestedByToggleList struct {
	suggestedByList []string
	uiResources     *res.UIResources
	container       *widget.Container
}

func getCheckboxWidget(label string, changedHandler widget.CheckboxChangedHandlerFunc, uiRes *res.UIResources) *widget.Checkbox {
	return widget.NewCheckbox(
		widget.CheckboxOpts.Spacing(uiRes.Checkbox.Spacing),
		widget.CheckboxOpts.Image(uiRes.Checkbox.Image),
		widget.CheckboxOpts.StateChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
			if changedHandler != nil {
				changedHandler(args)
			}
		}),
		widget.CheckboxOpts.Text(
			label,
			uiRes.LabelResources.Face,
			uiRes.LabelResources.Text,
		),
	)
}

func NewSuggestedByToggle(suggestedByList []string, uiRes *res.UIResources) *SuggestedByToggleList {
	suggestedByListContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			VerticalPosition: widget.GridLayoutPositionStart,
		})),
	)

	for _, suggestedBy := range suggestedByList {
		checkbox := getCheckboxWidget(suggestedBy, nil, uiRes)
		suggestedByListContainer.AddChild(checkbox)
	}

	return &SuggestedByToggleList{
		suggestedByList: suggestedByList,
		container:       suggestedByListContainer,
	}
}

func (s *SuggestedByToggleList) GetContainer() *widget.Container {
	return s.container
}
