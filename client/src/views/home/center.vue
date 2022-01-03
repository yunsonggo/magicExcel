<template>
    <div>
        <div class="page-container page-component">
            <section class="content element-doc content"  style="margin: 50px;border: 1px solid #f4f4f4;">
                <h4>修改密码</h4>
                <div style="width: 400px">
                    <el-form
                        :model="ruleForm"
                        status-icon
                        :rules="rules"
                        ref="ruleForm"
                        label-width="0px"
                        class="demo-ruleForm"
                    >
                        <el-form-item prop="oldPass" required>
                            <el-input
                                type="password"
                                v-model="ruleForm.oldPass"
                                autocomplete="off"
                                placeholder="原始密码"
                            ></el-input>
                        </el-form-item>
                        <el-form-item prop="pass" required>
                            <el-input
                                type="password"
                                v-model="ruleForm.pass"
                                autocomplete="off"
                                placeholder="新密码"
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
                        <el-form-item>
                            <el-button @click="resetPassword">提交</el-button>
                            <el-button @click="resetForm('ruleForm')">重置</el-button>
                        </el-form-item>
                    </el-form>
                </div>
                <div class="demo-block demo-zh-CN demo-button" style="margin-top: 50px">
                    <h4>数据备份</h4>
                    <el-button @click="backupSql" size="small" type="danger" plain> 点击备份数据 </el-button>
                </div>
            </section>
        </div>
    </div>
</template>

<script>
import {reqResetPassword} from "@/api"
export default {
name: "center",
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
            ruleForm: {
                pass: "",
                checkPass:"",
                oldPass: "",
            },
            rules: {
                pass: [{ validator: validatePass, trigger: "blur" }],
                checkPass: [{ validator: validatePass2, trigger: "blur" }],
                oldPass: [{ required: true, message: "请输入原始密码", trigger: "blur" }],
            },
            imgCaptcha: "",
        }
    },
    methods:{
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
        async resetPassword() {
            const result = await reqResetPassword(this.ruleForm.oldPass,this.ruleForm.pass,this.ruleForm.checkPass)
            console.log(result)
            if (result.code === 1000) {
                this.$message.success("修改成功,请重新登录")
                localStorage.removeItem('token')
                this.$router.push("/user/login")
            } else {
                this.$message.error(result.msg)
            }
        },
        backupSql() {
            this.$message.info("暂未开发此功能")
        }
    }
}
</script>

<style scoped>

</style>