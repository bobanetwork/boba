package pipeline

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum-optimism/optimism/op-chain-ops/script"

	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/state"
)

type DeployBobaScript struct {
	Run func(output common.Address) error
}

type DeployBobaOutput struct {
        BobaL1                         common.Address `json:"bobaL1Address"`
}

func DeployBobaL1(env *Env, intent *state.Intent, st *state.State) error {
	lgr := env.Logger.New("stage", "deploy-boba-l1")

	if !shouldDeployBobaL1(intent, st) {
		lgr.Info("Boba L1 deployment not needed")
		return nil
	}

	lgr.Info("Deploying Boba L1")

	host := env.L1ScriptHost;
	var output DeployBobaOutput
	outputAddr := host.NewScriptAddress()

	cleanupOutput, err := script.WithPrecompileAtAddress[*DeployBobaOutput](host, outputAddr, &output,
		script.WithFieldSetter[*DeployBobaOutput])
	if err != nil {
		return fmt.Errorf("failed to insert DeployBobaOutput precompile: %w", err)
	}
	defer cleanupOutput()

	implContract := "DeployBoba"
	deployScript, cleanupDeploy, err := script.WithScript[DeployBobaScript](host, "DeployBoba.s.sol", implContract)
	if err != nil {
		return fmt.Errorf("failed to load %s script: %w", implContract, err)
	}
	defer cleanupDeploy()

	if err := deployScript.Run(outputAddr); err != nil {
		return fmt.Errorf("failed to run %s script: %w", implContract, err)
	}

	st.L1Deployment = &state.L1Deployment{
		BobaL1: output.BobaL1,
	}

        lgr.Info("Boba L1 done", "st.L1Deployment", st.L1Deployment)

	return nil
}

func shouldDeployBobaL1(intent *state.Intent, st *state.State) bool {
	return st.L1Deployment == nil
}
