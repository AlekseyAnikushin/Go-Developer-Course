module github.com/AlekseyAnikushin/Go-Developer-Course/lesson_10/task4

go 1.22.1

replace (
	v100 => github.com/AlekseyAnikushin/first_module v1.0.0
	v110 => github.com/AlekseyAnikushin/first_module v1.1.0
)

require (
	v100 v1.0.0
	v110 v1.1.0
)