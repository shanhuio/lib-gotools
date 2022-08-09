module shanhu.io/gotools

go 1.18

require (
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/tools v0.1.12
	shanhu.io/dags v0.0.0-20220320061527-bb7abb042c8a
	shanhu.io/gcimporter v0.0.0-00010101000000-000000000000
	shanhu.io/misc v0.0.0-20220809022537-39c292211112
	shanhu.io/text v0.0.0-20220403174149-0195ecfdda87
)

require (
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sys v0.0.0-20220808155132-1c4a2a72c664 // indirect
)

replace (
	shanhu.io/dags => ../../lib/dags
	shanhu.io/gcimporter => ../gcimporter
)
