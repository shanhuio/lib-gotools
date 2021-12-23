module shanhu.io/tools

go 1.17

require (
	golang.org/x/tools v0.1.8
	shanhu.io/dags v0.0.0-00010101000000-000000000000
	shanhu.io/gcimporter v0.0.0-00010101000000-000000000000
	shanhu.io/misc v0.0.0-20211219232220-7c32d2d7e486
	shanhu.io/text v0.0.0-20211223054527-6b1ed066cd84
)

require (
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace (
	shanhu.io/dags => ../dags
	shanhu.io/gcimporter => ../gcimporter
)
