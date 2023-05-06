package role

// 将身份代号转为名称
// 0-申请队员 1-岗前培训 2-见习队员 3-正式队员 4-督导老师 30-普通队员 31-核心队员 32-区域负责人 33-组委会成员 34-组委会主任
func GetRoleName(role int64) string {
	switch role {
	case 0:
		return "申请队员"
	case 1:
		return "岗前培训"
	case 2:
		return "见习队员"
	case 3:
		return "正式队员"
	case 4:
		return "督导老师"
	case 30:
		return "普通队员"
	case 31:
		return "核心队员"
	case 32:
		return "区域负责人"
	case 33:
		return "组委会成员"
	case 34:
		return "组委会主任"
	default:
		return "申请队员"
	}
}
