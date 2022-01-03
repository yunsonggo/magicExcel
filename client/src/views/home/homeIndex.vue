<template>
  <div>
    <div class="homeParent">
      <el-form
          :model="ruleForm"
          :rules="rules"
          ref="ruleForm"
          label-width="130px"
      >
        <el-form-item label="数据日期类型" prop="dataTag" class="radioItem">
          <el-radio-group v-model="ruleForm.dataTag">
            <el-radio label="1" border size="medium" @change="checkDataTag">单月数据</el-radio>
            <el-radio label="2" border size="medium" @change="checkDataTag">多月数据</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item v-if="ruleForm.dataTag === '2'" label="选择日期区间" prop="monthPicker" class="radioItem">
          <span class="itemStration">不能跨年,否则程序不能处理</span>
          <el-date-picker
              v-model="ruleForm.monthPicker"
              type="monthrange"
              align="right"
              unlink-panels
              range-separator="至"
              start-placeholder="开始月份"
              end-placeholder="结束月份"
              :picker-options="pickerOptions"
              name="excel_file"
              value-format="yyyy_MM"
          >
          </el-date-picker>
        </el-form-item>

        <el-form-item v-else class="radioItem" label="选择日期" prop="monthString">
          <el-date-picker
              v-model="ruleForm.monthString"
              type="month"
              placeholder="选择月"
              value-format="yyyy_MM"
          >
          </el-date-picker>
        </el-form-item>

        <el-form-item class="radioItem" label="数据类型" prop="dataOption">
          <el-select v-model="ruleForm.dataOption" placeholder="请选择数据类型">
            <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item class="radioItem" label="上传Excel">
          <span class="itemStration">注意数据格式</span>
          <el-upload
              class="upload-demo"
              :show-file-list="false"
              :before-upload="beforeUpload"
              :on-success="onSuccess"
              :on-error="onError"
              action="http://192.168.1.136:8090/api/auth/excel/upload"
              :headers="uploadHeader"
              multiple="multiple"
              :data="uploadData"
          >
            <el-button size="small" type="primary">点击上传</el-button>
            <div slot="tip" class="el-upload__tip">只能上传xlsx文件</div>
          </el-upload>
        </el-form-item>
      </el-form>
      <div class="demo-block demo-zh-CN demo-divider" v-if="filePath">
        <div class="source">
          <span class="source">{{successMsg}}条数据</span>
          <el-divider></el-divider>
          <span class="source">查看原始数据:</span>
          <span class="source">{{ filePath }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {reqResubmit} from "../../api";
export default {
  name: "homeIndex",
  data() {
    return {
      uploadHeader:{},
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
      options:[
        {
          value:'1',
          label:'燃油数据',
        },
        {
          value:'2',
          label:'维修数据',
        }
      ],
      ruleForm:{
        dataTag:"",
        options:"",
        dataOption:"",
        monthString:"",
        monthPicker:[],
      },
      rules: {
        dataTag: [
          { required: true, message: "请选择数据日期类型", trigger: "blur" },
        ],
        options: [
          { required: true, message: "请选择数据时间", trigger: "change" },
        ],
        dataOption:[
          {required: true, message: "请选择数据类型", trigger: "change" }
        ]
      },
      filePath: "",
      successMsg:"",
      uploadData:{
        dataTag:"",
        dataOption:"",
        monthString:"",
        monthPicker:[],
      },
      tempTableName:""
    }
  },
  methods:{
    checkDataTag() {
      let text = ''
      if (this.ruleForm.dataTag === '1') {
        text = '单月数据'
      } else if (this.ruleForm.dataTag === '2') {
        text = '多月数据'
      } else {
        text = '出现错误'
      }
      const h = this.$createElement;
      this.$message({
        message: h('p',null,[
            h('span',null,'已选择:'),
            h('i',{style:'color:teal;margin-left:15px'},text)
        ])
      })
    },
    beforeUpload(file) {
      if (this.ruleForm.dataOption === "") {
        this.$message.error("上传之前,请先选择数据类型!");
      }
      if (this.ruleForm.dataTag === "") {
        this.$message.error("上传之前,请先选择数据时间!");
      }
      if (this.ruleForm.dataTag === "1") {
        if (this.monthString === "") {
          this.$message.error("上传之前,请先选择数据日期!");
        } else {
          this.ruleForm.monthPicker = []
        }
      }
      if (this.ruleForm.dataTag === "2") {
        if (this.ruleForm.monthPicker.length < 1) {
          this.$message.error("上传之前,请先选择数据日期!")
        } else {
          this.ruleForm.monthString = ""
        }
      }
      this.uploadData.dataTag = this.ruleForm.dataTag
      this.uploadData.dataOption = this.ruleForm.dataOption
      this.uploadData.monthString = this.ruleForm.monthString
      this.uploadData.monthPicker = this.ruleForm.monthPicker
        let tokenString = localStorage.getItem('token')
        let token = "Bearer " + tokenString
        this.uploadHeader.Authorization = token
      const extXls = file.name.split(".")[1] === "xls";
      const extXlsx = file.name.split(".")[1] === "xlsx";
      if (!extXls && !extXlsx) {
        this.$message.error("上传失败,上传只能是 xls、xlsx格式!");
      }
      return extXls || extXlsx;
    },
    onSuccess(res, file) {
      if (res.code === 1000) {
        this.$message({
          message: "上传成功",
          type: "success",
        });
        console.log(res);
        this.filePath = res.data
        this.successMsg = res.msg
      } else {
        console.log(res,file)
        if (res.code === 2012) {
          console.log("resubmit res msg",res.msg)
          this.$confirm('确定:将会覆盖已有数据,是否继续?','选择日期的数据已经存在',{
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            this.resubmit(res.msg)
          }).catch(() => {
            this.$message({
              type: 'info',
              message: '已取消'
            });
          })
        } else {
          this.$message.error("上传失败:"+res.error);
        }
      }
    },
    onError(err, file) {
      console.log(err,file)
      this.$message.error("上传失败");
    },
    async resubmit(info) {
      const result = await reqResubmit(info.table_name,info.table_option,info.file_path)
      console.log(result)
      if (result.code === 1000) {
        this.successMsg = result.msg
        this.filePath = result.data
      }
    }
  }
}
</script>

<style scoped>
  .homeParent {
    margin: 25px;
    padding: 0;
    display: block;
  }
  .radioItem {
    margin: 15px 15px;
  }
  .itemTitle {
    font-size: 16px;
    color: #333;
    line-height: 40px;
    height: 40px;
    margin: 0;
    padding: 0;
    text-decoration: none;
    display: block;
    position: relative;
    transition: .15s ease-out;
    font-weight: 700;
  }
  .itemStration {
    display: block;
    color: #8492a6;
    font-size: 14px;
    margin-bottom: 20px;
  }
  .uploadForm {

  }
</style>