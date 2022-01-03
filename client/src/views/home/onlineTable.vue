<template>
    <div class="content element-doc content">
        <div class="demo-block demo-zh-CN demo-table">
            <div class="source">
                <h3 class="tableTitle">在线数据</h3>
                <div class="options">
                    <div class="optionItem">
                        <el-radio-group v-model="tableTag" style="margin-right: 15px">
                            <el-radio label="1" border size="small" @change="checkTableTag" style="margin-right: unset">
                                单月数据
                            </el-radio>
                            <el-radio label="2" border size="small" @change="checkTableTag">多月数据</el-radio>
                        </el-radio-group>
                        <el-date-picker
                            v-if="tableTag === '2' "
                            v-model="monthPicker"
                            type="monthrange"
                            align="right"
                            unlink-panels
                            range-separator="至"
                            start-placeholder="开始月份"
                            end-placeholder="结束月份"
                            name="excel_file"
                            value-format="yyyy_MM"
                            size="small"
                            style="margin-right: 15px"
                            :picker-options="pickerOptions"
                        >
                        </el-date-picker>
                        <el-date-picker
                            v-else
                            v-model="monthString"
                            type="month"
                            placeholder="选择月"
                            value-format="yyyy_MM"
                            size="small"
                            style="margin-right: 15px"
                        >
                        </el-date-picker>
                        <el-select v-model="tableOption" placeholder="选择数据类型" size="small" style="width: 140px;margin-right: 15px">
                            <el-option
                                v-for="item in tableOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            >
                            </el-option>
                        </el-select>
                    </div>
                    <div class="optionItem"  v-if="classArray.length > 0">
                        <el-select
                            v-model="findClass"
                            multiple
                            placeholder="筛选部门"
                            size="small"
                            style="width: 140px"
                            @change="filterHandler"
                        >
                            <el-option
                                v-for="item in classArray"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            >
                            </el-option>
                        </el-select>
                        <span style="color: #767676;font-size: 14px;margin-left: 10px">(非当前表格部门数据为空)</span>
                    </div>
                    <div class="optionItem">
                        <el-button @click="clearSelect" class="button" size="small" type ="warning" plain>清除选择</el-button>
                        <el-button @click="submitSelect" class="button" size="small" type="primary" plain>获取数据</el-button>
                        <el-button @click="createLimitExcel" class="button" size="small" type="danger" plain>导出当前数据</el-button>
                        <el-button @click="createExcel" class="button" size="small" type="danger" plain>导出所有数据</el-button>
                        <el-button @click="clearFilter" class="button" size="small" type ="info" plain>清除数据</el-button>
                        <el-button v-if="tableData.length > 0" @click="handleChart" class="button" size="small" type ="primary" plain>生成图表</el-button>
                        <el-button v-if="downLoadURL" @click="downLoadExcel" class="button" size="small" type ="success" plain>下载已导出文件</el-button>
                    </div>
                </div>
                <el-table
                    ref="filterTable"
                    :data="tableData"
                    style="width: 100%"
                    :border="true"
                    empty-text="暂无数据"
                    :render-header="renderHeader"
                >
                    <el-table-column
                        v-for="(item,index) in tableHeader"
                        :key="index"
                        :label="item"
                        align="center"
                    >
                        <template slot-scope="scope">
                            <span v-for="(item2,index2) in scope.row" :key="index2" v-if="index2 === index">
                               {{item2}}
                            </span>
                        </template>
                    </el-table-column>
                </el-table>
                <div class="block">
                    <el-pagination
                        @size-change="handleSizeChange"
                        @current-change="handleCurrentChange"
                        :current-page="currentPage"
                        :page-sizes="pageOptions"
                        :page-size="pageOptions[0]"
                        layout="total, sizes, prev, pager, next, jumper"
                        :total="total"
                        :background="true"
                    >
                    </el-pagination>
                </div>
            </div>
        </div>
        <el-drawer
            title="在线图表"
            :visible.sync="drawer"
            :direction="direction"
            size="100%"
            :before-close="handleChartClose">
            <ChartPage :chartData="tableData" :classArray="classArray"></ChartPage>
        </el-drawer>
    </div>
</template>

