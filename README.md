## Installation

## Documentation for API Endpoints

All URIs are relative to *https://napi.arvancloud.com/ecc/v1*

| Class              | Method                                                                                                                      | HTTP request                                                  | Description                                                                  |
| ------------------ | --------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| _FloatIPsApi_      | [**RegionsRegionFloatIpsDetachPatch**](docs/FloatIPsApi.md#regionsregionfloatipsdetachpatch)                                | **Patch** /regions/{region}/float-ips/detach                  | Detach a Float IP from a Server.                                             |
| _FloatIPsApi_      | [**RegionsRegionFloatIpsGet**](docs/FloatIPsApi.md#regionsregionfloatipsget)                                                | **Get** /regions/{region}/float-ips                           | List of all Float IPs.                                                       |
| _FloatIPsApi_      | [**RegionsRegionFloatIpsIdAttachPatch**](docs/FloatIPsApi.md#regionsregionfloatipsidattachpatch)                            | **Patch** /regions/{region}/float-ips/{id}/attach             | Attach a Float IP to a Server.                                               |
| _FloatIPsApi_      | [**RegionsRegionFloatIpsIdDelete**](docs/FloatIPsApi.md#regionsregionfloatipsiddelete)                                      | **Delete** /regions/{region}/float-ips/{id}                   | Delete Float IP.                                                             |
| _FloatIPsApi_      | [**RegionsRegionFloatIpsPost**](docs/FloatIPsApi.md#regionsregionfloatipspost)                                              | **Post** /regions/{region}/float-ips                          | Create new Float IP.                                                         |
| _ImagesApi_        | [**RegionsRegionImagesFileDelete**](docs/ImagesApi.md#regionsregionimagesfiledelete)                                        | **Delete** /regions/{region}/images/{file}                    | Delete an Image.                                                             |
| _ImagesApi_        | [**RegionsRegionImagesFileHead**](docs/ImagesApi.md#regionsregionimagesfilehead)                                            | **Head** /regions/{region}/images/{file}                      | Get upload offset. See https://tus.io/ for more detail.                      |
| _ImagesApi_        | [**RegionsRegionImagesFilePatch**](docs/ImagesApi.md#regionsregionimagesfilepatch)                                          | **Patch** /regions/{region}/images/{file}                     | Upload and apply bytes to a Image file. See https://tus.io/ for more detail. |
| _ImagesApi_        | [**RegionsRegionImagesGet**](docs/ImagesApi.md#regionsregionimagesget)                                                      | **Get** /regions/{region}/images                              | List of all Images.                                                          |
| _ImagesApi_        | [**RegionsRegionImagesImportPost**](docs/ImagesApi.md#regionsregionimagesimportpost)                                        | **Post** /regions/{region}/images/import                      | Import image from url.                                                       |
| _ImagesApi_        | [**RegionsRegionImagesMarketplaceGet**](docs/ImagesApi.md#regionsregionimagesmarketplaceget)                                | **Get** /regions/{region}/images/marketplace                  | List of all MarketPlace Items.                                               |
| _ImagesApi_        | [**RegionsRegionImagesPost**](docs/ImagesApi.md#regionsregionimagespost)                                                    | **Post** /regions/{region}/images                             | Request a new upload Image. See https://tus.io/ for more detail.             |
| _NetworksApi_      | [**RegionsRegionNetworksGet**](docs/NetworksApi.md#regionsregionnetworksget)                                                | **Get** /regions/{region}/networks                            | List of Networks and details.                                                |
| _NetworksApi_      | [**RegionsRegionNetworksIdAttachPatch**](docs/NetworksApi.md#regionsregionnetworksidattachpatch)                            | **Patch** /regions/{region}/networks/{id}/attach              | Attach a Network to a Server.                                                |
| _NetworksApi_      | [**RegionsRegionNetworksIdDetachPatch**](docs/NetworksApi.md#regionsregionnetworksiddetachpatch)                            | **Patch** /regions/{region}/networks/{id}/detach              | Detach a Network from a Server.                                              |
| _NetworksApi_      | [**RegionsRegionSubnetsIdDelete**](docs/NetworksApi.md#regionsregionsubnetsiddelete)                                        | **Delete** /regions/{region}/subnets/{id}                     | Delete Subnet.                                                               |
| _NetworksApi_      | [**RegionsRegionSubnetsIdGet**](docs/NetworksApi.md#regionsregionsubnetsidget)                                              | **Get** /regions/{region}/subnets/{id}                        | Get Subnet details.                                                          |
| _NetworksApi_      | [**RegionsRegionSubnetsIdPatch**](docs/NetworksApi.md#regionsregionsubnetsidpatch)                                          | **Patch** /regions/{region}/subnets/{id}                      | Update the Subnet.                                                           |
| _NetworksApi_      | [**RegionsRegionSubnetsPost**](docs/NetworksApi.md#regionsregionsubnetspost)                                                | **Post** /regions/{region}/subnets                            | Create new Subnet.                                                           |
| _PortsApi_         | [**RegionsRegionPortsIdDisablePatch**](docs/PortsApi.md#regionsregionportsiddisablepatch)                                   | **Patch** /regions/{region}/ports/{id}/disable                | Disable port.                                                                |
| _PortsApi_         | [**RegionsRegionPortsIdEnablePatch**](docs/PortsApi.md#regionsregionportsidenablepatch)                                     | **Patch** /regions/{region}/ports/{id}/enable                 | Enable port.                                                                 |
| _PtrApi_           | [**RegionsRegionPtrIpDelete**](docs/PtrApi.md#regionsregionptripdelete)                                                     | **Delete** /regions/{region}/ptr/{ip}                         | Delete PTR record                                                            |
| _PtrApi_           | [**RegionsRegionPtrPost**](docs/PtrApi.md#regionsregionptrpost)                                                             | **Post** /regions/{region}/ptr                                | Add PTR record                                                               |
| _QuotaApi_         | [**RegionsRegionQuotaGet**](docs/QuotaApi.md#regionsregionquotaget)                                                         | **Get** /regions/{region}/quota                               | List of limits for the account.                                              |
| _RegionsApi_       | [**RegionsGet**](docs/RegionsApi.md#regionsget)                                                                             | **Get** /regions                                              | List of all regions.                                                         |
| _ReportsApi_       | [**RegionsRegionReportsIdGet**](docs/ReportsApi.md#regionsregionreportsidget)                                               | **Get** /regions/{region}/reports/{id}                        | List of all resource reports of a server.                                    |
| _ReportsApi_       | [**RegionsRegionReportsIdMetricGet**](docs/ReportsApi.md#regionsregionreportsidmetricget)                                   | **Get** /regions/{region}/reports/{id}/{metric}               | List of reports of a server for a special metric.                            |
| _SSHKeysApi_       | [**RegionsRegionSshKeysGet**](docs/SSHKeysApi.md#regionsregionsshkeysget)                                                   | **Get** /regions/{region}/ssh-keys/                           | List of all SSH Keys.                                                        |
| _SSHKeysApi_       | [**RegionsRegionSshKeysNameDelete**](docs/SSHKeysApi.md#regionsregionsshkeysnamedelete)                                     | **Delete** /regions/{region}/ssh-keys/{name}                  | Delete SSH Key.                                                              |
| _SSHKeysApi_       | [**RegionsRegionSshKeysPost**](docs/SSHKeysApi.md#regionsregionsshkeyspost)                                                 | **Post** /regions/{region}/ssh-keys                           | Create new SSH Key.                                                          |
| _SecurityGroupApi_ | [**RegionsRegionSecuritiesCdnPost**](docs/SecurityGroupApi.md#regionsregionsecuritiescdnpost)                               | **Post** /regions/{region}/securities/cdn                     | Create CDN Security Group.                                                   |
| _SecurityGroupApi_ | [**RegionsRegionSecuritiesGet**](docs/SecurityGroupApi.md#regionsregionsecuritiesget)                                       | **Get** /regions/{region}/securities                          | List of Security Group and Rules.                                            |
| _SecurityGroupApi_ | [**RegionsRegionSecuritiesIdDelete**](docs/SecurityGroupApi.md#regionsregionsecuritiesiddelete)                             | **Delete** /regions/{region}/securities/{id}                  | Delete the Security Group.                                                   |
| _SecurityGroupApi_ | [**RegionsRegionSecuritiesPost**](docs/SecurityGroupApi.md#regionsregionsecuritiespost)                                     | **Post** /regions/{region}/securities                         | Create new Security Group.                                                   |
| _SecurityGroupApi_ | [**RegionsRegionSecuritiesSecurityRulesIdDelete**](docs/SecurityGroupApi.md#regionsregionsecuritiessecurityrulesiddelete)   | **Delete** /regions/{region}/securities/security-rules/{id}   | Delete Security Group Rule.                                                  |
| _SecurityGroupApi_ | [**RegionsRegionSecuritiesSecurityRulesIdGet**](docs/SecurityGroupApi.md#regionsregionsecuritiessecurityrulesidget)         | **Get** /regions/{region}/securities/security-rules/{id}      | List of all Rules for a Security Groups.                                     |
| _SecurityGroupApi_ | [**RegionsRegionSecuritiesSecurityRulesIdPost**](docs/SecurityGroupApi.md#regionsregionsecuritiessecurityrulesidpost)       | **Post** /regions/{region}/securities/security-rules/{id}     | Create new Rule for a Security Group.                                        |
| _ServerActionsApi_ | [**RegionsRegionServersIdActionsGet**](docs/ServerActionsApi.md#regionsregionserversidactionsget)                           | **Get** /regions/{region}/servers/{id}/actions                | Actions list of server.                                                      |
| _ServerActionsApi_ | [**RegionsRegionServersIdAddSecurityGroupPost**](docs/ServerActionsApi.md#regionsregionserversidaddsecuritygrouppost)       | **Post** /regions/{region}/servers/{id}/add-security-group    | Assign Server to Security Group.                                             |
| _ServerActionsApi_ | [**RegionsRegionServersIdChangePublicIpPost**](docs/ServerActionsApi.md#regionsregionserversidchangepublicippost)           | **Post** /regions/{region}/servers/{id}/change-public-ip      | change public ips.                                                           |
| _ServerActionsApi_ | [**RegionsRegionServersIdHardRebootPost**](docs/ServerActionsApi.md#regionsregionserversidhardrebootpost)                   | **Post** /regions/{region}/servers/{id}/hard-reboot           | Apply hard reboot action for the Server.                                     |
| _ServerActionsApi_ | [**RegionsRegionServersIdPowerOffPost**](docs/ServerActionsApi.md#regionsregionserversidpoweroffpost)                       | **Post** /regions/{region}/servers/{id}/power-off             | Apply shutdown action for the Server.                                        |
| _ServerActionsApi_ | [**RegionsRegionServersIdPowerOnPost**](docs/ServerActionsApi.md#regionsregionserversidpoweronpost)                         | **Post** /regions/{region}/servers/{id}/power-on              | Apply turn on action for the Server.                                         |
| _ServerActionsApi_ | [**RegionsRegionServersIdRebootPost**](docs/ServerActionsApi.md#regionsregionserversidrebootpost)                           | **Post** /regions/{region}/servers/{id}/reboot                | Apply soft reboot action for the Server.                                     |
| _ServerActionsApi_ | [**RegionsRegionServersIdRebuildPost**](docs/ServerActionsApi.md#regionsregionserversidrebuildpost)                         | **Post** /regions/{region}/servers/{id}/rebuild               | Apply rebuild action for the Server.                                         |
| _ServerActionsApi_ | [**RegionsRegionServersIdRemoveSecurityGroupPost**](docs/ServerActionsApi.md#regionsregionserversidremovesecuritygrouppost) | **Post** /regions/{region}/servers/{id}/remove-security-group | Remove the Server from the Security Group.                                   |
| _ServerActionsApi_ | [**RegionsRegionServersIdRenamePost**](docs/ServerActionsApi.md#regionsregionserversidrenamepost)                           | **Post** /regions/{region}/servers/{id}/rename                | Rename the Server.                                                           |
| _ServerActionsApi_ | [**RegionsRegionServersIdRescuePost**](docs/ServerActionsApi.md#regionsregionserversidrescuepost)                           | **Post** /regions/{region}/servers/{id}/rescue                | Apply rescue action for the Server.                                          |
| _ServerActionsApi_ | [**RegionsRegionServersIdResetRootPasswordPost**](docs/ServerActionsApi.md#regionsregionserversidresetrootpasswordpost)     | **Post** /regions/{region}/servers/{id}/reset-root-password   | Apply reset root password action for the Server.                             |
| _ServerActionsApi_ | [**RegionsRegionServersIdResizePost**](docs/ServerActionsApi.md#regionsregionserversidresizepost)                           | **Post** /regions/{region}/servers/{id}/resize                | Apply new size plan for the Server.                                          |
| _ServerActionsApi_ | [**RegionsRegionServersIdSnapshotPost**](docs/ServerActionsApi.md#regionsregionserversidsnapshotpost)                       | **Post** /regions/{region}/servers/{id}/snapshot              | Create snapshot of the Server.                                               |
| _ServerActionsApi_ | [**RegionsRegionServersIdUnrescuePost**](docs/ServerActionsApi.md#regionsregionserversidunrescuepost)                       | **Post** /regions/{region}/servers/{id}/unrescue              | Apply unrescue action for the Server.                                        |
| _ServerActionsApi_ | [**RegionsRegionServersIdVncGet**](docs/ServerActionsApi.md#regionsregionserversidvncget)                                   | **Get** /regions/{region}/servers/{id}/vnc                    | Gets a VNC console for a Server.                                             |
| _ServersApi_       | [**RegionsRegionOptionsGet**](docs/ServersApi.md#regionsregionoptionsget)                                                   | **Get** /regions/{region}/options                             | Return a region options.                                                     |
| _ServersApi_       | [**RegionsRegionServersGet**](docs/ServersApi.md#regionsregionserversget)                                                   | **Get** /regions/{region}/servers                             | Return all Servers.                                                          |
| _ServersApi_       | [**RegionsRegionServersIdDelete**](docs/ServersApi.md#regionsregionserversiddelete)                                         | **Delete** /regions/{region}/servers/{id}                     | Delete a Server.                                                             |
| _ServersApi_       | [**RegionsRegionServersIdGet**](docs/ServersApi.md#regionsregionserversidget)                                               | **Get** /regions/{region}/servers/{id}                        | Shows details of a Server.                                                   |
| _ServersApi_       | [**RegionsRegionServersPost**](docs/ServersApi.md#regionsregionserverspost)                                                 | **Post** /regions/{region}/servers                            | Create new Server.                                                           |
| _SizesApi_         | [**RegionsRegionSizeIdGet**](docs/SizesApi.md#regionsregionsizeidget)                                                       | **Get** /regions/{region}/size/{id}                           | Shows details of a flavor.                                                   |
| _SizesApi_         | [**RegionsRegionSizesCustomGet**](docs/SizesApi.md#regionsregionsizescustomget)                                             | **Get** /regions/{region}/sizes/custom                        | List of custom sizes.                                                        |
| _SizesApi_         | [**RegionsRegionSizesGet**](docs/SizesApi.md#regionsregionsizesget)                                                         | **Get** /regions/{region}/sizes                               | List of general plans.                                                       |
| _TagsApi_          | [**RegionsRegionTagsBatchPost**](docs/TagsApi.md#regionsregiontagsbatchpost)                                                | **Post** /regions/{region}/tags/batch                         | Replace a list of tags with instance tags (for a list of instances)          |
| _TagsApi_          | [**RegionsRegionTagsGet**](docs/TagsApi.md#regionsregiontagsget)                                                            | **Get** /regions/{region}/tags                                | Return all user tags.                                                        |
| _TagsApi_          | [**RegionsRegionTagsIdAttachPost**](docs/TagsApi.md#regionsregiontagsidattachpost)                                          | **Post** /regions/{region}/tags/{id}/attach                   | Attach tag to an instance                                                    |
| _TagsApi_          | [**RegionsRegionTagsIdDelete**](docs/TagsApi.md#regionsregiontagsiddelete)                                                  | **Delete** /regions/{region}/tags/{id}                        | Delete a tag.                                                                |
| _TagsApi_          | [**RegionsRegionTagsIdDetachPost**](docs/TagsApi.md#regionsregiontagsiddetachpost)                                          | **Post** /regions/{region}/tags/{id}/detach                   | Detach tag from an instance                                                  |
| _TagsApi_          | [**RegionsRegionTagsIdPut**](docs/TagsApi.md#regionsregiontagsidput)                                                        | **Put** /regions/{region}/tags/{id}                           | Edit a tag.                                                                  |
| _TagsApi_          | [**RegionsRegionTagsPost**](docs/TagsApi.md#regionsregiontagspost)                                                          | **Post** /regions/{region}/tags                               | Create a new tag for user.                                                   |
| _VolumesApi_       | [**RegionsRegionVolumesAttachPatch**](docs/VolumesApi.md#regionsregionvolumesattachpatch)                                   | **Patch** /regions/{region}/volumes/attach                    | Attach a Volume to a Server.                                                 |
| _VolumesApi_       | [**RegionsRegionVolumesDetachPatch**](docs/VolumesApi.md#regionsregionvolumesdetachpatch)                                   | **Patch** /regions/{region}/volumes/detach                    | Detach a Volume from a Server.                                               |
| _VolumesApi_       | [**RegionsRegionVolumesGet**](docs/VolumesApi.md#regionsregionvolumesget)                                                   | **Get** /regions/{region}/volumes                             | List of Volumes.                                                             |
| _VolumesApi_       | [**RegionsRegionVolumesIdDelete**](docs/VolumesApi.md#regionsregionvolumesiddelete)                                         | **Delete** /regions/{region}/volumes/{id}                     | Delete the Volume.                                                           |
| _VolumesApi_       | [**RegionsRegionVolumesIdGet**](docs/VolumesApi.md#regionsregionvolumesidget)                                               | **Get** /regions/{region}/volumes/{id}                        | Show details of a Volume.                                                    |
| _VolumesApi_       | [**RegionsRegionVolumesIdPatch**](docs/VolumesApi.md#regionsregionvolumesidpatch)                                           | **Patch** /regions/{region}/volumes/{id}                      | Update a Volume.                                                             |
| _VolumesApi_       | [**RegionsRegionVolumesLimitsGet**](docs/VolumesApi.md#regionsregionvolumeslimitsget)                                       | **Get** /regions/{region}/volumes/limits                      | Show general limits of Volumes.                                              |
| _VolumesApi_       | [**RegionsRegionVolumesPost**](docs/VolumesApi.md#regionsregionvolumespost)                                                 | **Post** /regions/{region}/volumes                            | Create new Volume.                                                           |

## Documentation For Models

- [CustomSize](docs/CustomSize.md)
- [Distribution](docs/Distribution.md)
- [Flavor](docs/Flavor.md)
- [FloatIp](docs/FloatIp.md)
- [HostRoute](docs/HostRoute.md)
- [Image](docs/Image.md)
- [ImageContainerFormat](docs/ImageContainerFormat.md)
- [ImageDiskFormat](docs/ImageDiskFormat.md)
- [ImageDocument](docs/ImageDocument.md)
- [ImageMetaData](docs/ImageMetaData.md)
- [ImageStatus](docs/ImageStatus.md)
- [ImgListDistribution](docs/ImgListDistribution.md)
- [IpAllocationPool](docs/IpAllocationPool.md)
- [MarketPlaceItem](docs/MarketPlaceItem.md)
- [Network](docs/Network.md)
- [Options](docs/Options.md)
- [Quota](docs/Quota.md)
- [Region](docs/Region.md)
- [SecurityGroup](docs/SecurityGroup.md)
- [SecurityGroupRule](docs/SecurityGroupRule.md)
- [Server](docs/Server.md)
- [ServerAction](docs/ServerAction.md)
- [ServerFlavor](docs/ServerFlavor.md)
- [ServerQuota](docs/ServerQuota.md)
- [ServerSecurityGroup](docs/ServerSecurityGroup.md)
- [ServerSnapshot](docs/ServerSnapshot.md)
- [ServerStatus](docs/ServerStatus.md)
- [SshKey](docs/SshKey.md)
- [Subnet](docs/Subnet.md)
- [Tag](docs/Tag.md)
- [Volume](docs/Volume.md)
- [VolumeAttachment](docs/VolumeAttachment.md)
- [VolumeLimits](docs/VolumeLimits.md)
- [VolumeStatus](docs/VolumeStatus.md)

## Documentation For Authorization

## api_key

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author
