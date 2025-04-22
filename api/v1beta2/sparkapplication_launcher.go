package v1beta2

import (
	ctrlClient "sigs.k8s.io/controller-runtime/pkg/client"
)

type SparkAppLauncher interface {
	LaunchSparkApplication(app *SparkApplication, cl ctrlClient.Client) error
}
