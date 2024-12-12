use txnbi;

create table if not exists chart_example
(
    id           bigint auto_increment comment 'id' primary key,
    goal				 text  null comment '分析目标',
    `name`               varchar(128) null comment '图表名称',
    chartType	   varchar(128) null comment '图表类型',
    genChart		 text	 null comment '生成的图表数据',
    genResult		 text	 null comment '生成的分析结论',
    status       varchar(128) not null default 'wait' comment 'wait,running,succeed,failed',
    createTime   datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    updateTime   datetime     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
    ) comment '示例图表表' collate = utf8mb4_unicode_ci;



INSERT INTO chart_example (
    id,
    goal,
    `name`,
    chartType,
    genChart,
    genResult,
    status,
    createTime,
    updateTime
) VALUES
      (
       1,
          '分析一下网站用户量数据趋势',
          'xx网站用户量表',
          '折线图',
          '{
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
          }',
          '通过分析可知，网站用户量整体呈波动趋势。其中 4 号用户数达到峰值 90，而 5 号骤降至 0。其余日期用户数相对较为平稳，整体变化较为明显，需要进一步分析 4 号用户数暴增以及 5 号骤降的原因，以便更好地优化网站运营策略。',
          'succeed','2024-12-11 23:49:28','2024-12-11 23:49:28'
      ),
      (
       2,
          '分析一下网站用户量数据趋势',
          'xx网站用户量表',
          '柱状图',
          '{
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
          }',
          '通过对网站用户量数据的分析，可以看出在 4 号出现了用户量的高峰，达到 90。整体趋势有波动，5 号用户数骤降为 0 可能是特殊情况。其他日期用户数相对较为平稳，整体呈现出一定的不规律性。需要进一步分析 5 号用户量为 0 的原因以及如何保持或提升用户量。',
          'succeed','2024-12-11 15:59:32','2024-12-11 15:59:32'
      ),
      (
       3,
          '分析一下平台营业额数据趋势',
          'YY网站营业额表',
          '柱状图',
          '{
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
          }',
          '从数据可以看出，营业额呈现波动变化，整体较为不稳定。其中 4 日营业额较高为 60，3 日营业额为 0 较为特殊。访问量和观看量也有波动，但与营业额的关联不明显。整体趋势不太规律，可能受多种因素影响，如促销活动、市场环境等。',
          'succeed','2024-12-11 15:59:32','2024-12-11 15:59:32'
      ),
      (
       4,
          '分析一下平台整体数据趋势',
          'YY网站营业额表',
          '折线图',
          '{
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
          }',
          '通过对平台整体数据的分析，可以看出营业额整体呈现波动上升的趋势，但存在个别日期营业额为 0 的情况。访问量相对较为不稳定，有较大波动。观看量整体呈上升趋势，其中 5 日观看量达到峰值。需要进一步分析营业额为 0 的原因以及访问量波动的影响因素。',
          'succeed','2024-12-11 15:59:32','2024-12-11 15:59:32'
      );
