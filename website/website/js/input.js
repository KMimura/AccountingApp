const Url = "/accounting-api"


Vue.component('post', {
    template: `
	<div class="input-area">
		<table class="input_table">
			<tr>
				<th>入力項目</th>
				<th>入力欄</th>
				<th>入力状況</th>
			</tr>
			<tr>
				<td>金額</td>
                <td><input type = "number" v-model="amount" class="user_input_box"> </td>
				<td align="center">
					<div v-if="amountValid"><i class="far fa-check-circle fa-lg my-green"></i></div>
					<div v-if="!amountValid"><i class="far fa-times-circle fa-lg my-red" ></i></div>
				</td>
            </tr>
			<tr>
				<td>日付</td>
				<td><input type = "date" v-model="date" class="user_input_box"> </td>
				<td align="center">
					<div v-if="dateValid"><i class="far fa-check-circle fa-lg my-green"></i></div>
					<div v-if="!dateValid"><i class="far fa-times-circle fa-lg my-red" ></i></div>
				</td>
            </tr>
			<tr>
				<td>種別</td>
				<td><input type = "text" v-model="type" class="user_input_box"> </td>
				<td align="center">
					<div v-if="typeValid"><i class="far fa-check-circle fa-lg my-green"></i></div>
					<div v-if="!typeValid"><i class="far fa-times-circle fa-lg my-red" ></i></div>
				</td>
			</tr>
			<tr>
				<td>収入である</td>
				<td align="center">
					<select name="ifEarning" v-model="ifEarning" class="user_input_box">
						<option value="true">YES</option>
						<option value="false">NO</option>
					</select>
				</td>
				<td align="center">
					<div v-if="ifEarningValid"><i class="far fa-check-circle fa-lg my-green"></i></div>
					<div v-if="!ifEarningValid"><i class="far fa-times-circle fa-lg my-red" ></i></div>
				</td>
			</tr>
			<tr>
				<td>コメント</td>
				<td><input type="text" v-model="comment" class="user_input_box"> </td>
			</tr>
		</table>
		<button class="input-area-button pointable" @click="postData()">送信</button><button class="input-area-button pointable" :class={hidden:!showDeleteButton} @click="deleteData()">削除</button>
	</div>
    `,
    props:[
        "showndata"
    ],
    data: function () {
        return {
            "amount":null,
            "date":null,
            "type":"",
            "ifEarning":"false",
            "comment":"",
            "amountValid":false,
            "dateValid":false,
            "typeValid":false,
            "ifEarningValid":true,
            "showDeleteButton":false
        }
    },
    watch: {
        amount: function (newVal) {
            this.amountValid = this.validateNum(newVal);
        },
        date: function(newVal){
            this.dateValid = this.validateDate(newVal);
        },
        type: function(newVal){
            this.typeValid = this.validateText(newVal);
        },
        ifEarning: function(newVal){
            this.ifEarningValid = this.validateBool(newVal);
        }
    },
    methods: {
        // 数値の入力事項のチェック
        validateNum(num) {
            if(num <= 0){
                return false;
            }
            if(num % 1 != 0){
                return false;
            }
            return true;
        },
        // 文字列の入力事項のチェック
        validateText(num) {
            if (num.length < 2) {
                return false;
            }
            if (num.length > 32) {
                return false;
            }
            return true;
        },
        // 日付の入力事項のチェック
        validateDate(date) {
            try{
                date = new Date(date)
                if(date.getFullYear() <= 2018 || date.getFullYear() > 2050){
                    return false;
                };
                if(date.getMonth() < 0 || date.getMonth() >= 12){
                    return false;
                };
                if(date.getDate() < 1 || date.getDate() > 31){
                    return false;
                }
                return true;
            }catch(err){
                return false;
            }
        },
        // 真偽の入力事項のチェック
        validateBool(bool) {
            return true;
        },
        // データをポスト
        postData(){
            if(!this.amountValid || !this.dateValid || !this.ifEarningValid || !this.typeValid){
                alert("入力が未完です")
                return;
            }
            // 詳細画面から遷移してきた場合は、データをアップデートする
            if(this.showndata){
                axios.post(Url,{
                    amount:this.amount,
                    date:this.date,
                    type:this.type,
                    ifEarning:this.ifEarning,
                    comment:this.comment,
                    id:this.showndata.id
                }).then((response) => {
                    alert("データを更新しました")
                    this.$emit("success");
                })
            }else{
                axios.post(Url,{
                    amount:this.amount,
                    date:this.date,
                    type:this.type,
                    ifEarning:this.ifEarning,
                    comment:this.comment,
                    id:""
                }).then((response) => {
                    alert("データを追加しました")
                    this.$emit("success");
                })
            }
        },
        // データを削除
        deleteData(){
            if(!this.showndata){
                alert("削除するデータがありません")
                return
            }
            axios.delete(Url + "?id=" + this.showndata.id,{data:{id:this.showndata.id}}).then((response) => {
                alert("データを削除しました")
                this.$emit("success");
            })
        }
    },
    created(){
        if(this.showndata){
            // 詳細画面から遷移してきた場合は、「削除」ボタンを表示
            this.showDeleteButton = true;
            this.amount = this.showndata.amount;
            this.date = this.showndata.date;
            this.type = this.showndata.type;
            if(this.showndata.ifearning == 0){
                this.ifEarning = "false"
            }else{
                this.ifEarning = "true"
            }
            this.comment = this.showndata.comment    
        }else{
            this.showDeleteButton = false;
        }
    }
})
