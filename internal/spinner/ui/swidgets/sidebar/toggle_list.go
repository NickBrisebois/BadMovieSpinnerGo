package sidebar

import (
	res "NickBrisebois/BadMovieSpinnerGo/internal/spinner/ui/resources"

	"github.com/ebitenui/ebitenui/widget"
)

type ToggleCallback func(toggled []string, args *widget.CheckboxChangedEventArgs)

type SuggestedByToggleList struct {
	suggestedByList []string
	uiResources     *res.UIResources
	container       *widget.Container
	toggles         []*widget.Checkbox
}

func NewSuggestedByToggle(suggestedByList []string, uiRes *res.UIResources, toggleCallback ToggleCallback) *SuggestedByToggleList {
	suggestedByListContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.GridLayoutData{
			VerticalPosition: widget.GridLayoutPositionStart,
		})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Stretch([]bool{true, true}, nil),
			widget.GridLayoutOpts.Spacing(10, 10),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(10)),
		)),
	)

	toggleListWidget := &SuggestedByToggleList{
		suggestedByList: suggestedByList,
	}

	toggles := make([]*widget.Checkbox, 0, len(suggestedByList))
	for _, suggestedBy := range suggestedByList {
		checkbox := toggleListWidget.getCheckboxWidget(
			suggestedBy,
			uiRes,
			toggleCallback,
		)
		toggles = append(toggles, checkbox)
	}

	for _, toggle := range toggles {
		suggestedByListContainer.AddChild(toggle)
	}

	toggleListWidget.container = suggestedByListContainer
	toggleListWidget.toggles = toggles
	return toggleListWidget
}

func (s *SuggestedByToggleList) getToggled() []string {
	var toggled []string
	for _, toggle := range s.toggles {
		if toggle.State() == widget.WidgetChecked {
			toggled = append(toggled, toggle.Text().Label)
		}
	}
	return toggled
}

func (s *SuggestedByToggleList) getCheckboxWidget(label string, uiRes *res.UIResources, toggleCallback ToggleCallback) *widget.Checkbox {
	return widget.NewCheckbox(
		widget.CheckboxOpts.Spacing(uiRes.Checkbox.Spacing),
		widget.CheckboxOpts.Image(uiRes.Checkbox.Image),
		widget.CheckboxOpts.StateChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
			toggleCallback(s.getToggled(), args)
		}),
		widget.CheckboxOpts.Text(
			label,
			uiRes.LabelResources.Face,
			uiRes.LabelResources.Text,
		),
		widget.CheckboxOpts.InitialState(widget.WidgetChecked),
	)
}

func (s *SuggestedByToggleList) GetContainer() *widget.Container {
	return s.container
}
