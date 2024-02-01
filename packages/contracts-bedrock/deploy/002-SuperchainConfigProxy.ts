import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../scripts/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const proxyAdmin = await getDeploymentAddress(hre, 'ProxyAdmin')

  await deploy({
    hre,
    name: 'SuperchainConfigProxy',
    contract: 'Proxy',
    args: [proxyAdmin],
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'admin', proxyAdmin)
    },
  })
}

deployFn.tags = ['SuperchainConfigProxy', 'setup', 'l1']

export default deployFn