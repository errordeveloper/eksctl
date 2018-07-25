package manager

import (
	"github.com/pkg/errors"
	"github.com/weaveworks/eksctl/pkg/cfn/builder"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

func (c *StackCollection) stackNameCluster() string {
	return "eksctl-" + c.Spec.ClusterName + "-cluster"
}

func (c *StackCollection) CreateCluster(errs chan error) error {
	name := c.stackNameCluster()
	logger.Info("creating cluster stack %q", name)

	stack := builder.NewClusterResourceSet()
	stack.AddAllResources(c.Spec.AvailabilityZones)

	templateBody, err := stack.RenderJSON()
	if err != nil {
		return errors.Wrap(err, "rendering template for cluster stack")
	}

	logger.Debug("templateBody = %s", string(templateBody))

	stackChan := make(chan Stack)
	taskErrs := make(chan error)

	if err := c.CreateStack(name, templateBody, nil, true, stackChan, taskErrs); err != nil {
		return err
	}

	go func() {
		defer close(errs)

		if err := <-taskErrs; err != nil {
			errs <- err
			return
		}

		if err := stack.GetAllOutputs(stackChan, c.Spec); err != nil {
			errs <- errors.Wrap(err, "getting cluster stack outputs")
		}

		logger.Debug("clusterConfig = %#v", c.Spec)
		logger.Success("created cluster stack %q", name)

		errs <- nil
	}()
	return nil
}

func (c *StackCollection) DeleteCluster() error {
	return c.DeleteStack(c.stackNameCluster())
}
