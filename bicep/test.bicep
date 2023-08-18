resource symbolicname 'Microsoft.KeyVault/vaults@2022-07-01' = {
  name: 'hackme'
  location: 'eastus'
  properties: {
    publicNetworkAccess: 'enabled'
    sku: {
      family: 'A'
      name: 'standard'
    }
    tenantId: '[subscription().tenantId]'
  }
}