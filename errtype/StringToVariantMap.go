package errtype

func StringToVariantMap() map[string]Variation {
	return stringToVariantMapOnce.Value().(map[string]Variation)
}
