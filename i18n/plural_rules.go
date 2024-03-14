package i18n

// Standard pluralization rule keys.
const (
	DefaultRuleKey = "default"
)

const (
	zeroKey  = "zero"
	oneKey   = "one"
	otherKey = "other"
)

var rules = map[string]PluralRule{
	// Most languages can use this rule
	DefaultRuleKey: func(d *Dict, n int) string {
		if n == 0 {
			v := d.GetEntry(zeroKey)
			if v != nil {
				return v.value
			}
		}
		if n == 1 {
			return d.Get(oneKey)
		}
		return d.Get(otherKey)
	},
}

// GetRule provides the PluralRule for the given key.
func GetRule(key string) PluralRule {
	return rules[key]
}

// mapPluralRule is used to map a language code into a pluralization rule.
func mapPluralRule(_ Code) PluralRule {
	return rules[DefaultRuleKey]
}