<script>
import {reqOnlineSubmit,reqOutputTable} from "@/api"
import ChartPage from '@/components/common/chart'
export default {
    name: "onlineTable",
    components:{
        ChartPage
    },
    data() {
        return {
            pickerOptions: {
                shortcuts: [{
                    text: '本月',
                    onClick(picker) {
                        picker.$emit('pick', [new Date(), new Date()]);
                    }
                }, {
                    text: '今年至今',
                    onClick(picker) {
                        const end = new Date();
                        const start = new Date(new Date().getFullYear(), 0);
                        picker.$emit('pick', [start, end]);
                    }
                }, {
                    text: '最近六个月',
                    onClick(picker) {
                        const end = new Date();
                        const start = new Date();
                        start.setMonth(start.getMonth() - 6);
                        picker.$emit('pick', [start, end]);
                    }
                }]
            },
            tableOptions: [{
                value: "1",
                label: "燃油数据"
            }, {
                value: "2",
                label: "维修数据"
            }, {
                value: "3",
                label: "综合数据"
            }],
            tableOption: "",
            tableTag: "",
            monthString: "",
            monthPicker: [],
            queryData:[],
            tableHeader:{class:'部门',car:'车辆'},
            tableData: [],
            classArray:[],
            currentPage: 1,
            total:0,
            pageOptions:[20,50,100,200,500,1000],
            pageOption:0,
            findClass:[],
            downLoadURL:"",
            drawer: false,
            direction: 'rtl',
        }
    },
    computed:{
        savedData() {
            if (this.tableData.length == 0 ||this.$store.getters.queryData.length > 0) {
                this.queryData = this.$store.getters.queryData
                this.total = this.queryData.length
                this.getTableData()
            }
        }
        // this.$store.dispatch("setExample",data)
        // this.$store.getters....
    },
    methods: {
        checkTableTag() {
            let text = ''
            if (this.tableTag === '1') {
                text = '单月数据'
            } else if (this.tableTag === '2') {
                text = '多月数据'
            } else {
                text = '出现错误'
            }
            const h = this.$createElement;
            this.$message({
                message: h('p', null, [
                    h('span', null, '已选择:'),
                    h('i', {style: 'color:teal;margin-left:15px'}, text)
                ])
            })
        },
        clearFilter() {
            this.tableOption = ""
            this.tableTag = ""
            this.monthString = ""
            this.monthPicker = []
            this.currentPage = 1
            this.pageOption = this.pageOptions[0]
            this.total = 0
            this.tableData = []
            this.queryData = []
            this.$refs.filterTable.clearFilter();
            this.findClass = []
            this.tableHeader = {class:'部门',car:'车辆'}
            this.downLoadURL = ""
        },
        clearSelect() {
            this.tableOption = ""
            this.tableTag = ""
            this.monthString = ""
            this.monthPicker = []
            this.currentPage = 1
            this.findClass = []
            this.pageOption = this.pageOptions[0]
            this.$refs.filterTable.clearFilter();
        },
        // formatter(row, column) {
        //     return row.backupNum;
        // },
        filterTag(value, row) {
            return row.tag === value;
        },
        filterHandler() {
            this.getTableData()
        },
        handleSizeChange(val) {
            this.pageOption = val
            this.currentPage = 1
            this.getTableData()
        },
        handleCurrentChange(val) {
            this.currentPage = val
            this.getTableData()
        },
        // 获取数据
        async submitSelect() {
            let msg = ''
            let picker = ''
            if (this.tableTag === '') {
                msg += "请选择 '单月' 或者 '多月' 数据"
            }
            if (this.tableOption === '') {
                msg += "请选择 数据类型 "
            }
            if (this.tableTag === '1') {
                if  (this.monthString === '') {
                    msg += '请选择月份'
                } else {
                    this.monthString = this.tableOption + '_' + this.monthString
                }
            } else if  (this.tableTag === '2') {
                if (this.monthPicker.length === 0) {
                    msg += '请选择月份区间'
                } else {
                    picker = this.tableOption + '_' + this.monthPicker[0]+'_'+this.monthPicker[1]
                }
            } else {
                msg += "请选择 '单月' 或者 '多月' 数据"
            }
            if (msg !== '') {
                this.$message.error(msg)
                return
            }
            const result = await reqOnlineSubmit(this.tableTag,this.tableOption,this.monthString,picker)
            this.clearFilter()
            if (result.code === 1000) {
                Object.entries(result.data).forEach(([key,value]) => {
                    let classFilter = {text:key,value:key}
                    this.classArray.push(classFilter)
                    Object.entries(value).forEach(([k,v]) => {
                        let item = {
                            class:key,
                            car:k,
                        }
                        if (v.hasOwnProperty('backup_num')) {
                            item.backupNum=v.backup_num
                            if (!this.tableHeader.hasOwnProperty('backupNum')) {
                                this.tableHeader.backupNum = '初始里程'
                            }
                        }
                        if (v.hasOwnProperty('now_num')) {
                            item.nowNum=v.now_num
                            if (!this.tableHeader.hasOwnProperty('nowNum')) {
                                this.tableHeader.nowNum = '当前里程'
                            }
                        }
                       if (v.hasOwnProperty('oil_type')) {
                           item.oilType=v.oil_type
                           if (!this.tableHeader.hasOwnProperty('oilType')) {
                               this.tableHeader.oilType = '油品'
                           }
                       }
                        if (v.hasOwnProperty('oil_num')) {
                            item.oilNum=v.oil_num.toFixed(2)
                            if (!this.tableHeader.hasOwnProperty('oilNum')) {
                                this.tableHeader.oilNum = '加油数量'
                            }
                        }
                        if (v.hasOwnProperty('pay')) {
                            item.pay=v.pay.toFixed(2)
                            if (!this.tableHeader.hasOwnProperty('pay')) {
                                this.tableHeader.pay = '加油金额'
                            }
                        }
                        if (v.hasOwnProperty('date_string')) {
                            item.lastDate=v.date_string
                            if (!this.tableHeader.hasOwnProperty('lastDate')) {
                                this.tableHeader.lastDate = '最后加油日期'
                            }
                        }
                        if (v.hasOwnProperty('status')) {
                            item.status=v.status
                            if (!this.tableHeader.hasOwnProperty('status')) {
                                this.tableHeader.status = '加油审批'
                            }
                        }
                        if (v.hasOwnProperty('oil_num') ) {
                            item.per=(v.oil_num / (v.now_num - v.backup_num) *100).toFixed(2)
                            if (!this.tableHeader.hasOwnProperty('per')) {
                                this.tableHeader.per = '百公里油耗'
                            }
                        }
                        if (v.hasOwnProperty('repair_pay')) {
                            item.repair_pay=v.repair_pay.toFixed(2)
                            if (!this.tableHeader.hasOwnProperty('repair_pay')) {
                                this.tableHeader.repair_pay = '维修金额'
                            }
                        }
                        if (v.hasOwnProperty('repair_date_string')) {
                            item.repair_date_string = v.repair_date_string
                            if (!this.tableHeader.hasOwnProperty('repair_date_string')) {
                                this.tableHeader.repair_date_string = '维修日期'
                            }
                        }
                        if (v.hasOwnProperty('repair_status')) {
                            item.repair_status = v.repair_status
                            if (!this.tableHeader.hasOwnProperty('repair_status')) {
                                this.tableHeader.repair_status = '维修审批'
                            }
                        }
                        this.queryData .push(item)
                    })
                })
                this.$store.dispatch("setQueryData",this.queryData)
                this.total = this.queryData.length
                this.getTableData()
                this.$message({
                    message: '获取数据成功',
                    type: 'success'
                });
            } else  {
                this.$message.error('获取数据失败:' + result.msg)
            }
        },
        // 表格分页数据
        getTableData() {
            this.tableData = []
            let tempArr = []
            let minIndex = (this.currentPage - 1) * this.pageOption
            let maxIndex = this.currentPage * this.pageOption - 1
            if (maxIndex >= this.total) {
                maxIndex = this.total - 1
            }
            for  (let i = minIndex ; i <= maxIndex ; i ++ ) {
                tempArr.push(this.queryData[i])
            }
            if (this.findClass.length > 0 ) {
                tempArr.forEach((item) => {
                     this.findClass.forEach((cl) => {
                         if (item.class === cl) {
                             this.tableData .push(item)
                         }
                     })
                 })
            } else {
                this.tableData = tempArr
            }
        },
        // 导出所有数据
        async createExcel() {
            if (this.queryData.length ===  0) {
                this.$message.error('没有数据,无法执行导出')
                return
            }
            const result = await reqOutputTable(this.queryData)
            console.log(result)
            if (result.code === 1000) {
                this.downLoadURL = result.data
                this.$message.success("导出成功")
            } else {
                this.$message.error("导出数据失败,请刷新页面重试:"+result.msg)
            }
        },
        // 导出当前数据
        async createLimitExcel() {
            if (this.tableData.length ===  0) {
                this.$message.error('没有数据,无法执行导出')
                return
            }
            const result = await reqOutputTable(this.tableData)
            console.log(result)
            if (result.code === 1000) {
                this.downLoadURL = result.data
                this.$message.success("导出成功")
            } else {
                this.$message.error("导出数据失败,请刷新页面重试:"+result.msg)
            }
        },
        // 动态渲染表头/列
        renderHeader(h,{column,index}) {
            let l = column.label.length
            let f = 16;
            column.minWidth = f * l
            return h(
                'span',
                {
                    class:'table-head',
                    style:{width:'100%'}
                },
                [column.label]
            )
        },
        // 下载已导出文件
        downLoadExcel() {
            if (this.downLoadURL === "") {
                this.$message.error("下载地址失效,请重新导出数据")
            } else {
                const a = document.createElement('a')
                a.download = this.tableOption +'_'+this.monthString+'_'+this.monthPicker
                a.style.display = 'none'
                a.href = this.downLoadURL
                document.body.appendChild(a)
                a.click()
                document.body.removeChild(a)
            }
        },
        // 生成图表
        handleChart() {
           // this.$router.push('/home/chart')
            this.drawer = true
        },
        handleChartClose(done) {
            this.$confirm('确认关闭？')
                .then(_ => {
                    done();
                })
                .catch(_ => {});
        }
    }
}
</script>

<style scoped>
.content {
    border-radius: 3px;
}

.demo-block {
    margin-bottom: 24px;
    display: block;
    transition: .2s;
}

.source {
    padding: 24px;
}

.source .button {
    margin-bottom: 10px;
}

.source .tableTitle {
    display: block;
    font-size: 1.17em;
    margin-block-start: 1em;
    margin-block-end: 1em;
    margin-inline-start: 0px;
    margin-inline-end: 0px;
    font-weight: bold;
}

.block {
    padding: 8px;
    background-color: #ffffff;
}

.options {
    display: block;
    padding: 10px;
    border: 1px solid #ebebeb;
    background-color: #ffffff;
}

.options .optionItem {
    margin: 10px 15px;
}
</style>