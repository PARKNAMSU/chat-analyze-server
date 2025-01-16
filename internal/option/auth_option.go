package option

type MembershipOption struct {
	Permission int
	MaxChatNum int
}

var (
	Permissions = map[string]int{
		"all":          1 << 0,
		"oneToOneChat": 1 << 1,
		"groupChat":    1 << 2,
		"fileUpload":   1 << 3,
		"secretChat":   1 << 4,
	}

	MEMBERSHIP_PREMIUM = MembershipOption{
		Permission: Permissions["all"],
		MaxChatNum: 5000,
	}

	MEMBERSHIP_NORMAL = MembershipOption{
		Permission: Permissions["oneToOneChat"] +
			Permissions["groupChat"],
		MaxChatNum: 1000,
	}
)
