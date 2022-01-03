<template>
    <div class="navbar">
        <el-menu
            :default-active="activeIndex"
            class="el-menu-demo"
            mode="horizontal"
            @select="handleSelect"
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b">
            <el-menu-item index="/home/index">上传数据</el-menu-item>
            <el-menu-item index="/home/online">在线数据</el-menu-item>
            <el-menu-item index="/home/center">用户中心</el-menu-item>
<!--            <el-menu-item index="4">订单管理</el-menu-item>-->
<!--            <el-menu-item index="5" disabled>消息中心</el-menu-item>-->
        </el-menu>
    </div>
</template>

<script>
export default {
    name: "navBar",
    data() {
        return {
            activeIndex: '',
        }
    },
    watch:{
        $route(val) {
            this.activeIndex = val.path
        }
    },
    methods:{
        handleSelect(key) {
            this.$router.push(key)
        }
    }
}
</script>

<style scoped>

</style>