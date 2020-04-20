const baseUrl = "/accounting-api?"
const config = {headers: {'Content-Type': 'application/json'}}
const pieColors = ["#22AC38", "#009944", "#009B6B", "#009E96", "#00A0C1", "#00A0E9", "#0086D1", "#0068B7", "#00479D", "#1D2088"]

var app = new Vue({
    el:"#app",
    data:{
        showSideMenu:false,
        showSummary:true,
        showInput:false,
        showDetail:false,
        dataset:[],
        fromDate:null,
        toDate:null
    },
    methods:{
        // 変数を初期化
        init(){
            this.showSummary=false;
            this.showInput = false;
            this.showDetail=false;
        },
        // APIに情報を問い合わせる
        clicked(){
            // パラメータとして指定するためフォーマット整形(YYYY-MM-DD => YYYYMMDD) 
            const fromStr = this.fromDate.slice(0,4) + this.fromDate.slice(5,7) + this.fromDate.slice(8,10)
            const toStr = this.toDate.slice(0,4) + this.toDate.slice(5,7) + this.toDate.slice(8,10)
            const getUrl = baseUrl + "from=" + fromStr + "&to=" + toStr;
            axios.get(getUrl, config).then((response) => {
                let rawData = response.data;
                for (let data of rawData){
                    data['datetime_date'] = new Date(data['date'])
		    data['date'] = data['datetime_date'].getFullYear() + '/' + (data['datetime_date'].getMonth() + 1) + '/' + data['datetime_date'].getDate();
                }
                rawData.sort(function(a,b) {
                    return (a.datetime_date < b.datetime_date ? -1 : 1);
                });
                this.dataset = rawData;
                console.log("DEBUG")
                console.log(this.dataset)
                this.showSideMenu = false;
            })
        },
        // サイドメニューを開閉する
        clickBar(){
            this.showSideMenu = !this.showSideMenu;
        }
    },
    created(){
        const today = new Date()
        let MM = today.getMonth() + 1;
        if(today.getMonth() < 9){
            MM = "0" + (today.getMonth() + 1);
        }
        const fromDateTemp = new Date(today.getFullYear(), today.getMonth(), 1)
        this.fromDate = fromDateTemp.getFullYear() + "-" + MM + "-01";
        const nextMonth = today.getMonth() + 1
        const toDateTemp = new Date(today.getFullYear(), nextMonth, 0)
        this.toDate = toDateTemp.getFullYear() + "-" + MM + "-" + toDateTemp.getDate();
        this.clicked();
    },
})
