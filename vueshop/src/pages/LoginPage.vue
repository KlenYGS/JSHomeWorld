<template>
  <div class="loginBack">
    <div id="inputPage">
        <div class="inputs">
            <div class="put"><form>
                <table id="registerPage">
                    <tr class="loginTr">
                        <td class="loginTd"><label>手机号</label></td>
                        <td><input type="text" v-model="phone" class="loginInput" ref="registerPhone" @blur="registerPhoneTest()"> </td>
                    </tr>
                    <tr class="loginTr">
                        <td class="loginTd"><label>昵称</label></td>
                        <td><input type="text" v-model="name" class="loginInput" ref="registerName" @blur="registerNameTest()"></td>
                    </tr>
                    <tr class="loginTr">
                        <td class="loginTd"><label>密码</label></td>
                        <input type="password" v-model="password" class="loginInput" ref="registerPassword" @blur="registerPasswordTest()">
                    </tr>
                    <tr class="loginTr">
                        <td class="loginTd"><label>确认密码</label></td>
                        <td><input type="password" v-model="password1" class="loginInput" ref="registerPassword1" @blur="registerPassword1Test()"></td>
                    </tr>
                </table>
            </form>
            <div id="registerButtonDiv">
                <input type="radio"><label>同意<a href="#">《用户使用》</a>协议</label><br/>
                <br/>
                <button class="firstPageButton" id="RegisterButon" @click="register()">注册</button>
            </div></div>
        </div>
        <div class="inputs" id="loginInputs">
            <div class="put" id="loginPut">
                <form>
                    <table>
                        <tr>
                            <td><label>手机号</label></td>
                            <td><input type="text" v-model="phone" class="loginInput" ref="loginPhone" @blur="loginPhoneTest()"></td>
                        </tr>
                        <tr>
                            <td><label>密码</label></td>
                            <td><input type="password" v-model="password" class="loginInput" ref="loginPassword" @blur="loginPasswordTest()"></td>
                        </tr>
                    </table>
                </form>
                <button ref="loginButton" id="loginButton" @click="login()" @mouseenter="LoginMouseEnter()" @mouseleave="LoginMouseLeave()">登录</button>
            </div>
        </div>
        <div ref="cover" id="cover">
            <button id="toRegister" v-show="logining" @click="toRegister()">注册账号</button>
            <button id="toLogin" v-show="registing" @click="toLogin()">登录</button>
        </div>
    </div>

  </div>
</template>

<script>
import axios from 'axios'

export default {
    name: "LoginPage",
    data(){
        return{
            phone:"",
            name:"",
            password:"",
            password1:"",
            logining : true,
            registing: false,
        }
    },
    methods: {
        resetAll() {
            this.phone = ""
            this.name = ""
            this.password = ""
            this.password1 = ""
            this.$refs.registerPhone.style["border"] = ''
            this.$refs.registerName.style["border"] = ''
            this.$refs.registerPassword.style["border"] = ''
            this.$refs.registerPassword1.style["border"] = ''
            this.$refs.loginPhone.style["border"] = ''
            this.$refs.loginPassword.style["border"] = ''
        },
        register() {
            if (!(this.registerPhoneTest() && this.registerNameTest() && this.registerPasswordTest() && this.registerPassword1Test())) {
                return
            }
            axios({
                method:'Post',
                url:"http://127.0.0.1:8000/register",
                 data:{
                     phone:this.phone,
                     name:this.name,
                     password:this.password,
                 }
            }).then(res => {
                if (res.data.register == true) {
                    this.toLogin()
                }
            })
        },
        login(){
            if (!(this.loginPhoneTest() && this.loginPasswordTest())) {
                return
            }
            axios({
                method:"Post",
                url:"http://127.0.0.1:8000/login",
                withCredentials: true,
                data:{
                    phone: this.phone,
                    password: this.password
                },
            }).then(res => {
                console.log(document.cookie)
                if (res.data.Request == true) {
                    this.$router.push({name:"SelectPage",params:{Phone:this.phone}})
                }
            })

        },
        toRegister() {
            this.resetAll()
            this.logining = !this.logining
            this.registing = !this.registing
            let pos = 0
            let cover = this.$refs.cover
            let timer = requestAnimationFrame(function fn(){
                cover.style["left"] = String(pos) + 'px'
                if (pos < 50) {
                    pos += 10
                } else if (pos < 200) {
                    pos += 20
                } else if (pos < 300) {
                    pos += 30
                } else {
                    pos += 10
                }
                if (pos >= 390) {
                    cancelAnimationFrame(timer)
                } else {
                    timer = requestAnimationFrame(fn)
                }
            })
        },
        toLogin() {
            this.resetAll()
            this.logining = !this.logining
            this.registing = !this.registing
            let pos = 390
            let cover = this.$refs.cover
            let timer = requestAnimationFrame(function fn(){
                cover.style["left"] = String(pos) + 'px'
                if (pos > 340) {
                    pos -= 10
                } else if (pos > 190) {
                    pos -= 20
                } else if (pos > 90) {
                    pos -= 30
                } else {
                    pos -= 10
                }
                if (pos < 0) {
                    cancelAnimationFrame(timer)
                } else {
                    timer = requestAnimationFrame(fn)
                }
            })
        },
        registerPhoneTest() {
            if (this.phone.length != 11) {
                this.$refs.registerPhone.style["border"] = "2px groove red"
                return false
            } else {
                this.$refs.registerPhone.style["border"] = ''
                return true
            }
        },
        registerNameTest() {
            let rex = /^[a-zA-Z0-9]{1,10}$/
            if (rex.test(this.name)) {
                this.$refs.registerName.style["border"] = ''
                return true
            } else {
                this.$refs.registerName.style["border"] = "2px groove red"
                return false
            }
        },
        registerPasswordTest() {
            let rex = /^[a-zA-Z0-9]{6,15}$/
            if (rex.test(this.password)) {
                this.$refs.registerPassword.style["border"] = ''
                return true
            } else {
                this.$refs.registerPassword.style["border"] = "2px groove red"
                return false
            }
        },
        registerPassword1Test() {
            if (this.password != this.password1) {
                this.$refs.registerPassword1.style["border"] = "2px groove red"
                return false
            } else {
                this.$refs.registerPassword1.style["border"] = ''
                return true
            }
        },
        loginPhoneTest() {
            if (this.phone.length != 11) {
                this.$refs.loginPhone.style["border"] = "2px groove red"
                return false
            } else {
                this.$refs.loginPhone.style["border"] = ''
                return true
            }
        },
        loginPasswordTest() {
            let rex = /^[a-zA-Z0-9]{6,15}$/
            if (rex.test(this.password)) {
                this.$refs.loginPassword.style["border"] = ''
                return true
            } else {
                this.$refs.loginPassword.style["border"] = "2px groove red"
                return false
            }
        },
        LoginMouseEnter() {
            this.$refs.loginButton.style["background-color"] = "rgb(42, 120, 194)"

        },
        LoginMouseLeave() {
            this.$refs.loginButton.style["background-color"] = "rgb(14, 86, 159)"
        }
    }
}
</script>


