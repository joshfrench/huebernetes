package controllers

import (
	"context"

	. "gopkg.in/check.v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "huebernetes.dev/huebernetes/api/v1alpha1"
)

func (s *ControllerSuite) TestSanity(c *C) {
	l := &v1alpha1.Light{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "resources.huebernetes.dev/v1alpha1",
			Kind:       "Light",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "light-1",
			Namespace: "default",
		},
		Spec: v1alpha1.LightSpec{
			Id: "00000000-1111-2222-3333-444444444444",
			Metadata: v1alpha1.Metadata{
				Name: "light-1",
			},
			On: v1alpha1.On{
				On: false,
			},
		},
	}
	err := k8sClient.Create(context.Background(), l)
	c.Assert(err, IsNil)
}
