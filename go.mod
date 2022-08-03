module shanhu.io/smlrepo/tools

go 1.18

require (
	golang.org/x/tools v0.1.12
	shanhu.io/dags v0.0.0-20220320061527-bb7abb042c8a
	shanhu.io/lab/gcimporter v0.0.0-00010101000000-000000000000
	shanhu.io/misc v0.0.0-20220803070526-2da1b044a170
	shanhu.io/text v0.0.0-20220403174149-0195ecfdda87
)

require (
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sys v0.0.0-20220731174439-a90be440212d // indirect
)

replace (
	shanhu.io/dags => ../../lib/dags
	shanhu.io/lab/gcimporter => ../../lab/gcimporter
)
