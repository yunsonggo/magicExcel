import ajax from './ajax'
const BASE_URL = 'http://192.168.1.136:8080/api'

// export const reqCities = () => ajax(BASE_URL + "/consumer/cities");
// export const reqEmailCode = (consumer_email) => ajax(BASE_URL + '/consumer/email/captcha',{consumer_email},'POST')
// 获取验证码
export const reqCaptcha = () => ajax(BASE_URL + "/captcha/img")
// 用户注册
export const reqRegister = (name,password,check_pass,captcha_id,captcha) => ajax(BASE_URL + "/user/register",{name,password,check_pass,captcha_id,captcha},'POST')
// 用户登录
export const reqLogin = (name,password,captcha_id,captcha) => ajax(BASE_URL + "/user/login",{name,password,captcha_id,captcha},'POST')
// 覆盖已有数据
export const reqResubmit = (table_name,table_option,file_path) => ajax(BASE_URL + '/auth/excel/resubmit',{table_name,table_option,file_path},'POST')
// 获取在线数据
export const reqOnlineSubmit = (table_tag,table_option,month_string,month_picker) => ajax(BASE_URL + '/auth/online/list',{table_tag,table_option,month_string,month_picker},'POST')
// 导出燃油数据
export const reqOutputTable = (query_data) => ajax(BASE_URL + '/auth/online/output/list',{query_data},'POST')
// 修改密码
export const reqResetPassword = (old_pass,pass,check_pass) => ajax(BASE_URL + '/auth/user/reset',{old_pass,pass,check_pass},'POST')