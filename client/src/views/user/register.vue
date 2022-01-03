<template>
    <div>
        <div class="background" >
            <img :src="imgSrc" width="100%" height="100%" alt="" />
        </div>
        <div class="front">
            <el-card class="card">
                <div slot="header" class="title">
                    <span>注册用户</span>
                </div>
                <div>
                    <el-form
                        :model="ruleForm"
                        status-icon
                        :rules="rules"
                        ref="ruleForm"
                        label-width="0px"
                        class="demo-ruleForm"
                    >
                        <el-form-item  prop="name" required>
                            <el-input type="text" v-model="ruleForm.name" autofocus="true" autocomplete="off"  placeholder="用户名" ></el-input>
                        </el-form-item>
                        <el-form-item prop="pass" required>
                            <el-input
                                type="password"
                                v-model="ruleForm.pass"
                                autocomplete="off"
                                placeholder="密码"
                            ></el-input>
                        </el-form-item>
                        <el-form-item prop="checkPass" required>
                            <el-input
                                type="password"
                                v-model="ruleForm.checkPass"
                                autocomplete="off"
                                placeholder="校验密码"
                            ></el-input>
                        </el-form-item>
                        <el-form-item prop="captcha" required style="">
                            <el-input
                                v-model="ruleForm.captcha"
                                type="text"
                                style="width:40%;vertical-align: middle;margin-right:26px"
                                placeholder="验证码"
                            ></el-input>
                            <img
                                style="border:1px solid #d7d7d7;vertical-align: middle;background-color: #f4f4f4"
                                :src="imgCaptcha"
                                @click="handleNewCaptcha"
                            />
                        </el-form-item>
                        <el-form-item>
                            <el-button
                                type="primary"
                                @click="submitForm('ruleForm')"
                                style="width:150px"
                            >注册</el-button
                            >
                            <el-button @click="resetForm('ruleForm')">重置</el-button>
                            <el-button @click="loginButton">用户登录</el-button>
                        </el-form-item>
                    </el-form>
                </div>
                <div class="infomation">
                    <span>声明:公司内部自用,如需备案,请联系:tiandali@email.cn</span>
                    <span>版本:@2022.1.0</span>
                </div>
            </el-card>
        </div>
    </div>
</template>

<script>
import {reqCaptcha, reqRegister } from "@/api"
export default {
    name: "register",
    data() {
        let validatePass = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("请输入密码"));
            } else {
                if (this.ruleForm.checkPass !== "") {
                    this.$refs.ruleForm.validateField("checkPass");
                }
                callback();
            }
        };
        let validatePass2 = (rule, value, callback) => {
            if (value === "") {
                callback(new Error("请再次输入密码"));
            } else if (value !== this.ruleForm.pass) {
                callback(new Error("两次输入密码不一致!"));
            } else {
                callback();
            }
        };
        return {
            imgSrc:require('../../assets/register.jpg'),
            ruleForm: {
                pass: "",
                checkPass:"",
                name: "",
                captcha: "",
                captchaId: "",
            },
            rules: {
                pass: [{ validator: validatePass, trigger: "blur" }],
                checkPass: [{ validator: validatePass2, trigger: "blur" }],
                name: [{ required: true, message: "请输入账号", trigger: "blur" }],
                captcha: [
                    { required: true, message: "请输入验证码", trigger: "blur" },
                    { min: 4, max: 4, message: "请输入四位验证码", trigger: "blur" },
                ],
            },
            imgCaptcha: "",
        }
    },
    mounted() {
        this.getCaptcha();
    },
    methods:{
        // 图形验证码
        async getCaptcha() {
            const result = await reqCaptcha();
            if (result.code === 1000) {
                this.imgCaptcha = result.data.B64s;
                this.ruleForm.captchaId = result.data.Id;
            }
        },
        handleNewCaptcha() {
            this.getCaptcha();
        },
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    this.handleRegister()
                } else {
                    return false;
                }
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        },
        loginButton() {
            this.$router.push('/user/login')
        },
        async handleRegister() {
            const result = await reqRegister(this.ruleForm.name,this.ruleForm.pass,this.ruleForm.checkPass,this.ruleForm.captchaId,this.ruleForm.captcha)
            if (result.code === 1000) {
                this.$message.success("注册成功")
                setTimeout(() => {
                    this.$router.push("/user/login")
                },1000)
            } else {
                this.$message.error("注册失败")
            }
        }
    }
}
</script>

<style scoped>
.background {
    height: 100%;
    width: 100%;
    z-index: 1;
    padding: 0;
    margin: 0;
    position: absolute;
    left: 0;
    top: 0;
    bottom: 0;
    right: 0;
}
.front {
    z-index: 1;
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    color: #ffffff;
}
.card {
    width: 400px;
    background-color: rgba(
        0,
        0,
        0,
        0.4
    ); /**rgba中的a为alpha通道， 为不透明参数，.0即为完全透明*/
    border-color: rgba(0,0, 0, 0.5);
    border-radius: 18px;
    z-index: 2;
}
.title {
    text-align: center;
    font-size: 24px;
    color: #ffffff;
}
</style>