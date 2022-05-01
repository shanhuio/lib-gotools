module shanhu.io/smlrepo/tools

go 1.18

require (
	golang.org/x/tools v0.1.10
	shanhu.io/dags v0.0.0-20220320061527-bb7abb042c8a
	shanhu.io/lab/gcimporter v0.0.0-00010101000000-000000000000
	shanhu.io/misc v0.0.0-20220417204140-117e3c66ed14
	shanhu.io/text v0.0.0-20220403174149-0195ecfdda87
)

require (
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/sys v0.0.0-20220429233432-b5fbb4746d32 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
)

replace (
	shanhu.io/dags => ../../lib/dags
	shanhu.io/lab/gcimporter => ../../lab/gcimporter
)
