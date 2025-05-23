// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.865
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Login() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Hello, TEMPL</title><!-- tailwind css --><link href=\"/static/css/tailwind.css\" rel=\"stylesheet\"><!-- htmx --><script src=\"https://unpkg.com/htmx.org@1.9.10\"></script><!-- alpineJS --><script defer src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js\"></script></head><body><div class=\"flex min-h-full flex-col justify-center px-6 py-12 lg:px-8\"><div class=\"sm:mx-auto sm:w-full sm:max-w-sm\"><img class=\"mx-auto h-10 w-auto\" src=\"/static/img/mark.svg\" alt=\"Your Company\"><h2 class=\"mt-10 text-center text-2xl/9 font-bold tracking-tight text-gray-900\">Sign in to your account</h2></div><div class=\"mt-10 sm:mx-auto sm:w-full sm:max-w-sm\"><form class=\"space-y-6\" action=\"#\" method=\"POST\"><div><label for=\"username\" class=\"block text-sm/6 font-medium text-gray-900\">Username</label><div class=\"mt-2\"><input type=\"username\" name=\"username\" id=\"username\" required class=\"block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6\"></div></div><div><div class=\"flex items-center justify-between\"><label for=\"password\" class=\"block text-sm/6 font-medium text-gray-900\">Password</label></div><div class=\"mt-2\"><input type=\"password\" name=\"password\" id=\"password\" autocomplete=\"current-password\" required class=\"block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6\"></div></div><div><button type=\"submit\" class=\"flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600\">Sign in</button></div></form></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
