package ui

import "github.com/a-h/templ"

type Option interface {
	GetID() string
	GetLabel() string
}

type IconOption struct {
	root      Option
	startIcon string
	endIcon   string
}

func (i *IconOption) StartIcon(icon string) *IconOption {
	i.startIcon = icon
	return i
}

func (i *IconOption) EndIcon(icon string) *IconOption {
	i.endIcon = icon
	return i
}

func (i *IconOption) GetLabel() string {
	return i.root.GetLabel()
}

func (i *IconOption) GetID() string {
	return i.root.GetID()
}

func IconOpt(root Option) *IconOption {
	return &IconOption{
		root: root,
	}
}

type ActionOption struct {
	id     string
	label  string
	action string
}

func (i *ActionOption) GetLabel() string {
	return i.label
}

func (i *ActionOption) GetID() string {
	return i.id
}

func ActionOpt(label string) *ActionOption {
	return &ActionOption{
		id:    "",
		label: label,
	}
}

func (i *ActionOption) Action(action string) *ActionOption {
	i.action = action
	return i
}

type LinkOption struct {
	id    string
	label string
	link  templ.SafeURL
}

func (i *LinkOption) GetLabel() string {
	return i.label
}

func (i *LinkOption) GetID() string {
	return i.id
}

func (i *LinkOption) Link(link templ.SafeURL) *LinkOption {
	i.link = link
	return i
}

func LinkOpt(label string) *LinkOption {
	return &LinkOption{
		label: label,
	}
}

type serializedOption struct {
	Id    string `json:"id"`
	Label string `json:"label"`
}

func toSerializedOptions(options []Option) []serializedOption {
	result := []serializedOption{}

	for _, option := range options {
		result = append(result, serializedOption{
			Id:    option.GetID(),
			Label: option.GetLabel(),
		})
	}

	return result
}
