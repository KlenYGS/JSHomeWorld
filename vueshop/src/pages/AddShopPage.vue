<template>
  <div>
      <table>
          <tr>
              <td>商品名字</td>
              <td><input type="text" v-model="name"></td>
          </tr>
          <tr>
              <td><label>宣传标语</label></td>
              <td><input type="text" v-model="shopTitleText" /></td>
          </tr>
          <tr>
            <td>
                <label>商品宣传图</label>
                <input type="file" ref="TitleImage" @change="addImage($event)" />
            </td>
          </tr>
          <tr>
            <td>
                <label>店内简介</label>
            </td>
            <td>
                <textarea v-model="text"></textarea>
            </td>
          </tr>
          <tr>
              <td><label>店内宣传图</label></td>
              <td><input type="file" multiple="multiple" @change="addImages($event)"></td>
          </tr>
          <tr>
              <td><label>商品价格</label></td>
              <td><input type="text" v-model.number="price" /></td>
          </tr>
          <tr>
              <td><label>商品数量</label></td>
              <td><input type="text" v-model.number="having" /></td>
          </tr>
          <tr>
              <td><label>店内宣传视频</label></td>
              <td><input type="file" @change="addVideo($event)"></td>
          </tr>
          <tr>
              <td>发货地址</td>
              <td><input type="text" v-model="addr"></td>
          </tr>
          <tr>
              <td>选项名字</td>
              <td><input type="text" v-model="ChooseText"></td>
          </tr>
          <tr>
              <td>选项</td>
              <td>
                  <div v-for="_,index in chooses" :key="index">
                    <input type="text" v-model="chooses[index]">
                  </div>
              </td>
          </tr>
          <tr>
              <td><button @click="addChoose()">+</button></td>
          </tr>
          <tr>
              <img ref="Image" />
          </tr>
      </table>
      <button @click="update()">上传</button>
      <button @click="get()">get</button>
      
  </div>
</template>

<script>
import axios from 'axios'

export default {
    data() {
        return {
            name:"",
            shopTitleText:"",
            price:0,
            having:0,
            text:"",
            titleImageData: null,
            imagesData: null,
            videoData: null,
            ChooseText: "",
            chooses: [],
            addr:""
        }
    },
    mounted() {
        document.title = "增加商店"
    },
    methods: {
        get(){
            axios({
                method:"Get",
                url:"http://127.0.0.1:8000/addImageTest",
                responseType: "blob"
            }).then( res=> {
                
                this.$refs.Image.src =URL.createObjectURL(new Blob([res.data],{type:"image/*"}))
            } )

        },
        update() {
            axios({
                method:"Post",
                url:"http://127.0.0.1:8000/addShop",
                data: {
                    name: this.name,
                    shopTitleText: this.shopTitleText,
                    text: this.text,
                    having: this.having,
                    price: this.price,
                    Address: this.addr
                }
            }).then( res => {
                let shopId = res.data.id
                let data = new FormData()
                if (this.titleImageData != null) {
                    data.append("ImageData",this.titleImageData)
                    data.append("flag","true")
                    data.append("shopId",shopId)
                    axios.post("http://127.0.0.1:8000/addImage",
                    data,
                    )
                }
                 let imageDatas = new FormData()
                 imageDatas.append("flag","false")
                 imageDatas.append("shopId",shopId)
                 if (this.imagesData != null) {
                    for (let a = 0; a < this.imagesData.length; a++) {
                        imageDatas.set("ImageData",this.imagesData[a])
                        axios.post("http://127.0.0.1:8000/addImage",imageDatas)
                    }
                }
                 let videoData = new FormData()
                 videoData.append("video", this.videoData)
                 videoData.append("ShopId", shopId)
                 axios.post("http://127.0.0.1:8000/video/addVideo", videoData)
                 if (this.ChooseText != "") {
                axios.post("http://127.0.0.1:8000/shopChoose", {
                    ShopId: shopId,
                    ChooseTip: this.ChooseText,
                    ChooseTipList: JSON.stringify(this.chooses)
                })
            }
            } )
        },
        addImage(event) {
           this.titleImageData = event.target.files[0]
        },
        addImages(event) {
            this.imagesData = event.target.files
        },
        addVideo(event) {
            this.videoData = event.target.files[0]
        },
        addChoose() {
            this.chooses.push("")
        }
    }
}
</script>

<style>

</style>