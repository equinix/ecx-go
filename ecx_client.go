//Package ecx implements ECX fabric client
package ecx

//Client describes operations provided by ECX Fabric client module
type Client interface {
	GetUserPorts() ([]Port, error)
	GetL2Connection(uuid string) (*L2Connection, error)
	CreateL2Connection(conn L2Connection) (*L2Connection, error)
	CreateL2RedundantConnection(priConn L2Connection, secConn L2Connection) (*L2Connection, error)
	NewL2ConnectionUpdateRequest(uuid string) L2ConnectionUpdateRequest
	DeleteL2Connection(uuid string) error
	ConfirmL2Connection(uuid string, confirmConn L2ConnectionToConfirm) (*L2ConnectionConfirmation, error)

	GetL2SellerProfiles() ([]L2ServiceProfile, error)
	GetL2ServiceProfile(uuid string) (*L2ServiceProfile, error)
	CreateL2ServiceProfile(sp L2ServiceProfile) (*L2ServiceProfile, error)
	UpdateL2ServiceProfile(sp L2ServiceProfile) (*L2ServiceProfile, error)
	DeleteL2ServiceProfile(uuid string) error
}

//L2ConnectionUpdateRequest describes composite request to update given Layer2 connection
type L2ConnectionUpdateRequest interface {
	WithName(name string) L2ConnectionUpdateRequest
	WithBandwidth(speed int, speedUnit string) L2ConnectionUpdateRequest
	Execute() error
}

//Error describes ECX Fabric error that occurs during API call processing
type Error struct {
	//ErrorCode is short error identifier
	ErrorCode string
	//ErrorMessage is textual description of an error
	ErrorMessage string
}

//L2Connection describes layer 2 connection managed by ECX Fabric
type L2Connection struct {
	UUID                string
	Name                string
	ProfileUUID         string
	Speed               int
	SpeedUnit           string
	Status              string
	Notifications       []string
	PurchaseOrderNumber string
	PortUUID            string
	DeviceUUID          string
	VlanSTag            int
	VlanCTag            int
	NamedTag            string
	AdditionalInfo      []L2ConnectionAdditionalInfo
	ZSidePortUUID       string
	ZSideVlanSTag       int
	ZSideVlanCTag       int
	SellerRegion        string
	SellerMetroCode     string
	AuthorizationKey    string
	RedundantUUID       string
}

//L2ConnectionAdditionalInfo additional info object used in L2 connections
type L2ConnectionAdditionalInfo struct {
	Name  string
	Value string
}

//L2ConnectionToConfirm accepts the hosted connection in the seller side
type L2ConnectionToConfirm struct {
	AccessKey string
	SecretKey string
}

//L2ConnectionConfirmation describes a connection confirmed
type L2ConnectionConfirmation struct {
	PrimaryConnectionID string
	Message             string
}

//L2ServiceProfile describes layer 2 service profile managed by ECX Fabric
type L2ServiceProfile struct {
	UUID                                string
	State                               string
	AlertPercentage                     float64
	AllowCustomSpeed                    bool
	AllowOverSubscription               bool
	APIAvailable                        bool
	AuthKeyLabel                        string
	ConnectionNameLabel                 string
	CTagLabel                           string
	EnableAutoGenerateServiceKey        bool
	EquinixManagedPortAndVlan           bool
	Features                            L2ServiceProfileFeatures
	IntegrationID                       string
	Name                                string
	OnBandwidthThresholdNotification    []string
	OnProfileApprovalRejectNotification []string
	OnVcApprovalRejectionNotification   []string
	OverSubscription                    string
	Ports                               []L2ServiceProfilePort
	Private                             bool
	PrivateUserEmails                   []string
	RequiredRedundancy                  bool
	SpeedBands                          []L2ServiceProfileSpeedBand
	SpeedFromAPI                        bool
	TagType                             string
	VlanSameAsPrimary                   bool
	Description                         string
	Metros                              []L2SellerProfileMetro
	AdditionalInfos                     []L2SellerProfileAdditionalInfo
}

//L2ServiceProfilePort describes port used in L2 service profile
type L2ServiceProfilePort struct {
	ID        string
	MetroCode string
}

//L2ServiceProfileSpeedBand describes speed / bandwidth used in L2 service profile
type L2ServiceProfileSpeedBand struct {
	Speed     int
	SpeedUnit string
}

//L2ServiceProfileFeatures describes features used in L2 service profile
type L2ServiceProfileFeatures struct {
	CloudReach  bool
	TestProfile bool
}

//Port describes ECX Fabric's user port
type Port struct {
	UUID          string
	Name          string
	Region        string
	IBX           string
	MetroCode     string
	Priority      string
	Encapsulation string
	Buyout        bool
	Bandwidth     string
	Status        string
}

//L2SellerProfileMetro describces details of a metro in which service provices is present
type L2SellerProfileMetro struct {
	Code    string
	Name    string
	IBXes   []string
	Regions map[string]string
}

//L2SellerProfileAdditionalInfo describces additional information that might be provided by service buyer when using given seller profile
type L2SellerProfileAdditionalInfo struct {
	Name             string
	Description      string
	DataType         string
	IsMandatory      bool
	IsCaptureInEmail bool
}
