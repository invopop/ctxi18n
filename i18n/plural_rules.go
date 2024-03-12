package i18n

// Standard pluralization rule keys.
const (
	ZeroOneOtherRuleKey = "zero-one-other"
)

const (
	zeroKey  = "zero"
	oneKey   = "one"
	otherKey = "other"
)

var rules = map[string]PluralRule{
	// Zero, One, Other, is the cardinal pluralization rule for
	// most Latin based languages, including english.
	ZeroOneOtherRuleKey: func(d *Dict, num int) string {
		if num == 0 {
			v := d.getEntry(zeroKey)
			if v != nil {
				return v.value
			}
		}
		if num == 1 {
			return d.get(oneKey)
		}
		return d.get(otherKey)
	},
}

// GetRule provides the PluralRule for the given key.
func GetRule(key string) PluralRule {
	return rules[key]
}

// mapPluralRule is used to map a language code into a pluralization rule.
func mapPluralRule(_ Code) PluralRule {
	return rules[ZeroOneOtherRuleKey]
}
