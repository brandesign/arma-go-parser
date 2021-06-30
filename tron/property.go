package tron

type Property struct {
	properties map[string]interface{}
}

func (p *Property) set(key string, val interface{}) {
	if p.properties == nil {
		p.properties = map[string]interface{}{}
	}

	p.properties[key] = val
}

func (p *Property) get(key string) interface{} {
	if p.properties == nil {
		return nil
	}

	return p.properties[key]
}

func (p *Property) SetString(key string, val string) {
	p.set(key, val)
}

func (p *Property) GetString(key string) string {
	v := p.get(key)
	if v == nil {
		return ""
	}

	return v.(string)
}

func (p *Property) SetInt(key string, val int) {
	p.set(key, val)
}

func (p *Property) GetInt(key string) int {
	v := p.get(key)
	if v == nil {
		return 0
	}

	return v.(int)
}

func (p *Property) SetFloat(key string, val float64) {
	p.set(key, val)
}

func (p *Property) GetFloat(key string) float64 {
	v := p.get(key)
	if v == nil {
		return 0
	}

	return v.(float64)
}

func (p *Property) SetBool(key string, val bool) {
	p.set(key, val)
}

func (p *Property) GetBool(key string) bool {
	v := p.get(key)
	if v == nil {
		return false
	}

	return v.(bool)
}
