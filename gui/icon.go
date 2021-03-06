package gui

type Icon rune

const ( //  font awesome corresponding characters
	NoIcon = '\x00'

	QuitIcon = '\uea0f' // '\uea0d'
	FileIcon = '\uf4c4'
	OpenIcon = '\ue930'
	SaveIcon = '\uf4be'

	ExportIcon = '\uf4e4'
	ImportIcon = '\uf4e5'

	SettingsIcon = '\ue992'
	// SettingsVIcon = '\ue993'

	ReloadIcon = '\u27f2'

	PlayIcon  = '\u25b6' // '\uea15'
	PauseIcon = '\u25ae' // '\uea16'
	StopIcon  = '\u25a0' // '\uea17'

	BeginIcon = '\u23ea' // '\uea18'
	EndIcon   = '\u23e9' // '\uea19'

	NextStepIcon = '\ue966'
	PrevStepIcon = '\ue965'

	CenterOnIcon = '\ue3b4'
	CenterOffIcon = '\ue3b5'
	RotateIcon = '\ue1c1'
)

func AlwaysIcon(icon Icon) func() Icon {
	return func() Icon {
		return icon
	}
}

func True() bool {
	return true
}
