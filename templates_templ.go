// Code generated by templ@v0.2.364 DO NOT EDIT.

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "strconv"

func helloStruct(data HelloData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div>")
		if err != nil {
			return err
		}
		var_2 := `Hello, `
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		var var_3 string = data.Name
		_, err = templBuffer.WriteString(templ.EscapeString(var_3))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func layout(body templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/@picocss/pico@1/css/pico.min.css\"><title>")
		if err != nil {
			return err
		}
		var_5 := `Hello, world!`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title></head><body>")
		if err != nil {
			return err
		}
		err = body.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func table(table Table) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_6 := templ.GetChildren(ctx)
		if var_6 == nil {
			var_6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<p>")
		if err != nil {
			return err
		}
		var_7 := `Table: `
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		var var_8 string = table.TableName
		_, err = templBuffer.WriteString(templ.EscapeString(var_8))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><table><tr><th>")
		if err != nil {
			return err
		}
		var_9 := `Id`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th>")
		if err != nil {
			return err
		}
		var_10 := `Name`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th>")
		if err != nil {
			return err
		}
		var_11 := `Number`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th>")
		if err != nil {
			return err
		}
		var_12 := `Complex`
		_, err = templBuffer.WriteString(var_12)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th>")
		if err != nil {
			return err
		}
		var_13 := `Static`
		_, err = templBuffer.WriteString(var_13)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th></tr>")
		if err != nil {
			return err
		}
		for _, item := range table.Items {
			_, err = templBuffer.WriteString("<tr><td>")
			if err != nil {
				return err
			}
			var var_14 string = item.Id
			_, err = templBuffer.WriteString(templ.EscapeString(var_14))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</td><td>")
			if err != nil {
				return err
			}
			var var_15 string = item.Name
			_, err = templBuffer.WriteString(templ.EscapeString(var_15))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</td><td>")
			if err != nil {
				return err
			}
			var var_16 string = strconv.Itoa(item.Number)
			_, err = templBuffer.WriteString(templ.EscapeString(var_16))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</td><td><div><p>")
			if err != nil {
				return err
			}
			var var_17 string = item.Complex.Id
			_, err = templBuffer.WriteString(templ.EscapeString(var_17))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p><p>")
			if err != nil {
				return err
			}
			var var_18 string = item.Complex.Name
			_, err = templBuffer.WriteString(templ.EscapeString(var_18))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p><p>")
			if err != nil {
				return err
			}
			var var_19 string = strconv.Itoa(item.Complex.Number)
			_, err = templBuffer.WriteString(templ.EscapeString(var_19))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p></div></td><td>")
			if err != nil {
				return err
			}
			var_20 := `Static`
			_, err = templBuffer.WriteString(var_20)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</td></tr>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</table>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
