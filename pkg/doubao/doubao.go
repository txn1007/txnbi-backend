package doubao

import (
	"context"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"os"
	"strings"
	"unicode"
)

var cli *arkruntime.Client

func init() {
	cli = arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
}

func GenChart(dest string, data string, chartType string) (chartData string, analysis string, err error) {
	prompt := fmt.Sprintf("分析需求：\n%s\n原始数据：\n%s", dest, data)

	req := model.ChatCompletionRequest{
		Model: "ep-20241204185325-qkmbn",
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleSystem,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(fmt.Sprintf("你是一个数据分析师和前端开发专家，接下来我会按照以下固定格式给你提供内容："+
						"\n分析需求：\n{数据分析的需求或者目标}\n"+
						"原始数据：\n{csv格式的原始数据，用,作为分隔符}\n"+
						"请根据这两部分内容，按照以下指定格式生成内容（此外不要输出任何多余的开头、结尾、注释）\n"+
						"【【【【【\n{前端 Echarts V5 %s 的 option 配置对象json代码，合理地将数据进行可视化，用户可以选中特定字段进行查看，"+
						"不要生成任何多余的内容，比如注释、图表名。必须包含配置项legend,grid,xAxis,yAxis}\n"+
						"【【【【【\n{明确的数据分析结论、越详细越好，结论不大于200字}", chartType)),
				},
			},
			{
				Role: model.ChatMessageRoleUser,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String(prompt),
				},
			},
		},
	}

	ctx := context.Background()
	resp, err := cli.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", "", fmt.Errorf("standard chat error: %v\n", err)
	}

	respContent := *resp.Choices[0].Message.Content.StringValue
	spl := strings.Split(respContent, "【【【【【")
	if len(spl) < 3 {
		return "", "", fmt.Errorf("生成的字符串不合法")
	}

	// 对字符串切分结果进行分析
	for _, v := range spl {
		// 去除换行
		t := strings.Trim(v, "\n")
		if len(t) < 1 {
			continue
		}
		if t[0] == '{' {
			chartData = t
		}
		// 第一个字符为汉字或英文字母
		if unicode.Is(unicode.Han, []rune(t)[0]) || unicode.IsLetter([]rune(t)[0]) {
			analysis = t
		}
	}

	return chartData, analysis, nil
}
