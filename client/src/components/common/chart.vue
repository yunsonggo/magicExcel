<template>
    <div>
        <div>
            <span style="color: #767676;font-size: 14px;margin-left: 15px">筛选部门(非当前表格部门数据为空):</span>
            <el-select
                v-model="tempClass"
                multiple
                placeholder="筛选部门"
                size="small"
                style="width: 140px;margin-left: 15px"
                v-if="classArray.length > 0"
                @change="filterHandler"
                :clearable="true"
            >
                <el-option
                    v-for="item in classArray"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                >
                </el-option>
            </el-select>
        </div>
        <div id="myChart" class="myChartStyle"></div>
    </div>
</template>

<script>
export default {
    name: "chart",
    props:{
        chartData: [],
        classArray:[],
        findClass:[]
    },
    data () {
        return {
            msg: 'Welcome to Your Vue.js App',
            myChart:{},
            cars:[],
            mileages:[],
            oilNum:[],
            oilPay:[],
            repairPay:[],
            per:[],
            tempData:[],
            tempClass:[],
            labelStyle:{
                show:true,
                rotate:90,
                align:'center',
                verticalAlign: 'middle',
                position: 'top',
                distance: 50,
                formatter: '{c}  {name|{a}}',
                fontSize: 14,
                rich: {
                    name: {}
                }
            }
        }
    },
    mounted(){
        this.drawLine()
    },
    methods: {
        drawLine(){
            // 基于准备好的dom，初始化echarts实例
          this.myChart = this.$echarts.init(document.getElementById('myChart'),null)
            window.onresize = function () {
                this.myChart.resize()
            }
           this.filterHandler()
        },
        handleSetOption() {
            console.log(this.cars)
            // 绘制图表
            this.myChart.setOption({
                title: {  },
                tooltip: {},
                grid:{
                    y2:140
                },
                xAxis: {
                    type:'category',
                    data: this.cars,
                    axisLabel:{
                        interval:0,
                        rotate: -30
                    }
                },
                yAxis: [{
                    type:'value'
                }],
                dataZoom:[
                    {   // 这个dataZoom组件，也控制x轴。
                        type: 'inside', // 这个 dataZoom 组件是 inside 型 dataZoom 组件
                        start: 1,      // 左边在 10% 的位置。
                        end: 60         // 右边在 60% 的位置。
                    },
                    {
                        type: 'inside',
                        yAxisIndex: 0,
                        start: 1,
                        end: 80
                    },
                    // {   // 这个dataZoom组件，默认控制x轴。
                    //         type: 'slider', // 这个 dataZoom 组件是 slider 型 dataZoom 组件
                    //         start: 1,      // 左边在 10% 的位置。
                    //         end: 60         // 右边在 60% 的位置。
                    // },
                    // {
                    //         type: 'slider',
                    //         yAxisIndex: 0,
                    //         start: 1,
                    //         end: 80
                    // },
                ],
                series: [
                    {
                        name: '行驶里程',
                        type: 'bar',
                        label:this.labelStyle,
                        data: this.mileages
                    },
                    {
                        name: '加油量',
                        type: 'bar',
                        label:this.labelStyle,
                        data: this.oilNum
                    },
                    {
                        name: '加油金额',
                        type: 'bar',
                        label:this.labelStyle,
                        data: this.oilPay
                    },
                    {
                        name: '维修金额',
                        type: 'bar',
                        label:this.labelStyle,
                        data: this.repairPay
                    },
                    {
                        name:'油耗  ✖ ️1%',
                        type:'bar',
                        label:this.labelStyle,
                        data:this.per
                    }
                ]
            },true);
            window.onresize = function () {
                this.myChart.resize()
            }
        },
        handleDataArray() {
            this.clearCarsData()
            if (this.tempData.length > 0 ) {
                this.tempData.forEach((item) => {
                    if (item.hasOwnProperty('car')) {
                        this.cars.push(item.car)
                    }
                    if (item.hasOwnProperty('nowNum')) {
                        this.mileages.push(item.nowNum - item.backupNum)
                    }
                    if (item.hasOwnProperty('oilNum')) {
                        this.oilNum.push(item.oilNum)
                    }
                   if (item.hasOwnProperty('pay')) {
                       this.oilPay.push(item.pay)
                   }
                    if (item.hasOwnProperty('repair_pay')) {
                        this.repairPay.push(item.repair_pay)
                    }
                    if (item.hasOwnProperty('per')) {
                        this.per.push((item.per*100).toFixed(0))
                    }
                })
            }
            this.handleSetOption()
        },
        filterHandler() {
            this.tempData = []
            console.log(this.tempClass)
            if (this.tempClass.length > 0) {
                if (this.chartData.length > 0) {
                    this.tempClass.forEach((theClass) => {
                        this.chartData.forEach((item) => {
                            if (theClass === item.class) {
                                this.tempData.push(item)
                            }
                        })
                    })
                }
            } else {
                if (this.chartData.length > 0) {
                    this.tempData = this.chartData
                }
            }
            this.handleDataArray()
        },
        clearCarsData() {
            this.cars=[]
            this.mileages=[]
            this.oilNum=[]
            this.oilPay=[]
            this.repairPay=[]
            this.per=[]
        }
    }
}
</script>

<style scoped>
.myChartStyle {
    width: 100%;
    height: 800px;
}
</style>