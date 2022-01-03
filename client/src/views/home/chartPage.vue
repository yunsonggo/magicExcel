<template>
    <div>
        <Chart></Chart>
        <p @click="handleBack">fanhui</p>
    </div>
</template>

<script>
import Chart from "@/components/common/chart.vue"
export default {
    name: "chartPage",
    components:{
        Chart
    },
    methods:{
        handleBack() {
            this.$router.back()
        }
    }
}
</script>

<style scoped>

</style>