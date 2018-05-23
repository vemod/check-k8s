package main

import (
	"io"

	"github.com/benkeil/check-k8s/pkg/checks"

	"github.com/benkeil/check-k8s/cmd/api"
	"github.com/spf13/cobra"
	"k8s.io/api/apps/v1"
)

type (
	checkDeploymentAvailableReplicasCmd struct {
		out               io.Writer
		Deployment        *v1.Deployment
		Name              string
		Namespace         string
		ThresholdWarning  string
		ThresholdCritical string
	}
)

func newCheckDeploymentAvailableReplicasCmd(out io.Writer) *cobra.Command {
	c := &checkDeploymentAvailableReplicasCmd{out: out}

	cmd := &cobra.Command{
		Use:          "availableReplicas",
		Short:        "check if a k8s deployment has a minimum of available replicas",
		SilenceUsage: true,
		Args:         NameArgs(),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			c.Name = args[0]
			deployment, err := api.GetDeployment(settings, api.GetDeploymentOptions{Name: c.Name, Namespace: c.Namespace})
			if err != nil {
				return err
			}
			c.Deployment = deployment
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c.run()
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&c.Namespace, "namespace", "n", "", "the namespace of the deployment")
	cmd.Flags().StringVarP(&c.ThresholdCritical, "critical", "c", "2:", "critical threshold for minimum available replicas")
	cmd.Flags().StringVarP(&c.ThresholdWarning, "warning", "w", "2:", "warning threshold for minimum available replicas")

	return cmd
}

func (c *checkDeploymentAvailableReplicasCmd) run() {
	checkDeployment := checks.NewCheckDeployment(c.Deployment)
	result := checkDeployment.CheckAvailableReplicas(c.ThresholdWarning, c.ThresholdCritical)
	result.Exit()
}