<style>

d{
    color: rgb(42, 120, 194);
}

.loginBack {
    width: 100%;
    height: 800px;
    background-image: url(../../public/static/loginBack.jpg);
    background-repeat: no-repeat;
    background-position-x: center;
}

#inputPage {
    position: relative;
    width: 750px;
    height: 450px;
    top: 150px;
    left: 350px;
    background-color: rgb(250, 250, 250);
    border-radius: 5px;
    box-shadow: 0 1px 4px gray;
}

#loginInputs {
    right: 0px;
}

#loginPut {
    padding-left: 25px;
    padding-top: 40px;
}

.inputs {
    pointer-events: none;
    position: absolute;
    display: inline-block;
    width: 375px;
    height: 450px;
}

.put {
    padding-top: 20px;
    pointer-events: all;
}

#cover {
    width: 375px;
    height: 450px;
    position: absolute;
    background-color: rgb(236, 223, 223);
    background-image: url(../assets/weichat.png);
    background-repeat: no-repeat;
    background-position-x: center;
    background-position-y: center;
    background-size: auto;
}

#registerPage {
    margin-left: 50px;
}

.loginTd {
    text-align: right;
}

.loginInput {
    margin-top: 10px;
    margin-bottom: 10px;
    margin-left: 10px;
    height: 30px;
    width: 200px;
    border-radius: 12px;
    border: 2px groove rgb(228, 213, 189);
}

.loginInput:focus {
    margin-top: 10px;
    margin-bottom: 10px;
    margin-left: 10px;
    height: 30px;
    width: 200px;
    border-radius: 12px;
    border: 2px groove rgb(241, 224, 197);
    box-shadow: inset 0 0 5px gray;
}

#registerButtonDiv {
    margin-top: 30px;
    margin-left: 100px;
}

.firstPageButton {
    margin-left: 80px;
    width: 120px;
    height: 45px;
    background-color: rgb(141, 181, 200);
    border-radius: 15px;
    border: 1px groove rgb(191, 220, 226);
}

#loginButton, #RegisterButon {
    cursor: pointer;
    margin-left: 75px;
    margin-top: 30px;
    width: 150px;
    height: 40px;
    background-color: rgb(14, 86, 159);
    border-radius: 5px;
    border: 1px groove rgb(178, 236, 238);
    color: rgb(242, 238, 238);
    font-family:'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif
}

#RegisterButon {
    margin-top: 20px;
    margin-left: 50px;
}

#toRegister {
    cursor: pointer;
    position: relative;
    top: 410px;
    left: 30px;
    font-size: 15px;
    color: rgb(56, 100, 167);
    background-color: rgba(0, 0, 0, 0);
    border: 0 inset gray;
    outline: none;
}

#toRegister:hover {
    cursor: pointer;
    position: relative;
    top: 410px;
    left: 30px;
    font-size: 15px;
    text-decoration: underline;
    color: rgb(56, 100, 167);
    background-color: rgba(0, 0, 0, 0);
    border: 0 inset gray;
    outline: none;
}

#toLogin {
    cursor: pointer;
    position: relative;
    top: 410px;
    left: 310px;
    font-size: 15px;
    color: rgb(56, 100, 167);
    background-color: rgba(0, 0, 0, 0);
    border: 0 inset gray;
    outline: none;
}

#toLogin:hover {
    cursor: pointer;
    position: relative;
    top: 410px;
    left: 310px;
    font-size: 15px;
    text-decoration: underline;
    color: rgb(56, 100, 167);
    background-color: rgba(0, 0, 0, 0);
    border: 0 inset gray;
    outline: none;
}

</style>