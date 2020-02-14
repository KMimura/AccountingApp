Vue.component('chart', {
    template: `
    <div>
        <div class="summary-content">
            今月の収支：<p class="balance" :class="balancePositive ? 'positive-balance' : 'negative-balance'">{{totalSum}}円</p>
        </div>
        <line-chart :dataset="dataset" v-if="showingChartPosition == 0"></line-chart>
        <pie-chart :dataset="dataset" v-else-if="showingChartPosition == 1"></pie-chart>
        <div class="switchCharts">
            <button class="switchChartsButton pointable" @click="declimentChartPosition()"><<</button>
            <button class="switchChartsButton pointable" @click="inclimentChartPosition()">>></button>
        </div>
    </div>
    `,
    data:function(){
        return {
            "totalSum":0,
            "showingChartPosition":0,
            "balancePositive":false
        }
    },
    props:[
        "dataset"
    ],
    watch: { 
        dataset: function() {
            this.calcSum();
        },
    },
    methods:{
        /*
        * 期間中の収支を計算
        */
        calcSum(){
            this.totalSum = 0;
            for (let data of this.dataset){
                if (data.ifearning == 1){
                    this.totalSum += data.amount;
                }else if(data.ifearning == 0){
                    this.totalSum -= data.amount;
                }
            }    
            if(this.totalSum > 0){
                this.balancePositive = true;
                this.totalSum = "+" + this.totalSum;
            }
        },
        /*
        * チャート切り替えボタン「<<」押下時
        */
        declimentChartPosition(){
            if(this.showingChartPosition > 0){
                this.showingChartPosition -= 1;
            }else{
                this.showingChartPosition = 1;
            }
        },
        /*
        * チャート切り替えボタン「>>」押下時
        */
        inclimentChartPosition(){
        if(this.showingChartPosition < 1){
            this.showingChartPosition += 1;
        }else{
            this.showingChartPosition = 0;
        }
    }
    },
})
