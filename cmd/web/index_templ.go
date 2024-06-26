// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package web

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Index() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" data-theme=\"light\"><head><meta charset=\"utf-8\"><meta name=\"description\" content=\"Agenda Hora é a solução definitiva em software de agendamento, projetado para otimizar agendamentos e reservas de horários forma eficiente. Cadastre-se na lista de espera hoje!\"><meta name=\"keywords\" content=\"Agenda Hora, agendahora, software de agendamento, gestão de compromissos, ferramenta de reservas, calendário online, agendamento empresarial, agendamento autonomo\"><meta name=\"robots\" content=\"index, follow\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><link rel=\"icon\" href=\"data:image/svg+xml,&lt;svg xmlns=&#39;http://www.w3.org/2000/svg&#39; viewBox=&#39;0 0 100 100&#39;&gt;&lt;text y=&#39;0.9em&#39; font-size=&#39;90&#39;&gt;📘&lt;/text&gt;&lt;/svg&gt;\"><style>\n\t\t\t\t\t/* Estilos do toast */\n\t\t\t\t\t.toast {\n\t\t\t\t\t\t\tposition: fixed;\n\t\t\t\t\t\t\ttop: 20px;\n\t\t\t\t\t\t\tright: 20px;\n\t\t\t\t\t\t\tbackground-color: #4CAF50;\n\t\t\t\t\t\t\tcolor: white;\n\t\t\t\t\t\t\tpadding: 16px;\n\t\t\t\t\t\t\tz-index: 1;\n\t\t\t\t\t\t\tdisplay: none; /* Inicialmente oculto */\n\t\t\t\t\t}\n\t\t\t</style><title>Agenda hora | agedamento de serviços</title><link href=\"assets/css/output.css\" rel=\"stylesheet\"><link href=\"assets/css/daisyui.css\" rel=\"stylesheet\"><script src=\"assets/js/htmx.min.js\"></script></head><body><main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
