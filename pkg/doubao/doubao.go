package doubao

import (
	"context"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"os"
	"strings"
)

var cli *arkruntime.Client

func init() {
	cli = arkruntime.NewClientWithApiKey(os.Getenv("ARK_API_KEY"))
}

func GenChart(dest string, data string) (chartData string, analysis string, err error) {
	prompt := fmt.Sprintf("分析需求：\n%s\n原始数据：\n%s", dest, data)

	req := model.ChatCompletionRequest{
		Model: "ep-20241204185325-qkmbn",
		Messages: []*model.ChatCompletionMessage{
			{
				Role: model.ChatMessageRoleSystem,
				Content: &model.ChatCompletionMessageContent{
					StringValue: volcengine.String("你是一个数据分析师和前端开发专家，接下来我会按照以下固定格式给你提供内容：" +
						"\n分析需求：\n{数据分析的需求或者目标}\n" +
						"原始数据：\n{csv格式的原始数据，用,作为分隔符}\n" +
						"请根据这两部分内容，按照以下指定格式生成内容（此外不要输出任何多余的开头、结尾、注释）\n" +
						"【【【【【\n{前端 Echarts V5 的 option 配置对象json代码，合理地将数据进行可视化，不要生成任何多余的内容，比如注释}\n" +
						"【【【【【\n{明确的数据分析结论、越详细越好，结论不大于200字}"),
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

	return spl[1], spl[2], nil
}
