## 1.2.0 (Unreleased)

NOTES:

* this version of module started to use `equinix/rest-go` client
for any REST interactions with Equinix APIs
* ECX names were removed from descriptions and documentation in favor
of Equinix Fabric name

FEATURES:

* **L2Connection**: `func GetL2OutgoingConnections()` gives possibility to fetch
 all a-side (outgoing) connections for a customer account associated with
authenticated application

ENHANCEMENTS:

* **L2Connection** added additional attributes:
  * *Actions* provide details about pending actions to complete connection provisioning
  * *DeviceInterfaceID* indicates network interface identifier on a network device
  * *ProviderStatus* indicates connection status on a z-side
  * *RedundancyType* indicates whether connection is primary or secondary
  (for redundant connections)

## 1.1.0 (September 22, 2020)

FEATURES:

* **L2Connection** can be created with device identifier (in addition to port identifier)
 to allow interconnections with Network Edge devices

ENHANCEMENTS:

* **L2ServiceProfile** model and fetch logic was enriched with additional data
 useful when fetching seller profiles:
  * additional information that can be provided when creating connection
  * seller's metro locations
  * profile encapsulation
  * global organization and organization names

## 1.0.0 (July 31, 2020)

NOTES:

* first version of Equinix Cloud Exchange Fabric Go client

FEATURES:

* **L2ServiceProfile**: possibility to create, fetch, update (name and bandwidth),
 remove private and public service profiles
* **L2ServiceProfile**: possibility to fetch seller service profiles.
* **L2Connection**: possibility to create, fetch, update, remove ECX Fabric
 layer 2 connections
* **L2Connection**: possibility to approve layer2 connection with provider's
 access and secret keys (AWS use case)
* **UserPort**: possibly to fetch list of user ports
