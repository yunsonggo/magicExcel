import axios from 'axios'

export default function ajax(url = '',data = {}, type = 'GET',token = '') {
    return new Promise(function (resolve,reject) {
        if (token === '') {
            token = localStorage.getItem('token')
        }
        let promise
        if (type === 'GET') {
            let dataStr = ''
            Object.keys(data).forEach(key => {
                dataStr += key + '=' + data[key] + '&'
            })
            if (dataStr !== '') {
                dataStr = dataStr.substring(0,dataStr.lastIndexOf('&'))
                url = url + '?' + dataStr
            }
            promise = axios.get(url)
        } else {
            promise = axios.post(url,data,{
                headers:{
                    'Authorization':'Bearer '+ token
                }
            })
        }
        promise.then(response => {
            resolve(response.data)
        })
            .catch(error => {
                reject(error)
            })
    })
}