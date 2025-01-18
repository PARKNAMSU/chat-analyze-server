package platform_model

// 파트너 플랫폼 정보
type PartnerPlatform struct {
	platformId       int
	domain           string
	userKeyType      string
	name             string
	apiKey           string
	permissionBit    int
	maxChatNum       int
	maxAccessNum     int
	maxUserNum       int
	isExitUserDelete int
}

// PlatformId methods
func (p *PartnerPlatform) PlatformId() int {
	return p.platformId
}

func (p *PartnerPlatform) SetPlatformId(id int) {
	p.platformId = id
}

// Domain methods
func (p *PartnerPlatform) Domain() string {
	return p.domain
}

func (p *PartnerPlatform) SetDomain(domain string) {
	p.domain = domain
}

// UserKeyType methods
func (p *PartnerPlatform) UserKeyType() string {
	return p.userKeyType
}

func (p *PartnerPlatform) SetUserKeyType(keyType string) {
	p.userKeyType = keyType
}

// Name methods
func (p *PartnerPlatform) Name() string {
	return p.name
}

func (p *PartnerPlatform) SetName(name string) {
	p.name = name
}

// ApiKey methods
func (p *PartnerPlatform) ApiKey() string {
	return p.apiKey
}

func (p *PartnerPlatform) SetApiKey(apiKey string) {
	p.apiKey = apiKey
}

// PermissionBit methods
func (p *PartnerPlatform) PermissionBit() int {
	return p.permissionBit
}

func (p *PartnerPlatform) SetPermissionBit(bit int) {
	p.permissionBit = bit
}

// MaxChatNum methods
func (p *PartnerPlatform) MaxChatNum() int {
	return p.maxChatNum
}

func (p *PartnerPlatform) SetMaxChatNum(num int) {
	p.maxChatNum = num
}

// MaxAccessNum methods
func (p *PartnerPlatform) MaxAccessNum() int {
	return p.maxAccessNum
}

func (p *PartnerPlatform) SetMaxAccessNum(num int) {
	p.maxAccessNum = num
}

// MaxUserNum methods
func (p *PartnerPlatform) MaxUserNum() int {
	return p.maxUserNum
}

func (p *PartnerPlatform) SetMaxUserNum(num int) {
	p.maxUserNum = num
}

// IsExitUserDelete methods
func (p *PartnerPlatform) IsExitUserDelete() int {
	return p.isExitUserDelete
}

func (p *PartnerPlatform) SetIsExitUserDelete(flag int) {
	p.isExitUserDelete = flag
}
