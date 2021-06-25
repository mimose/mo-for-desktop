module mo-for-desktop

go 1.16

require (
	github.com/wailsapp/wails v1.16.3
	github.com/mimose/gcosy v0.0.0
)

replace (
	github.com/mimose/gcosy v0.0.0 => ../gcosy
)
