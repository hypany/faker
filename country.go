package faker

// CountryMatrix 国家母体
type CountryMatrix = EnumMatrix

// NewCountryMatrix 构造一个 CountryMatrix
func NewCountryMatrix() *CountryMatrix {
	vals := []string{"阿富汗", "奥兰群岛", "阿尔巴尼亚", "阿尔及利亚", "美属萨摩亚", "安道尔", "安哥拉", "安圭拉", "安提瓜和巴布达", "阿根廷", "亚美尼亚",
		"阿鲁巴", "澳大利亚", "奥地利", "阿塞拜疆", "孟加拉", "巴林", "巴哈马", "巴巴多斯", "白俄罗斯", "比利时", "伯利兹", "贝宁", "百慕大", "不丹", "玻利维亚",
		"波斯尼亚和黑塞哥维那", "博茨瓦纳", "布维岛", "巴西", "文莱", "保加利亚", "布基纳法索", "布隆迪", "柬埔寨", "喀麦隆", "加拿大", "佛得角", "中非", "乍得",
		"智利", "圣诞岛", "科科斯（基林）群岛", "哥伦比亚", "科摩罗", "刚果（金）", "刚果", "库克群岛", "哥斯达黎加", "科特迪瓦", "中国", "克罗地亚", "古巴",
		"捷克", "塞浦路斯", "丹麦", "吉布提", "多米尼加", "东帝汶", "厄瓜多尔", "埃及", "赤道几内亚", "厄立特里亚", "爱沙尼亚", "埃塞俄比亚", "法罗群岛", "斐济",
		"芬兰", "法国", "法国大都会", "法属圭亚那", "法属波利尼西亚", "加蓬", "冈比亚", "格鲁吉亚", "德国", "加纳", "直布罗陀", "希腊", "格林纳达", "瓜德罗普岛",
		"关岛", "危地马拉", "根西岛", "几内亚比绍", "几内亚", "圭亚那", "海地", "洪都拉斯", "匈牙利", "冰岛", "印度", "印度尼西亚", "伊朗", "伊拉克", "爱尔兰",
		"马恩岛", "以色列", "意大利", "牙买加", "日本", "泽西岛", "约旦", "哈萨克斯坦", "肯尼亚", "基里巴斯", "韩国", "朝鲜", "科威特", "吉尔吉斯斯坦", "老挝",
		"拉脱维亚", "黎巴嫩", "莱索托", "利比里亚", "利比亚", "列支敦士登", "立陶宛", "卢森堡", "马其顿", "马拉维", "马来西亚", "马达加斯加", "马尔代夫", "马里",
		"马耳他", "马绍尔群岛", "马提尼克岛", "毛里塔尼亚", "毛里求斯", "马约特", "墨西哥", "密克罗尼西亚", "摩尔多瓦", "摩纳哥", "蒙古", "黑山", "蒙特塞拉特",
		"摩洛哥", "莫桑比克", "缅甸", "纳米比亚", "瑙鲁", "尼泊尔", "荷兰", "新喀里多尼亚", "新西兰", "尼加拉瓜", "尼日尔", "尼日利亚", "纽埃", "诺福克岛",
		"挪威", "阿曼", "巴基斯坦", "帕劳", "巴勒斯坦", "巴拿马", "巴布亚新几内亚", "秘鲁", "菲律宾", "皮特凯恩群岛", "波兰", "葡萄牙", "波多黎各", "卡塔尔",
		"留尼汪岛", "罗马尼亚", "卢旺达", "俄罗斯联邦", "圣赫勒拿", "圣基茨和尼维斯", "圣卢西亚", "圣文森特和格林纳丁斯", "萨尔瓦多", "萨摩亚", "圣马力诺",
		"圣多美和普林西比", "沙特阿拉伯", "塞内加尔", "塞舌尔", "塞拉利昂", "新加坡", "塞尔维亚", "斯洛伐克", "斯洛文尼亚", "所罗门群岛", "索马里", "南非",
		"西班牙", "斯里兰卡", "苏丹", "苏里南", "斯威士兰", "瑞典", "瑞士", "叙利亚", "塔吉克斯坦", "坦桑尼亚", "泰国", "特立尼达和多巴哥", "东帝汶", "多哥",
		"托克劳", "汤加", "突尼斯", "土耳其", "土库曼斯坦", "图瓦卢", "乌干达", "乌克兰", "阿拉伯联合酋长国", "英国", "美国", "乌拉圭", "乌兹别克斯坦", "瓦努阿图",
		"梵蒂冈", "委内瑞拉", "越南", "瓦利斯群岛和富图纳群岛", "西撒哈拉", "也门", "南斯拉夫", "赞比亚", "津巴布韦"}
	matrix, _ := NewEnumMatrix(vals)
	return matrix
}
