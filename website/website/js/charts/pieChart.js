Vue.component('pie-chart', {
    template: `
    <div class="summary-content" style="position: relative; max-width:100vw; height:60vh">
        <canvas id="pieChart"></canvas>
    </div>
    `,
    data: function () {
        return {
            "dataList": [],
            "labelList": [],
            "colorList": []
        }
    },
    props: [
        "dataset"
    ],
    watch: {
        dataset: function () {
            this.init();
            this.editPieChartData();
        },
        labelList: function () {
            this.createPieChart();
        }
    },
    methods: {
        // 変数の初期化
        init() {
            this.dataList = [];
            this.labelList = [];
            this.colorList = [];
        },
        // パイチャート表示用のデータを作成
        editPieChartData() {
            rankedByAmount = [];
            for (let data of this.dataset) {
                if (data.ifearning == 0) {
                    if (data.type in rankedByAmount) {
                        rankedByAmount[data.type]["data"] += data.amount;
                    } else {
                        rankedByAmount[data.type] = { "data": data.amount, "name": data.type };
                    }
                }
            }
            rankedByAmount = Object.values(rankedByAmount)
            rankedByAmount.sort(function (a, b) {
                if (a.data < b.data) return 1;
                if (a.data > b.data) return -1;
                return 0;
            });
            count = 0;
            for (let data of rankedByAmount) {
                if (count < 10) {
                    this.dataList.push(data.data);
                    // ラベルの文字数が長ければ省略
                    if(data.name.length > 10){
                        this.labelList.push(data.name.slice(1,11) + "...");
                    }else{
                        this.labelList.push(data.name);
                    }
                    this.colorList.push(pieColors[count])
                } else {
                    let addedAmount = this.dataList[9] + data.amount;
                    this.dataList.splice(9, 1, addedAmount)
                    this.labelList[9].splice(9, 1, "その他")
                }
                count += 1
            }
        },
        // パイチャートの描画
        createPieChart() {
            //グラフ描画
            let config = {
                type: 'doughnut',
                data: {
                    labels: this.labelList,
                    datasets: [{
                        data: this.dataList,
                        backgroundColor: this.colorList
                    }],
                },
                options: {
                    legend: {
                        position: "bottom",
                        labels: {
                            fontSize: 16,
                        },
                        padding: 0
                    },
                    maintainAspectRatio: false,
                }
            };
            chart = new Chart(document.getElementById('pieChart').getContext('2d'), config);
        }
    },
    created() {
        this.init();
        this.editPieChartData();
        document.addEventListener("resize", this.createPieChart);
    },
})
