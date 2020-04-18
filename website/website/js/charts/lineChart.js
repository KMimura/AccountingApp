Vue.component('line-chart', {
    template: `
    <div class="summary-content" style="position: relative; max-width:80vw; height:40vh; margin:auto;">
        <canvas id="lineChart"></canvas>
    </div>
    `,
    data: function () {
        return {
            "dataList": [],
            "labelList": [],
        }
    },
    props: [
        "dataset"
    ],
    watch: {
        dataset: function () {
            this.init();
            this.editLineChartData();
        },
        labelList: function () {
            this.createLineChart();
        }
    },
    methods: {
        // 変数の初期化
        init() {
            this.dataList = [];
            this.labelList = [];
        },
        // パイチャート表示用のデータを作成
        editLineChartData() {
            let prevDate = ""
            let dataListTemp = []
            let labelListTemp = []

            for (let data of this.dataset){
                if(data["ifearning"] == 1){
                    continue;
                }
                //前のデータと同一の日付であれば、加算する
                if (prevDate == data["date"]){
                    dataListTemp[dataListTemp.length - 1] += data["amount"]
                }else{
                    dataListTemp.push(data["amount"])
                    labelListTemp.push((data["datetime_date"].getMonth() + 1 ) + "月"  + data["datetime_date"].getDate() + "日")
                }
                prevDate = data["date"]
            }
            this.dataList = dataListTemp;
            this.labelList = labelListTemp;
        },
        // パイチャートの描画
        createLineChart() {
            //グラフ描画
            let config = {
                type: 'line',
                data: {
                    labels: this.labelList,
                    datasets:[{
                        label:"支出",
                        borderColor: "#3cba9f",
                        data:this.dataList,
                        fill:true,
                        lineTension:0.1,
                        borderWidth:4,
                        pointRadius:1,
                        borderJoinStyle:"round"
                    }]
                },
                options: {
                    legend: {
                        display: false
                    },
                    maintainAspectRatio: false,
                    scales: {
                        xAxes: [{
                            ticks: {
                                fontSize: 10
                            }
                        }]
                    }
                }
            };
            chart = new Chart(document.getElementById('lineChart').getContext('2d'), config);
        }
    },
    created() {
        this.init();
        this.editLineChartData();
        document.addEventListener("resize", this.createPieChart);
    },
})
