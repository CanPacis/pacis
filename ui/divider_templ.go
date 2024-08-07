// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func divider(c *DividerComponent) templ.Component {
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
		if c.content != nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span role=\"separator\" aria-orientation=\"horizontal\" class=\"flex items-center my-2\"><span class=\"flex-1 border-t border-p-gray-400\"></span> <span class=\"h-8 px-3 text-sm border rounded-full size-fit border-p-gray-400\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = c.content.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"flex-1 border-t border-p-gray-400\"></span></span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			if has(c.label) {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span role=\"separator\" aria-orientation=\"horizontal\" class=\"flex items-center my-2\"><span class=\"flex-1 border-t border-p-gray-400\"></span> <span class=\"flex items-center justify-center h-8 px-3 text-sm border rounded-full size-fit border-p-gray-400\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(c.label)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/divider.templ`, Line: 17, Col: 14}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"flex-1 border-t border-p-gray-400\"></span></span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<hr class=\"my-2 border-t border-p-gray-400\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		return templ_7745c5c3_Err
	})
}

type DividerComponent struct {
	Component[*DividerComponent]
	label   string
	content templ.Component
}

func (c *DividerComponent) Label(label string) *DividerComponent {
	c.label = label
	return c
}

func (c *DividerComponent) Content(content templ.Component) *DividerComponent {
	c.content = content
	return c
}

func Divider() *DividerComponent {
	c := &DividerComponent{}
	c.Component = NewComponent(c, divider)
	return c
}
