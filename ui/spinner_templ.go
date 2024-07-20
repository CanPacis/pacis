// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func (c *SpinnerComponent) getSize() string {
	return *sw(
		c.variance().Size,
		[]Variant{SizeHuge, SizeLarge, SizeMedium, SizeSmall, SizeTiny},
		"32", "24", "16", "12", "8",
	)
}

func spinner(c *SpinnerComponent) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{class(c.classList, "relative", "size-fit", "animate-spin")}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/spinner.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" aria-busy=\"true\" aria-label=\"Loading\" role=\"status\" label=\"Spinner\" title=\"Spinner\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(c.ID())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/spinner.templ`, Line: 19, Col: 13}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, c.getAttrs("base"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><svg width=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(c.getSize())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/spinner.templ`, Line: 23, Col: 22}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" height=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(c.getSize())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/spinner.templ`, Line: 24, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" aria-labelledby=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(c.ID())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `ui/spinner.templ`, Line: 25, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M7.20001 0.8C7.20001 0.358172 7.55818 0 8.00001 0C8.44184 0 8.80001 0.358172 8.80001 0.8V3.2C8.80001 3.64183 8.44184 4 8.00001 4C7.55818 4 7.20001 3.64183 7.20001 3.2V0.8Z\" fill=\"#858585\"></path> <path d=\"M7.20001 12.8C7.20001 12.3582 7.55818 12 8.00001 12C8.44184 12 8.80001 12.3582 8.80001 12.8V15.2C8.80001 15.6418 8.44184 16 8.00001 16C7.55818 16 7.20001 15.6418 7.20001 15.2V12.8Z\" fill=\"#565656\"></path> <path d=\"M0.8 8.8C0.358172 8.8 1.93129e-08 8.44183 0 8C-1.93129e-08 7.55818 0.358172 7.2 0.8 7.2H3.2C3.64183 7.2 4 7.55818 4 8C4 8.44183 3.64183 8.8 3.2 8.8H0.8Z\" fill=\"#6E6E6E\"></path> <path d=\"M12.8 8.8C12.3582 8.8 12 8.44183 12 8C12 7.55818 12.3582 7.2 12.8 7.2H15.2C15.6418 7.2 16 7.55818 16 8C16 8.44183 15.6418 8.8 15.2 8.8H12.8Z\" fill=\"#3F3F3F\"></path> <path d=\"M2.16462 12.0928C1.78199 12.3137 1.29272 12.1826 1.0718 11.8C0.850888 11.4174 0.981988 10.9281 1.36462 10.7072L3.44308 9.50719C3.82572 9.28627 4.31499 9.41737 4.5359 9.80001C4.75682 10.1826 4.62572 10.6719 4.24308 10.8928L2.16462 12.0928Z\" fill=\"#666666\"></path> <path d=\"M12.557 6.09283C12.1743 6.31374 11.6851 6.18264 11.4641 5.80001C11.2432 5.41737 11.3743 4.9281 11.757 4.70719L13.8354 3.50719C14.2181 3.28627 14.7073 3.41737 14.9282 3.80001C15.1492 4.18264 15.0181 4.67191 14.6354 4.89283L12.557 6.09283Z\" fill=\"#373737\"></path> <path d=\"M3.70719 1.96463C3.48628 1.58199 3.61738 1.09272 4.00001 0.871805C4.38264 0.650891 4.87192 0.781991 5.09283 1.16463L6.29283 3.24309C6.51374 3.62572 6.38264 4.11499 6.00001 4.33591C5.61738 4.55682 5.1281 4.42572 4.90719 4.04309L3.70719 1.96463Z\" fill=\"#7D7D7D\"></path> <path d=\"M9.70719 12.3569C9.48628 11.9743 9.61738 11.485 10 11.2641C10.3826 11.0432 10.8719 11.1743 11.0928 11.5569L12.2928 13.6354C12.5137 14.018 12.3826 14.5073 12 14.7282C11.6174 14.9491 11.1281 14.818 10.9072 14.4354L9.70719 12.3569Z\" fill=\"#4F4F4F\"></path> <path d=\"M5.09281 14.4354C4.8719 14.818 4.38262 14.9491 3.99999 14.7282C3.61736 14.5073 3.48626 14.018 3.70717 13.6354L4.90717 11.5569C5.12808 11.1743 5.61736 11.0432 5.99999 11.2641C6.38262 11.485 6.51372 11.9743 6.29281 12.3569L5.09281 14.4354Z\" fill=\"#5E5E5E\"></path> <path d=\"M11.0928 4.04308C10.8719 4.42571 10.3826 4.55681 9.99999 4.3359C9.61736 4.11499 9.48626 3.62571 9.70717 3.24308L10.9072 1.16462C11.1281 0.781985 11.6174 0.650885 12 0.871799C12.3826 1.09271 12.5137 1.58199 12.2928 1.96462L11.0928 4.04308Z\" fill=\"#303030\"></path> <path d=\"M1.36463 4.89281C0.982 4.6719 0.8509 4.18263 1.07181 3.79999C1.29273 3.41736 1.782 3.28626 2.16463 3.50717L4.2431 4.70717C4.62573 4.92809 4.75683 5.41736 4.53592 5.79999C4.315 6.18263 3.82573 6.31373 3.4431 6.09281L1.36463 4.89281Z\" fill=\"#757575\"></path> <path d=\"M11.757 10.8928C11.3743 10.6719 11.2432 10.1826 11.4641 9.79999C11.6851 9.41736 12.1743 9.28626 12.557 9.50717L14.6354 10.7072C15.0181 10.9281 15.1492 11.4174 14.9282 11.8C14.7073 12.1826 14.2181 12.3137 13.8354 12.0928L11.757 10.8928Z\" fill=\"#474747\"></path></svg></span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

type SpinnerComponent struct {
	VariantComponent[*SpinnerComponent]
	variant Variant
}

func Spinner() *SpinnerComponent {
	c := &SpinnerComponent{}
	c.Component = NewComponent(c, spinner)
	return c
}
