package utils

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/blang/semver"
	"github.com/kballard/go-shellquote"
	"github.com/pkg/errors"

	"k8s.io/client-go/tools/clientcmd"

	"github.com/kubicorn/kubicorn/pkg/logger"
	"github.com/weaveworks/launcher/pkg/kubectl"
)

func fmtKubectlCmd(ktl *kubectl.LocalClient, cmds ...string) string {
	args := []string{kubectl.Command}
	args = append(args, ktl.GlobalArgs...)
	args = append(args, cmds...)
	return shellquote.Join(args...)
}

func CheckKubectlVersion(env []string) error {
	ktl := &kubectl.LocalClient{Env: env}
	kubectlPath, err := ktl.LookPath()
	if err != nil {
		return fmt.Errorf("kubectl not found, v1.10.0 or newever is required")
	}
	logger.Debug("kubectl: %q", kubectlPath)

	clientVersion, _, err := kubectl.GetVersionInfo(ktl)
	logger.Debug("kubectl version: %s", clientVersion)
	if err != nil {
		logger.Debug("ignored error: %s", err.Error())
	}

	version, err := semver.Parse(strings.TrimLeft(clientVersion, "v"))
	if err != nil {
		return errors.Wrapf(err, "parsing kubectl version string %q", version)
	}
	if version.Major == 1 && version.Minor < 10 {
		return fmt.Errorf("kubectl version %s was found at %q, minimum required version to use EKS is v1.10.0", clientVersion, kubectlPath)
	}
	return nil
}

func CheckHeptioAuthenticatorAWS() error {
	path, err := exec.LookPath("heptio-authenticator-aws")
	if err == nil {
		logger.Debug("heptio-authenticator-aws: %q", path)
	} else {
		return fmt.Errorf("heptio-authenticator-aws not installed")
	}
	return nil
}

func CheckAllCommands(kubeconfigPath string, isContextSet bool, contextName string, env []string) error {
	if err := CheckKubectlVersion(env); err != nil {
		return err
	}

	if err := CheckHeptioAuthenticatorAWS(); err != nil {
		return err
	}

	if kubeconfigPath != "" {
		ktl := &kubectl.LocalClient{
			GlobalArgs: []string{},
			Env:        env,
		}
		if kubeconfigPath != clientcmd.RecommendedHomeFile {
			ktl.GlobalArgs = append(ktl.GlobalArgs, fmt.Sprintf("--kubeconfig=%s", kubeconfigPath))
		}
		if !isContextSet {
			ktl.GlobalArgs = append(ktl.GlobalArgs, fmt.Sprintf("--context=%s", contextName))
		}

		suggestion := fmt.Sprintf("(check '%s')", fmtKubectlCmd(ktl, "version"))

		_, serverVersion, err := kubectl.GetVersionInfo(ktl)
		if err != nil {
			return errors.Wrapf(err, "unable to use kubectl with the EKS cluster %s", suggestion)
		}
		version, err := semver.Parse(strings.TrimLeft(serverVersion, "v"))
		if err != nil {
			return errors.Wrapf(err, "parsing Kubernetes version string %q return by the EKS API server", version)
		}
		if version.Major == 1 && version.Minor < 10 {
			return fmt.Errorf("Kubernetes version %s found, v1.10.0 or newer is expected with EKS %s", serverVersion, suggestion)
		}

		logger.Info("kubectl command should work with %q, try '%s'", kubeconfigPath, fmtKubectlCmd(ktl, "get", "nodes"))
	} else {
		logger.Debug("skipping kubectl integration ckecks, as writing kubeconfig file is disabled")
	}

	return nil
}
