package zerolog

// FieldsWithPrefix _
func (c Context) FieldsWithPrefix(fields map[string]interface{}, prefix string) Context {
	return c.Fields(prefixMapKeys(fields, prefix))
}

// FieldsWithUnderscorePrefix _
func (c Context) FieldsWithUnderscorePrefix(fields map[string]interface{}) Context {
	return c.Fields(prefixMapKeys(fields, "_"))
}

// FieldsWithPrefix _
func (e *Event) FieldsWithPrefix(fields map[string]interface{}, prefix string) *Event {
	return e.Fields(prefixMapKeys(fields, prefix))
}

// FieldsWithUnderscorePrefix _
func (e *Event) FieldsWithUnderscorePrefix(fields map[string]interface{}) *Event {
	return e.Fields(prefixMapKeys(fields, "_"))
}

func prefixMapKeys(fields map[string]interface{}, prefix string) map[string]interface{} {
	if prefix != "" {
		n := make(map[string]interface{}, len(fields))

		for k, v := range fields {
			n[prefix+k] = v
		}

		return n
	}
	return fields
}

// Event _
func (c Context) Event(event string) Context {
	return c.Str("_event", event)
}

// Event _
func (e *Event) Event(event string) *Event {
	return e.Str("_event", event)
}
