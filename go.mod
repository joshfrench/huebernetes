module github.com/joshfrench/huebernetes

go 1.16

replace huebernetes.dev => ./

require (
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	huebernetes.dev v0.0.0-00010101000000-000000000000 // indirect
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	sigs.k8s.io/controller-runtime v0.10.0
)
