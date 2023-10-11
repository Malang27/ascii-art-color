package main

// Color est une fonction qui renvoie une séquence d'échappement de la couleur nécessaire
func Color(color string) string {

	// Détecter la couleur spécifiée et renvoyer la séquence d'échappement appropriée
	switch color {
	case "white":
		return "\u001b[38;2;255;255;255m"
	case "black":
		return "\u001b[38;2;0;0;0m"
	case "red":
		return "\u001b[38;2;255;0;0m"
	case "green":
		return "\u001b[38;2;0;255;0m"
	case "blue":
		return "\u001b[38;2;0;0;255m"
	case "yellow":
		return "\u001b[38;2;255;255;0m"
	case "pink":
		return "\u001b[38;2;255;0;255m"
	case "grey":
		return "\u001b[38;2;128;128;128m"
	case "purple":
		return "\u001b[38;2;160;32;255m"
	case "brown":
		return "\u001b[38;2;160;128;96m"
	case "orange":
		return "\u001b[38;2;255;160;16m"
	case "cyan":
		return "\u001b[38;2;0;183;235m"
	}

	// Retourner le blanc par défaut
	return "\u001b[38;2;255;255;255m"
}
