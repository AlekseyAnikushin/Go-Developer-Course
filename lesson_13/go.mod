module lesson_13

go 1.22.1

replace (
	v100 => github.com/AlekseyAnikushin/module13 v1.0.0
	v110 => github.com/AlekseyAnikushin/module13 v1.1.0
	v200 v0.0.0 => github.com/AlekseyAnikushin/module13/v2 v2.0.0
	v210 v0.0.1 => github.com/AlekseyAnikushin/module13/v2 v2.1.0
)

require (
	v100 v1.0.0
	v110 v1.1.0
	v200 v0.0.0
	v210 v0.0.1
)
