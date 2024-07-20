package ui

import (
	"fmt"

	"github.com/a-h/templ"
)

type VariantComponent[T templ.Component] struct {
	Component[T]
	variant            Variant
	calculatedVariance *Variance
}

func (c *VariantComponent[T]) variance() Variance {
	if c.calculatedVariance != nil {
		return *c.calculatedVariance
	}

	v := &Variance{}
	v.Of(c.variant)
	c.calculatedVariance = v
	return *v
}

func (c *VariantComponent[T]) Variant(variant Variant) T {
	c.variant |= variant
	return c.self
}

type Variant uint16

func (v Variant) String() string {
	switch v {
	case ColorPrimary:
		return "ColorPrimary"
	case ColorSecondary:
		return "ColorSecondary"
	case SizeHuge:
		return "SizeHuge"
	case SizeLarge:
		return "SizeLarge"
	case SizeMedium:
		return "SizeMedium"
	case SizeSmall:
		return "SizeSmall"
	case SizeTiny:
		return "SizeTiny"
	case FormDisabled:
		return "FormDisabled"
	case FormEnabled:
		return "FormEnabled"
	case ElementLoading:
		return "ElementLoading"
	case ElementLoaded:
		return "ElementLoaded"
	case UndefinedVariant:
		return "Undefined"
	default:
		return fmt.Sprintf("Composed(%d)", v)
	}
}

type Variance struct {
	Color   Variant
	Size    Variant
	Form    Variant
	Loading Variant
}

func (v *Variance) Of(variant Variant) {
	v.Size = Variant(variant & 31)
	if v.Size == UndefinedVariant {
		v.Size = SizeMedium
	}
	v.Color = Variant(variant & 224)
	if v.Color == UndefinedVariant {
		v.Color = ColorPrimary
	}
	v.Form = Variant(variant & 768)
	if v.Form == UndefinedVariant {
		v.Form = FormEnabled
	}
	v.Loading = Variant(variant & 3072)
}

const (
	ElementLoading Variant = 2048
	ElementLoaded  Variant = 1024
)

const (
	FormEnabled  Variant = 512
	FormDisabled Variant = 256
)

const (
	ColorPrimary   Variant = 128
	ColorSecondary Variant = 64
	ColorSubtle    Variant = 32
)

const (
	SizeHuge   Variant = 16
	SizeLarge  Variant = 8
	SizeMedium Variant = 4
	SizeSmall  Variant = 2
	SizeTiny   Variant = 1
)

const UndefinedVariant Variant = 0

type Radius string

const (
	RoundedSmall  Radius = "rounded-sm"
	Rounded       Radius = "rounded"
	RoundedMedium Radius = "rounded-md"
	RoundedLarge  Radius = "rounded-lg"
	RoundedXLarge Radius = "rounded-xl"
	RoundedFull   Radius = "rounded-full"
)

type Alignment string

const (
	Horizontal Alignment = "horizontal"
	Vertical   Alignment = "vertical"
)
