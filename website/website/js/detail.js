Vue.component('detail', {
    template: `
    <div class="detail-area">
        <div class="detail-modal-window" :class="{ hidden: !showModal }">
            <div class="go-back-button"><i class="fas fa-times-circle pointable" @click="hideModal()"></i></div>
            <post v-if="showModal" :showndata="shownData" v-on:success="hideModal(); updateData()"></post>
        </div>
        <table class="detail-table">
            <tr>
                <th>日付</th>
                <th>金額</th>
                <th>種別</th>
            </tr>
            <tr v-for="data in dataset" :class="[data.ifearning == 0 ? 'outgo' : 'income']" @click="onClickItem(data)">
                <td><div class="detail-table-cell">{{data.date}}</div></td>
                <td><div class="detail-table-cell">{{data.amount}}</div></td>
                <td><div class="detail-table-cell">{{data.type}}</div></td>
            </tr>
        </table>
    </div>
    `,
    props: [
        "dataset"
    ],
    data: function () {
        return {
            "showModal":false,
            "shownData":{}
        }
    },
    methods:{
        /*
        * モーダルウィンドウを表示する
        * @params data {object} - 選択された取引データ
        */
        onClickItem(data){
            this.shownData = data;
            this.showModal = true;
        },
        /*
        * モーダルウィンドウを隠す
        */
        hideModal(){
            this.showModal = false;
        },
        /*
        * データの削除後、画面を更新する
        */
        updateData(){
            this.$emit("deleted");
        }
    }
})
