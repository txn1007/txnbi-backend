package store

import "txnbi-backend/api"

var ExampleChart = []api.ChartInfoV0{
	{
		ChartID:   61,
		ChartName: "xx网站用户量表",
		ChartGoal: "分析一下网站用户量数据趋势",
		ChartType: "折线图",
		ChartCode: `{
                "legend": {
                    "data": []
                },
                "grid": {
                    "left": "3%",
                    "right": "4%",
                    "bottom": "3%",
                    "containLabel": true
                },
                "xAxis": {
                    "type": "category",
                    "boundaryGap": false,
                    "data": ["1 号","2 号","3 号","4 号","5 号","6 号","7 号"]
                },
                "yAxis": {
                    "type": "value"
                },
                "series": [
                    {
                        "name": "用户数",
                        "type": "line",
                        "data": [10,20,30,90,0,10,20]
                    }
                ]
            }`,
		ChartResult: "通过分析可知，网站用户量整体呈波动趋势。其中 4 号用户数达到峰值 90，而 5 号骤降至 0。其余日期用户数相对较为平稳，整体变化较为明显，需要进一步分析 4 号用户数暴增以及 5 号骤降的原因，以便更好地优化网站运营策略。",
		UpdateTime:  "2024-12-11 23:49:28",
	},
	{
		ChartID:   62,
		ChartName: "xx网站用户量表",
		ChartGoal: "分析一下网站用户量数据趋势",
		ChartType: "柱状图",
		ChartCode: `{
                "legend": {
                    "data": ["用户数"]
                },
                "grid": {
                    "left": "3%",
                    "right": "4%",
                    "bottom": "3%",
                    "containLabel": true
                },
                "xAxis": {
                    "type": "category",
                    "boundaryGap": false,
                    "data": ["1 号","2 号","3 号","4 号","5 号","6 号","7 号"]
                },
                "yAxis": {
                    "type": "value"
                },
                "series": [
                    {
                        "name": "用户数",
                        "type": "bar",
                        "data": [10,20,30,90,0,10,20]
                    }
                ]
            }`,
		ChartResult: "通过对网站用户量数据的分析，可以看出在 4 号出现了用户量的高峰，达到 90。整体趋势有波动，5 号用户数骤降为 0 可能是特殊情况。其他日期用户数相对较为平稳，整体呈现出一定的不规律性。需要进一步分析 5 号用户量为 0 的原因以及如何保持或提升用户量。",
		UpdateTime:  "2024-12-11 15:59:32",
	},
	{
		ChartID:   63,
		ChartName: "YY网站营业额表",
		ChartGoal: "分析一下平台营业额数据趋势",
		ChartType: "柱状图",
		ChartCode: `{
                "legend": {
                    "data": []
                },
                "grid": {
                    "left": "3%",
                    "right": "4%",
                    "bottom": "3%",
                    "containLabel": true
                },
                "xAxis": {
                    "type": "category",
                    "boundaryGap": false,
                    "data": ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"]
                },
                "yAxis": {
                    "type": "value"
                },
                "series": [
                    {
                        "name": "营业额",
                        "type": "bar",
                        "data": [10, 30, 0, 60, 34, 44, 55, 11, 22, 33]
                    }
                ]
            }`,
		ChartResult: "从数据可以看出，营业额呈现波动变化，整体较为不稳定。其中 4 日营业额较高为 60，3 日营业额为 0 较为特殊。访问量和观看量也有波动，但与营业额的关联不明显。整体趋势不太规律，可能受多种因素影响，如促销活动、市场环境等。",
		UpdateTime:  "2024-12-11 15:59:32",
	},
	{
		ChartID:   65,
		ChartName: "YY网站营业额表",
		ChartGoal: "分析一下平台整体数据趋势",
		ChartType: "折线图",
		ChartCode: `{
                "legend": {
                    "data": ["营业额", "访问量", "观看量"]
                },
                "grid": {
                    "left": "3%",
                    "right": "4%",
                    "bottom": "3%",
                    "containLabel": true
                },
                "xAxis": {
                    "type": "category",
                    "data": ["1", "2", "3", "4", "5", "6", "7", "8", "9", "10"]
                },
                "yAxis": {
                    "type": "value"
                },
                "series": [
                    {
                        "name": "营业额",
                        "type": "line",
                        "data": [10, 30, 0, 60, 34, 44, 55, 11, 22, 33]
                    },
                    {
                        "name": "访问量",
                        "type": "line",
                        "data": [23, 3, 32, 33, 31, 23, 23, 41, 23, 413]
                    },
                    {
                        "name": "观看量",
                        "type": "line",
                        "data": [223, 123, 324, 132, 543, 123, 342, 324, 23, 321]
                    }
                ]
            }`,
		ChartResult: "通过对平台整体数据的分析，可以看出营业额整体呈现波动上升的趋势，但存在个别日期营业额为 0 的情况。访问量相对较为不稳定，有较大波动。观看量整体呈上升趋势，其中 5 日观看量达到峰值。需要进一步分析营业额为 0 的原因以及访问量波动的影响因素。",
		UpdateTime:  "2024-12-11 15:59:32",
	},
}
