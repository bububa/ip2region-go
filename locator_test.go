package ip2region

import (
	"reflect"
	"testing"
)

func TestIp2Long(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		ipstr string
		// Expected results.
		want uint
	}{
		{
			name:  "ip2long test",
			ipstr: "106.38.53.110",
			want:  1780888942,
		},
	}
	for _, tt := range tests {
		got := ip2long(tt.ipstr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ip2long() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSearh(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		ipstr string
		// Expected results.
		want    Location
		wantErr bool
	}{
		{
			name:  "Search 北京 test",
			ipstr: "106.38.53.110",
			want: Location{
				Country:  "中国",
				Region:   "华北",
				Province: "北京市",
				City:     "北京市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 上海 test",
			ipstr: "114.80.166.240",
			want: Location{
				Country:  "中国",
				Region:   "华东",
				Province: "上海市",
				City:     "上海市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 台湾 test",
			ipstr: "123.193.103.252",
			want: Location{
				Country:  "台湾",
				Region:   "0",
				Province: "台湾省",
				City:     "0",
				ISP:      "Taiwan Fixed Network",
			},
		},
		{
			name:  "Search 广东 test",
			ipstr: "113.111.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华南",
				Province: "广东省",
				City:     "广州市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 山东 test",
			ipstr: "27.223.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华东",
				Province: "山东省",
				City:     "青岛市",
				ISP:      "联通",
			},
		},
		{
			name:  "Search 浙江 test",
			ipstr: "183.159.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华东",
				Province: "浙江省",
				City:     "杭州市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 江苏 test",
			ipstr: "49.95.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华东",
				Province: "江苏省",
				City:     "南京市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 辽宁 test",
			ipstr: "175.175.255.252",
			want: Location{
				Country:  "中国",
				Region:   "东北",
				Province: "辽宁省",
				City:     "葫芦岛市",
				ISP:      "联通",
			},
		},
		{
			name:  "Search 四川 test",
			ipstr: "182.143.255.252",
			want: Location{
				Country:  "中国",
				Region:   "西南",
				Province: "四川省",
				City:     "成都市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 河南 test",
			ipstr: "115.63.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华中",
				Province: "河南省",
				City:     "焦作市",
				ISP:      "联通",
			},
		},
		{
			name:  "Search 湖北 test",
			ipstr: "27.31.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华中",
				Province: "湖北省",
				City:     "孝感市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 福建 test",
			ipstr: "120.39.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华东",
				Province: "福建省",
				City:     "三明市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 湖南 test",
			ipstr: "222.247.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华中",
				Province: "湖南省",
				City:     "长沙市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 河北 test",
			ipstr: "106.119.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华北",
				Province: "河北省",
				City:     "唐山市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 重庆 test",
			ipstr: "125.87.255.252",
			want: Location{
				Country:  "中国",
				Region:   "西南",
				Province: "重庆市",
				City:     "重庆市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 山西 test",
			ipstr: "118.79.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华北",
				Province: "山西省",
				City:     "忻州市",
				ISP:      "联通",
			},
		},
		{
			name:  "Search 江西 test",
			ipstr: "115.151.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华东",
				Province: "江西省",
				City:     "萍乡市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 陕西 test",
			ipstr: "117.34.255.252",
			want: Location{
				Country:  "中国",
				Region:   "西北",
				Province: "陕西省",
				City:     "商洛市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 安徽 test",
			ipstr: "60.175.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华东",
				Province: "安徽省",
				City:     "宿州市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 黑龙江 test",
			ipstr: "222.172.127.252",
			want: Location{
				Country:  "中国",
				Region:   "东北",
				Province: "黑龙江省",
				City:     "大庆市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 广西 test",
			ipstr: "171.39.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华南",
				Province: "广西壮族自治区",
				City:     "梧州市",
				ISP:      "联通",
			},
		},
		{
			name:  "Search 吉林 test",
			ipstr: "111.117.255.252",
			want: Location{
				Country:  "中国",
				Region:   "东北",
				Province: "吉林省",
				City:     "长春市",
				ISP:      "教育网",
			},
		},
		{
			name:  "Search 云南 test",
			ipstr: "182.247.255.252",
			want: Location{
				Country:  "中国",
				Region:   "西南",
				Province: "云南省",
				City:     "昆明市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 天津 test",
			ipstr: "180.213.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华北",
				Province: "天津市",
				City:     "天津市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 内蒙古 test",
			ipstr: "222.74.255.252",
			want: Location{
				Country:  "中国",
				Region:   "华北",
				Province: "内蒙古自治区",
				City:     "赤峰市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 新疆 test",
			ipstr: "124.119.255.252",
			want: Location{
				Country:  "中国",
				Region:   "西北",
				Province: "新疆维吾尔自治区",
				City:     "阿克苏地区",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 甘肃 test",
			ipstr: "210.26.0.71",
			want: Location{
				Country:  "中国",
				Region:   "西北",
				Province: "甘肃省",
				City:     "兰州市",
				ISP:      "教育网",
			},
		},
		{
			name:  "Search 贵州 test",
			ipstr: "114.139.255.252",
			want: Location{
				Country:  "中国",
				Region:   "西南",
				Province: "贵州省",
				City:     "黔南布依族苗族自治州",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 海南 test",
			ipstr: "121.58.127.252",
			want: Location{
				Country:  "中国",
				Region:   "华南",
				Province: "海南省",
				City:     "省直辖县级行政区划",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 宁夏 test",
			ipstr: "124.224.255.252",
			want: Location{
				Country:  "中国",
				Region:   "西北",
				Province: "宁夏回族自治区",
				City:     "固原市",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 青海 test",
			ipstr: "175.184.191.252",
			want: Location{
				Country:  "中国",
				Region:   "西北",
				Province: "青海省",
				City:     "西宁市",
				ISP:      "联通",
			},
		},
		{
			name:  "Search 西藏 test",
			ipstr: "27.98.255.252",
			want: Location{
				Country:  "中国",
				Region:   "西南",
				Province: "西藏自治区",
				City:     "拉萨市",
				ISP:      "联通",
			},
		},
		{
			name:  "Search 香港 test",
			ipstr: "202.97.111.252",
			want: Location{
				Country:  "香港",
				Region:   "0",
				Province: "香港特别行政区",
				City:     "0",
				ISP:      "电信",
			},
		},
		{
			name:  "Search 澳门 test",
			ipstr: "60.246.255.252",
			want: Location{
				Country:  "澳门",
				Region:   "0",
				Province: "澳门特别行政区",
				City:     "0",
				ISP:      "CTM",
			},
		},
	}
	locator, err := NewLocator("./data/ip2region.db")
	if err != nil {
		t.Errorf("Search() error = %v", err)
		return
	}
	for _, tt := range tests {
		got := locator.Search(tt.ipstr)
		if !reflect.DeepEqual(*got, tt.want) {
			t.Errorf("%q. Search() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
