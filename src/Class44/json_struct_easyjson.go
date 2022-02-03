// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package Class44

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonA35b7fcDecodeClass44(in *jlexer.Lexer, out *JobInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "skill":
			if in.IsNull() {
				in.Skip()
				out.Skill = nil
			} else {
				in.Delim('[')
				if out.Skill == nil {
					if !in.IsDelim(']') {
						out.Skill = make([]string, 0, 4)
					} else {
						out.Skill = []string{}
					}
				} else {
					out.Skill = (out.Skill)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Skill = append(out.Skill, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA35b7fcEncodeClass44(out *jwriter.Writer, in JobInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"skill\":"
		out.RawString(prefix[1:])
		if in.Skill == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Skill {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JobInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA35b7fcEncodeClass44(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JobInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA35b7fcEncodeClass44(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JobInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA35b7fcDecodeClass44(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JobInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA35b7fcDecodeClass44(l, v)
}
func easyjsonA35b7fcDecodeClass441(in *jlexer.Lexer, out *Employee) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "basic_info":
			(out.BasicInfo).UnmarshalEasyJSON(in)
		case "job_info":
			(out.JobInfo).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA35b7fcEncodeClass441(out *jwriter.Writer, in Employee) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"basic_info\":"
		out.RawString(prefix[1:])
		(in.BasicInfo).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"job_info\":"
		out.RawString(prefix)
		(in.JobInfo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Employee) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA35b7fcEncodeClass441(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Employee) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA35b7fcEncodeClass441(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Employee) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA35b7fcDecodeClass441(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Employee) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA35b7fcDecodeClass441(l, v)
}
func easyjsonA35b7fcDecodeClass442(in *jlexer.Lexer, out *BasicInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "age":
			out.Age = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA35b7fcEncodeClass442(out *jwriter.Writer, in BasicInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"age\":"
		out.RawString(prefix)
		out.Int(int(in.Age))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BasicInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA35b7fcEncodeClass442(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BasicInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA35b7fcEncodeClass442(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BasicInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA35b7fcDecodeClass442(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BasicInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA35b7fcDecodeClass442(l, v)
}
