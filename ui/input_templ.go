// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"strings"
)

func (c *InputComponent) getIconSize() int {
	switch c.variance().Size {
	case SizeHuge:
		return 20
	case SizeLarge, SizeMedium, SizeSmall:
		return 16
	case SizeTiny:
		return 14
	default:
		return 16
	}
}

func input(c *InputComponent) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{"input rounded-lg h-fit", kv(c.fullWidth, "w-full", "w-fit")}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(c.ID())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 23, Col: 13}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if c.getAttr("input", "type").(InputType) == InputTypePassword {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" x-data=\"{ type: &#39;password&#39;, value: &#39;&#39; }\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			if c.clearable {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" x-data=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("{ value: '%s' }", c.getAttr("input", "value")))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 28, Col: 72}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" x-data")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, c.getAttrs("label"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><div class=\"flex flex-col flex-auto gap-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if has(c.label) && !c.invisibleLabel {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"text-xs text-p-gray-900\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(c.label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 39, Col: 14}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		var templ_7745c5c3_Var7 = []any{"flex", "flex-auto", "flex-nowrap",
			kv(c.variance().Form == FormDisabled, "cursor-not-allowed"),
			*sw(
				c.variance().Size,
				[]Variant{SizeHuge, SizeLarge, SizeMedium, SizeSmall, SizeTiny},
				"text-lg h-16",
				"text-base h-12",
				"text-sm h-10",
				"text-sm h-8",
				"text-sm h-6",
			)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var7...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var7).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"flex flex-1 rounded-lg\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if c.prefix != nil {
			var templ_7745c5c3_Var9 = []any{
				"flex items-center h-full px-3 border text-p-gray-900 border-p-gray-400",
				kv(c.variance().Size == SizeTiny, "rounded-l-md", "rounded-l-lg"),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var9...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var9).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = c.prefix.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		var templ_7745c5c3_Var11 = []any{
			"h-full flex-1 flex justify-between items-center gap-2 px-3 bg-p-background-100 rounded-lg focusable border-p-gray-400 z-10 has-[:user-invalid]:error-ring",
			kv(has(c.error), "error-ring", "focus-within:focus-ring"),
			*sw(
				true,
				[]any{
					c.prefix == nil && c.suffix == nil,
					c.prefix != nil && c.suffix == nil,
					c.prefix == nil && c.suffix != nil,
					c.prefix != nil && c.suffix != nil,
				},
				kv(c.variance().Size == SizeTiny, "rounded-md border", "rounded-lg border"),
				kv(c.variance().Size == SizeTiny, "rounded-r-none border-t border-b border-r", "rounded-r-lg border-t border-b border-r"),
				kv(c.variance().Size == SizeTiny, "rounded-l-none border-t border-b border-l", "rounded-l-lg border-t border-b border-l"),
				"rounded-none border-t border-b",
			),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var11...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var11).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(strings.TrimSpace(c.startIcon)) > 0 {
			templ_7745c5c3_Err = Icon(c.startIcon).Size(c.getIconSize()).StrokeColor(Color(Gray, 700)).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		var templ_7745c5c3_Var13 = []any{"flex-1 h-full border-none outline-none bg-p-background-100 border-p-gray-400",
			kv(c.variance().Form == FormDisabled, "placeholder-p-gray-400 cursor-not-allowed text-p-gray-700", "placeholder-p-gray-700 text-inherit")}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var13...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if c.clearable {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" x-model:value=\"value\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var13).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" label=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(c.label)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 94, Col: 22}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" aria-label=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(c.label)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 95, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" aria-labelledby=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(c.ID())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 96, Col: 31}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if c.variance().Form == FormDisabled {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, c.getAttrs("input"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if c.getAttr("input", "type").(InputType) == InputTypePassword {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" x-bind:type=\"type\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if c.getAttr("input", "type").(InputType) == InputTypePassword {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"rounded-sm focus-visible:focus-ring\" type=\"button\" x-on:click.prevent=\"type = type === &#39;password&#39; ? &#39;text&#39; : &#39;password&#39;\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = Icon("eye").Size(c.getIconSize()).StrokeColor(Color(Gray, 700)).Attr("x-show", "type === 'password'").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = Icon("eye-closed").Size(c.getIconSize()).StrokeColor(Color(Gray, 700)).Attr("x-show", "type !== 'password'").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if c.clearable {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"rounded-sm focus-visible:focus-ring\" type=\"button\" x-on:click.prevent=\"value = &#39;&#39;\" x-show=\"value.trim().length !== 0\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = Icon("x").Size(c.getIconSize()).StrokeColor(Color(Gray, 700)).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if len(strings.TrimSpace(c.endIcon)) > 0 {
			templ_7745c5c3_Err = Icon(c.endIcon).Size(c.getIconSize()).StrokeColor(Color(Gray, 700)).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if c.suffix != nil {
			var templ_7745c5c3_Var18 = []any{
				"flex items-center h-full px-3 border text-p-gray-900 border-p-gray-400",
				kv(c.variance().Size == SizeTiny, "rounded-r-md", "rounded-r-lg"),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var18...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var19 string
			templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var18).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = c.suffix.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if has(c.error) {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex gap-2 text-xs text-p-red-900\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = Icon("alert-octagon").Size(c.getIconSize()).StrokeColor(Color(Red, 900)).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var20 string
			templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(c.error)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 136, Col: 15}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if has(c.formError) {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"hidden gap-2 text-xs text-p-red-900 form-error\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = Icon("alert-octagon").Size(c.getIconSize()).StrokeColor(Color(Red, 900)).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var21 string
			templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(c.formError)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/input.templ`, Line: 144, Col: 19}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></label>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

type InputComponent struct {
	VariantComponent[*InputComponent]
	fullWidth      bool
	label          string
	invisibleLabel bool
	clearable      bool
	startIcon      string
	endIcon        string
	prefix         templ.Component
	suffix         templ.Component
	error          string
	formError      string
}

type InputType string

const (
	InputTypeText     InputType = "text"
	InputTypePassword InputType = "password"
	InputTypeEmail    InputType = "email"
	InputTypeNumber   InputType = "number"
	InputTypeTel      InputType = "tel"
	InputTypeUrl      InputType = "url"
	InputTypeSearch   InputType = "search"
	InputTypeDate     InputType = "date"
	InputTypeTime     InputType = "time"
	InputTypeWeek     InputType = "week"
	InputTypeMonth    InputType = "month"
	InputTypeDatetime InputType = "datetime"
	InputTypeColor    InputType = "color"
)

func (c *InputComponent) Type(typ InputType) *InputComponent {
	c.setAttr("type", typ, "input")
	return c
}

func (c *InputComponent) Disabled(disabled bool) *InputComponent {
	if disabled {
		return c.Variant(FormDisabled)
	} else {
		return c.Variant(FormEnabled)
	}
}

func (c *InputComponent) Placeholder(placeholder string) *InputComponent {
	c.setAttr("placeholder", placeholder, "input")
	return c
}

func (c *InputComponent) Required(required bool) *InputComponent {
	c.setAttr("required", required, "input")
	return c
}

func (c *InputComponent) Readonly(readonly bool) *InputComponent {
	c.setAttr("readonly", readonly, "input")
	return c
}

func (c *InputComponent) Autofocus(autofocus bool) *InputComponent {
	c.setAttr("autofocus", autofocus, "input")
	return c
}

func (c *InputComponent) FullWidth() *InputComponent {
	c.fullWidth = true
	return c
}

func (c *InputComponent) Label(label string) *InputComponent {
	c.label = label
	return c
}

func (c *InputComponent) Clearable() *InputComponent {
	c.clearable = true
	return c
}

func (c *InputComponent) StartIcon(icon string) *InputComponent {
	c.startIcon = icon
	return c
}

func (c *InputComponent) EndIcon(icon string) *InputComponent {
	c.endIcon = icon
	return c
}

func (c *InputComponent) Prefix(prefix templ.Component) *InputComponent {
	c.prefix = prefix
	return c
}

func (c *InputComponent) Suffix(suffix templ.Component) *InputComponent {
	c.suffix = suffix
	return c
}

func (c *InputComponent) PrefixText(prefix string) *InputComponent {
	c.prefix = textRenderer{data: prefix}
	return c
}

func (c *InputComponent) SuffixText(suffix string) *InputComponent {
	c.suffix = textRenderer{data: suffix}
	return c
}

func (c *InputComponent) Error(err string) *InputComponent {
	c.error = err
	return c
}

func (c *InputComponent) Eerror(err error) *InputComponent {
	if err != nil {
		c.error = err.Error()
	}
	return c
}

func (c *InputComponent) FormError(err string) *InputComponent {
	c.formError = err
	return c
}

func (c *InputComponent) Value(value string) *InputComponent {
	c.setAttr("value", value, "input")
	return c
}

func Input() *InputComponent {
	c := &InputComponent{}
	c.Component = NewComponent(c, input)
	c.setAttr("type", InputTypeText, "input")
	c.setAttr("value", "", "input")
	c.setAttr("placeholder", "", "input")

	return c
}
